// Copyright 2023-2024 Buf Technologies, Inc.
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

// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: buf/registry/plugin/v1beta1/collection_service.proto

package pluginv1beta1connect

import (
	v1beta1 "buf.build/gen/go/bufbuild/registry/protocolbuffers/go/buf/registry/plugin/v1beta1"
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// CollectionServiceName is the fully-qualified name of the CollectionService service.
	CollectionServiceName = "buf.registry.plugin.v1beta1.CollectionService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// CollectionServiceGetCollectionsProcedure is the fully-qualified name of the CollectionService's
	// GetCollections RPC.
	CollectionServiceGetCollectionsProcedure = "/buf.registry.plugin.v1beta1.CollectionService/GetCollections"
	// CollectionServiceListCollectionsProcedure is the fully-qualified name of the CollectionService's
	// ListCollections RPC.
	CollectionServiceListCollectionsProcedure = "/buf.registry.plugin.v1beta1.CollectionService/ListCollections"
	// CollectionServiceGetPluginCollectionAssociationsProcedure is the fully-qualified name of the
	// CollectionService's GetPluginCollectionAssociations RPC.
	CollectionServiceGetPluginCollectionAssociationsProcedure = "/buf.registry.plugin.v1beta1.CollectionService/GetPluginCollectionAssociations"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	collectionServiceServiceDescriptor                               = v1beta1.File_buf_registry_plugin_v1beta1_collection_service_proto.Services().ByName("CollectionService")
	collectionServiceGetCollectionsMethodDescriptor                  = collectionServiceServiceDescriptor.Methods().ByName("GetCollections")
	collectionServiceListCollectionsMethodDescriptor                 = collectionServiceServiceDescriptor.Methods().ByName("ListCollections")
	collectionServiceGetPluginCollectionAssociationsMethodDescriptor = collectionServiceServiceDescriptor.Methods().ByName("GetPluginCollectionAssociations")
)

// CollectionServiceClient is a client for the buf.registry.plugin.v1beta1.CollectionService
// service.
type CollectionServiceClient interface {
	// Get Collections.
	GetCollections(context.Context, *connect.Request[v1beta1.GetCollectionsRequest]) (*connect.Response[v1beta1.GetCollectionsResponse], error)
	// List Collections for a given Plugin.
	ListCollections(context.Context, *connect.Request[v1beta1.ListCollectionsRequest]) (*connect.Response[v1beta1.ListCollectionsResponse], error)
	// Get the Collections for the given Plugins.
	GetPluginCollectionAssociations(context.Context, *connect.Request[v1beta1.GetPluginCollectionAssociationsRequest]) (*connect.Response[v1beta1.GetPluginCollectionAssociationsResponse], error)
}

// NewCollectionServiceClient constructs a client for the
// buf.registry.plugin.v1beta1.CollectionService service. By default, it uses the Connect protocol
// with the binary Protobuf Codec, asks for gzipped responses, and sends uncompressed requests. To
// use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or connect.WithGRPCWeb()
// options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewCollectionServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) CollectionServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &collectionServiceClient{
		getCollections: connect.NewClient[v1beta1.GetCollectionsRequest, v1beta1.GetCollectionsResponse](
			httpClient,
			baseURL+CollectionServiceGetCollectionsProcedure,
			connect.WithSchema(collectionServiceGetCollectionsMethodDescriptor),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
		listCollections: connect.NewClient[v1beta1.ListCollectionsRequest, v1beta1.ListCollectionsResponse](
			httpClient,
			baseURL+CollectionServiceListCollectionsProcedure,
			connect.WithSchema(collectionServiceListCollectionsMethodDescriptor),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
		getPluginCollectionAssociations: connect.NewClient[v1beta1.GetPluginCollectionAssociationsRequest, v1beta1.GetPluginCollectionAssociationsResponse](
			httpClient,
			baseURL+CollectionServiceGetPluginCollectionAssociationsProcedure,
			connect.WithSchema(collectionServiceGetPluginCollectionAssociationsMethodDescriptor),
			connect.WithIdempotency(connect.IdempotencyNoSideEffects),
			connect.WithClientOptions(opts...),
		),
	}
}

// collectionServiceClient implements CollectionServiceClient.
type collectionServiceClient struct {
	getCollections                  *connect.Client[v1beta1.GetCollectionsRequest, v1beta1.GetCollectionsResponse]
	listCollections                 *connect.Client[v1beta1.ListCollectionsRequest, v1beta1.ListCollectionsResponse]
	getPluginCollectionAssociations *connect.Client[v1beta1.GetPluginCollectionAssociationsRequest, v1beta1.GetPluginCollectionAssociationsResponse]
}

