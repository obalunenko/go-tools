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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        (unknown)
// source: buf/registry/plugin/v1beta1/download_service.proto

package pluginv1beta1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DownloadRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The references to get contents for.
	Values        []*DownloadRequest_Value `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DownloadRequest) Reset() {
	*x = DownloadRequest{}
	mi := &file_buf_registry_plugin_v1beta1_download_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DownloadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadRequest) ProtoMessage() {}

func (x *DownloadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_buf_registry_plugin_v1beta1_download_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadRequest.ProtoReflect.Descriptor instead.
func (*DownloadRequest) Descriptor() ([]byte, []int) {
	return file_buf_registry_plugin_v1beta1_download_service_proto_rawDescGZIP(), []int{0}
}

func (x *DownloadRequest) GetValues() []*DownloadRequest_Value {
	if x != nil {
		return x.Values
	}
	return nil
}

type DownloadResponse struct {
	state         protoimpl.MessageState      `protogen:"open.v1"`
	Contents      []*DownloadResponse_Content `protobuf:"bytes,1,rep,name=contents,proto3" json:"contents,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DownloadResponse) Reset() {
	*x = DownloadResponse{}
	mi := &file_buf_registry_plugin_v1beta1_download_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DownloadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadResponse) ProtoMessage() {}

func (x *DownloadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_buf_registry_plugin_v1beta1_download_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadResponse.ProtoReflect.Descriptor instead.
func (*DownloadResponse) Descriptor() ([]byte, []int) {
	return file_buf_registry_plugin_v1beta1_download_service_proto_rawDescGZIP(), []int{1}
}

func (x *DownloadResponse) GetContents() []*DownloadResponse_Content {
	if x != nil {
		return x.Contents
	}
	return nil
}

// A request for content for a single version of a Plugin.
type DownloadRequest_Value struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The reference to get content for.
	//
	// See the documentation on Reference for reference resolution details.
	//
	// Once the resource is resolved, the following content is returned:
	//   - If a Plugin is referenced, the content of the latest commit of the default label is
	//     returned.
	//   - If a Label is referenced, the content of the Commit of this Label is returned.
	//   - If a Commit is referenced, the content for this Commit is returned.
	ResourceRef   *ResourceRef `protobuf:"bytes,1,opt,name=resource_ref,json=resourceRef,proto3" json:"resource_ref,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DownloadRequest_Value) Reset() {
	*x = DownloadRequest_Value{}
	mi := &file_buf_registry_plugin_v1beta1_download_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DownloadRequest_Value) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadRequest_Value) ProtoMessage() {}

func (x *DownloadRequest_Value) ProtoReflect() protoreflect.Message {
	mi := &file_buf_registry_plugin_v1beta1_download_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadRequest_Value.ProtoReflect.Descriptor instead.
func (*DownloadRequest_Value) Descriptor() ([]byte, []int) {
	return file_buf_registry_plugin_v1beta1_download_service_proto_rawDescGZIP(), []int{0, 0}
}

func (x *DownloadRequest_Value) GetResourceRef() *ResourceRef {
	if x != nil {
		return x.ResourceRef
	}
	return nil
}

// Content for a single version of a Plugin.
type DownloadResponse_Content struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The Commit associated with the Content.
	Commit *Commit `protobuf:"bytes,1,opt,name=commit,proto3" json:"commit,omitempty"`
	// The compression type.
	CompressionType CompressionType `protobuf:"varint,2,opt,name=compression_type,json=compressionType,proto3,enum=buf.registry.plugin.v1beta1.CompressionType" json:"compression_type,omitempty"`
	// The content.
	Content       []byte `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DownloadResponse_Content) Reset() {
	*x = DownloadResponse_Content{}
	mi := &file_buf_registry_plugin_v1beta1_download_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DownloadResponse_Content) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadResponse_Content) ProtoMessage() {}

