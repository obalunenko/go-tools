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
// source: buf/registry/plugin/v1beta1/collection_service.proto

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

// The list order.
type ListCollectionsRequest_Order int32

const (
	ListCollectionsRequest_ORDER_UNSPECIFIED ListCollectionsRequest_Order = 0
	// Order by create_time newest to oldest.
	ListCollectionsRequest_ORDER_CREATE_TIME_DESC ListCollectionsRequest_Order = 1
	// Order by create_time oldest to newest.
	ListCollectionsRequest_ORDER_CREATE_TIME_ASC ListCollectionsRequest_Order = 2
)

// Enum value maps for ListCollectionsRequest_Order.
var (
	ListCollectionsRequest_Order_name = map[int32]string{
		0: "ORDER_UNSPECIFIED",
		1: "ORDER_CREATE_TIME_DESC",
		2: "ORDER_CREATE_TIME_ASC",
	}
	ListCollectionsRequest_Order_value = map[string]int32{
		"ORDER_UNSPECIFIED":      0,
		"ORDER_CREATE_TIME_DESC": 1,
		"ORDER_CREATE_TIME_ASC":  2,
	}
)

func (x ListCollectionsRequest_Order) Enum() *ListCollectionsRequest_Order {
	p := new(ListCollectionsRequest_Order)
	*p = x
	return p
}

func (x ListCollectionsRequest_Order) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ListCollectionsRequest_Order) Descriptor() protoreflect.EnumDescriptor {
	return file_buf_registry_plugin_v1beta1_collection_service_proto_enumTypes[0].Descriptor()
}

func (ListCollectionsRequest_Order) Type() protoreflect.EnumType {
	return &file_buf_registry_plugin_v1beta1_collection_service_proto_enumTypes[0]
}

func (x ListCollectionsRequest_Order) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ListCollectionsRequest_Order.Descriptor instead.
func (ListCollectionsRequest_Order) EnumDescriptor() ([]byte, []int) {
	return file_buf_registry_plugin_v1beta1_collection_service_proto_rawDescGZIP(), []int{2, 0}
}

type GetCollectionsRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The Collections to request.
	CollectionRefs []*CollectionRef `protobuf:"bytes,1,rep,name=collection_refs,json=collectionRefs,proto3" json:"collection_refs,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *GetCollectionsRequest) Reset() {
	*x = GetCollectionsRequest{}
	mi := &file_buf_registry_plugin_v1beta1_collection_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCollectionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCollectionsRequest) ProtoMessage() {}

func (x *GetCollectionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_buf_registry_plugin_v1beta1_collection_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCollectionsRequest.ProtoReflect.Descriptor instead.
func (*GetCollectionsRequest) Descriptor() ([]byte, []int) {
	return file_buf_registry_plugin_v1beta1_collection_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetCollectionsRequest) GetCollectionRefs() []*CollectionRef {
	if x != nil {
		return x.CollectionRefs
	}
	return nil
}

type GetCollectionsResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The retrieved Collections in the same order as requested.
	Collections   []*Collection `protobuf:"bytes,1,rep,name=collections,proto3" json:"collections,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetCollectionsResponse) Reset() {
	*x = GetCollectionsResponse{}
	mi := &file_buf_registry_plugin_v1beta1_collection_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCollectionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCollectionsResponse) ProtoMessage() {}

