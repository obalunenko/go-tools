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
// source: buf/registry/module/v1/file.proto

package modulev1

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

// A specific file type.
type FileType int32

const (
	FileType_FILE_TYPE_UNSPECIFIED FileType = 0
	// A .proto file.
	FileType_FILE_TYPE_PROTO FileType = 1
	// A documentation file.
	//
	// Documentation files are always named README.md, README.markdown, or buf.md.
	FileType_FILE_TYPE_DOC FileType = 2
	// A license file.
	//
	// License files are always named LICENSE.
	FileType_FILE_TYPE_LICENSE FileType = 3
)

// Enum value maps for FileType.
var (
	FileType_name = map[int32]string{
		0: "FILE_TYPE_UNSPECIFIED",
		1: "FILE_TYPE_PROTO",
		2: "FILE_TYPE_DOC",
		3: "FILE_TYPE_LICENSE",
	}
	FileType_value = map[string]int32{
		"FILE_TYPE_UNSPECIFIED": 0,
		"FILE_TYPE_PROTO":       1,
		"FILE_TYPE_DOC":         2,
		"FILE_TYPE_LICENSE":     3,
	}
)

func (x FileType) Enum() *FileType {
	p := new(FileType)
	*p = x
	return p
}

func (x FileType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FileType) Descriptor() protoreflect.EnumDescriptor {
	return file_buf_registry_module_v1_file_proto_enumTypes[0].Descriptor()
}

func (FileType) Type() protoreflect.EnumType {
	return &file_buf_registry_module_v1_file_proto_enumTypes[0]
}

func (x FileType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FileType.Descriptor instead.
func (FileType) EnumDescriptor() ([]byte, []int) {
	return file_buf_registry_module_v1_file_proto_rawDescGZIP(), []int{0}
}

// A file that can be read or written to from disk.
//
// A File includes a path and associated content.
// Files are purposefully simple, and do not include attributes such as permissions.
type File struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The path of the File.
	//
	// The path must be relative, and cannot contain any "." or ".." components.
	// The separator "/" must be used.
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// The content of the File.
	//
	// May be empty.
	Content       []byte `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *File) Reset() {
	*x = File{}
	mi := &file_buf_registry_module_v1_file_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*File) ProtoMessage() {}

func (x *File) ProtoReflect() protoreflect.Message {
	mi := &file_buf_registry_module_v1_file_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use File.ProtoReflect.Descriptor instead.
func (*File) Descriptor() ([]byte, []int) {
	return file_buf_registry_module_v1_file_proto_rawDescGZIP(), []int{0}
}

func (x *File) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *File) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

var File_buf_registry_module_v1_file_proto protoreflect.FileDescriptor

var file_buf_registry_module_v1_file_proto_rawDesc = []byte{
	0x0a, 0x21, 0x62, 0x75, 0x66, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x6d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x16, 0x62, 0x75, 0x66, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x79, 0x2e, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x62, 0x75, 0x66,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8b, 0x01, 0x0a, 0x04, 0x46, 0x69, 0x6c,
	0x65, 0x12, 0x69, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x55, 0xba, 0x48, 0x52, 0xc8, 0x01, 0x01, 0x72, 0x4d, 0x18, 0x80, 0x20, 0x32, 0x44, 0x5e, 0x28,
	0x5b, 0x5e, 0x2f, 0x2e, 0x5d, 0x5b, 0x5e, 0x2f, 0x5d, 0x3f, 0x7c, 0x5b, 0x5e, 0x2f, 0x5d, 0x5b,
	0x5e, 0x2f, 0x2e, 0x5d, 0x7c, 0x5b, 0x5e, 0x2f, 0x5d, 0x7b, 0x33, 0x2c, 0x7d, 0x29, 0x28, 0x2f,
	0x28, 0x5b, 0x5e, 0x2f, 0x2e, 0x5d, 0x5b, 0x5e, 0x2f, 0x5d, 0x3f, 0x7c, 0x5b, 0x5e, 0x2f, 0x5d,
	0x5b, 0x5e, 0x2f, 0x2e, 0x5d, 0x7c, 0x5b, 0x5e, 0x2f, 0x5d, 0x7b, 0x33, 0x2c, 0x7d, 0x29, 0x29,
	0x2a, 0x24, 0xba, 0x01, 0x01, 0x5c, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2a, 0x64, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x19, 0x0a, 0x15, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x13, 0x0a,
	0x0f, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x52, 0x4f, 0x54, 0x4f,
	0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x44, 0x4f, 0x43, 0x10, 0x02, 0x12, 0x15, 0x0a, 0x11, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x4c, 0x49, 0x43, 0x45, 0x4e, 0x53, 0x45, 0x10, 0x03, 0x42, 0x57, 0x5a, 0x55,
	0x62, 0x75, 0x66, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f,
	0x2f, 0x62, 0x75, 0x66, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x75, 0x66, 0x66, 0x65,
	0x72, 0x73, 0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x75, 0x66, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x79, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x6d, 0x6f, 0x64,
	0x75, 0x6c, 0x65, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_buf_registry_module_v1_file_proto_rawDescOnce sync.Once
	file_buf_registry_module_v1_file_proto_rawDescData = file_buf_registry_module_v1_file_proto_rawDesc
)

func file_buf_registry_module_v1_file_proto_rawDescGZIP() []byte {
	file_buf_registry_module_v1_file_proto_rawDescOnce.Do(func() {
		file_buf_registry_module_v1_file_proto_rawDescData = protoimpl.X.CompressGZIP(file_buf_registry_module_v1_file_proto_rawDescData)
	})
	return file_buf_registry_module_v1_file_proto_rawDescData
}

var file_buf_registry_module_v1_file_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_buf_registry_module_v1_file_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_buf_registry_module_v1_file_proto_goTypes = []any{
	(FileType)(0), // 0: buf.registry.module.v1.FileType
	(*File)(nil),  // 1: buf.registry.module.v1.File
}
var file_buf_registry_module_v1_file_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_buf_registry_module_v1_file_proto_init() }
func file_buf_registry_module_v1_file_proto_init() {
	if File_buf_registry_module_v1_file_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_buf_registry_module_v1_file_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_buf_registry_module_v1_file_proto_goTypes,
		DependencyIndexes: file_buf_registry_module_v1_file_proto_depIdxs,
		EnumInfos:         file_buf_registry_module_v1_file_proto_enumTypes,
		MessageInfos:      file_buf_registry_module_v1_file_proto_msgTypes,
	}.Build()
	File_buf_registry_module_v1_file_proto = out.File
	file_buf_registry_module_v1_file_proto_rawDesc = nil
	file_buf_registry_module_v1_file_proto_goTypes = nil
	file_buf_registry_module_v1_file_proto_depIdxs = nil
}