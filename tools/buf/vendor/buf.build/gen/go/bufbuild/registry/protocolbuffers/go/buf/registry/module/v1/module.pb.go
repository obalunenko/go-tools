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
// source: buf/registry/module/v1/module.proto

//go:build !protoopaque

package modulev1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	_ "buf.build/gen/go/bufbuild/registry/protocolbuffers/go/buf/registry/priv/extension/v1beta1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The visibility of a Module, currently either public or private.
type ModuleVisibility int32

const (
	ModuleVisibility_MODULE_VISIBILITY_UNSPECIFIED ModuleVisibility = 0
	// MODULE_VISIBILITY_PUBLIC says that the module is publicly available.
	ModuleVisibility_MODULE_VISIBILITY_PUBLIC ModuleVisibility = 1
	// MODULE_VISIBILITY_PRIVATE says that the module is private.
	ModuleVisibility_MODULE_VISIBILITY_PRIVATE ModuleVisibility = 2
)

// Enum value maps for ModuleVisibility.
var (
	ModuleVisibility_name = map[int32]string{
		0: "MODULE_VISIBILITY_UNSPECIFIED",
		1: "MODULE_VISIBILITY_PUBLIC",
		2: "MODULE_VISIBILITY_PRIVATE",
	}
	ModuleVisibility_value = map[string]int32{
		"MODULE_VISIBILITY_UNSPECIFIED": 0,
		"MODULE_VISIBILITY_PUBLIC":      1,
		"MODULE_VISIBILITY_PRIVATE":     2,
	}
)

func (x ModuleVisibility) Enum() *ModuleVisibility {
	p := new(ModuleVisibility)
	*p = x
	return p
}

func (x ModuleVisibility) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ModuleVisibility) Descriptor() protoreflect.EnumDescriptor {
	return file_buf_registry_module_v1_module_proto_enumTypes[0].Descriptor()
}

func (ModuleVisibility) Type() protoreflect.EnumType {
	return &file_buf_registry_module_v1_module_proto_enumTypes[0]
}

func (x ModuleVisibility) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// The state of a Module, currently either active or deprecated.
type ModuleState int32

const (
	ModuleState_MODULE_STATE_UNSPECIFIED ModuleState = 0
	// MODULE_STATE_ACTIVE says that the Module is currently active.
	ModuleState_MODULE_STATE_ACTIVE ModuleState = 1
	// MODULE_STATE_DEPRECATED says that the Module has been deprecated and should not longer be
	// used.
	ModuleState_MODULE_STATE_DEPRECATED ModuleState = 2
)

// Enum value maps for ModuleState.
var (
	ModuleState_name = map[int32]string{
		0: "MODULE_STATE_UNSPECIFIED",
		1: "MODULE_STATE_ACTIVE",
		2: "MODULE_STATE_DEPRECATED",
	}
	ModuleState_value = map[string]int32{
		"MODULE_STATE_UNSPECIFIED": 0,
		"MODULE_STATE_ACTIVE":      1,
		"MODULE_STATE_DEPRECATED":  2,
	}
)

func (x ModuleState) Enum() *ModuleState {
	p := new(ModuleState)
	*p = x
	return p
}

func (x ModuleState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ModuleState) Descriptor() protoreflect.EnumDescriptor {
	return file_buf_registry_module_v1_module_proto_enumTypes[1].Descriptor()
}

func (ModuleState) Type() protoreflect.EnumType {
	return &file_buf_registry_module_v1_module_proto_enumTypes[1]
}

