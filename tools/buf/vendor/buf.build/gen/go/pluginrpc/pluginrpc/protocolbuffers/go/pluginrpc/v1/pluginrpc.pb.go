// Copyright 2024 Buf Technologies, Inc.
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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        (unknown)
// source: pluginrpc/v1/pluginrpc.proto

// Interfaces for the pluginrpc system.
//
// These interfaces define how requests and responses and sent over stdin and stdout, as well
// as define how procedures are specified.
//
// These interfaces will work if a given plugin responds to `--protocol` with `1`.

//go:build !protoopaque

package pluginrpcv1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// A code for an error.
//
// Matches with google.rpc.Code.
type Code int32

const (
	// The zero value.
	Code_CODE_UNSPECIFIED Code = 0
	// Indicates that the operation was canceled, typically by the
	Code_CODE_CANCELED Code = 1
	// Indicates that the operation failed for an unknown reason.
	Code_CODE_UNKNOWN Code = 2
	// Indicates that client supplied an invalid argument.
	Code_CODE_INVALID_ARGUMENT Code = 3
	// Indicates that deadline expired before the operation could complete.
	Code_CODE_DEADLINE_EXCEEDED Code = 4
	// Indicates that some requested entity (for example, a file or directory)
	// was not found.
	Code_CODE_NOT_FOUND Code = 5
	// Indicates that client attempted to create an entity (for example, a file
	// or directory) that already exists.
	Code_CODE_ALREADY_EXISTS Code = 6
	// Indicates that the caller doesn't have permission to execute the
	// specified operation.
	Code_CODE_PERMISSION_DENIED Code = 7
	// Indicates that some resource has been exhausted. For example, a per-user
	// quota may be exhausted or the entire file system may be full.
	Code_CODE_RESOURCE_EXHAUSTED Code = 8
	// Indicates that the system is not in a state required for the operation's execution.
	Code_CODE_FAILED_PRECONDITION Code = 9
	// Indicates that operation was aborted by the system, usually because of a
	// concurrency issue such as a sequencer check failure or transaction abort.
	Code_CODE_ABORTED Code = 10
	// Indicates that the operation was attempted past the valid range (for example,
	// seeking past end-of-file).
	Code_CODE_OUT_OF_RANGE Code = 11
	// Indicates that the operation isn't implemented, supported, or enabled in this service.
	Code_CODE_UNIMPLEMENTED Code = 12
	// Indicates that some invariants expected by the underlying system have been broken.
	// This code is reserved for serious errors.
	Code_CODE_INTERNAL Code = 13
	// Indicates that the service is currently unavailable. This is usually temporary, so
	// clients can back off and retry idempotent operations.
	Code_CODE_UNAVAILABLE Code = 14
	// Indicates that the operation has resulted in unrecoverable data loss or corruption.
	Code_CODE_DATA_LOSS Code = 15
	// Indicates that the request does not have valid authentication credentials for the operation.
	Code_CODE_UNAUTHENTICATED Code = 16
)

// Enum value maps for Code.
var (
	Code_name = map[int32]string{
		0:  "CODE_UNSPECIFIED",
		1:  "CODE_CANCELED",
		2:  "CODE_UNKNOWN",
		3:  "CODE_INVALID_ARGUMENT",
		4:  "CODE_DEADLINE_EXCEEDED",
		5:  "CODE_NOT_FOUND",
		6:  "CODE_ALREADY_EXISTS",
		7:  "CODE_PERMISSION_DENIED",
		8:  "CODE_RESOURCE_EXHAUSTED",
		9:  "CODE_FAILED_PRECONDITION",
		10: "CODE_ABORTED",
		11: "CODE_OUT_OF_RANGE",
		12: "CODE_UNIMPLEMENTED",
		13: "CODE_INTERNAL",
		14: "CODE_UNAVAILABLE",
		15: "CODE_DATA_LOSS",
		16: "CODE_UNAUTHENTICATED",
	}
	Code_value = map[string]int32{
		"CODE_UNSPECIFIED":         0,
		"CODE_CANCELED":            1,
		"CODE_UNKNOWN":             2,
		"CODE_INVALID_ARGUMENT":    3,
		"CODE_DEADLINE_EXCEEDED":   4,
		"CODE_NOT_FOUND":           5,
		"CODE_ALREADY_EXISTS":      6,
		"CODE_PERMISSION_DENIED":   7,
		"CODE_RESOURCE_EXHAUSTED":  8,
		"CODE_FAILED_PRECONDITION": 9,
		"CODE_ABORTED":             10,
		"CODE_OUT_OF_RANGE":        11,
		"CODE_UNIMPLEMENTED":       12,
		"CODE_INTERNAL":            13,
		"CODE_UNAVAILABLE":         14,
		"CODE_DATA_LOSS":           15,
		"CODE_UNAUTHENTICATED":     16,
	}
)

