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
// source: buf/registry/plugin/v1beta1/collection.proto

package pluginv1beta1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	_ "buf.build/gen/go/bufbuild/registry/protocolbuffers/go/buf/registry/priv/extension/v1beta1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// A collection for a Plugin.
//
// These collections help organize plugins based on their functionality or ecosystem.
type Collection struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The id of the Collection.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The time the Collection was created on the BSR.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// The last time the Collection was updated on the BSR.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// The name of the Collection.
	//
	// Unique within a BSR instance.
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// The configurable description of the Collection.
	Description   string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Collection) Reset() {
	*x = Collection{}
	mi := &file_buf_registry_plugin_v1beta1_collection_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Collection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Collection) ProtoMessage() {}

func (x *Collection) ProtoReflect() protoreflect.Message {
	mi := &file_buf_registry_plugin_v1beta1_collection_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Collection.ProtoReflect.Descriptor instead.
func (*Collection) Descriptor() ([]byte, []int) {
	return file_buf_registry_plugin_v1beta1_collection_proto_rawDescGZIP(), []int{0}
}

func (x *Collection) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Collection) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Collection) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *Collection) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Collection) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

// CollectionRef is a reference to a Collection, either an id or a name.
type CollectionRef struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Value:
	//
	//	*CollectionRef_Id
	//	*CollectionRef_Name
	Value         isCollectionRef_Value `protobuf_oneof:"value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CollectionRef) Reset() {
	*x = CollectionRef{}
	mi := &file_buf_registry_plugin_v1beta1_collection_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CollectionRef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CollectionRef) ProtoMessage() {}

func (x *CollectionRef) ProtoReflect() protoreflect.Message {
	mi := &file_buf_registry_plugin_v1beta1_collection_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CollectionRef.ProtoReflect.Descriptor instead.
func (*CollectionRef) Descriptor() ([]byte, []int) {
	return file_buf_registry_plugin_v1beta1_collection_proto_rawDescGZIP(), []int{1}
}

func (x *CollectionRef) GetValue() isCollectionRef_Value {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *CollectionRef) GetId() string {
	if x != nil {
		if x, ok := x.Value.(*CollectionRef_Id); ok {
			return x.Id
		}
	}
	return ""
}

func (x *CollectionRef) GetName() string {
	if x != nil {
		if x, ok := x.Value.(*CollectionRef_Name); ok {
			return x.Name
		}
	}
	return ""
}

type isCollectionRef_Value interface {
	isCollectionRef_Value()
}

type CollectionRef_Id struct {
	// The id of the Collection.
	Id string `protobuf:"bytes,1,opt,name=id,proto3,oneof"`
}

type CollectionRef_Name struct {
	// The name of the Collection.
	Name string `protobuf:"bytes,2,opt,name=name,proto3,oneof"`
}

func (*CollectionRef_Id) isCollectionRef_Value() {}

func (*CollectionRef_Name) isCollectionRef_Value() {}

var File_buf_registry_plugin_v1beta1_collection_proto protoreflect.FileDescriptor

var file_buf_registry_plugin_v1beta1_collection_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x62, 0x75, 0x66, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x70,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x63, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b,
	0x62, 0x75, 0x66, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x33, 0x62, 0x75, 0x66,
	0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x2f, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x88,
	0x02, 0x0a, 0x0a, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0xba, 0x48, 0x08, 0xc8, 0x01,
	0x01, 0x72, 0x03, 0x88, 0x02, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x43, 0x0a, 0x0b, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x06, 0xba, 0x48, 0x03,
	0xc8, 0x01, 0x01, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x43, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x0b, 0xba, 0x48, 0x08, 0xc8, 0x01, 0x01, 0x72, 0x03, 0x18, 0xfa, 0x01, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0x48, 0x05, 0x72,
	0x03, 0x18, 0xde, 0x02, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x3a, 0x06, 0xea, 0xc5, 0x2b, 0x02, 0x10, 0x01, 0x22, 0x64, 0x0a, 0x0d, 0x43, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x66, 0x12, 0x1a, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0x48, 0x05, 0x72, 0x03, 0x88, 0x02, 0x01,
	0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xba, 0x48, 0x06, 0x72, 0x04, 0x10, 0x02, 0x18, 0x64, 0x48,
	0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x3a, 0x06, 0xea, 0xc5, 0x2b, 0x02, 0x08, 0x01, 0x42,
	0x0e, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x05, 0xba, 0x48, 0x02, 0x08, 0x01, 0x42,
	0x61, 0x5a, 0x5f, 0x62, 0x75, 0x66, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x75, 0x66, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x72, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x75,
	0x66, 0x66, 0x65, 0x72, 0x73, 0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x75, 0x66, 0x2f, 0x72, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x3b, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_buf_registry_plugin_v1beta1_collection_proto_rawDescOnce sync.Once
	file_buf_registry_plugin_v1beta1_collection_proto_rawDescData = file_buf_registry_plugin_v1beta1_collection_proto_rawDesc
)

func file_buf_registry_plugin_v1beta1_collection_proto_rawDescGZIP() []byte {
	file_buf_registry_plugin_v1beta1_collection_proto_rawDescOnce.Do(func() {
		file_buf_registry_plugin_v1beta1_collection_proto_rawDescData = protoimpl.X.CompressGZIP(file_buf_registry_plugin_v1beta1_collection_proto_rawDescData)
	})
	return file_buf_registry_plugin_v1beta1_collection_proto_rawDescData
}

var file_buf_registry_plugin_v1beta1_collection_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_buf_registry_plugin_v1beta1_collection_proto_goTypes = []any{
	(*Collection)(nil),            // 0: buf.registry.plugin.v1beta1.Collection
	(*CollectionRef)(nil),         // 1: buf.registry.plugin.v1beta1.CollectionRef
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_buf_registry_plugin_v1beta1_collection_proto_depIdxs = []int32{
	2, // 0: buf.registry.plugin.v1beta1.Collection.create_time:type_name -> google.protobuf.Timestamp
	2, // 1: buf.registry.plugin.v1beta1.Collection.update_time:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_buf_registry_plugin_v1beta1_collection_proto_init() }
func file_buf_registry_plugin_v1beta1_collection_proto_init() {
	if File_buf_registry_plugin_v1beta1_collection_proto != nil {
		return
	}
	file_buf_registry_plugin_v1beta1_collection_proto_msgTypes[1].OneofWrappers = []any{
		(*CollectionRef_Id)(nil),
		(*CollectionRef_Name)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_buf_registry_plugin_v1beta1_collection_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_buf_registry_plugin_v1beta1_collection_proto_goTypes,
		DependencyIndexes: file_buf_registry_plugin_v1beta1_collection_proto_depIdxs,
		MessageInfos:      file_buf_registry_plugin_v1beta1_collection_proto_msgTypes,
	}.Build()
	File_buf_registry_plugin_v1beta1_collection_proto = out.File
	file_buf_registry_plugin_v1beta1_collection_proto_rawDesc = nil
	file_buf_registry_plugin_v1beta1_collection_proto_goTypes = nil
	file_buf_registry_plugin_v1beta1_collection_proto_depIdxs = nil
}