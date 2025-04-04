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
// source: buf/registry/module/v1beta1/upload_service.proto

//go:build protoopaque

package modulev1beta1

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

type UploadRequest struct {
	state               protoimpl.MessageState    `protogen:"opaque.v1"`
	xxx_hidden_Contents *[]*UploadRequest_Content `protobuf:"bytes,1,rep,name=contents,proto3"`
	xxx_hidden_DepRefs  *[]*UploadRequest_DepRef  `protobuf:"bytes,2,rep,name=dep_refs,json=depRefs,proto3"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *UploadRequest) Reset() {
	*x = UploadRequest{}
	mi := &file_buf_registry_module_v1beta1_upload_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UploadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadRequest) ProtoMessage() {}

func (x *UploadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_buf_registry_module_v1beta1_upload_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *UploadRequest) GetContents() []*UploadRequest_Content {
	if x != nil {
		if x.xxx_hidden_Contents != nil {
			return *x.xxx_hidden_Contents
		}
	}
	return nil
}

func (x *UploadRequest) GetDepRefs() []*UploadRequest_DepRef {
	if x != nil {
		if x.xxx_hidden_DepRefs != nil {
			return *x.xxx_hidden_DepRefs
		}
	}
	return nil
}

func (x *UploadRequest) SetContents(v []*UploadRequest_Content) {
	x.xxx_hidden_Contents = &v
}

func (x *UploadRequest) SetDepRefs(v []*UploadRequest_DepRef) {
	x.xxx_hidden_DepRefs = &v
}

type UploadRequest_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// The Contents of all references.
	Contents []*UploadRequest_Content
	// The dependencies of the references specified by Contents.
	//
	// This will include all transitive dependencies.
	//
	// Dependencies between Contents are implicit and do not need to be specified. The BSR will detect
	// dependencies between Contents via .proto imports.
	//
	// Commits should be unique by Module, that is no two dep_refs should have the same Module but
	// different Commit IDs.
	DepRefs []*UploadRequest_DepRef
}

func (b0 UploadRequest_builder) Build() *UploadRequest {
	m0 := &UploadRequest{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Contents = &b.Contents
	x.xxx_hidden_DepRefs = &b.DepRefs
	return m0
}

// See the package documentation for more details. You should likely use buf.registry.module.v1beta1
// and not this package.
type UploadResponse struct {
	state              protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Commits *[]*Commit             `protobuf:"bytes,1,rep,name=commits,proto3"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *UploadResponse) Reset() {
	*x = UploadResponse{}
	mi := &file_buf_registry_module_v1beta1_upload_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UploadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadResponse) ProtoMessage() {}

func (x *UploadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_buf_registry_module_v1beta1_upload_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *UploadResponse) GetCommits() []*Commit {
	if x != nil {
		if x.xxx_hidden_Commits != nil {
			return *x.xxx_hidden_Commits
		}
	}
	return nil
}

func (x *UploadResponse) SetCommits(v []*Commit) {
	x.xxx_hidden_Commits = &v
}

type UploadResponse_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// The Commits for each reference in the same order as given on the request.
	//
	// A single Commit will be returned for each reference. These Commits may or may not be new.
	// If nothing changed for a given reference, the existing Commit will be returned.
	Commits []*Commit
}

