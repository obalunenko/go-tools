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
// source: buf/registry/owner/v1/organization.proto

package ownerv1

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

// The verification status of an Organization.
type OrganizationVerificationStatus int32

const (
	OrganizationVerificationStatus_ORGANIZATION_VERIFICATION_STATUS_UNSPECIFIED OrganizationVerificationStatus = 0
	// The Organization is unverified.
	OrganizationVerificationStatus_ORGANIZATION_VERIFICATION_STATUS_UNVERIFIED OrganizationVerificationStatus = 1
	// The Organization is verified.
	OrganizationVerificationStatus_ORGANIZATION_VERIFICATION_STATUS_VERIFIED OrganizationVerificationStatus = 2
	// The Organization is an official organization of the BSR owner.
	OrganizationVerificationStatus_ORGANIZATION_VERIFICATION_STATUS_OFFICIAL OrganizationVerificationStatus = 3
)

// Enum value maps for OrganizationVerificationStatus.
var (
	OrganizationVerificationStatus_name = map[int32]string{
		0: "ORGANIZATION_VERIFICATION_STATUS_UNSPECIFIED",
		1: "ORGANIZATION_VERIFICATION_STATUS_UNVERIFIED",
		2: "ORGANIZATION_VERIFICATION_STATUS_VERIFIED",
		3: "ORGANIZATION_VERIFICATION_STATUS_OFFICIAL",
	}
	OrganizationVerificationStatus_value = map[string]int32{
		"ORGANIZATION_VERIFICATION_STATUS_UNSPECIFIED": 0,
		"ORGANIZATION_VERIFICATION_STATUS_UNVERIFIED":  1,
		"ORGANIZATION_VERIFICATION_STATUS_VERIFIED":    2,
		"ORGANIZATION_VERIFICATION_STATUS_OFFICIAL":    3,
	}
)

func (x OrganizationVerificationStatus) Enum() *OrganizationVerificationStatus {
	p := new(OrganizationVerificationStatus)
	*p = x
	return p
}

func (x OrganizationVerificationStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrganizationVerificationStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_buf_registry_owner_v1_organization_proto_enumTypes[0].Descriptor()
}

func (OrganizationVerificationStatus) Type() protoreflect.EnumType {
	return &file_buf_registry_owner_v1_organization_proto_enumTypes[0]
}

func (x OrganizationVerificationStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrganizationVerificationStatus.Descriptor instead.
func (OrganizationVerificationStatus) EnumDescriptor() ([]byte, []int) {
	return file_buf_registry_owner_v1_organization_proto_rawDescGZIP(), []int{0}
}

// Organization is an organization on the BSR.
//
// A name uniquely identifies an Organization, however name is mutable.
type Organization struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The id for the Organization.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The time the Organization was created.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// The last time the Organization was updated.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// The name of the Organization.
	//
	// A name uniquely identifies an Organization, however name is mutable.
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// The configurable description of the Organization.
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// The configurable URL that represents the homepage for an Organization.
	Url string `protobuf:"bytes,6,opt,name=url,proto3" json:"url,omitempty"`
	// The verification status of the Organization.
	VerificationStatus OrganizationVerificationStatus `protobuf:"varint,7,opt,name=verification_status,json=verificationStatus,proto3,enum=buf.registry.owner.v1.OrganizationVerificationStatus" json:"verification_status,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *Organization) Reset() {
	*x = Organization{}
	mi := &file_buf_registry_owner_v1_organization_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Organization) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Organization) ProtoMessage() {}

func (x *Organization) ProtoReflect() protoreflect.Message {
	mi := &file_buf_registry_owner_v1_organization_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Organization.ProtoReflect.Descriptor instead.
func (*Organization) Descriptor() ([]byte, []int) {
	return file_buf_registry_owner_v1_organization_proto_rawDescGZIP(), []int{0}
}

func (x *Organization) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Organization) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Organization) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *Organization) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Organization) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Organization) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Organization) GetVerificationStatus() OrganizationVerificationStatus {
	if x != nil {
		return x.VerificationStatus
	}
	return OrganizationVerificationStatus_ORGANIZATION_VERIFICATION_STATUS_UNSPECIFIED
}

// OrganizationRef is a reference to an Organization, either an id or a name.
//
// This is used in requests.
type OrganizationRef struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Value:
	//
	//	*OrganizationRef_Id
	//	*OrganizationRef_Name
	Value         isOrganizationRef_Value `protobuf_oneof:"value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OrganizationRef) Reset() {
	*x = OrganizationRef{}
	mi := &file_buf_registry_owner_v1_organization_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OrganizationRef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrganizationRef) ProtoMessage() {}

