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
// source: buf/registry/plugin/v1beta1/upload_service.proto

package pluginv1beta1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	_ "buf.build/gen/go/bufbuild/registry/protocolbuffers/go/buf/registry/priv/extension/v1beta1"
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

type UploadRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The Contents of all references.
	Contents      []*UploadRequest_Content `protobuf:"bytes,1,rep,name=contents,proto3" json:"contents,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UploadRequest) Reset() {
	*x = UploadRequest{}
	mi := &file_buf_registry_plugin_v1beta1_upload_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UploadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadRequest) ProtoMessage() {}

func (x *UploadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_buf_registry_plugin_v1beta1_upload_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadRequest.ProtoReflect.Descriptor instead.
func (*UploadRequest) Descriptor() ([]byte, []int) {
	return file_buf_registry_plugin_v1beta1_upload_service_proto_rawDescGZIP(), []int{0}
}

func (x *UploadRequest) GetContents() []*UploadRequest_Content {
	if x != nil {
		return x.Contents
	}
	return nil
}

type UploadResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Commits       []*Commit              `protobuf:"bytes,1,rep,name=commits,proto3" json:"commits,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UploadResponse) Reset() {
	*x = UploadResponse{}
	mi := &file_buf_registry_plugin_v1beta1_upload_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UploadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadResponse) ProtoMessage() {}

func (x *UploadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_buf_registry_plugin_v1beta1_upload_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadResponse.ProtoReflect.Descriptor instead.
func (*UploadResponse) Descriptor() ([]byte, []int) {
	return file_buf_registry_plugin_v1beta1_upload_service_proto_rawDescGZIP(), []int{1}
}

func (x *UploadResponse) GetCommits() []*Commit {
	if x != nil {
		return x.Commits
	}
	return nil
}

// Content to upload.
type UploadRequest_Content struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The Plugin of the reference.
	PluginRef *PluginRef `protobuf:"bytes,1,opt,name=plugin_ref,json=pluginRef,proto3" json:"plugin_ref,omitempty"`
	// Compression type of the content.
	CompressionType CompressionType `protobuf:"varint,2,opt,name=compression_type,json=compressionType,proto3,enum=buf.registry.plugin.v1beta1.CompressionType" json:"compression_type,omitempty"`
	// The content to upload.
	Content []byte `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	// The labels to associate with the Commit for the Content.
	//
	// If an id is set, this id must represent a Label that already exists and is
	// owned by the Plugin. The Label will point to the newly-created Commits for the References,
	// or will be updated to point to the pre-existing Commit for the Reference.
	//
	// If no labels are referenced, the default Label for the Plugin is used.
	//
	// If the Labels do not exist, they will be created. If the Labels were archived, they will be
	// unarchived.
	ScopedLabelRefs []*ScopedLabelRef `protobuf:"bytes,4,rep,name=scoped_label_refs,json=scopedLabelRefs,proto3" json:"scoped_label_refs,omitempty"`
	// The URL of the source control commit to associate with the Commit for this Content.
	//
	// BSR users can navigate to this link to find source control information that is relevant to
	// this Commit (e.g. commit description, PR discussion, authors, approvers, etc.).
	SourceControlUrl string `protobuf:"bytes,5,opt,name=source_control_url,json=sourceControlUrl,proto3" json:"source_control_url,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *UploadRequest_Content) Reset() {
	*x = UploadRequest_Content{}
	mi := &file_buf_registry_plugin_v1beta1_upload_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UploadRequest_Content) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadRequest_Content) ProtoMessage() {}

func (x *UploadRequest_Content) ProtoReflect() protoreflect.Message {
	mi := &file_buf_registry_plugin_v1beta1_upload_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadRequest_Content.ProtoReflect.Descriptor instead.
func (*UploadRequest_Content) Descriptor() ([]byte, []int) {
	return file_buf_registry_plugin_v1beta1_upload_service_proto_rawDescGZIP(), []int{0, 0}
}

func (x *UploadRequest_Content) GetPluginRef() *PluginRef {
	if x != nil {
		return x.PluginRef
	}
	return nil
}

func (x *UploadRequest_Content) GetCompressionType() CompressionType {
	if x != nil {
		return x.CompressionType
	}
	return CompressionType_COMPRESSION_TYPE_UNSPECIFIED
}

func (x *UploadRequest_Content) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *UploadRequest_Content) GetScopedLabelRefs() []*ScopedLabelRef {
	if x != nil {
		return x.ScopedLabelRefs
	}
	return nil
}

func (x *UploadRequest_Content) GetSourceControlUrl() string {
	if x != nil {
		return x.SourceControlUrl
	}
	return ""
}

var File_buf_registry_plugin_v1beta1_upload_service_proto protoreflect.FileDescriptor