func (x ModuleState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// A module within the BSR.
type Module struct {
	state protoimpl.MessageState `protogen:"hybrid.v1"`
	// The id of the Module.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The time the Module was created on the BSR.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// The last time the Module was updated on the BSR.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// The name of the Module.
	//
	// Unique within a given User or Organization.
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// The id of the User or Organization that owns the Module.
	OwnerId string `protobuf:"bytes,5,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	// The Module's visibility, either public or private.
	Visibility ModuleVisibility `protobuf:"varint,6,opt,name=visibility,proto3,enum=buf.registry.module.v1.ModuleVisibility" json:"visibility,omitempty"`
	// The Module state, either active or deprecated.
	State ModuleState `protobuf:"varint,7,opt,name=state,proto3,enum=buf.registry.module.v1.ModuleState" json:"state,omitempty"`
	// The configurable description of the Module.
	Description string `protobuf:"bytes,8,opt,name=description,proto3" json:"description,omitempty"`
	// The configurable URL in the description of the Module,
	Url string `protobuf:"bytes,9,opt,name=url,proto3" json:"url,omitempty"`
	// The name of the default Label of the Module.
	//
	// This Label may not yet exist. When a Module is created, it has no Commits, and Labels
	// must have a Commit, so this Label is not created when a Module is created. Additionally,
	// a User may modify the name of the default Label without this Label yet being created.
	//
	// This could also be the name of an archived Label.
	DefaultLabelName string `protobuf:"bytes,10,opt,name=default_label_name,json=defaultLabelName,proto3" json:"default_label_name,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *Module) Reset() {
	*x = Module{}
	mi := &file_buf_registry_module_v1_module_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Module) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Module) ProtoMessage() {}

func (x *Module) ProtoReflect() protoreflect.Message {
	mi := &file_buf_registry_module_v1_module_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *Module) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Module) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Module) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *Module) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Module) GetOwnerId() string {
	if x != nil {
		return x.OwnerId
	}
	return ""
}

func (x *Module) GetVisibility() ModuleVisibility {
	if x != nil {
		return x.Visibility
	}
	return ModuleVisibility_MODULE_VISIBILITY_UNSPECIFIED
}

func (x *Module) GetState() ModuleState {
	if x != nil {
		return x.State
	}
	return ModuleState_MODULE_STATE_UNSPECIFIED
}

func (x *Module) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Module) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Module) GetDefaultLabelName() string {
	if x != nil {
		return x.DefaultLabelName
	}
	return ""
}

func (x *Module) SetId(v string) {
	x.Id = v
}

func (x *Module) SetCreateTime(v *timestamppb.Timestamp) {
	x.CreateTime = v
}

func (x *Module) SetUpdateTime(v *timestamppb.Timestamp) {
	x.UpdateTime = v
}

func (x *Module) SetName(v string) {
	x.Name = v
}

func (x *Module) SetOwnerId(v string) {
	x.OwnerId = v
}

func (x *Module) SetVisibility(v ModuleVisibility) {
	x.Visibility = v
}

func (x *Module) SetState(v ModuleState) {
	x.State = v
}

func (x *Module) SetDescription(v string) {
	x.Description = v
}

func (x *Module) SetUrl(v string) {
	x.Url = v
}

func (x *Module) SetDefaultLabelName(v string) {
	x.DefaultLabelName = v
}

func (x *Module) HasCreateTime() bool {
	if x == nil {
		return false
	}
	return x.CreateTime != nil
}

func (x *Module) HasUpdateTime() bool {
	if x == nil {
		return false
	}
	return x.UpdateTime != nil
}

func (x *Module) ClearCreateTime() {
	x.CreateTime = nil
}

func (x *Module) ClearUpdateTime() {
	x.UpdateTime = nil
}

type Module_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// The id of the Module.
	Id string
	// The time the Module was created on the BSR.
	CreateTime *timestamppb.Timestamp
	// The last time the Module was updated on the BSR.
	UpdateTime *timestamppb.Timestamp
	// The name of the Module.
	//
	// Unique within a given User or Organization.
	Name string
	// The id of the User or Organization that owns the Module.
	OwnerId string
	// The Module's visibility, either public or private.
	Visibility ModuleVisibility
	// The Module state, either active or deprecated.
	State ModuleState
	// The configurable description of the Module.
	Description string
	// The configurable URL in the description of the Module,
	Url string
	// The name of the default Label of the Module.
	//
	// This Label may not yet exist. When a Module is created, it has no Commits, and Labels
	// must have a Commit, so this Label is not created when a Module is created. Additionally,
	// a User may modify the name of the default Label without this Label yet being created.
	//
	// This could also be the name of an archived Label.
	DefaultLabelName string
}

