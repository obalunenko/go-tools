// Copyright 2020-2025 Buf Technologies, Inc.
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
// source: buf/alpha/registry/v1alpha1/download.proto

package registryv1alpha1

import (
	v1alpha1 "github.com/bufbuild/buf/private/gen/proto/go/buf/alpha/module/v1alpha1"
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

// DownloadRequest specifies the module to download.
type DownloadRequest struct {
	state                 protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Owner      string                 `protobuf:"bytes,1,opt,name=owner,proto3"`
	xxx_hidden_Repository string                 `protobuf:"bytes,2,opt,name=repository,proto3"`
	xxx_hidden_Reference  string                 `protobuf:"bytes,3,opt,name=reference,proto3"`
	unknownFields         protoimpl.UnknownFields
	sizeCache             protoimpl.SizeCache
}

func (x *DownloadRequest) Reset() {
	*x = DownloadRequest{}
	mi := &file_buf_alpha_registry_v1alpha1_download_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DownloadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadRequest) ProtoMessage() {}

func (x *DownloadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_buf_alpha_registry_v1alpha1_download_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *DownloadRequest) GetOwner() string {
	if x != nil {
		return x.xxx_hidden_Owner
	}
	return ""
}

func (x *DownloadRequest) GetRepository() string {
	if x != nil {
		return x.xxx_hidden_Repository
	}
	return ""
}

func (x *DownloadRequest) GetReference() string {
	if x != nil {
		return x.xxx_hidden_Reference
	}
	return ""
}

func (x *DownloadRequest) SetOwner(v string) {
	x.xxx_hidden_Owner = v
}

func (x *DownloadRequest) SetRepository(v string) {
	x.xxx_hidden_Repository = v
}

func (x *DownloadRequest) SetReference(v string) {
	x.xxx_hidden_Reference = v
}

type DownloadRequest_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Owner      string
	Repository string
	// Optional reference (if unspecified, will use the repository's default_branch).
	Reference string
}

func (b0 DownloadRequest_builder) Build() *DownloadRequest {
	m0 := &DownloadRequest{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Owner = b.Owner
	x.xxx_hidden_Repository = b.Repository
	x.xxx_hidden_Reference = b.Reference
	return m0
}

// DownloadResponse contains the remote module.
type DownloadResponse struct {
	state             protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Module *v1alpha1.Module       `protobuf:"bytes,1,opt,name=module,proto3"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *DownloadResponse) Reset() {
	*x = DownloadResponse{}
	mi := &file_buf_alpha_registry_v1alpha1_download_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DownloadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadResponse) ProtoMessage() {}

func (x *DownloadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_buf_alpha_registry_v1alpha1_download_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *DownloadResponse) GetModule() *v1alpha1.Module {
	if x != nil {
		return x.xxx_hidden_Module
	}
	return nil
}

func (x *DownloadResponse) SetModule(v *v1alpha1.Module) {
	x.xxx_hidden_Module = v
}

func (x *DownloadResponse) HasModule() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Module != nil
}

func (x *DownloadResponse) ClearModule() {
	x.xxx_hidden_Module = nil
}

type DownloadResponse_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Module *v1alpha1.Module
}

func (b0 DownloadResponse_builder) Build() *DownloadResponse {
	m0 := &DownloadResponse{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Module = b.Module
	return m0
}

// DownloadManifestAndBlobsRequest specifies the module to download.
type DownloadManifestAndBlobsRequest struct {
	state                 protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Owner      string                 `protobuf:"bytes,1,opt,name=owner,proto3"`
	xxx_hidden_Repository string                 `protobuf:"bytes,2,opt,name=repository,proto3"`
	xxx_hidden_Reference  string                 `protobuf:"bytes,3,opt,name=reference,proto3"`
	unknownFields         protoimpl.UnknownFields
	sizeCache             protoimpl.SizeCache
}

func (x *DownloadManifestAndBlobsRequest) Reset() {
	*x = DownloadManifestAndBlobsRequest{}
	mi := &file_buf_alpha_registry_v1alpha1_download_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DownloadManifestAndBlobsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadManifestAndBlobsRequest) ProtoMessage() {}

func (x *DownloadManifestAndBlobsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_buf_alpha_registry_v1alpha1_download_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *DownloadManifestAndBlobsRequest) GetOwner() string {
	if x != nil {
		return x.xxx_hidden_Owner
	}
	return ""
}

func (x *DownloadManifestAndBlobsRequest) GetRepository() string {
	if x != nil {
		return x.xxx_hidden_Repository
	}
	return ""
}

func (x *DownloadManifestAndBlobsRequest) GetReference() string {
	if x != nil {
		return x.xxx_hidden_Reference
	}
	return ""
}

func (x *DownloadManifestAndBlobsRequest) SetOwner(v string) {
	x.xxx_hidden_Owner = v
}

func (x *DownloadManifestAndBlobsRequest) SetRepository(v string) {
	x.xxx_hidden_Repository = v
}

func (x *DownloadManifestAndBlobsRequest) SetReference(v string) {
	x.xxx_hidden_Reference = v
}

type DownloadManifestAndBlobsRequest_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Owner      string
	Repository string
	// Optional reference (if unspecified, will use the repository's default_branch).
	Reference string
}

func (b0 DownloadManifestAndBlobsRequest_builder) Build() *DownloadManifestAndBlobsRequest {
	m0 := &DownloadManifestAndBlobsRequest{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Owner = b.Owner
	x.xxx_hidden_Repository = b.Repository
	x.xxx_hidden_Reference = b.Reference
	return m0
}

// DownloadManifestAndBlobsResponse is the returned resolved remote module.
type DownloadManifestAndBlobsResponse struct {
	state               protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Manifest *v1alpha1.Blob         `protobuf:"bytes,1,opt,name=manifest,proto3"`
	xxx_hidden_Blobs    *[]*v1alpha1.Blob      `protobuf:"bytes,2,rep,name=blobs,proto3"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *DownloadManifestAndBlobsResponse) Reset() {
	*x = DownloadManifestAndBlobsResponse{}
	mi := &file_buf_alpha_registry_v1alpha1_download_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DownloadManifestAndBlobsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadManifestAndBlobsResponse) ProtoMessage() {}

func (x *DownloadManifestAndBlobsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_buf_alpha_registry_v1alpha1_download_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *DownloadManifestAndBlobsResponse) GetManifest() *v1alpha1.Blob {
	if x != nil {
		return x.xxx_hidden_Manifest
	}
	return nil
}

func (x *DownloadManifestAndBlobsResponse) GetBlobs() []*v1alpha1.Blob {
	if x != nil {
		if x.xxx_hidden_Blobs != nil {
			return *x.xxx_hidden_Blobs
		}
	}
	return nil
}

func (x *DownloadManifestAndBlobsResponse) SetManifest(v *v1alpha1.Blob) {
	x.xxx_hidden_Manifest = v
}

func (x *DownloadManifestAndBlobsResponse) SetBlobs(v []*v1alpha1.Blob) {
	x.xxx_hidden_Blobs = &v
}

func (x *DownloadManifestAndBlobsResponse) HasManifest() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Manifest != nil
}

