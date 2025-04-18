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
// source: buf/registry/module/v1/graph.proto

//go:build !protoopaque

package modulev1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	_ "buf.build/gen/go/bufbuild/registry/protocolbuffers/go/buf/registry/priv/extension/v1beta1"
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

// A dependency graph.
type Graph struct {
	state protoimpl.MessageState `protogen:"hybrid.v1"`
	// The nodes of the graph, each of which are Commits.
	Commits []*Commit `protobuf:"bytes,1,rep,name=commits,proto3" json:"commits,omitempty"`
	// The edges of the graph.
	Edges         []*Graph_Edge `protobuf:"bytes,2,rep,name=edges,proto3" json:"edges,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Graph) Reset() {
	*x = Graph{}
	mi := &file_buf_registry_module_v1_graph_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Graph) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Graph) ProtoMessage() {}

func (x *Graph) ProtoReflect() protoreflect.Message {
	mi := &file_buf_registry_module_v1_graph_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *Graph) GetCommits() []*Commit {
	if x != nil {
		return x.Commits
	}
	return nil
}

func (x *Graph) GetEdges() []*Graph_Edge {
	if x != nil {
		return x.Edges
	}
	return nil
}

func (x *Graph) SetCommits(v []*Commit) {
	x.Commits = v
}

func (x *Graph) SetEdges(v []*Graph_Edge) {
	x.Edges = v
}

type Graph_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// The nodes of the graph, each of which are Commits.
	Commits []*Commit
	// The edges of the graph.
	Edges []*Graph_Edge
}

func (b0 Graph_builder) Build() *Graph {
	m0 := &Graph{}
	b, x := &b0, m0
	_, _ = b, x
	x.Commits = b.Commits
	x.Edges = b.Edges
	return m0
}

// A node in the dependency graph.
type Graph_Node struct {
	state protoimpl.MessageState `protogen:"hybrid.v1"`
	// The commit of the node.
	CommitId      string `protobuf:"bytes,1,opt,name=commit_id,json=commitId,proto3" json:"commit_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Graph_Node) Reset() {
	*x = Graph_Node{}
	mi := &file_buf_registry_module_v1_graph_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Graph_Node) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Graph_Node) ProtoMessage() {}

func (x *Graph_Node) ProtoReflect() protoreflect.Message {
	mi := &file_buf_registry_module_v1_graph_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *Graph_Node) GetCommitId() string {
	if x != nil {
		return x.CommitId
	}
	return ""
}

func (x *Graph_Node) SetCommitId(v string) {
	x.CommitId = v
}

type Graph_Node_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// The commit of the node.
	CommitId string
}

func (b0 Graph_Node_builder) Build() *Graph_Node {
	m0 := &Graph_Node{}
	b, x := &b0, m0
	_, _ = b, x
	x.CommitId = b.CommitId
	return m0
}

