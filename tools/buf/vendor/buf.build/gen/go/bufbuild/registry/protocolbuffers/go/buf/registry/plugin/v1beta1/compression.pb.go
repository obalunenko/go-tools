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
// source: buf/registry/plugin/v1beta1/compression.proto

package pluginv1beta1

import (
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

// The type of compression.
type CompressionType int32

const (
	CompressionType_COMPRESSION_TYPE_UNSPECIFIED CompressionType = 0
	// No compression.
	CompressionType_COMPRESSION_TYPE_NONE CompressionType = 1
	// Zstandard compression.
	CompressionType_COMPRESSION_TYPE_ZSTD CompressionType = 2
)

// Enum value maps for CompressionType.
var (
	CompressionType_name = map[int32]string{
		0: "COMPRESSION_TYPE_UNSPECIFIED",
		1: "COMPRESSION_TYPE_NONE",
		2: "COMPRESSION_TYPE_ZSTD",
	}
	CompressionType_value = map[string]int32{
		"COMPRESSION_TYPE_UNSPECIFIED": 0,
		"COMPRESSION_TYPE_NONE":        1,
		"COMPRESSION_TYPE_ZSTD":        2,
	}
)

func (x CompressionType) Enum() *CompressionType {
	p := new(CompressionType)
	*p = x
	return p
}

func (x CompressionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CompressionType) Descriptor() protoreflect.EnumDescriptor {
	return file_buf_registry_plugin_v1beta1_compression_proto_enumTypes[0].Descriptor()
}

func (CompressionType) Type() protoreflect.EnumType {
	return &file_buf_registry_plugin_v1beta1_compression_proto_enumTypes[0]
}

func (x CompressionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CompressionType.Descriptor instead.
func (CompressionType) EnumDescriptor() ([]byte, []int) {
	return file_buf_registry_plugin_v1beta1_compression_proto_rawDescGZIP(), []int{0}
}

var File_buf_registry_plugin_v1beta1_compression_proto protoreflect.FileDescriptor

var file_buf_registry_plugin_v1beta1_compression_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x62, 0x75, 0x66, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x70,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x63, 0x6f,
	0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1b, 0x62, 0x75, 0x66, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2a, 0x69, 0x0a, 0x0f,
	0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x20, 0x0a, 0x1c, 0x43, 0x4f, 0x4d, 0x50, 0x52, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x19, 0x0a, 0x15, 0x43, 0x4f, 0x4d, 0x50, 0x52, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x01, 0x12, 0x19, 0x0a, 0x15,
	0x43, 0x4f, 0x4d, 0x50, 0x52, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x5a, 0x53, 0x54, 0x44, 0x10, 0x02, 0x42, 0x61, 0x5a, 0x5f, 0x62, 0x75, 0x66, 0x2e, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x75, 0x66, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x73, 0x2f, 0x67, 0x6f,
	0x2f, 0x62, 0x75, 0x66, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x70, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_buf_registry_plugin_v1beta1_compression_proto_rawDescOnce sync.Once
	file_buf_registry_plugin_v1beta1_compression_proto_rawDescData = file_buf_registry_plugin_v1beta1_compression_proto_rawDesc
)

func file_buf_registry_plugin_v1beta1_compression_proto_rawDescGZIP() []byte {
	file_buf_registry_plugin_v1beta1_compression_proto_rawDescOnce.Do(func() {
		file_buf_registry_plugin_v1beta1_compression_proto_rawDescData = protoimpl.X.CompressGZIP(file_buf_registry_plugin_v1beta1_compression_proto_rawDescData)
	})
	return file_buf_registry_plugin_v1beta1_compression_proto_rawDescData
}

var file_buf_registry_plugin_v1beta1_compression_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_buf_registry_plugin_v1beta1_compression_proto_goTypes = []any{
	(CompressionType)(0), // 0: buf.registry.plugin.v1beta1.CompressionType
}
var file_buf_registry_plugin_v1beta1_compression_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_buf_registry_plugin_v1beta1_compression_proto_init() }
func file_buf_registry_plugin_v1beta1_compression_proto_init() {
	if File_buf_registry_plugin_v1beta1_compression_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_buf_registry_plugin_v1beta1_compression_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_buf_registry_plugin_v1beta1_compression_proto_goTypes,
		DependencyIndexes: file_buf_registry_plugin_v1beta1_compression_proto_depIdxs,
		EnumInfos:         file_buf_registry_plugin_v1beta1_compression_proto_enumTypes,
	}.Build()
	File_buf_registry_plugin_v1beta1_compression_proto = out.File
	file_buf_registry_plugin_v1beta1_compression_proto_rawDesc = nil
	file_buf_registry_plugin_v1beta1_compression_proto_goTypes = nil
	file_buf_registry_plugin_v1beta1_compression_proto_depIdxs = nil
}