func (x *DownloadManifestAndBlobsResponse) ClearManifest() {
	x.xxx_hidden_Manifest = nil
}

type DownloadManifestAndBlobsResponse_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// manifest is the manifest of the module's content.
	// The content of the manifest blob is a text encoding of an ordered list of unique paths, each path encoded as:
	//
	//	<digest_type>:<digest>[SP][SP]<path>[LF]
	//
	// The only supported digest type is 'shake256'. The shake256 digest consists of 64 bytes of lowercase hex
	// encoded output of SHAKE256.
	Manifest *v1alpha1.Blob
	// blobs is a set of blobs that closes on the module's manifest to form the
	// complete module's content.
	Blobs []*v1alpha1.Blob
}

func (b0 DownloadManifestAndBlobsResponse_builder) Build() *DownloadManifestAndBlobsResponse {
	m0 := &DownloadManifestAndBlobsResponse{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Manifest = b.Manifest
	x.xxx_hidden_Blobs = &b.Blobs
	return m0
}

var File_buf_alpha_registry_v1alpha1_download_proto protoreflect.FileDescriptor

const file_buf_alpha_registry_v1alpha1_download_proto_rawDesc = "" +
	"\n" +
	"*buf/alpha/registry/v1alpha1/download.proto\x12\x1bbuf.alpha.registry.v1alpha1\x1a&buf/alpha/module/v1alpha1/module.proto\"e\n" +
	"\x0fDownloadRequest\x12\x14\n" +
	"\x05owner\x18\x01 \x01(\tR\x05owner\x12\x1e\n" +
	"\n" +
	"repository\x18\x02 \x01(\tR\n" +
	"repository\x12\x1c\n" +
	"\treference\x18\x03 \x01(\tR\treference\"M\n" +
	"\x10DownloadResponse\x129\n" +
	"\x06module\x18\x01 \x01(\v2!.buf.alpha.module.v1alpha1.ModuleR\x06module\"u\n" +
	"\x1fDownloadManifestAndBlobsRequest\x12\x14\n" +
	"\x05owner\x18\x01 \x01(\tR\x05owner\x12\x1e\n" +
	"\n" +
	"repository\x18\x02 \x01(\tR\n" +
	"repository\x12\x1c\n" +
	"\treference\x18\x03 \x01(\tR\treference\"\x96\x01\n" +
	" DownloadManifestAndBlobsResponse\x12;\n" +
	"\bmanifest\x18\x01 \x01(\v2\x1f.buf.alpha.module.v1alpha1.BlobR\bmanifest\x125\n" +
	"\x05blobs\x18\x02 \x03(\v2\x1f.buf.alpha.module.v1alpha1.BlobR\x05blobs2\x9e\x02\n" +
	"\x0fDownloadService\x12l\n" +
	"\bDownload\x12,.buf.alpha.registry.v1alpha1.DownloadRequest\x1a-.buf.alpha.registry.v1alpha1.DownloadResponse\"\x03\x90\x02\x01\x12\x9c\x01\n" +
	"\x18DownloadManifestAndBlobs\x12<.buf.alpha.registry.v1alpha1.DownloadManifestAndBlobsRequest\x1a=.buf.alpha.registry.v1alpha1.DownloadManifestAndBlobsResponse\"\x03\x90\x02\x01B\x9a\x02\n" +
	"\x1fcom.buf.alpha.registry.v1alpha1B\rDownloadProtoP\x01ZYgithub.com/bufbuild/buf/private/gen/proto/go/buf/alpha/registry/v1alpha1;registryv1alpha1\xa2\x02\x03BAR\xaa\x02\x1bBuf.Alpha.Registry.V1alpha1\xca\x02\x1bBuf\\Alpha\\Registry\\V1alpha1\xe2\x02'Buf\\Alpha\\Registry\\V1alpha1\\GPBMetadata\xea\x02\x1eBuf::Alpha::Registry::V1alpha1b\x06proto3"

var file_buf_alpha_registry_v1alpha1_download_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_buf_alpha_registry_v1alpha1_download_proto_goTypes = []any{
	(*DownloadRequest)(nil),                  // 0: buf.alpha.registry.v1alpha1.DownloadRequest
	(*DownloadResponse)(nil),                 // 1: buf.alpha.registry.v1alpha1.DownloadResponse
	(*DownloadManifestAndBlobsRequest)(nil),  // 2: buf.alpha.registry.v1alpha1.DownloadManifestAndBlobsRequest
	(*DownloadManifestAndBlobsResponse)(nil), // 3: buf.alpha.registry.v1alpha1.DownloadManifestAndBlobsResponse
	(*v1alpha1.Module)(nil),                  // 4: buf.alpha.module.v1alpha1.Module
	(*v1alpha1.Blob)(nil),                    // 5: buf.alpha.module.v1alpha1.Blob
}
var file_buf_alpha_registry_v1alpha1_download_proto_depIdxs = []int32{
	4, // 0: buf.alpha.registry.v1alpha1.DownloadResponse.module:type_name -> buf.alpha.module.v1alpha1.Module
	5, // 1: buf.alpha.registry.v1alpha1.DownloadManifestAndBlobsResponse.manifest:type_name -> buf.alpha.module.v1alpha1.Blob
	5, // 2: buf.alpha.registry.v1alpha1.DownloadManifestAndBlobsResponse.blobs:type_name -> buf.alpha.module.v1alpha1.Blob
	0, // 3: buf.alpha.registry.v1alpha1.DownloadService.Download:input_type -> buf.alpha.registry.v1alpha1.DownloadRequest
	2, // 4: buf.alpha.registry.v1alpha1.DownloadService.DownloadManifestAndBlobs:input_type -> buf.alpha.registry.v1alpha1.DownloadManifestAndBlobsRequest
	1, // 5: buf.alpha.registry.v1alpha1.DownloadService.Download:output_type -> buf.alpha.registry.v1alpha1.DownloadResponse
	3, // 6: buf.alpha.registry.v1alpha1.DownloadService.DownloadManifestAndBlobs:output_type -> buf.alpha.registry.v1alpha1.DownloadManifestAndBlobsResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_buf_alpha_registry_v1alpha1_download_proto_init() }
func file_buf_alpha_registry_v1alpha1_download_proto_init() {
	if File_buf_alpha_registry_v1alpha1_download_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_buf_alpha_registry_v1alpha1_download_proto_rawDesc), len(file_buf_alpha_registry_v1alpha1_download_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_buf_alpha_registry_v1alpha1_download_proto_goTypes,
		DependencyIndexes: file_buf_alpha_registry_v1alpha1_download_proto_depIdxs,
		MessageInfos:      file_buf_alpha_registry_v1alpha1_download_proto_msgTypes,
	}.Build()
	File_buf_alpha_registry_v1alpha1_download_proto = out.File
	file_buf_alpha_registry_v1alpha1_download_proto_goTypes = nil
	file_buf_alpha_registry_v1alpha1_download_proto_depIdxs = nil
}