func (x *GetCollectionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_buf_registry_plugin_v1beta1_collection_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCollectionsResponse.ProtoReflect.Descriptor instead.
func (*GetCollectionsResponse) Descriptor() ([]byte, []int) {
	return file_buf_registry_plugin_v1beta1_collection_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetCollectionsResponse) GetCollections() []*Collection {
	if x != nil {
		return x.Collections
	}
	return nil
}

type ListCollectionsRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The maximum number of items to return.
	//
	// The default value is 10.
	PageSize uint32 `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// The page to start from.
	//
	// If empty, the first page is returned.
	PageToken string `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// The order to return the Plugins.
	//
	// If not specified, defaults to ORDER_CREATE_TIME_DESC.
	Order         ListCollectionsRequest_Order `protobuf:"varint,3,opt,name=order,proto3,enum=buf.registry.plugin.v1beta1.ListCollectionsRequest_Order" json:"order,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListCollectionsRequest) Reset() {
	*x = ListCollectionsRequest{}
	mi := &file_buf_registry_plugin_v1beta1_collection_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListCollectionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCollectionsRequest) ProtoMessage() {}

func (x *ListCollectionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_buf_registry_plugin_v1beta1_collection_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCollectionsRequest.ProtoReflect.Descriptor instead.
func (*ListCollectionsRequest) Descriptor() ([]byte, []int) {
	return file_buf_registry_plugin_v1beta1_collection_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListCollectionsRequest) GetPageSize() uint32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListCollectionsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *ListCollectionsRequest) GetOrder() ListCollectionsRequest_Order {
	if x != nil {
		return x.Order
	}
	return ListCollectionsRequest_ORDER_UNSPECIFIED
}

type ListCollectionsResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The next page token.
	//
	// If empty, there are no more pages.
	NextPageToken string `protobuf:"bytes,1,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	// The listed Collections.
	Collections   []*Collection `protobuf:"bytes,2,rep,name=collections,proto3" json:"collections,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListCollectionsResponse) Reset() {
	*x = ListCollectionsResponse{}
	mi := &file_buf_registry_plugin_v1beta1_collection_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListCollectionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCollectionsResponse) ProtoMessage() {}

func (x *ListCollectionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_buf_registry_plugin_v1beta1_collection_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCollectionsResponse.ProtoReflect.Descriptor instead.
func (*ListCollectionsResponse) Descriptor() ([]byte, []int) {
	return file_buf_registry_plugin_v1beta1_collection_service_proto_rawDescGZIP(), []int{3}
}

func (x *ListCollectionsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

func (x *ListCollectionsResponse) GetCollections() []*Collection {
	if x != nil {
		return x.Collections
	}
	return nil
}

type GetPluginCollectionAssociationsRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The Plugins to request Collections for.
	PluginRefs    []*PluginRef `protobuf:"bytes,1,rep,name=plugin_refs,json=pluginRefs,proto3" json:"plugin_refs,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPluginCollectionAssociationsRequest) Reset() {
	*x = GetPluginCollectionAssociationsRequest{}
	mi := &file_buf_registry_plugin_v1beta1_collection_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPluginCollectionAssociationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPluginCollectionAssociationsRequest) ProtoMessage() {}

func (x *GetPluginCollectionAssociationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_buf_registry_plugin_v1beta1_collection_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPluginCollectionAssociationsRequest.ProtoReflect.Descriptor instead.
func (*GetPluginCollectionAssociationsRequest) Descriptor() ([]byte, []int) {
	return file_buf_registry_plugin_v1beta1_collection_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetPluginCollectionAssociationsRequest) GetPluginRefs() []*PluginRef {
	if x != nil {
		return x.PluginRefs
	}
	return nil
}

type GetPluginCollectionAssociationsResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The Associations for the requested Plugins in the same order as requested.
	Associations  []*GetPluginCollectionAssociationsResponse_Association `protobuf:"bytes,1,rep,name=associations,proto3" json:"associations,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPluginCollectionAssociationsResponse) Reset() {
	*x = GetPluginCollectionAssociationsResponse{}
	mi := &file_buf_registry_plugin_v1beta1_collection_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPluginCollectionAssociationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPluginCollectionAssociationsResponse) ProtoMessage() {}

func (x *GetPluginCollectionAssociationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_buf_registry_plugin_v1beta1_collection_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPluginCollectionAssociationsResponse.ProtoReflect.Descriptor instead.
func (*GetPluginCollectionAssociationsResponse) Descriptor() ([]byte, []int) {
	return file_buf_registry_plugin_v1beta1_collection_service_proto_rawDescGZIP(), []int{5}
}

func (x *GetPluginCollectionAssociationsResponse) GetAssociations() []*GetPluginCollectionAssociationsResponse_Association {
	if x != nil {
		return x.Associations
	}
	return nil
}

// The Associations for the requested Plugins.
type GetPluginCollectionAssociationsResponse_Association struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The id of the Plugin.
	PluginId string `protobuf:"bytes,1,opt,name=plugin_id,json=pluginId,proto3" json:"plugin_id,omitempty"`
	// The collection ids associated with the Plugin.
	CollectionIds []string `protobuf:"bytes,2,rep,name=collection_ids,json=collectionIds,proto3" json:"collection_ids,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPluginCollectionAssociationsResponse_Association) Reset() {
	*x = GetPluginCollectionAssociationsResponse_Association{}
	mi := &file_buf_registry_plugin_v1beta1_collection_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPluginCollectionAssociationsResponse_Association) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPluginCollectionAssociationsResponse_Association) ProtoMessage() {}

