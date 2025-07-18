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
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: buf/registry/module/v1/download_service.proto

//go:build !protoopaque

package modulev1

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

type DownloadRequest struct {
	state protoimpl.MessageState `protogen:"hybrid.v1"`
	// The references to get contents for.
	Values        []*DownloadRequest_Value `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DownloadRequest) Reset() {
	*x = DownloadRequest{}
	mi := &file_buf_registry_module_v1_download_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DownloadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadRequest) ProtoMessage() {}

func (x *DownloadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_buf_registry_module_v1_download_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *DownloadRequest) GetValues() []*DownloadRequest_Value {
	if x != nil {
		return x.Values
	}
	return nil
}

func (x *DownloadRequest) SetValues(v []*DownloadRequest_Value) {
	x.Values = v
}

type DownloadRequest_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// The references to get contents for.
	Values []*DownloadRequest_Value
}

func (b0 DownloadRequest_builder) Build() *DownloadRequest {
	m0 := &DownloadRequest{}
	b, x := &b0, m0
	_, _ = b, x
	x.Values = b.Values
	return m0
}

type DownloadResponse struct {
	state protoimpl.MessageState `protogen:"hybrid.v1"`
	// The Contents of the references in the same order as requested.
	Contents      []*DownloadResponse_Content `protobuf:"bytes,1,rep,name=contents,proto3" json:"contents,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DownloadResponse) Reset() {
	*x = DownloadResponse{}
	mi := &file_buf_registry_module_v1_download_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DownloadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadResponse) ProtoMessage() {}

func (x *DownloadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_buf_registry_module_v1_download_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *DownloadResponse) GetContents() []*DownloadResponse_Content {
	if x != nil {
		return x.Contents
	}
	return nil
}

func (x *DownloadResponse) SetContents(v []*DownloadResponse_Content) {
	x.Contents = v
}

type DownloadResponse_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// The Contents of the references in the same order as requested.
	Contents []*DownloadResponse_Content
}

func (b0 DownloadResponse_builder) Build() *DownloadResponse {
	m0 := &DownloadResponse{}
	b, x := &b0, m0
	_, _ = b, x
	x.Contents = b.Contents
	return m0
}

