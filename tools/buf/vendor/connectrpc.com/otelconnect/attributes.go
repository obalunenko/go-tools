// Copyright 2022-2025 The Connect Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package otelconnect

import (
	"fmt"
	"net"
	"net/http"
	"strconv"
	"strings"

	"connectrpc.com/connect"
	"go.opentelemetry.io/otel/attribute"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
)

// AttributeFilter is used to filter attributes out based on the [connect.Spec]
// and [attribute.KeyValue]. If the filter returns true the attribute will be
// kept else it will be removed. AttributeFilter must be safe to call concurrently.
type AttributeFilter func(connect.Spec, attribute.KeyValue) bool

func (filter AttributeFilter) filter(spec connect.Spec, values ...attribute.KeyValue) []attribute.KeyValue {
	if filter == nil {
		return values
	}
	// Assign a new slice of zero length with the same underlying
	// array as the values slice. This avoids unnecessary memory allocations.
	filteredValues := values[:0]
	for _, attr := range values {
		if filter(spec, attr) {
			filteredValues = append(filteredValues, attr)
		}
	}
	clear(values[len(filteredValues):])
	return filteredValues
}

func addProcedureAttributes(attrs []attribute.KeyValue, procedure string) []attribute.KeyValue {
	svc, method, ok := strings.Cut(procedure, "/")
	if !ok {
		// fall back to treating the whole string as the method
		return append(attrs, semconv.RPCMethodKey.String(procedure))
	}
	if svc != "" {
		attrs = append(attrs, semconv.RPCServiceKey.String(svc))
	}
	if method != "" {
		attrs = append(attrs, semconv.RPCMethodKey.String(method))
	}
	return attrs
}

func addRequestAttributes(attrs []attribute.KeyValue, spec connect.Spec, peer connect.Peer) []attribute.KeyValue {
	if addr := peer.Addr; addr != "" {
		attrs = addAddressAttributes(attrs, addr)
	}
	name := strings.TrimLeft(spec.Procedure, "/")
	protocol := protocolToSemConv(peer.Protocol)
	attrs = append(attrs, semconv.RPCSystemKey.String(protocol))
	attrs = addProcedureAttributes(attrs, name)
	return attrs
}

func addAddressAttributes(attrs []attribute.KeyValue, address string) []attribute.KeyValue {
	if host, port, err := net.SplitHostPort(address); err == nil {
		portInt, err := strconv.Atoi(port)
		if err == nil {
			return append(attrs,
				semconv.NetPeerNameKey.String(host),
				semconv.NetPeerPortKey.Int(portInt),
			)
		}
	}
	return append(attrs, semconv.NetPeerNameKey.String(address))
}

func statusCodeAttribute(protocol string, serverErr error) (attribute.KeyValue, bool) {
	// Following the respective specifications, use integers and "status_code" for
	// gRPC codes in contrast to strings and "error_code" for Connect codes.
	switch protocol {
	case grpcProtocol, grpcwebProtocol:
		codeKey := attribute.Key("rpc." + protocol + ".status_code")
		if serverErr != nil {
			return codeKey.Int64(int64(connect.CodeOf(serverErr))), true
		}
		return codeKey.Int64(0), true // gRPC uses 0 for success
	case connectProtocol:
		if connect.IsNotModifiedError(serverErr) {
			// A "not modified" error is special: it's code is technically "unknown" but
			// it would be misleading to label it as an unknown error since it's not really
			// an error, but rather a sentinel to trigger a "304 Not Modified" HTTP status.
			return semconv.HTTPStatusCodeKey.Int(http.StatusNotModified), true
		}
		if serverErr != nil {
			codeKey := attribute.Key("rpc." + protocol + ".error_code")
			return codeKey.String(connect.CodeOf(serverErr).String()), true
		}
	}
	return attribute.KeyValue{}, false
}

func headerAttributes(protocol, eventType string, metadata http.Header, allowedKeys []string) []attribute.KeyValue {
	attributes := make([]attribute.KeyValue, 0, len(allowedKeys))
	return addHeaderAttributes(attributes, protocol, eventType, metadata, allowedKeys)
}

func addHeaderAttributes(attributes []attribute.KeyValue, protocol, eventType string, metadata http.Header, allowedKeys []string) []attribute.KeyValue {
	for _, allowedKey := range allowedKeys {
		if val, ok := metadata[allowedKey]; ok {
			keyValue := attribute.StringSlice(
				formatHeaderAttributeKey(protocol, eventType, allowedKey),
				val,
			)
			attributes = append(attributes, keyValue)
		}
	}
	return attributes
}

// formatHeaderAttributeKey formats header attributes as suggested by the OpenTelemetry specification:
// https://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/trace/semantic_conventions/rpc.md#grpc-request-and-response-metadata
func formatHeaderAttributeKey(protocol, eventType, key string) string {
	key = strings.ReplaceAll(strings.ToLower(key), "-", "_")
	return fmt.Sprintf("rpc.%s.%s.metadata.%s", protocol, eventType, key)
}