func (x Code) Enum() *Code {
	p := new(Code)
	*p = x
	return p
}

func (x Code) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Code) Descriptor() protoreflect.EnumDescriptor {
	return file_pluginrpc_v1_pluginrpc_proto_enumTypes[0].Descriptor()
}

func (Code) Type() protoreflect.EnumType {
	return &file_pluginrpc_v1_pluginrpc_proto_enumTypes[0]
}

func (x Code) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// A specification of a single procedure that can be invoked within a plugin.
//
// A prodecure has a path, and the args used to invoke it via the plugin.
type Procedure struct {
	state protoimpl.MessageState `protogen:"hybrid.v1"`
	// The path of the procedure.
	//
	// When mapped to a Protobuf method, this will be `/fully.qualified.package.Service/Method`.
	//
	// Example:
	//
	//	package buf.plugin.check.v1;
	//
	//	service LintService {
	//	  rpc Lint(LintRequest) returns (LintResponse);
	//	}
	//
	// The path would be `/buf.plugin.check.v1.LintService/Lint`.
	//
	// The path must be a valid URI.
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// The args to invoke on the plugin to invoke this method per the protocol specification.
	//
	// Example: If the args are `["plugin, "lint"]`, this method would be invoked by calling
	// the binary with the args `plugin lint`.
	//
	// This is optional. If not set, the default is that the procedure can be called
	// by calling the path as the only argument.
	//
	// Args must be at least of length 2, may only consist of characters in [a-zA-Z0-9_-], and may not
	// start or end with a dash or underscore.
	Args          []string `protobuf:"bytes,2,rep,name=args,proto3" json:"args,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Procedure) Reset() {
	*x = Procedure{}
	mi := &file_pluginrpc_v1_pluginrpc_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Procedure) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Procedure) ProtoMessage() {}

func (x *Procedure) ProtoReflect() protoreflect.Message {
	mi := &file_pluginrpc_v1_pluginrpc_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *Procedure) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Procedure) GetArgs() []string {
	if x != nil {
		return x.Args
	}
	return nil
}

func (x *Procedure) SetPath(v string) {
	x.Path = v
}

func (x *Procedure) SetArgs(v []string) {
	x.Args = v
}

type Procedure_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// The path of the procedure.
	//
	// When mapped to a Protobuf method, this will be `/fully.qualified.package.Service/Method`.
	//
	// Example:
	//
	//	package buf.plugin.check.v1;
	//
	//	service LintService {
	//	  rpc Lint(LintRequest) returns (LintResponse);
	//	}
	//
	// The path would be `/buf.plugin.check.v1.LintService/Lint`.
	//
	// The path must be a valid URI.
	Path string
	// The args to invoke on the plugin to invoke this method per the protocol specification.
	//
	// Example: If the args are `["plugin, "lint"]`, this method would be invoked by calling
	// the binary with the args `plugin lint`.
	//
	// This is optional. If not set, the default is that the procedure can be called
	// by calling the path as the only argument.
	//
	// Args must be at least of length 2, may only consist of characters in [a-zA-Z0-9_-], and may not
	// start or end with a dash or underscore.
	Args []string
}

func (b0 Procedure_builder) Build() *Procedure {
	m0 := &Procedure{}
	b, x := &b0, m0
	_, _ = b, x
	x.Path = b.Path
	x.Args = b.Args
	return m0
}

// The response given when the `--spec` flag is passed to the plugin.
type Spec struct {
	state protoimpl.MessageState `protogen:"hybrid.v1"`
	// The specifications of the Procedures implemented.
	//
	// All Procedures must have a unique path and args. The latter is unenforceable
	// via protovalidate, but users should assume that pluginrpc implementations
	// will enforce that args are unique.
	Procedures    []*Procedure `protobuf:"bytes,1,rep,name=procedures,proto3" json:"procedures,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Spec) Reset() {
	*x = Spec{}
	mi := &file_pluginrpc_v1_pluginrpc_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Spec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Spec) ProtoMessage() {}

func (x *Spec) ProtoReflect() protoreflect.Message {
	mi := &file_pluginrpc_v1_pluginrpc_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *Spec) GetProcedures() []*Procedure {
	if x != nil {
		return x.Procedures
	}
	return nil
}

func (x *Spec) SetProcedures(v []*Procedure) {
	x.Procedures = v
}

type Spec_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// The specifications of the Procedures implemented.
	//
	// All Procedures must have a unique path and args. The latter is unenforceable
	// via protovalidate, but users should assume that pluginrpc implementations
	// will enforce that args are unique.
	Procedures []*Procedure
}

func (b0 Spec_builder) Build() *Spec {
	m0 := &Spec{}
	b, x := &b0, m0
	_, _ = b, x
	x.Procedures = b.Procedures
	return m0
}

// A request sent over a transport.
type Request struct {
	state protoimpl.MessageState `protogen:"hybrid.v1"`
	// The request value.
	//
	// May not be present.
	Value         *anypb.Any `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Request) Reset() {
	*x = Request{}
	mi := &file_pluginrpc_v1_pluginrpc_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_pluginrpc_v1_pluginrpc_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *Request) GetValue() *anypb.Any {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *Request) SetValue(v *anypb.Any) {
	x.Value = v
}

func (x *Request) HasValue() bool {
	if x == nil {
		return false
	}
	return x.Value != nil
}

func (x *Request) ClearValue() {
	x.Value = nil
}

type Request_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// The request value.
	//
	// May not be present.
	Value *anypb.Any
}