func (x *GetPluginCollectionAssociationsResponse_Association) ProtoReflect() protoreflect.Message {
	mi := &file_buf_registry_plugin_v1beta1_collection_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPluginCollectionAssociationsResponse_Association.ProtoReflect.Descriptor instead.
func (*GetPluginCollectionAssociationsResponse_Association) Descriptor() ([]byte, []int) {
	return file_buf_registry_plugin_v1beta1_collection_service_proto_rawDescGZIP(), []int{5, 0}
}

func (x *GetPluginCollectionAssociationsResponse_Association) GetPluginId() string {
	if x != nil {
		return x.PluginId
	}
	return ""
}

func (x *GetPluginCollectionAssociationsResponse_Association) GetCollectionIds() []string {
	if x != nil {
		return x.CollectionIds
	}
	return nil
}

var File_buf_registry_plugin_v1beta1_collection_service_proto protoreflect.FileDescriptor

var file_buf_registry_plugin_v1beta1_collection_service_proto_rawDesc = []byte{
	0x0a, 0x34, 0x62, 0x75, 0x66, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x70,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x63, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x62, 0x75, 0x66, 0x2e, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x1a, 0x2c, 0x62, 0x75, 0x66, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x79, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x28, 0x62, 0x75, 0x66, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x70,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x62, 0x75, 0x66,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x79, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x43,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x60, 0x0a, 0x0f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x72, 0x65, 0x66, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x62, 0x75, 0x66,
	0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x66, 0x42, 0x0b, 0xba, 0x48, 0x08, 0x92, 0x01, 0x05, 0x08, 0x01,
	0x10, 0xfa, 0x01, 0x52, 0x0e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x66, 0x73, 0x22, 0x6d, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x53, 0x0a,
	0x0b, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x27, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x79, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x08, 0xba, 0x48, 0x05,
	0x92, 0x01, 0x02, 0x08, 0x01, 0x52, 0x0b, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x22, 0x9a, 0x02, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a,
	0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x42, 0x08, 0xba, 0x48, 0x05, 0x2a, 0x03, 0x18, 0xfa, 0x01, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x12, 0x27, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0x48, 0x05, 0x72, 0x03, 0x18,
	0x80, 0x20, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x59, 0x0a,
	0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x39, 0x2e, 0x62,
	0x75, 0x66, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x08, 0xba, 0x48, 0x05, 0x82, 0x01, 0x02, 0x10,
	0x01, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x55, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x12, 0x15, 0x0a, 0x11, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1a, 0x0a, 0x16, 0x4f, 0x52, 0x44, 0x45,
	0x52, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x44, 0x45,
	0x53, 0x43, 0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x43, 0x52,
	0x45, 0x41, 0x54, 0x45, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x41, 0x53, 0x43, 0x10, 0x02, 0x22,
	0x96, 0x01, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x0f, 0x6e,
	0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0x48, 0x05, 0x72, 0x03, 0x18, 0x80, 0x20, 0x52, 0x0d,
	0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x49, 0x0a,
	0x0b, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x27, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x79, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x63, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x7e, 0x0a, 0x26, 0x47, 0x65, 0x74, 0x50,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x41,
	0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x54, 0x0a, 0x0b, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x5f, 0x72, 0x65, 0x66,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x66, 0x42,
	0x0b, 0xba, 0x48, 0x08, 0x92, 0x01, 0x05, 0x08, 0x01, 0x10, 0xfa, 0x01, 0x52, 0x0a, 0x70, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x66, 0x73, 0x22, 0xf2, 0x01, 0x0a, 0x27, 0x47, 0x65, 0x74,
	0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x41, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x74, 0x0a, 0x0c, 0x61, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x50, 0x2e, 0x62, 0x75, 0x66,
	0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x73, 0x73, 0x6f,
	0x63, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x41, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x61, 0x73,
	0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x51, 0x0a, 0x0b, 0x41, 0x73,
	0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d,
	0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x73, 0x32, 0xcb, 0x03,
	0x0a, 0x11, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x7e, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x32, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x62, 0x75, 0x66, 0x2e,
	0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x03,
	0x90, 0x02, 0x01, 0x12, 0x81, 0x01, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x33, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x62,
	0x75, 0x66, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x03, 0x90, 0x02, 0x01, 0x12, 0xb1, 0x01, 0x0a, 0x1f, 0x47, 0x65, 0x74, 0x50,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x41,
	0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x43, 0x2e, 0x62, 0x75,
	0x66, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x73, 0x73,
	0x6f, 0x63, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x44, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x41, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x03, 0x90, 0x02, 0x01, 0x42, 0x61, 0x5a, 0x5f, 0x62,
	0x75, 0x66, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f,
	0x62, 0x75, 0x66, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72,
	0x73, 0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x75, 0x66, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x79, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x3b, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_buf_registry_plugin_v1beta1_collection_service_proto_rawDescOnce sync.Once
	file_buf_registry_plugin_v1beta1_collection_service_proto_rawDescData = file_buf_registry_plugin_v1beta1_collection_service_proto_rawDesc
)

func file_buf_registry_plugin_v1beta1_collection_service_proto_rawDescGZIP() []byte {
	file_buf_registry_plugin_v1beta1_collection_service_proto_rawDescOnce.Do(func() {
		file_buf_registry_plugin_v1beta1_collection_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_buf_registry_plugin_v1beta1_collection_service_proto_rawDescData)
	})
	return file_buf_registry_plugin_v1beta1_collection_service_proto_rawDescData
}

