// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: tensorflow/core/framework/node_def.proto

package framework

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

type NodeDef struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name given to this operator. Used for naming inputs,
	// logging, visualization, etc.  Unique within a single GraphDef.
	// Must match the regexp "[A-Za-z0-9.][A-Za-z0-9_>./]*".
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The operation name.  There may be custom parameters in attrs.
	// Op names starting with an underscore are reserved for internal use.
	Op string `protobuf:"bytes,2,opt,name=op,proto3" json:"op,omitempty"`
	// Each input is "node:src_output" with "node" being a string name and
	// "src_output" indicating which output tensor to use from "node". If
	// "src_output" is 0 the ":0" suffix can be omitted.  Regular inputs
	// may optionally be followed by control inputs that have the format
	// "^node".
	Input []string `protobuf:"bytes,3,rep,name=input,proto3" json:"input,omitempty"`
	// A (possibly partial) specification for the device on which this
	// node should be placed.
	// The expected syntax for this string is as follows:
	//
	// DEVICE_SPEC ::= PARTIAL_SPEC
	//
	// PARTIAL_SPEC ::= ("/" CONSTRAINT) *
	// CONSTRAINT ::= ("job:" JOB_NAME)
	//
	//	| ("replica:" [1-9][0-9]*)
	//	| ("task:" [1-9][0-9]*)
	//	| ("device:" [A-Za-z]* ":" ([1-9][0-9]* | "*") )
	//
	// Valid values for this string include:
	// * "/job:worker/replica:0/task:1/device:GPU:3"  (full specification)
	// * "/job:worker/device:GPU:3"                   (partial specification)
	// * ""                                    (no specification)
	//
	// If the constraints do not resolve to a single device (or if this
	// field is empty or not present), the runtime will attempt to
	// choose a device automatically.
	Device string `protobuf:"bytes,4,opt,name=device,proto3" json:"device,omitempty"`
	// Operation-specific graph-construction-time configuration.
	// Note that this should include all attrs defined in the
	// corresponding OpDef, including those with a value matching
	// the default -- this allows the default to change and makes
	// NodeDefs easier to interpret on their own.  However, if
	// an attr with a default is not specified in this list, the
	// default will be used.
	// The "names" (keys) must match the regexp "[a-z][a-z0-9_]+" (and
	// one of the names from the corresponding OpDef's attr field).
	// The values must have a type matching the corresponding OpDef
	// attr's type field.
	// TODO(josh11b): Add some examples here showing best practices.
	Attr map[string]*AttrValue `protobuf:"bytes,5,rep,name=attr,proto3" json:"attr,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// This stores debug information associated with the node.
	ExperimentalDebugInfo *NodeDef_ExperimentalDebugInfo `protobuf:"bytes,6,opt,name=experimental_debug_info,json=experimentalDebugInfo,proto3" json:"experimental_debug_info,omitempty"`
	// The complete type of this node. Experimental and subject to change.
	// Currently, the field only contains the return types of the node. That will
	// extend in the future to contain the entire signature of the node, as a
	// function type.
	ExperimentalType *FullTypeDef `protobuf:"bytes,7,opt,name=experimental_type,json=experimentalType,proto3" json:"experimental_type,omitempty"`
}

func (x *NodeDef) Reset() {
	*x = NodeDef{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_framework_node_def_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeDef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeDef) ProtoMessage() {}

func (x *NodeDef) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_framework_node_def_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeDef.ProtoReflect.Descriptor instead.
func (*NodeDef) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_framework_node_def_proto_rawDescGZIP(), []int{0}
}

func (x *NodeDef) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NodeDef) GetOp() string {
	if x != nil {
		return x.Op
	}
	return ""
}

func (x *NodeDef) GetInput() []string {
	if x != nil {
		return x.Input
	}
	return nil
}

func (x *NodeDef) GetDevice() string {
	if x != nil {
		return x.Device
	}
	return ""
}

func (x *NodeDef) GetAttr() map[string]*AttrValue {
	if x != nil {
		return x.Attr
	}
	return nil
}

func (x *NodeDef) GetExperimentalDebugInfo() *NodeDef_ExperimentalDebugInfo {
	if x != nil {
		return x.ExperimentalDebugInfo
	}
	return nil
}

func (x *NodeDef) GetExperimentalType() *FullTypeDef {
	if x != nil {
		return x.ExperimentalType
	}
	return nil
}

type NodeDef_ExperimentalDebugInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Opaque string inserted into error messages created by the runtime.
	//
	// This is intended to store the list of names of the nodes from the
	// original graph that this node was derived. For example if this node, say
	// C, was result of a fusion of 2 nodes A and B, then 'original_node' would
	// be {A, B}. This information can be used to map errors originating at the
	// current node to some top level source code.
	OriginalNodeNames []string `protobuf:"bytes,1,rep,name=original_node_names,json=originalNodeNames,proto3" json:"original_node_names,omitempty"`
	// This is intended to store the list of names of the functions from the
	// original graph that this node was derived. For example if this node, say
	// C, was result of a fusion of node A in function FA and node B in function
	// FB, then `original_funcs` would be {FA, FB}. If the node is in the top
	// level graph, the `original_func` is empty. This information, with the
	// `original_node_names` can be used to map errors originating at the
	// current ndoe to some top level source code.
	OriginalFuncNames []string `protobuf:"bytes,2,rep,name=original_func_names,json=originalFuncNames,proto3" json:"original_func_names,omitempty"`
}

func (x *NodeDef_ExperimentalDebugInfo) Reset() {
	*x = NodeDef_ExperimentalDebugInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_framework_node_def_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeDef_ExperimentalDebugInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeDef_ExperimentalDebugInfo) ProtoMessage() {}

func (x *NodeDef_ExperimentalDebugInfo) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_framework_node_def_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeDef_ExperimentalDebugInfo.ProtoReflect.Descriptor instead.
func (*NodeDef_ExperimentalDebugInfo) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_framework_node_def_proto_rawDescGZIP(), []int{0, 1}
}

func (x *NodeDef_ExperimentalDebugInfo) GetOriginalNodeNames() []string {
	if x != nil {
		return x.OriginalNodeNames
	}
	return nil
}

func (x *NodeDef_ExperimentalDebugInfo) GetOriginalFuncNames() []string {
	if x != nil {
		return x.OriginalFuncNames
	}
	return nil
}

var File_tensorflow_core_framework_node_def_proto protoreflect.FileDescriptor

var file_tensorflow_core_framework_node_def_proto_rawDesc = []byte{
	0x0a, 0x28, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x6e, 0x6f, 0x64, 0x65,
	0x5f, 0x64, 0x65, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x74, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x1a, 0x2a, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c,
	0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72,
	0x6b, 0x2f, 0x61, 0x74, 0x74, 0x72, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x29, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x66, 0x75,
	0x6c, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x80, 0x04,
	0x0a, 0x07, 0x4e, 0x6f, 0x64, 0x65, 0x44, 0x65, 0x66, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x6f, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x6f, 0x70, 0x12, 0x14, 0x0a,
	0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6e,
	0x70, 0x75, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x31, 0x0a, 0x04, 0x61,
	0x74, 0x74, 0x72, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x74, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x44, 0x65, 0x66, 0x2e, 0x41,
	0x74, 0x74, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x61, 0x74, 0x74, 0x72, 0x12, 0x61,
	0x0a, 0x17, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x5f, 0x64,
	0x65, 0x62, 0x75, 0x67, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x29, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x4e, 0x6f, 0x64,
	0x65, 0x44, 0x65, 0x66, 0x2e, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61,
	0x6c, 0x44, 0x65, 0x62, 0x75, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x15, 0x65, 0x78, 0x70, 0x65,
	0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x44, 0x65, 0x62, 0x75, 0x67, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x44, 0x0a, 0x11, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61,
	0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x46, 0x75, 0x6c, 0x6c, 0x54, 0x79,
	0x70, 0x65, 0x44, 0x65, 0x66, 0x52, 0x10, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e,
	0x74, 0x61, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x1a, 0x4e, 0x0a, 0x09, 0x41, 0x74, 0x74, 0x72, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2b, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c,
	0x6f, 0x77, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x77, 0x0a, 0x15, 0x45, 0x78, 0x70, 0x65, 0x72,
	0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x44, 0x65, 0x62, 0x75, 0x67, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x2e, 0x0a, 0x13, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x6e, 0x6f, 0x64,
	0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x11, 0x6f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x4e, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73,
	0x12, 0x2e, 0x0a, 0x13, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x66, 0x75, 0x6e,
	0x63, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x11, 0x6f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x46, 0x75, 0x6e, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x73,
	0x42, 0xbb, 0x01, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x42, 0x0c, 0x4e, 0x6f, 0x64, 0x65, 0x44, 0x65, 0x66, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x50, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x70, 0x61, 0x6e, 0x74, 0x61, 0x70, 0x61, 0x73, 0x65, 0x6e, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2d, 0x61, 0x70,
	0x69, 0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x66, 0x72, 0x61, 0x6d,
	0x65, 0x77, 0x6f, 0x72, 0x6b, 0xf8, 0x01, 0x01, 0xa2, 0x02, 0x03, 0x54, 0x58, 0x58, 0xaa, 0x02,
	0x0a, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0xca, 0x02, 0x0a, 0x54, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0xe2, 0x02, 0x16, 0x54, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x0a, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_core_framework_node_def_proto_rawDescOnce sync.Once
	file_tensorflow_core_framework_node_def_proto_rawDescData = file_tensorflow_core_framework_node_def_proto_rawDesc
)

func file_tensorflow_core_framework_node_def_proto_rawDescGZIP() []byte {
	file_tensorflow_core_framework_node_def_proto_rawDescOnce.Do(func() {
		file_tensorflow_core_framework_node_def_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_core_framework_node_def_proto_rawDescData)
	})
	return file_tensorflow_core_framework_node_def_proto_rawDescData
}

var file_tensorflow_core_framework_node_def_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_tensorflow_core_framework_node_def_proto_goTypes = []interface{}{
	(*NodeDef)(nil),                       // 0: tensorflow.NodeDef
	nil,                                   // 1: tensorflow.NodeDef.AttrEntry
	(*NodeDef_ExperimentalDebugInfo)(nil), // 2: tensorflow.NodeDef.ExperimentalDebugInfo
	(*FullTypeDef)(nil),                   // 3: tensorflow.FullTypeDef
	(*AttrValue)(nil),                     // 4: tensorflow.AttrValue
}
var file_tensorflow_core_framework_node_def_proto_depIdxs = []int32{
	1, // 0: tensorflow.NodeDef.attr:type_name -> tensorflow.NodeDef.AttrEntry
	2, // 1: tensorflow.NodeDef.experimental_debug_info:type_name -> tensorflow.NodeDef.ExperimentalDebugInfo
	3, // 2: tensorflow.NodeDef.experimental_type:type_name -> tensorflow.FullTypeDef
	4, // 3: tensorflow.NodeDef.AttrEntry.value:type_name -> tensorflow.AttrValue
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_tensorflow_core_framework_node_def_proto_init() }
func file_tensorflow_core_framework_node_def_proto_init() {
	if File_tensorflow_core_framework_node_def_proto != nil {
		return
	}
	file_tensorflow_core_framework_attr_value_proto_init()
	file_tensorflow_core_framework_full_type_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_core_framework_node_def_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeDef); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tensorflow_core_framework_node_def_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeDef_ExperimentalDebugInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tensorflow_core_framework_node_def_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_core_framework_node_def_proto_goTypes,
		DependencyIndexes: file_tensorflow_core_framework_node_def_proto_depIdxs,
		MessageInfos:      file_tensorflow_core_framework_node_def_proto_msgTypes,
	}.Build()
	File_tensorflow_core_framework_node_def_proto = out.File
	file_tensorflow_core_framework_node_def_proto_rawDesc = nil
	file_tensorflow_core_framework_node_def_proto_goTypes = nil
	file_tensorflow_core_framework_node_def_proto_depIdxs = nil
}