func (x *OrganizationRef) ProtoReflect() protoreflect.Message {
	mi := &file_buf_registry_owner_v1_organization_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrganizationRef.ProtoReflect.Descriptor instead.
func (*OrganizationRef) Descriptor() ([]byte, []int) {
	return file_buf_registry_owner_v1_organization_proto_rawDescGZIP(), []int{1}
}

func (x *OrganizationRef) GetValue() isOrganizationRef_Value {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *OrganizationRef) GetId() string {
	if x != nil {
		if x, ok := x.Value.(*OrganizationRef_Id); ok {
			return x.Id
		}
	}
	return ""
}

func (x *OrganizationRef) GetName() string {
	if x != nil {
		if x, ok := x.Value.(*OrganizationRef_Name); ok {
			return x.Name
		}
	}
	return ""
}

type isOrganizationRef_Value interface {
	isOrganizationRef_Value()
}

type OrganizationRef_Id struct {
	// The id of the Organization.
	Id string `protobuf:"bytes,1,opt,name=id,proto3,oneof"`
}

type OrganizationRef_Name struct {
	// The name of the Organization.
	Name string `protobuf:"bytes,2,opt,name=name,proto3,oneof"`
}

func (*OrganizationRef_Id) isOrganizationRef_Value() {}

func (*OrganizationRef_Name) isOrganizationRef_Value() {}

var File_buf_registry_owner_v1_organization_proto protoreflect.FileDescriptor

var file_buf_registry_owner_v1_organization_proto_rawDesc = []byte{
	0x0a, 0x28, 0x62, 0x75, 0x66, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x62, 0x75, 0x66, 0x2e,
	0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x1a, 0x33, 0x62, 0x75, 0x66, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f,
	0x70, 0x72, 0x69, 0x76, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbd, 0x03, 0x0a, 0x0c, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x0b, 0xba, 0x48, 0x08, 0xc8, 0x01, 0x01, 0x72, 0x03, 0x88, 0x02, 0x01, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x43, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x43, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01,
	0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x27, 0xba, 0x48, 0x24, 0xc8,
	0x01, 0x01, 0x72, 0x1f, 0x10, 0x02, 0x18, 0x20, 0x32, 0x19, 0x5e, 0x5b, 0x61, 0x2d, 0x7a, 0x5d,
	0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x2d, 0x5d, 0x2a, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d,
	0x39, 0x5d, 0x24, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08,
	0xba, 0x48, 0x05, 0x72, 0x03, 0x18, 0xde, 0x02, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x0e, 0xba, 0x48, 0x0b, 0xd8, 0x01, 0x01, 0x72, 0x06, 0x18, 0xff, 0x01, 0x88,
	0x01, 0x01, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x73, 0x0a, 0x13, 0x76, 0x65, 0x72, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x35, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x79, 0x2e, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x67,
	0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x0b, 0xba, 0x48, 0x08,
	0xc8, 0x01, 0x01, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x12, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x3a, 0x06, 0xea, 0xc5,
	0x2b, 0x02, 0x10, 0x01, 0x22, 0x81, 0x01, 0x0a, 0x0f, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x66, 0x12, 0x1a, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0x48, 0x05, 0x72, 0x03, 0x88, 0x02, 0x01, 0x48, 0x00,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x3a, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x24, 0xba, 0x48, 0x21, 0x72, 0x1f, 0x10, 0x02, 0x18, 0x20, 0x32, 0x19, 0x5e,
	0x5b, 0x61, 0x2d, 0x7a, 0x5d, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x2d, 0x5d, 0x2a, 0x5b,
	0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x24, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x3a, 0x06, 0xea, 0xc5, 0x2b, 0x02, 0x08, 0x01, 0x42, 0x0e, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x05, 0xba, 0x48, 0x02, 0x08, 0x01, 0x2a, 0xe1, 0x01, 0x0a, 0x1e, 0x4f, 0x72, 0x67,
	0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x30, 0x0a, 0x2c, 0x4f,
	0x52, 0x47, 0x41, 0x4e, 0x49, 0x5a, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x56, 0x45, 0x52, 0x49,
	0x46, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x2f, 0x0a,
	0x2b, 0x4f, 0x52, 0x47, 0x41, 0x4e, 0x49, 0x5a, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x56, 0x45,
	0x52, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x5f, 0x55, 0x4e, 0x56, 0x45, 0x52, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x01, 0x12, 0x2d,
	0x0a, 0x29, 0x4f, 0x52, 0x47, 0x41, 0x4e, 0x49, 0x5a, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x56,
	0x45, 0x52, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x56, 0x45, 0x52, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x02, 0x12, 0x2d, 0x0a,
	0x29, 0x4f, 0x52, 0x47, 0x41, 0x4e, 0x49, 0x5a, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x56, 0x45,
	0x52, 0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x5f, 0x4f, 0x46, 0x46, 0x49, 0x43, 0x49, 0x41, 0x4c, 0x10, 0x03, 0x42, 0x55, 0x5a, 0x53,
	0x62, 0x75, 0x66, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f,
	0x2f, 0x62, 0x75, 0x66, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x75, 0x66, 0x66, 0x65,
	0x72, 0x73, 0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x75, 0x66, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x79, 0x2f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_buf_registry_owner_v1_organization_proto_rawDescOnce sync.Once
	file_buf_registry_owner_v1_organization_proto_rawDescData = file_buf_registry_owner_v1_organization_proto_rawDesc
)

func file_buf_registry_owner_v1_organization_proto_rawDescGZIP() []byte {
	file_buf_registry_owner_v1_organization_proto_rawDescOnce.Do(func() {
		file_buf_registry_owner_v1_organization_proto_rawDescData = protoimpl.X.CompressGZIP(file_buf_registry_owner_v1_organization_proto_rawDescData)
	})
	return file_buf_registry_owner_v1_organization_proto_rawDescData
}

var file_buf_registry_owner_v1_organization_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_buf_registry_owner_v1_organization_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_buf_registry_owner_v1_organization_proto_goTypes = []any{
	(OrganizationVerificationStatus)(0), // 0: buf.registry.owner.v1.OrganizationVerificationStatus
	(*Organization)(nil),                // 1: buf.registry.owner.v1.Organization
	(*OrganizationRef)(nil),             // 2: buf.registry.owner.v1.OrganizationRef
	(*timestamppb.Timestamp)(nil),       // 3: google.protobuf.Timestamp
}
var file_buf_registry_owner_v1_organization_proto_depIdxs = []int32{
	3, // 0: buf.registry.owner.v1.Organization.create_time:type_name -> google.protobuf.Timestamp
	3, // 1: buf.registry.owner.v1.Organization.update_time:type_name -> google.protobuf.Timestamp
	0, // 2: buf.registry.owner.v1.Organization.verification_status:type_name -> buf.registry.owner.v1.OrganizationVerificationStatus
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_buf_registry_owner_v1_organization_proto_init() }
func file_buf_registry_owner_v1_organization_proto_init() {
	if File_buf_registry_owner_v1_organization_proto != nil {
		return
	}
	file_buf_registry_owner_v1_organization_proto_msgTypes[1].OneofWrappers = []any{
		(*OrganizationRef_Id)(nil),
		(*OrganizationRef_Name)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_buf_registry_owner_v1_organization_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_buf_registry_owner_v1_organization_proto_goTypes,
		DependencyIndexes: file_buf_registry_owner_v1_organization_proto_depIdxs,
		EnumInfos:         file_buf_registry_owner_v1_organization_proto_enumTypes,
		MessageInfos:      file_buf_registry_owner_v1_organization_proto_msgTypes,
	}.Build()
	File_buf_registry_owner_v1_organization_proto = out.File
	file_buf_registry_owner_v1_organization_proto_rawDesc = nil
	file_buf_registry_owner_v1_organization_proto_goTypes = nil
	file_buf_registry_owner_v1_organization_proto_depIdxs = nil
}