// GetCollections calls buf.registry.plugin.v1beta1.CollectionService.GetCollections.
func (c *collectionServiceClient) GetCollections(ctx context.Context, req *connect.Request[v1beta1.GetCollectionsRequest]) (*connect.Response[v1beta1.GetCollectionsResponse], error) {
	return c.getCollections.CallUnary(ctx, req)
}

// ListCollections calls buf.registry.plugin.v1beta1.CollectionService.ListCollections.
func (c *collectionServiceClient) ListCollections(ctx context.Context, req *connect.Request[v1beta1.ListCollectionsRequest]) (*connect.Response[v1beta1.ListCollectionsResponse], error) {
	return c.listCollections.CallUnary(ctx, req)
}

// GetPluginCollectionAssociations calls
// buf.registry.plugin.v1beta1.CollectionService.GetPluginCollectionAssociations.
func (c *collectionServiceClient) GetPluginCollectionAssociations(ctx context.Context, req *connect.Request[v1beta1.GetPluginCollectionAssociationsRequest]) (*connect.Response[v1beta1.GetPluginCollectionAssociationsResponse], error) {
	return c.getPluginCollectionAssociations.CallUnary(ctx, req)
}

// CollectionServiceHandler is an implementation of the
// buf.registry.plugin.v1beta1.CollectionService service.
type CollectionServiceHandler interface {
	// Get Collections.
	GetCollections(context.Context, *connect.Request[v1beta1.GetCollectionsRequest]) (*connect.Response[v1beta1.GetCollectionsResponse], error)
	// List Collections for a given Plugin.
	ListCollections(context.Context, *connect.Request[v1beta1.ListCollectionsRequest]) (*connect.Response[v1beta1.ListCollectionsResponse], error)
	// Get the Collections for the given Plugins.
	GetPluginCollectionAssociations(context.Context, *connect.Request[v1beta1.GetPluginCollectionAssociationsRequest]) (*connect.Response[v1beta1.GetPluginCollectionAssociationsResponse], error)
}

// NewCollectionServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewCollectionServiceHandler(svc CollectionServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	collectionServiceGetCollectionsHandler := connect.NewUnaryHandler(
		CollectionServiceGetCollectionsProcedure,
		svc.GetCollections,
		connect.WithSchema(collectionServiceGetCollectionsMethodDescriptor),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	collectionServiceListCollectionsHandler := connect.NewUnaryHandler(
		CollectionServiceListCollectionsProcedure,
		svc.ListCollections,
		connect.WithSchema(collectionServiceListCollectionsMethodDescriptor),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	collectionServiceGetPluginCollectionAssociationsHandler := connect.NewUnaryHandler(
		CollectionServiceGetPluginCollectionAssociationsProcedure,
		svc.GetPluginCollectionAssociations,
		connect.WithSchema(collectionServiceGetPluginCollectionAssociationsMethodDescriptor),
		connect.WithIdempotency(connect.IdempotencyNoSideEffects),
		connect.WithHandlerOptions(opts...),
	)
	return "/buf.registry.plugin.v1beta1.CollectionService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case CollectionServiceGetCollectionsProcedure:
			collectionServiceGetCollectionsHandler.ServeHTTP(w, r)
		case CollectionServiceListCollectionsProcedure:
			collectionServiceListCollectionsHandler.ServeHTTP(w, r)
		case CollectionServiceGetPluginCollectionAssociationsProcedure:
			collectionServiceGetPluginCollectionAssociationsHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedCollectionServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedCollectionServiceHandler struct{}

func (UnimplementedCollectionServiceHandler) GetCollections(context.Context, *connect.Request[v1beta1.GetCollectionsRequest]) (*connect.Response[v1beta1.GetCollectionsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("buf.registry.plugin.v1beta1.CollectionService.GetCollections is not implemented"))
}

func (UnimplementedCollectionServiceHandler) ListCollections(context.Context, *connect.Request[v1beta1.ListCollectionsRequest]) (*connect.Response[v1beta1.ListCollectionsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("buf.registry.plugin.v1beta1.CollectionService.ListCollections is not implemented"))
}

func (UnimplementedCollectionServiceHandler) GetPluginCollectionAssociations(context.Context, *connect.Request[v1beta1.GetPluginCollectionAssociationsRequest]) (*connect.Response[v1beta1.GetPluginCollectionAssociationsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("buf.registry.plugin.v1beta1.CollectionService.GetPluginCollectionAssociations is not implemented"))
}