var file_buf_registry_plugin_v1beta1_collection_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_buf_registry_plugin_v1beta1_collection_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_buf_registry_plugin_v1beta1_collection_service_proto_goTypes = []any{
	(ListCollectionsRequest_Order)(0),                           // 0: buf.registry.plugin.v1beta1.ListCollectionsRequest.Order
	(*GetCollectionsRequest)(nil),                               // 1: buf.registry.plugin.v1beta1.GetCollectionsRequest
	(*GetCollectionsResponse)(nil),                              // 2: buf.registry.plugin.v1beta1.GetCollectionsResponse
	(*ListCollectionsRequest)(nil),                              // 3: buf.registry.plugin.v1beta1.ListCollectionsRequest
	(*ListCollectionsResponse)(nil),                             // 4: buf.registry.plugin.v1beta1.ListCollectionsResponse
	(*GetPluginCollectionAssociationsRequest)(nil),              // 5: buf.registry.plugin.v1beta1.GetPluginCollectionAssociationsRequest
	(*GetPluginCollectionAssociationsResponse)(nil),             // 6: buf.registry.plugin.v1beta1.GetPluginCollectionAssociationsResponse
	(*GetPluginCollectionAssociationsResponse_Association)(nil), // 7: buf.registry.plugin.v1beta1.GetPluginCollectionAssociationsResponse.Association
	(*CollectionRef)(nil),                                       // 8: buf.registry.plugin.v1beta1.CollectionRef
	(*Collection)(nil),                                          // 9: buf.registry.plugin.v1beta1.Collection
	(*PluginRef)(nil),                                           // 10: buf.registry.plugin.v1beta1.PluginRef
}
var file_buf_registry_plugin_v1beta1_collection_service_proto_depIdxs = []int32{
	8,  // 0: buf.registry.plugin.v1beta1.GetCollectionsRequest.collection_refs:type_name -> buf.registry.plugin.v1beta1.CollectionRef
	9,  // 1: buf.registry.plugin.v1beta1.GetCollectionsResponse.collections:type_name -> buf.registry.plugin.v1beta1.Collection
	0,  // 2: buf.registry.plugin.v1beta1.ListCollectionsRequest.order:type_name -> buf.registry.plugin.v1beta1.ListCollectionsRequest.Order
	9,  // 3: buf.registry.plugin.v1beta1.ListCollectionsResponse.collections:type_name -> buf.registry.plugin.v1beta1.Collection
	10, // 4: buf.registry.plugin.v1beta1.GetPluginCollectionAssociationsRequest.plugin_refs:type_name -> buf.registry.plugin.v1beta1.PluginRef
	7,  // 5: buf.registry.plugin.v1beta1.GetPluginCollectionAssociationsResponse.associations:type_name -> buf.registry.plugin.v1beta1.GetPluginCollectionAssociationsResponse.Association
	1,  // 6: buf.registry.plugin.v1beta1.CollectionService.GetCollections:input_type -> buf.registry.plugin.v1beta1.GetCollectionsRequest
	3,  // 7: buf.registry.plugin.v1beta1.CollectionService.ListCollections:input_type -> buf.registry.plugin.v1beta1.ListCollectionsRequest
	5,  // 8: buf.registry.plugin.v1beta1.CollectionService.GetPluginCollectionAssociations:input_type -> buf.registry.plugin.v1beta1.GetPluginCollectionAssociationsRequest
	2,  // 9: buf.registry.plugin.v1beta1.CollectionService.GetCollections:output_type -> buf.registry.plugin.v1beta1.GetCollectionsResponse
	4,  // 10: buf.registry.plugin.v1beta1.CollectionService.ListCollections:output_type -> buf.registry.plugin.v1beta1.ListCollectionsResponse
	6,  // 11: buf.registry.plugin.v1beta1.CollectionService.GetPluginCollectionAssociations:output_type -> buf.registry.plugin.v1beta1.GetPluginCollectionAssociationsResponse
	9,  // [9:12] is the sub-list for method output_type
	6,  // [6:9] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_buf_registry_plugin_v1beta1_collection_service_proto_init() }
func file_buf_registry_plugin_v1beta1_collection_service_proto_init() {
	if File_buf_registry_plugin_v1beta1_collection_service_proto != nil {
		return
	}
	file_buf_registry_plugin_v1beta1_collection_proto_init()
	file_buf_registry_plugin_v1beta1_plugin_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_buf_registry_plugin_v1beta1_collection_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_buf_registry_plugin_v1beta1_collection_service_proto_goTypes,
		DependencyIndexes: file_buf_registry_plugin_v1beta1_collection_service_proto_depIdxs,
		EnumInfos:         file_buf_registry_plugin_v1beta1_collection_service_proto_enumTypes,
		MessageInfos:      file_buf_registry_plugin_v1beta1_collection_service_proto_msgTypes,
	}.Build()
	File_buf_registry_plugin_v1beta1_collection_service_proto = out.File
	file_buf_registry_plugin_v1beta1_collection_service_proto_rawDesc = nil
	file_buf_registry_plugin_v1beta1_collection_service_proto_goTypes = nil
	file_buf_registry_plugin_v1beta1_collection_service_proto_depIdxs = nil
}