// A request for content for a single reference.
type DownloadRequest_Value struct {
	state protoimpl.MessageState `protogen:"hybrid.v1"`
	// The reference to get content for.
	//
	// See the documentation on ResourceRef for resource resolution details.
	//
	// Once the resource is resolved, the following content is returned:
	//   - If a Module is referenced, the content of the Commit of the default Label is returned.
	//   - If a Label is referenced, the content of the Commit of this Label is returned.
	//   - If a Commit is referenced, the content for this Commit is returned.
	ResourceRef *ResourceRef `protobuf:"bytes,1,opt,name=resource_ref,json=resourceRef,proto3" json:"resource_ref,omitempty"`
	// Specific file types to request.
	//
	// If not set, all file types are returned.
	FileTypes []FileType `protobuf:"varint,2,rep,packed,name=file_types,json=fileTypes,proto3,enum=buf.registry.module.v1.FileType" json:"file_types,omitempty"`
	// Specific file paths to retrieve.
	//
	// May be directories. For example, path "foo/bar" will result in files "foo/bar/baz.proto",
	// "foo/bar/LICENSE" being downloaded.
	//
	// If empty, all file paths for the given reference are retrieved.
	//
	// If no paths match, an empty Files list will be returned, however the call may still
	// be successful if paths_allow_not_exist is set (the dependency list may still be on
	// the response). If a directory "foo/bar" is specified but this directory has no files,
	// this is considered to be a non-match.
	//
	// This field also interacts with file_types - if file_types is set, a path only matches
	// if it is also of the file type, and if there are no matching paths for the given FileTypes,
	// an error is returned unless paths_not_allow_exist is set.
	//
	// The path must be relative, and cannot contain any "." or ".." components
	// The separator "/" must be used.
	Paths []string `protobuf:"bytes,3,rep,name=paths,proto3" json:"paths,omitempty"`
	// Whether to allow file paths not to exist within the given module.
	//
	// For example, one may want to retrieve the file paths "buf.md" and "README.md",
	// but only expect one to actually exist.
	//
	// If false, it is an error to specify non-existent file paths.
	PathsAllowNotExist bool `protobuf:"varint,4,opt,name=paths_allow_not_exist,json=pathsAllowNotExist,proto3" json:"paths_allow_not_exist,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *DownloadRequest_Value) Reset() {
	*x = DownloadRequest_Value{}
	mi := &file_buf_registry_module_v1_download_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DownloadRequest_Value) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadRequest_Value) ProtoMessage() {}

func (x *DownloadRequest_Value) ProtoReflect() protoreflect.Message {
	mi := &file_buf_registry_module_v1_download_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *DownloadRequest_Value) GetResourceRef() *ResourceRef {
	if x != nil {
		return x.ResourceRef
	}
	return nil
}

func (x *DownloadRequest_Value) GetFileTypes() []FileType {
	if x != nil {
		return x.FileTypes
	}
	return nil
}

func (x *DownloadRequest_Value) GetPaths() []string {
	if x != nil {
		return x.Paths
	}
	return nil
}

func (x *DownloadRequest_Value) GetPathsAllowNotExist() bool {
	if x != nil {
		return x.PathsAllowNotExist
	}
	return false
}

func (x *DownloadRequest_Value) SetResourceRef(v *ResourceRef) {
	x.ResourceRef = v
}

func (x *DownloadRequest_Value) SetFileTypes(v []FileType) {
	x.FileTypes = v
}

func (x *DownloadRequest_Value) SetPaths(v []string) {
	x.Paths = v
}

func (x *DownloadRequest_Value) SetPathsAllowNotExist(v bool) {
	x.PathsAllowNotExist = v
}

func (x *DownloadRequest_Value) HasResourceRef() bool {
	if x == nil {
		return false
	}
	return x.ResourceRef != nil
}

func (x *DownloadRequest_Value) ClearResourceRef() {
	x.ResourceRef = nil
}

type DownloadRequest_Value_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// The reference to get content for.
	//
	// See the documentation on ResourceRef for resource resolution details.
	//
	// Once the resource is resolved, the following content is returned:
	//   - If a Module is referenced, the content of the Commit of the default Label is returned.
	//   - If a Label is referenced, the content of the Commit of this Label is returned.
	//   - If a Commit is referenced, the content for this Commit is returned.
	ResourceRef *ResourceRef
	// Specific file types to request.
	//
	// If not set, all file types are returned.
	FileTypes []FileType
	// Specific file paths to retrieve.
	//
	// May be directories. For example, path "foo/bar" will result in files "foo/bar/baz.proto",
	// "foo/bar/LICENSE" being downloaded.
	//
	// If empty, all file paths for the given reference are retrieved.
	//
	// If no paths match, an empty Files list will be returned, however the call may still
	// be successful if paths_allow_not_exist is set (the dependency list may still be on
	// the response). If a directory "foo/bar" is specified but this directory has no files,
	// this is considered to be a non-match.
	//
	// This field also interacts with file_types - if file_types is set, a path only matches
	// if it is also of the file type, and if there are no matching paths for the given FileTypes,
	// an error is returned unless paths_not_allow_exist is set.
	//
	// The path must be relative, and cannot contain any "." or ".." components
	// The separator "/" must be used.
	Paths []string
	// Whether to allow file paths not to exist within the given module.
	//
	// For example, one may want to retrieve the file paths "buf.md" and "README.md",
	// but only expect one to actually exist.
	//
	// If false, it is an error to specify non-existent file paths.
	PathsAllowNotExist bool
}

func (b0 DownloadRequest_Value_builder) Build() *DownloadRequest_Value {
	m0 := &DownloadRequest_Value{}
	b, x := &b0, m0
	_, _ = b, x
	x.ResourceRef = b.ResourceRef
	x.FileTypes = b.FileTypes
	x.Paths = b.Paths
	x.PathsAllowNotExist = b.PathsAllowNotExist
	return m0
}

// Content for a single Commit.
type DownloadResponse_Content struct {
	state protoimpl.MessageState `protogen:"hybrid.v1"`
	// The Commit associated with the Content.
	//
	// The Commit associated with this ID will be present in the commits field.
	//
	// The Commit will use the DigestType specified in the request value.
	Commit *Commit `protobuf:"bytes,1,opt,name=commit,proto3" json:"commit,omitempty"`
	// The Files of the content.
	//
	// This will consist of the .proto files, license files, and documentation files.
	//
	// If no paths match and paths_allow_not_exist is set, this may be empty.
	Files         []*File `protobuf:"bytes,2,rep,name=files,proto3" json:"files,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DownloadResponse_Content) Reset() {
	*x = DownloadResponse_Content{}
	mi := &file_buf_registry_module_v1_download_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DownloadResponse_Content) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadResponse_Content) ProtoMessage() {}