// An edge in the dependency graph.
type Graph_Edge struct {
	state protoimpl.MessageState `protogen:"hybrid.v1"`
	// The Node of the start of the edge.
	FromNode *Graph_Node `protobuf:"bytes,1,opt,name=from_node,json=fromNode,proto3" json:"from_node,omitempty"`
	// The Node of the end of the edge.
	ToNode        *Graph_Node `protobuf:"bytes,2,opt,name=to_node,json=toNode,proto3" json:"to_node,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Graph_Edge) Reset() {
	*x = Graph_Edge{}
	mi := &file_buf_registry_module_v1_graph_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Graph_Edge) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Graph_Edge) ProtoMessage() {}

func (x *Graph_Edge) ProtoReflect() protoreflect.Message {
	mi := &file_buf_registry_module_v1_graph_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *Graph_Edge) GetFromNode() *Graph_Node {
	if x != nil {
		return x.FromNode
	}
	return nil
}

func (x *Graph_Edge) GetToNode() *Graph_Node {
	if x != nil {
		return x.ToNode
	}
	return nil
}

func (x *Graph_Edge) SetFromNode(v *Graph_Node) {
	x.FromNode = v
}

func (x *Graph_Edge) SetToNode(v *Graph_Node) {
	x.ToNode = v
}

func (x *Graph_Edge) HasFromNode() bool {
	if x == nil {
		return false
	}
	return x.FromNode != nil
}

func (x *Graph_Edge) HasToNode() bool {
	if x == nil {
		return false
	}
	return x.ToNode != nil
}

func (x *Graph_Edge) ClearFromNode() {
	x.FromNode = nil
}

func (x *Graph_Edge) ClearToNode() {
	x.ToNode = nil
}

type Graph_Edge_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// The Node of the start of the edge.
	FromNode *Graph_Node
	// The Node of the end of the edge.
	ToNode *Graph_Node
}

func (b0 Graph_Edge_builder) Build() *Graph_Edge {
	m0 := &Graph_Edge{}
	b, x := &b0, m0
	_, _ = b, x
	x.FromNode = b.FromNode
	x.ToNode = b.ToNode
	return m0
}

var File_buf_registry_module_v1_graph_proto protoreflect.FileDescriptor

const file_buf_registry_module_v1_graph_proto_rawDesc = "" +
	"\n" +
	"\"buf/registry/module/v1/graph.proto\x12\x16buf.registry.module.v1\x1a#buf/registry/module/v1/commit.proto\x1a3buf/registry/priv/extension/v1beta1/extension.proto\x1a\x1bbuf/validate/validate.proto\"\xd6\x02\n" +
	"\x05Graph\x12B\n" +
	"\acommits\x18\x01 \x03(\v2\x1e.buf.registry.module.v1.CommitB\b\xbaH\x05\x92\x01\x02\b\x01R\acommits\x128\n" +
	"\x05edges\x18\x02 \x03(\v2\".buf.registry.module.v1.Graph.EdgeR\x05edges\x1a0\n" +
	"\x04Node\x12(\n" +
	"\tcommit_id\x18\x01 \x01(\tB\v\xbaH\b\xc8\x01\x01r\x03\x88\x02\x01R\bcommitId\x1a\x94\x01\n" +
	"\x04Edge\x12G\n" +
	"\tfrom_node\x18\x01 \x01(\v2\".buf.registry.module.v1.Graph.NodeB\x06\xbaH\x03\xc8\x01\x01R\bfromNode\x12C\n" +
	"\ato_node\x18\x02 \x01(\v2\".buf.registry.module.v1.Graph.NodeB\x06\xbaH\x03\xc8\x01\x01R\x06toNode:\x06\xea\xc5+\x02\x10\x01BWZUbuf.build/gen/go/bufbuild/registry/protocolbuffers/go/buf/registry/module/v1;modulev1b\x06proto3"

var file_buf_registry_module_v1_graph_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_buf_registry_module_v1_graph_proto_goTypes = []any{
	(*Graph)(nil),      // 0: buf.registry.module.v1.Graph
	(*Graph_Node)(nil), // 1: buf.registry.module.v1.Graph.Node
	(*Graph_Edge)(nil), // 2: buf.registry.module.v1.Graph.Edge
	(*Commit)(nil),     // 3: buf.registry.module.v1.Commit
}
var file_buf_registry_module_v1_graph_proto_depIdxs = []int32{
	3, // 0: buf.registry.module.v1.Graph.commits:type_name -> buf.registry.module.v1.Commit
	2, // 1: buf.registry.module.v1.Graph.edges:type_name -> buf.registry.module.v1.Graph.Edge
	1, // 2: buf.registry.module.v1.Graph.Edge.from_node:type_name -> buf.registry.module.v1.Graph.Node
	1, // 3: buf.registry.module.v1.Graph.Edge.to_node:type_name -> buf.registry.module.v1.Graph.Node
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_buf_registry_module_v1_graph_proto_init() }
func file_buf_registry_module_v1_graph_proto_init() {
	if File_buf_registry_module_v1_graph_proto != nil {
		return
	}
	file_buf_registry_module_v1_commit_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_buf_registry_module_v1_graph_proto_rawDesc), len(file_buf_registry_module_v1_graph_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_buf_registry_module_v1_graph_proto_goTypes,
		DependencyIndexes: file_buf_registry_module_v1_graph_proto_depIdxs,
		MessageInfos:      file_buf_registry_module_v1_graph_proto_msgTypes,
	}.Build()
	File_buf_registry_module_v1_graph_proto = out.File
	file_buf_registry_module_v1_graph_proto_goTypes = nil
	file_buf_registry_module_v1_graph_proto_depIdxs = nil
}
