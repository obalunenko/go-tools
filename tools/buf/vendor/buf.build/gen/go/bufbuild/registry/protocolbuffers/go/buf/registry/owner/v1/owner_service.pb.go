// Copyright 2023-2025 Buf Technologies, Inc.
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
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: buf/registry/owner/v1/owner_service.proto

//go:build !protoopaque

package ownerv1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetOwnersRequest struct {
	state protoimpl.MessageState `protogen:"hybrid.v1"`
	// The Users or Organizations to request.
	OwnerRefs     []*OwnerRef `protobuf:"bytes,1,rep,name=owner_refs,json=ownerRefs,proto3" json:"owner_refs,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetOwnersRequest) Reset() {
	*x = GetOwnersRequest{}
	mi := &file_buf_registry_owner_v1_owner_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetOwnersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOwnersRequest) ProtoMessage() {}

func (x *GetOwnersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_buf_registry_owner_v1_owner_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *GetOwnersRequest) GetOwnerRefs() []*OwnerRef {
	if x != nil {
		return x.OwnerRefs
	}
	return nil
}

func (x *GetOwnersRequest) SetOwnerRefs(v []*OwnerRef) {
	x.OwnerRefs = v
}

type GetOwnersRequest_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// The Users or Organizations to request.
	OwnerRefs []*OwnerRef
}

func (b0 GetOwnersRequest_builder) Build() *GetOwnersRequest {
	m0 := &GetOwnersRequest{}
	b, x := &b0, m0
	_, _ = b, x
	x.OwnerRefs = b.OwnerRefs
	return m0
}

type GetOwnersResponse struct {
	state protoimpl.MessageState `protogen:"hybrid.v1"`
	// The retrieved Users or Organizations in the same order as requested.
	Owners        []*Owner `protobuf:"bytes,1,rep,name=owners,proto3" json:"owners,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetOwnersResponse) Reset() {
	*x = GetOwnersResponse{}
	mi := &file_buf_registry_owner_v1_owner_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetOwnersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOwnersResponse) ProtoMessage() {}

func (x *GetOwnersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_buf_registry_owner_v1_owner_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *GetOwnersResponse) GetOwners() []*Owner {
	if x != nil {
		return x.Owners
	}
	return nil
}

func (x *GetOwnersResponse) SetOwners(v []*Owner) {
	x.Owners = v
}

type GetOwnersResponse_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// The retrieved Users or Organizations in the same order as requested.
	Owners []*Owner
}

func (b0 GetOwnersResponse_builder) Build() *GetOwnersResponse {
	m0 := &GetOwnersResponse{}
	b, x := &b0, m0
	_, _ = b, x
	x.Owners = b.Owners
	return m0
}

var File_buf_registry_owner_v1_owner_service_proto protoreflect.FileDescriptor

var file_buf_registry_owner_v1_owner_service_proto_rawDesc = string([]byte{
	0x0a, 0x29, 0x62, 0x75, 0x66, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x62, 0x75, 0x66,
	0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x1a, 0x21, 0x62, 0x75, 0x66, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x2f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x5f, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4b, 0x0a, 0x0a, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f,
	0x72, 0x65, 0x66, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x62, 0x75, 0x66,
	0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x66, 0x42, 0x0b, 0xba, 0x48, 0x08,
	0x92, 0x01, 0x05, 0x08, 0x01, 0x10, 0xfa, 0x01, 0x52, 0x09, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x52,
	0x65, 0x66, 0x73, 0x22, 0x53, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x06, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x42, 0x08, 0xba, 0x48, 0x05, 0x92, 0x01, 0x02, 0x08, 0x01,
	0x52, 0x06, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x32, 0x73, 0x0a, 0x0c, 0x4f, 0x77, 0x6e, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x63, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x4f,
	0x77, 0x6e, 0x65, 0x72, 0x73, 0x12, 0x27, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x79, 0x2e, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28,
	0x2e, 0x62, 0x75, 0x66, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x6f, 0x77,
	0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x03, 0x90, 0x02, 0x01, 0x42, 0x55, 0x5a,
	0x53, 0x62, 0x75, 0x66, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67,
	0x6f, 0x2f, 0x62, 0x75, 0x66, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x75, 0x66, 0x66,
	0x65, 0x72, 0x73, 0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x75, 0x66, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x79, 0x2f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x6f, 0x77, 0x6e,
	0x65, 0x72, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var file_buf_registry_owner_v1_owner_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_buf_registry_owner_v1_owner_service_proto_goTypes = []any{
	(*GetOwnersRequest)(nil),  // 0: buf.registry.owner.v1.GetOwnersRequest
	(*GetOwnersResponse)(nil), // 1: buf.registry.owner.v1.GetOwnersResponse
	(*OwnerRef)(nil),          // 2: buf.registry.owner.v1.OwnerRef
	(*Owner)(nil),             // 3: buf.registry.owner.v1.Owner
}
var file_buf_registry_owner_v1_owner_service_proto_depIdxs = []int32{
	2, // 0: buf.registry.owner.v1.GetOwnersRequest.owner_refs:type_name -> buf.registry.owner.v1.OwnerRef
	3, // 1: buf.registry.owner.v1.GetOwnersResponse.owners:type_name -> buf.registry.owner.v1.Owner
	0, // 2: buf.registry.owner.v1.OwnerService.GetOwners:input_type -> buf.registry.owner.v1.GetOwnersRequest
	1, // 3: buf.registry.owner.v1.OwnerService.GetOwners:output_type -> buf.registry.owner.v1.GetOwnersResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_buf_registry_owner_v1_owner_service_proto_init() }
func file_buf_registry_owner_v1_owner_service_proto_init() {
	if File_buf_registry_owner_v1_owner_service_proto != nil {
		return
	}
	file_buf_registry_owner_v1_owner_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_buf_registry_owner_v1_owner_service_proto_rawDesc), len(file_buf_registry_owner_v1_owner_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_buf_registry_owner_v1_owner_service_proto_goTypes,
		DependencyIndexes: file_buf_registry_owner_v1_owner_service_proto_depIdxs,
		MessageInfos:      file_buf_registry_owner_v1_owner_service_proto_msgTypes,
	}.Build()
	File_buf_registry_owner_v1_owner_service_proto = out.File
	file_buf_registry_owner_v1_owner_service_proto_goTypes = nil
	file_buf_registry_owner_v1_owner_service_proto_depIdxs = nil
}
