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
// source: buf/registry/plugin/v1beta1/label.proto

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

// A label on a specific Plugin.
//
// Many Labels can be associated with one Commit.
type Label struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The id of the Label.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The time the Label was created on the BSR.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// The last time the Label was updated on the BSR.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// The time the Label was archived if it is currently archived.
	//
	// If this field is not set, the Label is not currently archived.
	ArchiveTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=archive_time,json=archiveTime,proto3" json:"archive_time,omitempty"`
	// The name of the Label.
	//
	// Unique within a given Plugin.
	Name string `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	// The id of the Organization that owns the Plugin that the Label is associated with.
	OwnerId string `protobuf:"bytes,6,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	// The id of the Plugin that the Label is associated with.
	PluginId string `protobuf:"bytes,7,opt,name=plugin_id,json=pluginId,proto3" json:"plugin_id,omitempty"`
	// The id of the Commit currently associated with the Label.
	//
	// If policy checks are enabled, this will point to the most recent Commit that passed or was
	// approved. To get the history of the Commits that have been associated with a Label, use
	// ListLabelHistory.
	CommitId string `protobuf:"bytes,8,opt,name=commit_id,json=commitId,proto3" json:"commit_id,omitempty"`
	// The id of the User that last updated this Label on the BSR.
	//
	// May be empty if the User is no longer available.
	UpdatedByUserId string `protobuf:"bytes,9,opt,name=updated_by_user_id,json=updatedByUserId,proto3" json:"updated_by_user_id,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *Label) Reset() {
	*x = Label{}
	mi := &file_buf_registry_plugin_v1beta1_label_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Label) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Label) ProtoMessage() {}

func (x *Label) ProtoReflect() protoreflect.Message {
	mi := &file_buf_registry_plugin_v1beta1_label_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Label.ProtoReflect.Descriptor instead.
func (*Label) Descriptor() ([]byte, []int) {
	return file_buf_registry_plugin_v1beta1_label_proto_rawDescGZIP(), []int{0}
}

func (x *Label) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Label) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Label) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *Label) GetArchiveTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ArchiveTime
	}
	return nil
}

func (x *Label) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Label) GetOwnerId() string {
	if x != nil {
		return x.OwnerId
	}
	return ""
}

func (x *Label) GetPluginId() string {
	if x != nil {
		return x.PluginId
	}
	return ""
}

func (x *Label) GetCommitId() string {
	if x != nil {
		return x.CommitId
	}
	return ""
}

func (x *Label) GetUpdatedByUserId() string {
	if x != nil {
		return x.UpdatedByUserId
	}
	return ""
}

// LabelRef is a reference to a Label, either an id or a fully-qualified name.
//
// This is used in requests.
type LabelRef struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Value:
	//
	//	*LabelRef_Id
	//	*LabelRef_Name_
	Value         isLabelRef_Value `protobuf_oneof:"value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LabelRef) Reset() {
	*x = LabelRef{}
	mi := &file_buf_registry_plugin_v1beta1_label_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LabelRef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LabelRef) ProtoMessage() {}