func (x *DownloadResponse_Content) ProtoReflect() protoreflect.Message {
	mi := &file_buf_registry_module_v1_download_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *DownloadResponse_Content) GetCommit() *Commit {
	if x != nil {
		return x.Commit
	}
	return nil
}

func (x *DownloadResponse_Content) GetFiles() []*File {
	if x != nil {
		return x.Files
	}
	return nil
}

func (x *DownloadResponse_Content) SetCommit(v *Commit) {
	x.Commit = v
}

func (x *DownloadResponse_Content) SetFiles(v []*File) {
	x.Files = v
}

func (x *DownloadResponse_Content) HasCommit() bool {
	if x == nil {
		return false
	}
	return x.Commit != nil
}

func (x *DownloadResponse_Content) ClearCommit() {
	x.Commit = nil
}

type DownloadResponse_Content_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// The Commit associated with the Content.
	//
	// The Commit associated with this ID will be present in the commits field.
	//
	// The Commit will use the DigestType specified in the request value.
	Commit *Commit
	// The Files of the content.
	//
	// This will consist of the .proto files, license files, and documentation files.
	//
	// If no paths match and paths_allow_not_exist is set, this may be empty.
	Files []*File
}

func (b0 DownloadResponse_Content_builder) Build() *DownloadResponse_Content {
	m0 := &DownloadResponse_Content{}
	b, x := &b0, m0
	_, _ = b, x
	x.Commit = b.Commit
	x.Files = b.Files
	return m0
}

var File_buf_registry_module_v1_download_service_proto protoreflect.FileDescriptor

const file_buf_registry_module_v1_download_service_proto_rawDesc = "" +
	"\n" +
	"-buf/registry/module/v1/download_service.proto\x12\x16buf.registry.module.v1\x1a#buf/registry/module/v1/commit.proto\x1a!buf/registry/module/v1/file.proto\x1a%buf/registry/module/v1/resource.proto\x1a\x1bbuf/validate/validate.proto\"\xb5\x03\n" +
	"\x0fDownloadRequest\x12R\n" +
	"\x06values\x18\x01 \x03(\v2-.buf.registry.module.v1.DownloadRequest.ValueB\v\xbaH\b\x92\x01\x05\b\x01\x10\xfa\x01R\x06values\x1a\xcd\x02\n" +
	"\x05Value\x12N\n" +
	"\fresource_ref\x18\x01 \x01(\v2#.buf.registry.module.v1.ResourceRefB\x06\xbaH\x03\xc8\x01\x01R\vresourceRef\x12R\n" +
	"\n" +
	"file_types\x18\x02 \x03(\x0e2 .buf.registry.module.v1.FileTypeB\x11\xbaH\x0e\x92\x01\v\x18\x01\"\a\x82\x01\x04\x10\x01 \x00R\tfileTypes\x12m\n" +
	"\x05paths\x18\x03 \x03(\tBW\xbaHT\x92\x01Q\"OrM\x18\x80 2D^([^/.][^/]?|[^/][^/.]|[^/]{3,})(/([^/.][^/]?|[^/][^/.]|[^/]{3,}))*$\xba\x01\x01\\R\x05paths\x121\n" +
	"\x15paths_allow_not_exist\x18\x04 \x01(\bR\x12pathsAllowNotExist\"\xe9\x01\n" +
	"\x10DownloadResponse\x12V\n" +
	"\bcontents\x18\x01 \x03(\v20.buf.registry.module.v1.DownloadResponse.ContentB\b\xbaH\x05\x92\x01\x02\b\x01R\bcontents\x1a}\n" +
	"\aContent\x12>\n" +
	"\x06commit\x18\x01 \x01(\v2\x1e.buf.registry.module.v1.CommitB\x06\xbaH\x03\xc8\x01\x01R\x06commit\x122\n" +
	"\x05files\x18\x02 \x03(\v2\x1c.buf.registry.module.v1.FileR\x05files2u\n" +
	"\x0fDownloadService\x12b\n" +
	"\bDownload\x12'.buf.registry.module.v1.DownloadRequest\x1a(.buf.registry.module.v1.DownloadResponse\"\x03\x90\x02\x01BWZUbuf.build/gen/go/bufbuild/registry/protocolbuffers/go/buf/registry/module/v1;modulev1b\x06proto3"