func (b0 UploadResponse_builder) Build() *UploadResponse {
	m0 := &UploadResponse{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Commits = &b.Commits
	return m0
}

// A dependency of one or more references specified by Contents.
//
// Dependencies between Contents are implicit and do not need to be specified. The BSR will detect
// dependencies between Contents via .proto imports.
type UploadRequest_DepRef struct {
	state               protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_CommitId string                 `protobuf:"bytes,1,opt,name=commit_id,json=commitId,proto3"`
	xxx_hidden_Registry string                 `protobuf:"bytes,2,opt,name=registry,proto3"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *UploadRequest_DepRef) Reset() {
	*x = UploadRequest_DepRef{}
	mi := &file_buf_registry_module_v1beta1_upload_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UploadRequest_DepRef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadRequest_DepRef) ProtoMessage() {}

func (x *UploadRequest_DepRef) ProtoReflect() protoreflect.Message {
	mi := &file_buf_registry_module_v1beta1_upload_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *UploadRequest_DepRef) GetCommitId() string {
	if x != nil {
		return x.xxx_hidden_CommitId
	}
	return ""
}

func (x *UploadRequest_DepRef) GetRegistry() string {
	if x != nil {
		return x.xxx_hidden_Registry
	}
	return ""
}

func (x *UploadRequest_DepRef) SetCommitId(v string) {
	x.xxx_hidden_CommitId = v
}

func (x *UploadRequest_DepRef) SetRegistry(v string) {
	x.xxx_hidden_Registry = v
}

type UploadRequest_DepRef_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// The commit_id of the dependency.
	CommitId string
	// The registry hostname of the dependency.
	Registry string
}

func (b0 UploadRequest_DepRef_builder) Build() *UploadRequest_DepRef {
	m0 := &UploadRequest_DepRef{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_CommitId = b.CommitId
	x.xxx_hidden_Registry = b.Registry
	return m0
}

// Content to upload for a given reference.
type UploadRequest_Content struct {
	state                       protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_ModuleRef        *ModuleRef             `protobuf:"bytes,1,opt,name=module_ref,json=moduleRef,proto3"`
	xxx_hidden_Files            *[]*File               `protobuf:"bytes,2,rep,name=files,proto3"`
	xxx_hidden_ScopedLabelRefs  *[]*ScopedLabelRef     `protobuf:"bytes,3,rep,name=scoped_label_refs,json=scopedLabelRefs,proto3"`
	xxx_hidden_SourceControlUrl string                 `protobuf:"bytes,4,opt,name=source_control_url,json=sourceControlUrl,proto3"`
	unknownFields               protoimpl.UnknownFields
	sizeCache                   protoimpl.SizeCache
}

func (x *UploadRequest_Content) Reset() {
	*x = UploadRequest_Content{}
	mi := &file_buf_registry_module_v1beta1_upload_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UploadRequest_Content) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadRequest_Content) ProtoMessage() {}

func (x *UploadRequest_Content) ProtoReflect() protoreflect.Message {
	mi := &file_buf_registry_module_v1beta1_upload_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *UploadRequest_Content) GetModuleRef() *ModuleRef {
	if x != nil {
		return x.xxx_hidden_ModuleRef
	}
	return nil
}

func (x *UploadRequest_Content) GetFiles() []*File {
	if x != nil {
		if x.xxx_hidden_Files != nil {
			return *x.xxx_hidden_Files
		}
	}
	return nil
}

func (x *UploadRequest_Content) GetScopedLabelRefs() []*ScopedLabelRef {
	if x != nil {
		if x.xxx_hidden_ScopedLabelRefs != nil {
			return *x.xxx_hidden_ScopedLabelRefs
		}
	}
	return nil
}

func (x *UploadRequest_Content) GetSourceControlUrl() string {
	if x != nil {
		return x.xxx_hidden_SourceControlUrl
	}
	return ""
}

func (x *UploadRequest_Content) SetModuleRef(v *ModuleRef) {
	x.xxx_hidden_ModuleRef = v
}

func (x *UploadRequest_Content) SetFiles(v []*File) {
	x.xxx_hidden_Files = &v
}

func (x *UploadRequest_Content) SetScopedLabelRefs(v []*ScopedLabelRef) {
	x.xxx_hidden_ScopedLabelRefs = &v
}

func (x *UploadRequest_Content) SetSourceControlUrl(v string) {
	x.xxx_hidden_SourceControlUrl = v
}

func (x *UploadRequest_Content) HasModuleRef() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_ModuleRef != nil
}

func (x *UploadRequest_Content) ClearModuleRef() {
	x.xxx_hidden_ModuleRef = nil
}

type UploadRequest_Content_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// The Module of the reference.
	ModuleRef *ModuleRef
	// The Files of the Content.
	//
	// This will consist of the .proto files, license files, and documentation files.
	Files []*File
	// The labels to associate with the Commit for the Content.
	//
	// If an id is set, this id must represent a Label that already exists and is
	// owned by the Module. The Label will point to the newly-created Commits for the References,
	// or will be updated to point to the pre-existing Commit for the Reference.
	//
	// If no labels are referenced, the default Label for the Module is used.
	//
	// If the Labels do not exist, they will be created.
	// If the Labels were archived, they will be unarchived.
	ScopedLabelRefs []*ScopedLabelRef
	// The URL of the source control commit to associate with the Commit for this Content.
	//
	// BSR users can navigate to this link to find source control information that is relevant to this Commit
	// (e.g. commit description, PR discussion, authors, approvers, etc.).
	SourceControlUrl string
}

func (b0 UploadRequest_Content_builder) Build() *UploadRequest_Content {
	m0 := &UploadRequest_Content{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_ModuleRef = b.ModuleRef
	x.xxx_hidden_Files = &b.Files
	x.xxx_hidden_ScopedLabelRefs = &b.ScopedLabelRefs
	x.xxx_hidden_SourceControlUrl = b.SourceControlUrl
	return m0
}

var File_buf_registry_module_v1beta1_upload_service_proto protoreflect.FileDescriptor

var file_buf_registry_module_v1beta1_upload_service_proto_rawDesc = string([]byte{
	0x0a, 0x30, 0x62, 0x75, 0x66, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x6d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x75, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1b, 0x62, 0x75, 0x66, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x2e, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a,
	0x28, 0x62, 0x75, 0x66, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x6d, 0x6f,
	0x64, 0x75, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x62, 0x75, 0x66, 0x2f, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x27, 0x62, 0x75, 0x66, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f,
	0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x62, 0x75, 0x66, 0x2f,
	0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xc4, 0x04, 0x0a, 0x0d, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x58, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x79, 0x2e, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x42, 0x08, 0xba, 0x48, 0x05, 0x92, 0x01,
	0x02, 0x08, 0x01, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x4c, 0x0a,
	0x08, 0x64, 0x65, 0x70, 0x5f, 0x72, 0x65, 0x66, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x31, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x6d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x44, 0x65, 0x70, 0x52,
	0x65, 0x66, 0x52, 0x07, 0x64, 0x65, 0x70, 0x52, 0x65, 0x66, 0x73, 0x1a, 0x56, 0x0a, 0x06, 0x44,
	0x65, 0x70, 0x52, 0x65, 0x66, 0x12, 0x28, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0xba, 0x48, 0x08, 0xc8, 0x01, 0x01,
	0x72, 0x03, 0x88, 0x02, 0x01, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x49, 0x64, 0x12,
	0x22, 0x0a, 0x08, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x08, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x79, 0x1a, 0xb2, 0x02, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12,
	0x4d, 0x0a, 0x0a, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x72, 0x65, 0x66, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x79, 0x2e, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x66, 0x42, 0x06, 0xba, 0x48, 0x03,
	0xc8, 0x01, 0x01, 0x52, 0x09, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x66, 0x12, 0x41,
	0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x62, 0x75, 0x66, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x6d, 0x6f, 0x64,
	0x75, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x46, 0x69, 0x6c, 0x65,
	0x42, 0x08, 0xba, 0x48, 0x05, 0x92, 0x01, 0x02, 0x08, 0x01, 0x52, 0x05, 0x66, 0x69, 0x6c, 0x65,
	0x73, 0x12, 0x57, 0x0a, 0x11, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x64, 0x5f, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x5f, 0x72, 0x65, 0x66, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x62,
	0x75, 0x66, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x6d, 0x6f, 0x64, 0x75,
	0x6c, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x53, 0x63, 0x6f, 0x70, 0x65,
	0x64, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x65, 0x66, 0x52, 0x0f, 0x73, 0x63, 0x6f, 0x70, 0x65,
	0x64, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x65, 0x66, 0x73, 0x12, 0x3c, 0x0a, 0x12, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5f, 0x75, 0x72, 0x6c,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0e, 0xba, 0x48, 0x0b, 0xd8, 0x01, 0x01, 0x72, 0x06,
	0x18, 0xff, 0x01, 0x88, 0x01, 0x01, 0x52, 0x10, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x55, 0x72, 0x6c, 0x22, 0x59, 0x0a, 0x0e, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x07, 0x63, 0x6f,
	0x6d, 0x6d, 0x69, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x62, 0x75,
	0x66, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x6d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74,
	0x42, 0x08, 0xba, 0x48, 0x05, 0x92, 0x01, 0x02, 0x08, 0x01, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d,
	0x69, 0x74, 0x73, 0x32, 0x74, 0x0a, 0x0d, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x63, 0x0a, 0x06, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x2a,
	0x2e, 0x62, 0x75, 0x66, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x6d, 0x6f,
	0x64, 0x75, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x62, 0x75, 0x66,
	0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x61, 0x5a, 0x5f, 0x62, 0x75, 0x66,
	0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x75,
	0x66, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x73, 0x2f,
	0x67, 0x6f, 0x2f, 0x62, 0x75, 0x66, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f,
	0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x6d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
})

var file_buf_registry_module_v1beta1_upload_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_buf_registry_module_v1beta1_upload_service_proto_goTypes = []any{
	(*UploadRequest)(nil),         // 0: buf.registry.module.v1beta1.UploadRequest
	(*UploadResponse)(nil),        // 1: buf.registry.module.v1beta1.UploadResponse
	(*UploadRequest_DepRef)(nil),  // 2: buf.registry.module.v1beta1.UploadRequest.DepRef
	(*UploadRequest_Content)(nil), // 3: buf.registry.module.v1beta1.UploadRequest.Content
	(*Commit)(nil),                // 4: buf.registry.module.v1beta1.Commit
	(*ModuleRef)(nil),             // 5: buf.registry.module.v1beta1.ModuleRef
	(*File)(nil),                  // 6: buf.registry.module.v1beta1.File
	(*ScopedLabelRef)(nil),        // 7: buf.registry.module.v1beta1.ScopedLabelRef
}
var file_buf_registry_module_v1beta1_upload_service_proto_depIdxs = []int32{
	3, // 0: buf.registry.module.v1beta1.UploadRequest.contents:type_name -> buf.registry.module.v1beta1.UploadRequest.Content
	2, // 1: buf.registry.module.v1beta1.UploadRequest.dep_refs:type_name -> buf.registry.module.v1beta1.UploadRequest.DepRef
	4, // 2: buf.registry.module.v1beta1.UploadResponse.commits:type_name -> buf.registry.module.v1beta1.Commit
	5, // 3: buf.registry.module.v1beta1.UploadRequest.Content.module_ref:type_name -> buf.registry.module.v1beta1.ModuleRef
	6, // 4: buf.registry.module.v1beta1.UploadRequest.Content.files:type_name -> buf.registry.module.v1beta1.File
	7, // 5: buf.registry.module.v1beta1.UploadRequest.Content.scoped_label_refs:type_name -> buf.registry.module.v1beta1.ScopedLabelRef
	0, // 6: buf.registry.module.v1beta1.UploadService.Upload:input_type -> buf.registry.module.v1beta1.UploadRequest
	1, // 7: buf.registry.module.v1beta1.UploadService.Upload:output_type -> buf.registry.module.v1beta1.UploadResponse
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_buf_registry_module_v1beta1_upload_service_proto_init() }
func file_buf_registry_module_v1beta1_upload_service_proto_init() {
	if File_buf_registry_module_v1beta1_upload_service_proto != nil {
		return
	}
	file_buf_registry_module_v1beta1_commit_proto_init()
	file_buf_registry_module_v1beta1_file_proto_init()
	file_buf_registry_module_v1beta1_label_proto_init()
	file_buf_registry_module_v1beta1_module_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_buf_registry_module_v1beta1_upload_service_proto_rawDesc), len(file_buf_registry_module_v1beta1_upload_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_buf_registry_module_v1beta1_upload_service_proto_goTypes,
		DependencyIndexes: file_buf_registry_module_v1beta1_upload_service_proto_depIdxs,
		MessageInfos:      file_buf_registry_module_v1beta1_upload_service_proto_msgTypes,
	}.Build()
	File_buf_registry_module_v1beta1_upload_service_proto = out.File
	file_buf_registry_module_v1beta1_upload_service_proto_goTypes = nil
	file_buf_registry_module_v1beta1_upload_service_proto_depIdxs = nil
}