func (b0 Module_builder) Build() *Module {
	m0 := &Module{}
	b, x := &b0, m0
	_, _ = b, x
	x.Id = b.Id
	x.CreateTime = b.CreateTime
	x.UpdateTime = b.UpdateTime
	x.Name = b.Name
	x.OwnerId = b.OwnerId
	x.Visibility = b.Visibility
	x.State = b.State
	x.Description = b.Description
	x.Url = b.Url
	x.DefaultLabelName = b.DefaultLabelName
	return m0
}

// ModuleRef is a reference to a Module, either an id or a fully-qualified name.
//
// This is used in requests.
type ModuleRef struct {
	state protoimpl.MessageState `protogen:"hybrid.v1"`
	// Types that are valid to be assigned to Value:
	//
	//	*ModuleRef_Id
	//	*ModuleRef_Name_
	Value         isModuleRef_Value `protobuf_oneof:"value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ModuleRef) Reset() {
	*x = ModuleRef{}
	mi := &file_buf_registry_module_v1_module_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ModuleRef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModuleRef) ProtoMessage() {}

func (x *ModuleRef) ProtoReflect() protoreflect.Message {
	mi := &file_buf_registry_module_v1_module_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *ModuleRef) GetValue() isModuleRef_Value {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *ModuleRef) GetId() string {
	if x != nil {
		if x, ok := x.Value.(*ModuleRef_Id); ok {
			return x.Id
		}
	}
	return ""
}

func (x *ModuleRef) GetName() *ModuleRef_Name {
	if x != nil {
		if x, ok := x.Value.(*ModuleRef_Name_); ok {
			return x.Name
		}
	}
	return nil
}

func (x *ModuleRef) SetId(v string) {
	x.Value = &ModuleRef_Id{v}
}

func (x *ModuleRef) SetName(v *ModuleRef_Name) {
	if v == nil {
		x.Value = nil
		return
	}
	x.Value = &ModuleRef_Name_{v}
}

func (x *ModuleRef) HasValue() bool {
	if x == nil {
		return false
	}
	return x.Value != nil
}

func (x *ModuleRef) HasId() bool {
	if x == nil {
		return false
	}
	_, ok := x.Value.(*ModuleRef_Id)
	return ok
}

func (x *ModuleRef) HasName() bool {
	if x == nil {
		return false
	}
	_, ok := x.Value.(*ModuleRef_Name_)
	return ok
}

func (x *ModuleRef) ClearValue() {
	x.Value = nil
}

func (x *ModuleRef) ClearId() {
	if _, ok := x.Value.(*ModuleRef_Id); ok {
		x.Value = nil
	}
}

func (x *ModuleRef) ClearName() {
	if _, ok := x.Value.(*ModuleRef_Name_); ok {
		x.Value = nil
	}
}

const ModuleRef_Value_not_set_case case_ModuleRef_Value = 0
const ModuleRef_Id_case case_ModuleRef_Value = 1
const ModuleRef_Name_case case_ModuleRef_Value = 2

func (x *ModuleRef) WhichValue() case_ModuleRef_Value {
	if x == nil {
		return ModuleRef_Value_not_set_case
	}
	switch x.Value.(type) {
	case *ModuleRef_Id:
		return ModuleRef_Id_case
	case *ModuleRef_Name_:
		return ModuleRef_Name_case
	default:
		return ModuleRef_Value_not_set_case
	}
}

type ModuleRef_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// Fields of oneof Value:
	// The id of the Module.
	Id *string
	// The fully-qualified name of the Module.
	Name *ModuleRef_Name
	// -- end of Value
}

func (b0 ModuleRef_builder) Build() *ModuleRef {
	m0 := &ModuleRef{}
	b, x := &b0, m0
	_, _ = b, x
	if b.Id != nil {
		x.Value = &ModuleRef_Id{*b.Id}
	}
	if b.Name != nil {
		x.Value = &ModuleRef_Name_{b.Name}
	}
	return m0
}

type case_ModuleRef_Value protoreflect.FieldNumber

func (x case_ModuleRef_Value) String() string {
	md := file_buf_registry_module_v1_module_proto_msgTypes[1].Descriptor()
	if x == 0 {
		return "not set"
	}
	return protoimpl.X.MessageFieldStringOf(md, protoreflect.FieldNumber(x))
}

type isModuleRef_Value interface {
	isModuleRef_Value()
}

type ModuleRef_Id struct {
	// The id of the Module.
	Id string `protobuf:"bytes,1,opt,name=id,proto3,oneof"`
}

type ModuleRef_Name_ struct {
	// The fully-qualified name of the Module.
	Name *ModuleRef_Name `protobuf:"bytes,2,opt,name=name,proto3,oneof"`
}

func (*ModuleRef_Id) isModuleRef_Value() {}

func (*ModuleRef_Name_) isModuleRef_Value() {}

// The fully-qualified name of a Module within a BSR instance.
//
// A Name uniquely identifies a Module.
// This is used for requests when a caller only has the module name and not the ID.
type ModuleRef_Name struct {
	state protoimpl.MessageState `protogen:"hybrid.v1"`
	// The name of the owner of the Module, either a User or Organization.
	Owner string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	// The name of the Module.
	Module        string `protobuf:"bytes,2,opt,name=module,proto3" json:"module,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ModuleRef_Name) Reset() {
	*x = ModuleRef_Name{}
	mi := &file_buf_registry_module_v1_module_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ModuleRef_Name) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModuleRef_Name) ProtoMessage() {}