func (x *LabelRef) ProtoReflect() protoreflect.Message {
	mi := &file_buf_registry_plugin_v1beta1_label_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LabelRef.ProtoReflect.Descriptor instead.
func (*LabelRef) Descriptor() ([]byte, []int) {
	return file_buf_registry_plugin_v1beta1_label_proto_rawDescGZIP(), []int{1}
}

func (x *LabelRef) GetValue() isLabelRef_Value {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *LabelRef) GetId() string {
	if x != nil {
		if x, ok := x.Value.(*LabelRef_Id); ok {
			return x.Id
		}
	}
	return ""
}

func (x *LabelRef) GetName() *LabelRef_Name {
	if x != nil {
		if x, ok := x.Value.(*LabelRef_Name_); ok {
			return x.Name
		}
	}
	return nil
}

type isLabelRef_Value interface {
	isLabelRef_Value()
}

type LabelRef_Id struct {
	// The id of the Label.
	Id string `protobuf:"bytes,1,opt,name=id,proto3,oneof"`
}

type LabelRef_Name_ struct {
	// The fully-qualified name of the Label.
	Name *LabelRef_Name `protobuf:"bytes,2,opt,name=name,proto3,oneof"`
}

func (*LabelRef_Id) isLabelRef_Value() {}

func (*LabelRef_Name_) isLabelRef_Value() {}

// A reference to a Label scoped to a Plugin, either an id or a name.
//
// This is used in requests.
type ScopedLabelRef struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Value:
	//
	//	*ScopedLabelRef_Id
	//	*ScopedLabelRef_Name
	Value         isScopedLabelRef_Value `protobuf_oneof:"value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ScopedLabelRef) Reset() {
	*x = ScopedLabelRef{}
	mi := &file_buf_registry_plugin_v1beta1_label_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ScopedLabelRef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScopedLabelRef) ProtoMessage() {}

func (x *ScopedLabelRef) ProtoReflect() protoreflect.Message {
	mi := &file_buf_registry_plugin_v1beta1_label_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScopedLabelRef.ProtoReflect.Descriptor instead.
func (*ScopedLabelRef) Descriptor() ([]byte, []int) {
	return file_buf_registry_plugin_v1beta1_label_proto_rawDescGZIP(), []int{2}
}

func (x *ScopedLabelRef) GetValue() isScopedLabelRef_Value {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *ScopedLabelRef) GetId() string {
	if x != nil {
		if x, ok := x.Value.(*ScopedLabelRef_Id); ok {
			return x.Id
		}
	}
	return ""
}

func (x *ScopedLabelRef) GetName() string {
	if x != nil {
		if x, ok := x.Value.(*ScopedLabelRef_Name); ok {
			return x.Name
		}
	}
	return ""
}

type isScopedLabelRef_Value interface {
	isScopedLabelRef_Value()
}

type ScopedLabelRef_Id struct {
	// The id of the Label.
	Id string `protobuf:"bytes,1,opt,name=id,proto3,oneof"`
}

type ScopedLabelRef_Name struct {
	// The name of the Label.
	Name string `protobuf:"bytes,2,opt,name=name,proto3,oneof"`
}

func (*ScopedLabelRef_Id) isScopedLabelRef_Value() {}

func (*ScopedLabelRef_Name) isScopedLabelRef_Value() {}

// The fully-qualified name of a Label within a BSR instance.
//
// A Name uniquely identifies a Label. This is used for requests when a caller only has the label
// name and not the ID.
type LabelRef_Name struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The name of the owner of the Plugin that contains the Label, an Organization.
	Owner string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	// The name of the Plugin that contains the Label.
	Plugin string `protobuf:"bytes,2,opt,name=plugin,proto3" json:"plugin,omitempty"`
	// The Label name.
	Label         string `protobuf:"bytes,3,opt,name=label,proto3" json:"label,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LabelRef_Name) Reset() {
	*x = LabelRef_Name{}
	mi := &file_buf_registry_plugin_v1beta1_label_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LabelRef_Name) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LabelRef_Name) ProtoMessage() {}

func (x *LabelRef_Name) ProtoReflect() protoreflect.Message {
	mi := &file_buf_registry_plugin_v1beta1_label_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LabelRef_Name.ProtoReflect.Descriptor instead.
func (*LabelRef_Name) Descriptor() ([]byte, []int) {
	return file_buf_registry_plugin_v1beta1_label_proto_rawDescGZIP(), []int{1, 0}
}

func (x *LabelRef_Name) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *LabelRef_Name) GetPlugin() string {
	if x != nil {
		return x.Plugin
	}
	return ""
}

func (x *LabelRef_Name) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

var File_buf_registry_plugin_v1beta1_label_proto protoreflect.FileDescriptor

var file_buf_registry_plugin_v1beta1_label_proto_rawDesc = []byte{
	0x0a, 0x27, 0x62, 0x75, 0x66, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x70,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x62, 0x75, 0x66, 0x2e, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x33, 0x62, 0x75, 0x66, 0x2f, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x79, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x62, 0x75, 0x66,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcc, 0x03, 0x0a, 0x05, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x12, 0x1b, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x0b, 0xba, 0x48, 0x08, 0xc8, 0x01, 0x01, 0x72, 0x03, 0x88, 0x02, 0x01, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x43, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x43, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3d, 0x0a, 0x0c, 0x61, 0x72,
	0x63, 0x68, 0x69, 0x76, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x61, 0x72,
	0x63, 0x68, 0x69, 0x76, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0xba, 0x48, 0x08, 0xc8, 0x01, 0x01, 0x72,
	0x03, 0x18, 0xfa, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x08, 0x6f, 0x77,
	0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0xba, 0x48,
	0x08, 0xc8, 0x01, 0x01, 0x72, 0x03, 0x88, 0x02, 0x01, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x28, 0x0a, 0x09, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0xba, 0x48, 0x08, 0xc8, 0x01, 0x01, 0x72, 0x03, 0x88,
	0x02, 0x01, 0x52, 0x08, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x09,
	0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x0b, 0xba, 0x48, 0x08, 0xc8, 0x01, 0x01, 0x72, 0x03, 0x88, 0x02, 0x01, 0x52, 0x08, 0x63, 0x6f,
	0x6d, 0x6d, 0x69, 0x74, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x12, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x62, 0x79, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x0b, 0xba, 0x48, 0x08, 0xd8, 0x01, 0x01, 0x72, 0x03, 0x88, 0x02, 0x01, 0x52,
	0x0f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x3a, 0x06, 0xea, 0xc5, 0x2b, 0x02, 0x10, 0x01, 0x22, 0xf0, 0x01, 0x0a, 0x08, 0x4c, 0x61, 0x62,
	0x65, 0x6c, 0x52, 0x65, 0x66, 0x12, 0x1a, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x08, 0xba, 0x48, 0x05, 0x72, 0x03, 0x88, 0x02, 0x01, 0x48, 0x00, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x40, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2a, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x70,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x52, 0x65, 0x66, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x48, 0x00, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x1a, 0x6e, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x05, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xba, 0x48, 0x07, 0xc8,
	0x01, 0x01, 0x72, 0x02, 0x18, 0x20, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x21, 0x0a,
	0x06, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xba,
	0x48, 0x06, 0x72, 0x04, 0x10, 0x02, 0x18, 0x64, 0x52, 0x06, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x12, 0x21, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x0b, 0xba, 0x48, 0x08, 0xc8, 0x01, 0x01, 0x72, 0x03, 0x18, 0xfa, 0x01, 0x52, 0x05, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x3a, 0x06, 0xea, 0xc5, 0x2b, 0x02, 0x08, 0x01, 0x42, 0x0e, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x05, 0xba, 0x48, 0x02, 0x08, 0x01, 0x22, 0x5a, 0x0a, 0x0e, 0x53,
	0x63, 0x6f, 0x70, 0x65, 0x64, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x65, 0x66, 0x12, 0x1a, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0x48, 0x05, 0x72, 0x03,
	0x88, 0x02, 0x01, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x3a,
	0x06, 0xea, 0xc5, 0x2b, 0x02, 0x08, 0x01, 0x42, 0x0e, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x05, 0xba, 0x48, 0x02, 0x08, 0x01, 0x42, 0x61, 0x5a, 0x5f, 0x62, 0x75, 0x66, 0x2e, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x75, 0x66, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x73, 0x2f, 0x67, 0x6f,
	0x2f, 0x62, 0x75, 0x66, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x70, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_buf_registry_plugin_v1beta1_label_proto_rawDescOnce sync.Once
	file_buf_registry_plugin_v1beta1_label_proto_rawDescData = file_buf_registry_plugin_v1beta1_label_proto_rawDesc
)