func (b0 Request_builder) Build() *Request {
	m0 := &Request{}
	b, x := &b0, m0
	_, _ = b, x
	x.Value = b.Value
	return m0
}

// A response received over a transport.
type Response struct {
	state protoimpl.MessageState `protogen:"hybrid.v1"`
	// The response value.
	//
	// May or may not be present, regardless of if there is an Error.
	Value *anypb.Any `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	// An error.
	//
	// May or may not be present, regardless of if there is a response value.
	Error         *Error `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Response) Reset() {
	*x = Response{}
	mi := &file_pluginrpc_v1_pluginrpc_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_pluginrpc_v1_pluginrpc_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *Response) GetValue() *anypb.Any {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *Response) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *Response) SetValue(v *anypb.Any) {
	x.Value = v
}

func (x *Response) SetError(v *Error) {
	x.Error = v
}

func (x *Response) HasValue() bool {
	if x == nil {
		return false
	}
	return x.Value != nil
}

func (x *Response) HasError() bool {
	if x == nil {
		return false
	}
	return x.Error != nil
}

func (x *Response) ClearValue() {
	x.Value = nil
}

func (x *Response) ClearError() {
	x.Error = nil
}

type Response_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// The response value.
	//
	// May or may not be present, regardless of if there is an Error.
	Value *anypb.Any
	// An error.
	//
	// May or may not be present, regardless of if there is a response value.
	Error *Error
}

func (b0 Response_builder) Build() *Response {
	m0 := &Response{}
	b, x := &b0, m0
	_, _ = b, x
	x.Value = b.Value
	x.Error = b.Error
	return m0
}

// An error received over a transport as part of a Response.
type Error struct {
	state protoimpl.MessageState `protogen:"hybrid.v1"`
	// The error code.
	Code Code `protobuf:"varint,1,opt,name=code,proto3,enum=pluginrpc.v1.Code" json:"code,omitempty"`
	// The message of the error.
	Message       string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Error) Reset() {
	*x = Error{}
	mi := &file_pluginrpc_v1_pluginrpc_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_pluginrpc_v1_pluginrpc_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *Error) GetCode() Code {
	if x != nil {
		return x.Code
	}
	return Code_CODE_UNSPECIFIED
}

func (x *Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Error) SetCode(v Code) {
	x.Code = v
}

func (x *Error) SetMessage(v string) {
	x.Message = v
}

type Error_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// The error code.
	Code Code
	// The message of the error.
	Message string
}

func (b0 Error_builder) Build() *Error {
	m0 := &Error{}
	b, x := &b0, m0
	_, _ = b, x
	x.Code = b.Code
	x.Message = b.Message
	return m0
}