var file_buf_registry_module_v1_download_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_buf_registry_module_v1_download_service_proto_goTypes = []any{
	(*DownloadRequest)(nil),          // 0: buf.registry.module.v1.DownloadRequest
	(*DownloadResponse)(nil),         // 1: buf.registry.module.v1.DownloadResponse
	(*DownloadRequest_Value)(nil),    // 2: buf.registry.module.v1.DownloadRequest.Value
	(*DownloadResponse_Content)(nil), // 3: buf.registry.module.v1.DownloadResponse.Content
	(*ResourceRef)(nil),              // 4: buf.registry.module.v1.ResourceRef
	(FileType)(0),                    // 5: buf.registry.module.v1.FileType
	(*Commit)(nil),                   // 6: buf.registry.module.v1.Commit
	(*File)(nil),                     // 7: buf.registry.module.v1.File
}
var file_buf_registry_module_v1_download_service_proto_depIdxs = []int32{
	2, // 0: buf.registry.module.v1.DownloadRequest.values:type_name -> buf.registry.module.v1.DownloadRequest.Value
	3, // 1: buf.registry.module.v1.DownloadResponse.contents:type_name -> buf.registry.module.v1.DownloadResponse.Content
	4, // 2: buf.registry.module.v1.DownloadRequest.Value.resource_ref:type_name -> buf.registry.module.v1.ResourceRef
	5, // 3: buf.registry.module.v1.DownloadRequest.Value.file_types:type_name -> buf.registry.module.v1.FileType
	6, // 4: buf.registry.module.v1.DownloadResponse.Content.commit:type_name -> buf.registry.module.v1.Commit
	7, // 5: buf.registry.module.v1.DownloadResponse.Content.files:type_name -> buf.registry.module.v1.File
	0, // 6: buf.registry.module.v1.DownloadService.Download:input_type -> buf.registry.module.v1.DownloadRequest
	1, // 7: buf.registry.module.v1.DownloadService.Download:output_type -> buf.registry.module.v1.DownloadResponse
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_buf_registry_module_v1_download_service_proto_init() }
func file_buf_registry_module_v1_download_service_proto_init() {
	if File_buf_registry_module_v1_download_service_proto != nil {
		return
	}
	file_buf_registry_module_v1_commit_proto_init()
	file_buf_registry_module_v1_file_proto_init()
	file_buf_registry_module_v1_resource_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_buf_registry_module_v1_download_service_proto_rawDesc), len(file_buf_registry_module_v1_download_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_buf_registry_module_v1_download_service_proto_goTypes,
		DependencyIndexes: file_buf_registry_module_v1_download_service_proto_depIdxs,
		MessageInfos:      file_buf_registry_module_v1_download_service_proto_msgTypes,
	}.Build()
	File_buf_registry_module_v1_download_service_proto = out.File
	file_buf_registry_module_v1_download_service_proto_goTypes = nil
	file_buf_registry_module_v1_download_service_proto_depIdxs = nil
}