var file_buf_registry_plugin_v1beta1_upload_service_proto_rawDesc = []byte{
	0x0a, 0x30, 0x62, 0x75, 0x66, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x70,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x75, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1b, 0x62, 0x75, 0x66, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a,
	0x28, 0x62, 0x75, 0x66, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x70, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2d, 0x62, 0x75, 0x66, 0x2f, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x62, 0x75, 0x66, 0x2f, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x28, 0x62, 0x75, 0x66, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x70,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x33, 0x62, 0x75, 0x66,
	0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x2f, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xed, 0x03,
	0x0a, 0x0d, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x5a, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x32, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x42, 0x0a, 0xba, 0x48, 0x07, 0x92, 0x01, 0x04, 0x08, 0x01, 0x10,
	0x0a, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0xf7, 0x02, 0x0a, 0x07,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x4d, 0x0a, 0x0a, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x5f, 0x72, 0x65, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x62, 0x75,
	0x66, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x52, 0x65, 0x66, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x09, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x52, 0x65, 0x66, 0x12, 0x64, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x2c, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43,
	0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x42, 0x0b,
	0xba, 0x48, 0x08, 0xc8, 0x01, 0x01, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x0f, 0x63, 0x6f, 0x6d,
	0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x06, 0xba,
	0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x57,
	0x0a, 0x11, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x64, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f, 0x72,
	0x65, 0x66, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x62, 0x75, 0x66, 0x2e,
	0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x64, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x52, 0x65, 0x66, 0x52, 0x0f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x64, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x52, 0x65, 0x66, 0x73, 0x12, 0x3c, 0x0a, 0x12, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x0e, 0xba, 0x48, 0x0b, 0xd8, 0x01, 0x01, 0x72, 0x06, 0x18, 0xff, 0x01,
	0x88, 0x01, 0x01, 0x52, 0x10, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x55, 0x72, 0x6c, 0x3a, 0x06, 0xea, 0xc5, 0x2b, 0x02, 0x08, 0x01, 0x22, 0x61, 0x0a,
	0x0e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x47, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x23, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x42, 0x08, 0xba, 0x48, 0x05, 0x92, 0x01, 0x02, 0x08, 0x01, 0x52,
	0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x73, 0x3a, 0x06, 0xea, 0xc5, 0x2b, 0x02, 0x10, 0x01,
	0x32, 0x74, 0x0a, 0x0d, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x63, 0x0a, 0x06, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x2a, 0x2e, 0x62, 0x75,
	0x66, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x61, 0x5a, 0x5f, 0x62, 0x75, 0x66, 0x2e, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x75, 0x66, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x73, 0x2f, 0x67, 0x6f, 0x2f,
	0x62, 0x75, 0x66, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x70, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_buf_registry_plugin_v1beta1_upload_service_proto_rawDescOnce sync.Once
	file_buf_registry_plugin_v1beta1_upload_service_proto_rawDescData = file_buf_registry_plugin_v1beta1_upload_service_proto_rawDesc
)

func file_buf_registry_plugin_v1beta1_upload_service_proto_rawDescGZIP() []byte {
	file_buf_registry_plugin_v1beta1_upload_service_proto_rawDescOnce.Do(func() {
		file_buf_registry_plugin_v1beta1_upload_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_buf_registry_plugin_v1beta1_upload_service_proto_rawDescData)
	})
	return file_buf_registry_plugin_v1beta1_upload_service_proto_rawDescData
}

var file_buf_registry_plugin_v1beta1_upload_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_buf_registry_plugin_v1beta1_upload_service_proto_goTypes = []any{
	(*UploadRequest)(nil),         // 0: buf.registry.plugin.v1beta1.UploadRequest
	(*UploadResponse)(nil),        // 1: buf.registry.plugin.v1beta1.UploadResponse
	(*UploadRequest_Content)(nil), // 2: buf.registry.plugin.v1beta1.UploadRequest.Content
	(*Commit)(nil),                // 3: buf.registry.plugin.v1beta1.Commit
	(*PluginRef)(nil),             // 4: buf.registry.plugin.v1beta1.PluginRef
	(CompressionType)(0),          // 5: buf.registry.plugin.v1beta1.CompressionType
	(*ScopedLabelRef)(nil),        // 6: buf.registry.plugin.v1beta1.ScopedLabelRef
}
var file_buf_registry_plugin_v1beta1_upload_service_proto_depIdxs = []int32{
	2, // 0: buf.registry.plugin.v1beta1.UploadRequest.contents:type_name -> buf.registry.plugin.v1beta1.UploadRequest.Content
	3, // 1: buf.registry.plugin.v1beta1.UploadResponse.commits:type_name -> buf.registry.plugin.v1beta1.Commit
	4, // 2: buf.registry.plugin.v1beta1.UploadRequest.Content.plugin_ref:type_name -> buf.registry.plugin.v1beta1.PluginRef
	5, // 3: buf.registry.plugin.v1beta1.UploadRequest.Content.compression_type:type_name -> buf.registry.plugin.v1beta1.CompressionType
	6, // 4: buf.registry.plugin.v1beta1.UploadRequest.Content.scoped_label_refs:type_name -> buf.registry.plugin.v1beta1.ScopedLabelRef
	0, // 5: buf.registry.plugin.v1beta1.UploadService.Upload:input_type -> buf.registry.plugin.v1beta1.UploadRequest
	1, // 6: buf.registry.plugin.v1beta1.UploadService.Upload:output_type -> buf.registry.plugin.v1beta1.UploadResponse
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_buf_registry_plugin_v1beta1_upload_service_proto_init() }
func file_buf_registry_plugin_v1beta1_upload_service_proto_init() {
	if File_buf_registry_plugin_v1beta1_upload_service_proto != nil {
		return
	}
	file_buf_registry_plugin_v1beta1_commit_proto_init()
	file_buf_registry_plugin_v1beta1_compression_proto_init()
	file_buf_registry_plugin_v1beta1_label_proto_init()
	file_buf_registry_plugin_v1beta1_plugin_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_buf_registry_plugin_v1beta1_upload_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_buf_registry_plugin_v1beta1_upload_service_proto_goTypes,
		DependencyIndexes: file_buf_registry_plugin_v1beta1_upload_service_proto_depIdxs,
		MessageInfos:      file_buf_registry_plugin_v1beta1_upload_service_proto_msgTypes,
	}.Build()
	File_buf_registry_plugin_v1beta1_upload_service_proto = out.File
	file_buf_registry_plugin_v1beta1_upload_service_proto_rawDesc = nil
	file_buf_registry_plugin_v1beta1_upload_service_proto_goTypes = nil
	file_buf_registry_plugin_v1beta1_upload_service_proto_depIdxs = nil
}