var File_pluginrpc_v1_pluginrpc_proto protoreflect.FileDescriptor

var file_pluginrpc_v1_pluginrpc_proto_rawDesc = string([]byte{
	0x0a, 0x1c, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x70,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x72, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x62, 0x75,
	0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x74, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x64, 0x75, 0x72,
	0x65, 0x12, 0x1f, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x0b, 0xba, 0x48, 0x08, 0xc8, 0x01, 0x01, 0x72, 0x03, 0x88, 0x01, 0x01, 0x52, 0x04, 0x70, 0x61,
	0x74, 0x68, 0x12, 0x46, 0x0a, 0x04, 0x61, 0x72, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09,
	0x42, 0x32, 0xba, 0x48, 0x2f, 0x92, 0x01, 0x2c, 0x22, 0x2a, 0x72, 0x28, 0x32, 0x26, 0x5e, 0x5b,
	0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a, 0x30, 0x2d, 0x39, 0x5d, 0x5b, 0x61, 0x2d, 0x7a, 0x41, 0x2d,
	0x5a, 0x30, 0x2d, 0x39, 0x5f, 0x2d, 0x5d, 0x2a, 0x5b, 0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a, 0x30,
	0x2d, 0x39, 0x5d, 0x24, 0x52, 0x04, 0x61, 0x72, 0x67, 0x73, 0x22, 0xa8, 0x01, 0x0a, 0x04, 0x53,
	0x70, 0x65, 0x63, 0x12, 0x9f, 0x01, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x64, 0x75, 0x72,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x72, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x64, 0x75, 0x72,
	0x65, 0x42, 0x66, 0xba, 0x48, 0x63, 0xba, 0x01, 0x5b, 0x0a, 0x0b, 0x70, 0x61, 0x74, 0x68, 0x5f,
	0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x12, 0x1e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x64, 0x75, 0x72,
	0x65, 0x20, 0x70, 0x61, 0x74, 0x68, 0x73, 0x20, 0x6d, 0x75, 0x73, 0x74, 0x20, 0x62, 0x65, 0x20,
	0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x1a, 0x2c, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x6d, 0x61, 0x70,
	0x28, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x64, 0x75, 0x72, 0x65, 0x2c, 0x20, 0x70, 0x72, 0x6f, 0x63,
	0x65, 0x64, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x61, 0x74, 0x68, 0x29, 0x2e, 0x75, 0x6e, 0x69, 0x71,
	0x75, 0x65, 0x28, 0x29, 0x92, 0x01, 0x02, 0x08, 0x01, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x63, 0x65,
	0x64, 0x75, 0x72, 0x65, 0x73, 0x22, 0x35, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x61, 0x0a, 0x08,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x29, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x72, 0x70, 0x63, 0x2e,
	0x76, 0x31, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22,
	0x5c, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x31, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x72,
	0x70, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x42, 0x09, 0xba, 0x48, 0x06, 0x82,
	0x01, 0x03, 0x22, 0x01, 0x00, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48,
	0x03, 0xc8, 0x01, 0x01, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2a, 0x94, 0x03,
	0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d,
	0x43, 0x4f, 0x44, 0x45, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x12,
	0x10, 0x0a, 0x0c, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10,
	0x02, 0x12, 0x19, 0x0a, 0x15, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49,
	0x44, 0x5f, 0x41, 0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x03, 0x12, 0x1a, 0x0a, 0x16,
	0x43, 0x4f, 0x44, 0x45, 0x5f, 0x44, 0x45, 0x41, 0x44, 0x4c, 0x49, 0x4e, 0x45, 0x5f, 0x45, 0x58,
	0x43, 0x45, 0x45, 0x44, 0x45, 0x44, 0x10, 0x04, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x4f, 0x44, 0x45,
	0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x05, 0x12, 0x17, 0x0a, 0x13,
	0x43, 0x4f, 0x44, 0x45, 0x5f, 0x41, 0x4c, 0x52, 0x45, 0x41, 0x44, 0x59, 0x5f, 0x45, 0x58, 0x49,
	0x53, 0x54, 0x53, 0x10, 0x06, 0x12, 0x1a, 0x0a, 0x16, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x50, 0x45,
	0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x44, 0x45, 0x4e, 0x49, 0x45, 0x44, 0x10,
	0x07, 0x12, 0x1b, 0x0a, 0x17, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52,
	0x43, 0x45, 0x5f, 0x45, 0x58, 0x48, 0x41, 0x55, 0x53, 0x54, 0x45, 0x44, 0x10, 0x08, 0x12, 0x1c,
	0x0a, 0x18, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x5f, 0x50, 0x52,
	0x45, 0x43, 0x4f, 0x4e, 0x44, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x09, 0x12, 0x10, 0x0a, 0x0c,
	0x43, 0x4f, 0x44, 0x45, 0x5f, 0x41, 0x42, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x10, 0x0a, 0x12, 0x15,
	0x0a, 0x11, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x4f, 0x55, 0x54, 0x5f, 0x4f, 0x46, 0x5f, 0x52, 0x41,
	0x4e, 0x47, 0x45, 0x10, 0x0b, 0x12, 0x16, 0x0a, 0x12, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x55, 0x4e,
	0x49, 0x4d, 0x50, 0x4c, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x45, 0x44, 0x10, 0x0c, 0x12, 0x11, 0x0a,
	0x0d, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x10, 0x0d,
	0x12, 0x14, 0x0a, 0x10, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x55, 0x4e, 0x41, 0x56, 0x41, 0x49, 0x4c,
	0x41, 0x42, 0x4c, 0x45, 0x10, 0x0e, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x44,
	0x41, 0x54, 0x41, 0x5f, 0x4c, 0x4f, 0x53, 0x53, 0x10, 0x0f, 0x12, 0x18, 0x0a, 0x14, 0x43, 0x4f,
	0x44, 0x45, 0x5f, 0x55, 0x4e, 0x41, 0x55, 0x54, 0x48, 0x45, 0x4e, 0x54, 0x49, 0x43, 0x41, 0x54,
	0x45, 0x44, 0x10, 0x10, 0x42, 0x52, 0x5a, 0x50, 0x62, 0x75, 0x66, 0x2e, 0x62, 0x75, 0x69, 0x6c,
	0x64, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x72,
	0x70, 0x63, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x73, 0x2f, 0x67, 0x6f, 0x2f,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x72, 0x70, 0x63, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var file_pluginrpc_v1_pluginrpc_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pluginrpc_v1_pluginrpc_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pluginrpc_v1_pluginrpc_proto_goTypes = []any{
	(Code)(0),         // 0: pluginrpc.v1.Code
	(*Procedure)(nil), // 1: pluginrpc.v1.Procedure
	(*Spec)(nil),      // 2: pluginrpc.v1.Spec
	(*Request)(nil),   // 3: pluginrpc.v1.Request
	(*Response)(nil),  // 4: pluginrpc.v1.Response
	(*Error)(nil),     // 5: pluginrpc.v1.Error
	(*anypb.Any)(nil), // 6: google.protobuf.Any
}
var file_pluginrpc_v1_pluginrpc_proto_depIdxs = []int32{
	1, // 0: pluginrpc.v1.Spec.procedures:type_name -> pluginrpc.v1.Procedure
	6, // 1: pluginrpc.v1.Request.value:type_name -> google.protobuf.Any
	6, // 2: pluginrpc.v1.Response.value:type_name -> google.protobuf.Any
	5, // 3: pluginrpc.v1.Response.error:type_name -> pluginrpc.v1.Error
	0, // 4: pluginrpc.v1.Error.code:type_name -> pluginrpc.v1.Code
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_pluginrpc_v1_pluginrpc_proto_init() }
func file_pluginrpc_v1_pluginrpc_proto_init() {
	if File_pluginrpc_v1_pluginrpc_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_pluginrpc_v1_pluginrpc_proto_rawDesc), len(file_pluginrpc_v1_pluginrpc_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pluginrpc_v1_pluginrpc_proto_goTypes,
		DependencyIndexes: file_pluginrpc_v1_pluginrpc_proto_depIdxs,
		EnumInfos:         file_pluginrpc_v1_pluginrpc_proto_enumTypes,
		MessageInfos:      file_pluginrpc_v1_pluginrpc_proto_msgTypes,
	}.Build()
	File_pluginrpc_v1_pluginrpc_proto = out.File
	file_pluginrpc_v1_pluginrpc_proto_goTypes = nil
	file_pluginrpc_v1_pluginrpc_proto_depIdxs = nil
}
