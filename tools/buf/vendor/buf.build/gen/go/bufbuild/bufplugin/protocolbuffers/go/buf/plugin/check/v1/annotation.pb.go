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
// 	protoc-gen-go v1.36.0
// 	protoc        (unknown)
// source: buf/plugin/check/v1/annotation.proto

package checkv1

import (
	v1 "buf.build/gen/go/bufbuild/bufplugin/protocolbuffers/go/buf/plugin/descriptor/v1"
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

// An annotation representing a Rule failure.
//
// Annotations are propagated back to Buf and returned as lint or breaking change failures.
type Annotation struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The ID of the Rule that failed.
	//
	// Required.
	//
	// This must match an ID that the plugin has declared via its list rules RPC.
	RuleId string `protobuf:"bytes,1,opt,name=rule_id,json=ruleId,proto3" json:"rule_id,omitempty"`
	// A user-displayable message that explains the rule failure.
	//
	// Optional.
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// The location of the failure in a FileDescriptor within the file_descriptors list.
	//
	// Optional.
	FileLocation *v1.FileLocation `protobuf:"bytes,3,opt,name=file_location,json=fileLocation,proto3" json:"file_location,omitempty"`
	// The location of the failure in a FileDescriptor in the against_file_descriptors list.
	//
	// Optional.
	//
	// This may be present even if file_location is not present. For example, if a file was deleted,
	// this may reference the deleted FileDescriptor in against_file_descriptors, while file_location
	// will not be present.
	AgainstFileLocation *v1.FileLocation `protobuf:"bytes,4,opt,name=against_file_location,json=againstFileLocation,proto3" json:"against_file_location,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *Annotation) Reset() {
	*x = Annotation{}
	mi := &file_buf_plugin_check_v1_annotation_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Annotation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Annotation) ProtoMessage() {}

func (x *Annotation) ProtoReflect() protoreflect.Message {
	mi := &file_buf_plugin_check_v1_annotation_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Annotation.ProtoReflect.Descriptor instead.
func (*Annotation) Descriptor() ([]byte, []int) {
	return file_buf_plugin_check_v1_annotation_proto_rawDescGZIP(), []int{0}
}

func (x *Annotation) GetRuleId() string {
	if x != nil {
		return x.RuleId
	}
	return ""
}

func (x *Annotation) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Annotation) GetFileLocation() *v1.FileLocation {
	if x != nil {
		return x.FileLocation
	}
	return nil
}

func (x *Annotation) GetAgainstFileLocation() *v1.FileLocation {
	if x != nil {
		return x.AgainstFileLocation
	}
	return nil
}

var File_buf_plugin_check_v1_annotation_proto protoreflect.FileDescriptor

var file_buf_plugin_check_v1_annotation_proto_rawDesc = []byte{
	0x0a, 0x24, 0x62, 0x75, 0x66, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x62, 0x75, 0x66, 0x2e, 0x70, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x1a, 0x2c, 0x62, 0x75, 0x66,
	0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x6f, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x94, 0x02, 0x0a, 0x0a, 0x41, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x43, 0x0a, 0x07, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2a, 0xba, 0x48, 0x27, 0xc8, 0x01, 0x01, 0x72, 0x22,
	0x10, 0x03, 0x18, 0x40, 0x32, 0x1c, 0x5e, 0x5b, 0x41, 0x2d, 0x5a, 0x30, 0x2d, 0x39, 0x5d, 0x5b,
	0x41, 0x2d, 0x5a, 0x30, 0x2d, 0x39, 0x5f, 0x5d, 0x2a, 0x5b, 0x41, 0x2d, 0x5a, 0x30, 0x2d, 0x39,
	0x5d, 0x24, 0x52, 0x06, 0x72, 0x75, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x4b, 0x0a, 0x0d, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x62, 0x75,
	0x66, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x66, 0x69, 0x6c, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x5a, 0x0a, 0x15, 0x61, 0x67, 0x61, 0x69, 0x6e, 0x73, 0x74, 0x5f, 0x66, 0x69, 0x6c,
	0x65, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x26, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6c, 0x65,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x13, 0x61, 0x67, 0x61, 0x69, 0x6e, 0x73,
	0x74, 0x46, 0x69, 0x6c, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x54, 0x5a,
	0x52, 0x62, 0x75, 0x66, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67,
	0x6f, 0x2f, 0x62, 0x75, 0x66, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x62, 0x75, 0x66, 0x70, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x75, 0x66,
	0x66, 0x65, 0x72, 0x73, 0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x75, 0x66, 0x2f, 0x70, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x68, 0x65, 0x63,
	0x6b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_buf_plugin_check_v1_annotation_proto_rawDescOnce sync.Once
	file_buf_plugin_check_v1_annotation_proto_rawDescData = file_buf_plugin_check_v1_annotation_proto_rawDesc
)

func file_buf_plugin_check_v1_annotation_proto_rawDescGZIP() []byte {
	file_buf_plugin_check_v1_annotation_proto_rawDescOnce.Do(func() {
		file_buf_plugin_check_v1_annotation_proto_rawDescData = protoimpl.X.CompressGZIP(file_buf_plugin_check_v1_annotation_proto_rawDescData)
	})
	return file_buf_plugin_check_v1_annotation_proto_rawDescData
}

var file_buf_plugin_check_v1_annotation_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_buf_plugin_check_v1_annotation_proto_goTypes = []any{
	(*Annotation)(nil),      // 0: buf.plugin.check.v1.Annotation
	(*v1.FileLocation)(nil), // 1: buf.plugin.descriptor.v1.FileLocation
}
var file_buf_plugin_check_v1_annotation_proto_depIdxs = []int32{
	1, // 0: buf.plugin.check.v1.Annotation.file_location:type_name -> buf.plugin.descriptor.v1.FileLocation
	1, // 1: buf.plugin.check.v1.Annotation.against_file_location:type_name -> buf.plugin.descriptor.v1.FileLocation
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_buf_plugin_check_v1_annotation_proto_init() }
func file_buf_plugin_check_v1_annotation_proto_init() {
	if File_buf_plugin_check_v1_annotation_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_buf_plugin_check_v1_annotation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_buf_plugin_check_v1_annotation_proto_goTypes,
		DependencyIndexes: file_buf_plugin_check_v1_annotation_proto_depIdxs,
		MessageInfos:      file_buf_plugin_check_v1_annotation_proto_msgTypes,
	}.Build()
	File_buf_plugin_check_v1_annotation_proto = out.File
	file_buf_plugin_check_v1_annotation_proto_rawDesc = nil
	file_buf_plugin_check_v1_annotation_proto_goTypes = nil
	file_buf_plugin_check_v1_annotation_proto_depIdxs = nil
}