func (x *ModuleRef_Name) ProtoReflect() protoreflect.Message {
	mi := &file_buf_registry_module_v1_module_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *ModuleRef_Name) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *ModuleRef_Name) GetModule() string {
	if x != nil {
		return x.Module
	}
	return ""
}

func (x *ModuleRef_Name) SetOwner(v string) {
	x.Owner = v
}

func (x *ModuleRef_Name) SetModule(v string) {
	x.Module = v
}

type ModuleRef_Name_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// The name of the owner of the Module, either a User or Organization.
	Owner string
	// The name of the Module.
	Module string
}

func (b0 ModuleRef_Name_builder) Build() *ModuleRef_Name {
	m0 := &ModuleRef_Name{}
	b, x := &b0, m0
	_, _ = b, x
	x.Owner = b.Owner
	x.Module = b.Module
	return m0
}

var File_buf_registry_module_v1_module_proto protoreflect.FileDescriptor

const file_buf_registry_module_v1_module_proto_rawDesc = "" +
	"\n" +
	"#buf/registry/module/v1/module.proto\x12\x16buf.registry.module.v1\x1a3buf/registry/priv/extension/v1beta1/extension.proto\x1a\x1bbuf/validate/validate.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"\xa6\x04\n" +
	"\x06Module\x12\x1b\n" +
	"\x02id\x18\x01 \x01(\tB\v\xbaH\b\xc8\x01\x01r\x03\x88\x02\x01R\x02id\x12C\n" +
	"\vcreate_time\x18\x02 \x01(\v2\x1a.google.protobuf.TimestampB\x06\xbaH\x03\xc8\x01\x01R\n" +
	"createTime\x12C\n" +
	"\vupdate_time\x18\x03 \x01(\v2\x1a.google.protobuf.TimestampB\x06\xbaH\x03\xc8\x01\x01R\n" +
	"updateTime\x12\x1d\n" +
	"\x04name\x18\x04 \x01(\tB\t\xbaH\x06r\x04\x10\x02\x18dR\x04name\x12&\n" +
	"\bowner_id\x18\x05 \x01(\tB\v\xbaH\b\xc8\x01\x01r\x03\x88\x02\x01R\aownerId\x12U\n" +
	"\n" +
	"visibility\x18\x06 \x01(\x0e2(.buf.registry.module.v1.ModuleVisibilityB\v\xbaH\b\xc8\x01\x01\x82\x01\x02\x10\x01R\n" +
	"visibility\x12F\n" +
	"\x05state\x18\a \x01(\x0e2#.buf.registry.module.v1.ModuleStateB\v\xbaH\b\xc8\x01\x01\x82\x01\x02\x10\x01R\x05state\x12*\n" +
	"\vdescription\x18\b \x01(\tB\b\xbaH\x05r\x03\x18\xde\x02R\vdescription\x12 \n" +
	"\x03url\x18\t \x01(\tB\x0e\xbaH\v\xd8\x01\x01r\x06\x18\xff\x01\x88\x01\x01R\x03url\x129\n" +
	"\x12default_label_name\x18\n" +
	" \x01(\tB\v\xbaH\b\xc8\x01\x01r\x03\x18\xfa\x01R\x10defaultLabelName:\x06\xea\xc5+\x02\x10\x01\"\xca\x01\n" +
	"\tModuleRef\x12\x1a\n" +
	"\x02id\x18\x01 \x01(\tB\b\xbaH\x05r\x03\x88\x02\x01H\x00R\x02id\x12<\n" +
	"\x04name\x18\x02 \x01(\v2&.buf.registry.module.v1.ModuleRef.NameH\x00R\x04name\x1aK\n" +
	"\x04Name\x12 \n" +
	"\x05owner\x18\x01 \x01(\tB\n" +
	"\xbaH\a\xc8\x01\x01r\x02\x18 R\x05owner\x12!\n" +
	"\x06module\x18\x02 \x01(\tB\t\xbaH\x06r\x04\x10\x02\x18dR\x06module:\x06\xea\xc5+\x02\b\x01B\x0e\n" +
	"\x05value\x12\x05\xbaH\x02\b\x01*r\n" +
	"\x10ModuleVisibility\x12!\n" +
	"\x1dMODULE_VISIBILITY_UNSPECIFIED\x10\x00\x12\x1c\n" +
	"\x18MODULE_VISIBILITY_PUBLIC\x10\x01\x12\x1d\n" +
	"\x19MODULE_VISIBILITY_PRIVATE\x10\x02*a\n" +
	"\vModuleState\x12\x1c\n" +
	"\x18MODULE_STATE_UNSPECIFIED\x10\x00\x12\x17\n" +
	"\x13MODULE_STATE_ACTIVE\x10\x01\x12\x1b\n" +
	"\x17MODULE_STATE_DEPRECATED\x10\x02BWZUbuf.build/gen/go/bufbuild/registry/protocolbuffers/go/buf/registry/module/v1;modulev1b\x06proto3"