func file_buf_registry_plugin_v1beta1_label_proto_rawDescGZIP() []byte {
	file_buf_registry_plugin_v1beta1_label_proto_rawDescOnce.Do(func() {
		file_buf_registry_plugin_v1beta1_label_proto_rawDescData = protoimpl.X.CompressGZIP(file_buf_registry_plugin_v1beta1_label_proto_rawDescData)
	})
	return file_buf_registry_plugin_v1beta1_label_proto_rawDescData
}

var file_buf_registry_plugin_v1beta1_label_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_buf_registry_plugin_v1beta1_label_proto_goTypes = []any{
	(*Label)(nil),                 // 0: buf.registry.plugin.v1beta1.Label
	(*LabelRef)(nil),              // 1: buf.registry.plugin.v1beta1.LabelRef
	(*ScopedLabelRef)(nil),        // 2: buf.registry.plugin.v1beta1.ScopedLabelRef
	(*LabelRef_Name)(nil),         // 3: buf.registry.plugin.v1beta1.LabelRef.Name
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_buf_registry_plugin_v1beta1_label_proto_depIdxs = []int32{
	4, // 0: buf.registry.plugin.v1beta1.Label.create_time:type_name -> google.protobuf.Timestamp
	4, // 1: buf.registry.plugin.v1beta1.Label.update_time:type_name -> google.protobuf.Timestamp
	4, // 2: buf.registry.plugin.v1beta1.Label.archive_time:type_name -> google.protobuf.Timestamp
	3, // 3: buf.registry.plugin.v1beta1.LabelRef.name:type_name -> buf.registry.plugin.v1beta1.LabelRef.Name
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_buf_registry_plugin_v1beta1_label_proto_init() }
func file_buf_registry_plugin_v1beta1_label_proto_init() {
	if File_buf_registry_plugin_v1beta1_label_proto != nil {
		return
	}
	file_buf_registry_plugin_v1beta1_label_proto_msgTypes[1].OneofWrappers = []any{
		(*LabelRef_Id)(nil),
		(*LabelRef_Name_)(nil),
	}
	file_buf_registry_plugin_v1beta1_label_proto_msgTypes[2].OneofWrappers = []any{
		(*ScopedLabelRef_Id)(nil),
		(*ScopedLabelRef_Name)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_buf_registry_plugin_v1beta1_label_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_buf_registry_plugin_v1beta1_label_proto_goTypes,
		DependencyIndexes: file_buf_registry_plugin_v1beta1_label_proto_depIdxs,
		MessageInfos:      file_buf_registry_plugin_v1beta1_label_proto_msgTypes,
	}.Build()
	File_buf_registry_plugin_v1beta1_label_proto = out.File
	file_buf_registry_plugin_v1beta1_label_proto_rawDesc = nil
	file_buf_registry_plugin_v1beta1_label_proto_goTypes = nil
	file_buf_registry_plugin_v1beta1_label_proto_depIdxs = nil
}