func (x *DownloadResponse_Content) ProtoReflect() protoreflect.Message {
	mi := &file_buf_registry_plugin_v1beta1_download_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadResponse_Content.ProtoReflect.Descriptor instead.
func (*DownloadResponse_Content) Descriptor() ([]byte, []int) {
	return file_buf_registry_plugin_v1beta1_download_service_proto_rawDescGZIP(), []int{1, 0}
}

func (x *DownloadResponse_Content) GetCommit() *Commit {
	if x != nil {
		return x.Commit
	}
	return nil
}

func (x *DownloadResponse_Content) GetCompressionType() CompressionType {
	if x != nil {
		return x.CompressionType
	}
	return CompressionType_COMPRESSION_TYPE_UNSPECIFIED
}

func (x *DownloadResponse_Content) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

var File_buf_registry_plugin_v1beta1_download_service_proto protoreflect.FileDescriptor

var file_buf_registry_plugin_v1beta1_download_service_proto_rawDesc = []byte{
	0x0a, 0x32, 0x62, 0x75, 0x66, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x70,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x64, 0x6f,
	0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x62, 0x75, 0x66, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x79, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x1a, 0x28, 0x62, 0x75, 0x66, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2d, 0x62, 0x75, 0x66,
	0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x62, 0x75, 0x66, 0x2f,
	0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xc8, 0x01, 0x0a, 0x0f, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x57, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x0b, 0xba, 0x48, 0x08,
	0x92, 0x01, 0x05, 0x08, 0x01, 0x10, 0xfa, 0x01, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x1a, 0x5c, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x53, 0x0a, 0x0c, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x72, 0x65, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x28, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x70,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x66, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01,
	0x01, 0x52, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x66, 0x22, 0xc8,
	0x02, 0x0a, 0x10, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x5b, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x42, 0x08, 0xba, 0x48,
	0x05, 0x92, 0x01, 0x02, 0x08, 0x01, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73,
	0x1a, 0xd6, 0x01, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x43, 0x0a, 0x06,
	0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x62,
	0x75, 0x66, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69,
	0x74, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x69,
	0x74, 0x12, 0x64, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x62, 0x75,
	0x66, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x42, 0x0b, 0xba, 0x48, 0x08, 0xc8, 0x01,
	0x01, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x0f, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x32, 0x7f, 0x0a, 0x0f, 0x44, 0x6f, 0x77,
	0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6c, 0x0a, 0x08,
	0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x2c, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x72, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x03, 0x90, 0x02, 0x01, 0x42, 0x61, 0x5a, 0x5f, 0x62, 0x75,
	0x66, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x62,
	0x75, 0x66, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x73,
	0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x75, 0x66, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_buf_registry_plugin_v1beta1_download_service_proto_rawDescOnce sync.Once
	file_buf_registry_plugin_v1beta1_download_service_proto_rawDescData = file_buf_registry_plugin_v1beta1_download_service_proto_rawDesc
)

func file_buf_registry_plugin_v1beta1_download_service_proto_rawDescGZIP() []byte {
	file_buf_registry_plugin_v1beta1_download_service_proto_rawDescOnce.Do(func() {
		file_buf_registry_plugin_v1beta1_download_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_buf_registry_plugin_v1beta1_download_service_proto_rawDescData)
	})
	return file_buf_registry_plugin_v1beta1_download_service_proto_rawDescData
}

var file_buf_registry_plugin_v1beta1_download_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_buf_registry_plugin_v1beta1_download_service_proto_goTypes = []any{
	(*DownloadRequest)(nil),          // 0: buf.registry.plugin.v1beta1.DownloadRequest
	(*DownloadResponse)(nil),         // 1: buf.registry.plugin.v1beta1.DownloadResponse
	(*DownloadRequest_Value)(nil),    // 2: buf.registry.plugin.v1beta1.DownloadRequest.Value
	(*DownloadResponse_Content)(nil), // 3: buf.registry.plugin.v1beta1.DownloadResponse.Content
	(*ResourceRef)(nil),              // 4: buf.registry.plugin.v1beta1.ResourceRef
	(*Commit)(nil),                   // 5: buf.registry.plugin.v1beta1.Commit
	(CompressionType)(0),             // 6: buf.registry.plugin.v1beta1.CompressionType
}
var file_buf_registry_plugin_v1beta1_download_service_proto_depIdxs = []int32{
	2, // 0: buf.registry.plugin.v1beta1.DownloadRequest.values:type_name -> buf.registry.plugin.v1beta1.DownloadRequest.Value
	3, // 1: buf.registry.plugin.v1beta1.DownloadResponse.contents:type_name -> buf.registry.plugin.v1beta1.DownloadResponse.Content
	4, // 2: buf.registry.plugin.v1beta1.DownloadRequest.Value.resource_ref:type_name -> buf.registry.plugin.v1beta1.ResourceRef
	5, // 3: buf.registry.plugin.v1beta1.DownloadResponse.Content.commit:type_name -> buf.registry.plugin.v1beta1.Commit
	6, // 4: buf.registry.plugin.v1beta1.DownloadResponse.Content.compression_type:type_name -> buf.registry.plugin.v1beta1.CompressionType
	0, // 5: buf.registry.plugin.v1beta1.DownloadService.Download:input_type -> buf.registry.plugin.v1beta1.DownloadRequest
	1, // 6: buf.registry.plugin.v1beta1.DownloadService.Download:output_type -> buf.registry.plugin.v1beta1.DownloadResponse
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_buf_registry_plugin_v1beta1_download_service_proto_init() }
func file_buf_registry_plugin_v1beta1_download_service_proto_init() {
	if File_buf_registry_plugin_v1beta1_download_service_proto != nil {
		return
	}
	file_buf_registry_plugin_v1beta1_commit_proto_init()
	file_buf_registry_plugin_v1beta1_compression_proto_init()
	file_buf_registry_plugin_v1beta1_resource_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_buf_registry_plugin_v1beta1_download_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_buf_registry_plugin_v1beta1_download_service_proto_goTypes,
		DependencyIndexes: file_buf_registry_plugin_v1beta1_download_service_proto_depIdxs,
		MessageInfos:      file_buf_registry_plugin_v1beta1_download_service_proto_msgTypes,
	}.Build()
	File_buf_registry_plugin_v1beta1_download_service_proto = out.File
	file_buf_registry_plugin_v1beta1_download_service_proto_rawDesc = nil
	file_buf_registry_plugin_v1beta1_download_service_proto_goTypes = nil
	file_buf_registry_plugin_v1beta1_download_service_proto_depIdxs = nil
}