var file_buf_registry_module_v1_module_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_buf_registry_module_v1_module_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_buf_registry_module_v1_module_proto_goTypes = []any{
	(ModuleVisibility)(0),         // 0: buf.registry.module.v1.ModuleVisibility
	(ModuleState)(0),              // 1: buf.registry.module.v1.ModuleState
	(*Module)(nil),                // 2: buf.registry.module.v1.Module
	(*ModuleRef)(nil),             // 3: buf.registry.module.v1.ModuleRef
	(*ModuleRef_Name)(nil),        // 4: buf.registry.module.v1.ModuleRef.Name
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_buf_registry_module_v1_module_proto_depIdxs = []int32{
	5, // 0: buf.registry.module.v1.Module.create_time:type_name -> google.protobuf.Timestamp
	5, // 1: buf.registry.module.v1.Module.update_time:type_name -> google.protobuf.Timestamp
	0, // 2: buf.registry.module.v1.Module.visibility:type_name -> buf.registry.module.v1.ModuleVisibility
	1, // 3: buf.registry.module.v1.Module.state:type_name -> buf.registry.module.v1.ModuleState
	4, // 4: buf.registry.module.v1.ModuleRef.name:type_name -> buf.registry.module.v1.ModuleRef.Name
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_buf_registry_module_v1_module_proto_init() }
func file_buf_registry_module_v1_module_proto_init() {
	if File_buf_registry_module_v1_module_proto != nil {
		return
	}
	file_buf_registry_module_v1_module_proto_msgTypes[1].OneofWrappers = []any{
		(*ModuleRef_Id)(nil),
		(*ModuleRef_Name_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_buf_registry_module_v1_module_proto_rawDesc), len(file_buf_registry_module_v1_module_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_buf_registry_module_v1_module_proto_goTypes,
		DependencyIndexes: file_buf_registry_module_v1_module_proto_depIdxs,
		EnumInfos:         file_buf_registry_module_v1_module_proto_enumTypes,
		MessageInfos:      file_buf_registry_module_v1_module_proto_msgTypes,
	}.Build()
	File_buf_registry_module_v1_module_proto = out.File
	file_buf_registry_module_v1_module_proto_goTypes = nil
	file_buf_registry_module_v1_module_proto_depIdxs = nil
}
