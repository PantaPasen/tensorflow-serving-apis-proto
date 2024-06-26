// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: tensorflow_serving/apis/get_model_status.proto

package apis

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

// States that map to ManagerState enum in
// tensorflow_serving/core/servable_state.h
type ModelVersionStatus_State int32

const (
	// Default value.
	ModelVersionStatus_UNKNOWN ModelVersionStatus_State = 0
	// The manager is tracking this servable, but has not initiated any action
	// pertaining to it.
	ModelVersionStatus_START ModelVersionStatus_State = 10
	// The manager has decided to load this servable. In particular, checks
	// around resource availability and other aspects have passed, and the
	// manager is about to invoke the loader's Load() method.
	ModelVersionStatus_LOADING ModelVersionStatus_State = 20
	// The manager has successfully loaded this servable and made it available
	// for serving (i.e. GetServableHandle(id) will succeed). To avoid races,
	// this state is not reported until *after* the servable is made
	// available.
	ModelVersionStatus_AVAILABLE ModelVersionStatus_State = 30
	// The manager has decided to make this servable unavailable, and unload
	// it. To avoid races, this state is reported *before* the servable is
	// made unavailable.
	ModelVersionStatus_UNLOADING ModelVersionStatus_State = 40
	// This servable has reached the end of its journey in the manager. Either
	// it loaded and ultimately unloaded successfully, or it hit an error at
	// some point in its lifecycle.
	ModelVersionStatus_END ModelVersionStatus_State = 50
)

// Enum value maps for ModelVersionStatus_State.
var (
	ModelVersionStatus_State_name = map[int32]string{
		0:  "UNKNOWN",
		10: "START",
		20: "LOADING",
		30: "AVAILABLE",
		40: "UNLOADING",
		50: "END",
	}
	ModelVersionStatus_State_value = map[string]int32{
		"UNKNOWN":   0,
		"START":     10,
		"LOADING":   20,
		"AVAILABLE": 30,
		"UNLOADING": 40,
		"END":       50,
	}
)

func (x ModelVersionStatus_State) Enum() *ModelVersionStatus_State {
	p := new(ModelVersionStatus_State)
	*p = x
	return p
}

func (x ModelVersionStatus_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ModelVersionStatus_State) Descriptor() protoreflect.EnumDescriptor {
	return file_tensorflow_serving_apis_get_model_status_proto_enumTypes[0].Descriptor()
}

func (ModelVersionStatus_State) Type() protoreflect.EnumType {
	return &file_tensorflow_serving_apis_get_model_status_proto_enumTypes[0]
}

func (x ModelVersionStatus_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ModelVersionStatus_State.Descriptor instead.
func (ModelVersionStatus_State) EnumDescriptor() ([]byte, []int) {
	return file_tensorflow_serving_apis_get_model_status_proto_rawDescGZIP(), []int{1, 0}
}

// GetModelStatusRequest contains a ModelSpec indicating the model for which
// to get status.
type GetModelStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Model Specification. If version is not specified, information about all
	// versions of the model will be returned. If a version is specified, the
	// status of only that version will be returned.
	ModelSpec *ModelSpec `protobuf:"bytes,1,opt,name=model_spec,json=modelSpec,proto3" json:"model_spec,omitempty"`
}

func (x *GetModelStatusRequest) Reset() {
	*x = GetModelStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_serving_apis_get_model_status_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetModelStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetModelStatusRequest) ProtoMessage() {}

func (x *GetModelStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_serving_apis_get_model_status_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetModelStatusRequest.ProtoReflect.Descriptor instead.
func (*GetModelStatusRequest) Descriptor() ([]byte, []int) {
	return file_tensorflow_serving_apis_get_model_status_proto_rawDescGZIP(), []int{0}
}

func (x *GetModelStatusRequest) GetModelSpec() *ModelSpec {
	if x != nil {
		return x.ModelSpec
	}
	return nil
}

// Version number, state, and status for a single version of a model.
type ModelVersionStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Model version.
	Version int64 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	// Model state.
	State ModelVersionStatus_State `protobuf:"varint,2,opt,name=state,proto3,enum=tensorflow.serving.ModelVersionStatus_State" json:"state,omitempty"`
	// Model status.
	Status *StatusProto `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ModelVersionStatus) Reset() {
	*x = ModelVersionStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_serving_apis_get_model_status_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModelVersionStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModelVersionStatus) ProtoMessage() {}

func (x *ModelVersionStatus) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_serving_apis_get_model_status_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModelVersionStatus.ProtoReflect.Descriptor instead.
func (*ModelVersionStatus) Descriptor() ([]byte, []int) {
	return file_tensorflow_serving_apis_get_model_status_proto_rawDescGZIP(), []int{1}
}

func (x *ModelVersionStatus) GetVersion() int64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *ModelVersionStatus) GetState() ModelVersionStatus_State {
	if x != nil {
		return x.State
	}
	return ModelVersionStatus_UNKNOWN
}

func (x *ModelVersionStatus) GetStatus() *StatusProto {
	if x != nil {
		return x.Status
	}
	return nil
}

// Response for ModelStatusRequest on successful run.
type GetModelStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Version number and status information for applicable model version(s).
	ModelVersionStatus []*ModelVersionStatus `protobuf:"bytes,1,rep,name=model_version_status,proto3" json:"model_version_status,omitempty"`
}

func (x *GetModelStatusResponse) Reset() {
	*x = GetModelStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_serving_apis_get_model_status_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetModelStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetModelStatusResponse) ProtoMessage() {}

func (x *GetModelStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_serving_apis_get_model_status_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetModelStatusResponse.ProtoReflect.Descriptor instead.
func (*GetModelStatusResponse) Descriptor() ([]byte, []int) {
	return file_tensorflow_serving_apis_get_model_status_proto_rawDescGZIP(), []int{2}
}

func (x *GetModelStatusResponse) GetModelVersionStatus() []*ModelVersionStatus {
	if x != nil {
		return x.ModelVersionStatus
	}
	return nil
}

var File_tensorflow_serving_apis_get_model_status_proto protoreflect.FileDescriptor

var file_tensorflow_serving_apis_get_model_status_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x67, 0x65, 0x74, 0x5f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x12, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x6e, 0x67, 0x1a, 0x23, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x74, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x55, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3c, 0x0a, 0x0a, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e,
	0x67, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x53, 0x70, 0x65, 0x63, 0x52, 0x09, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x53, 0x70, 0x65, 0x63, 0x22, 0x80, 0x02, 0x0a, 0x12, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x42, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x37, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x74, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0x53, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0b, 0x0a,
	0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x53, 0x54,
	0x41, 0x52, 0x54, 0x10, 0x0a, 0x12, 0x0b, 0x0a, 0x07, 0x4c, 0x4f, 0x41, 0x44, 0x49, 0x4e, 0x47,
	0x10, 0x14, 0x12, 0x0d, 0x0a, 0x09, 0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x10,
	0x1e, 0x12, 0x0d, 0x0a, 0x09, 0x55, 0x4e, 0x4c, 0x4f, 0x41, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x28,
	0x12, 0x07, 0x0a, 0x03, 0x45, 0x4e, 0x44, 0x10, 0x32, 0x22, 0x74, 0x0a, 0x16, 0x47, 0x65, 0x74,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x5a, 0x0a, 0x14, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x26, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x14, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42,
	0xe9, 0x01, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c,
	0x6f, 0x77, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x42, 0x13, 0x47, 0x65, 0x74, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x4e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x61,
	0x6e, 0x74, 0x61, 0x70, 0x61, 0x73, 0x65, 0x6e, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2d, 0x61, 0x70, 0x69, 0x73,
	0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69,
	0x73, 0xf8, 0x01, 0x01, 0xa2, 0x02, 0x03, 0x54, 0x53, 0x58, 0xaa, 0x02, 0x12, 0x54, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0xca,
	0x02, 0x12, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5c, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x6e, 0x67, 0xe2, 0x02, 0x1e, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f,
	0x77, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x13, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c,
	0x6f, 0x77, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_serving_apis_get_model_status_proto_rawDescOnce sync.Once
	file_tensorflow_serving_apis_get_model_status_proto_rawDescData = file_tensorflow_serving_apis_get_model_status_proto_rawDesc
)

func file_tensorflow_serving_apis_get_model_status_proto_rawDescGZIP() []byte {
	file_tensorflow_serving_apis_get_model_status_proto_rawDescOnce.Do(func() {
		file_tensorflow_serving_apis_get_model_status_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_serving_apis_get_model_status_proto_rawDescData)
	})
	return file_tensorflow_serving_apis_get_model_status_proto_rawDescData
}

var file_tensorflow_serving_apis_get_model_status_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_tensorflow_serving_apis_get_model_status_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_tensorflow_serving_apis_get_model_status_proto_goTypes = []interface{}{
	(ModelVersionStatus_State)(0),  // 0: tensorflow.serving.ModelVersionStatus.State
	(*GetModelStatusRequest)(nil),  // 1: tensorflow.serving.GetModelStatusRequest
	(*ModelVersionStatus)(nil),     // 2: tensorflow.serving.ModelVersionStatus
	(*GetModelStatusResponse)(nil), // 3: tensorflow.serving.GetModelStatusResponse
	(*ModelSpec)(nil),              // 4: tensorflow.serving.ModelSpec
	(*StatusProto)(nil),            // 5: tensorflow.serving.StatusProto
}
var file_tensorflow_serving_apis_get_model_status_proto_depIdxs = []int32{
	4, // 0: tensorflow.serving.GetModelStatusRequest.model_spec:type_name -> tensorflow.serving.ModelSpec
	0, // 1: tensorflow.serving.ModelVersionStatus.state:type_name -> tensorflow.serving.ModelVersionStatus.State
	5, // 2: tensorflow.serving.ModelVersionStatus.status:type_name -> tensorflow.serving.StatusProto
	2, // 3: tensorflow.serving.GetModelStatusResponse.model_version_status:type_name -> tensorflow.serving.ModelVersionStatus
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_tensorflow_serving_apis_get_model_status_proto_init() }
func file_tensorflow_serving_apis_get_model_status_proto_init() {
	if File_tensorflow_serving_apis_get_model_status_proto != nil {
		return
	}
	file_tensorflow_serving_apis_model_proto_init()
	file_tensorflow_serving_apis_status_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_serving_apis_get_model_status_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetModelStatusRequest); i {
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
		file_tensorflow_serving_apis_get_model_status_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModelVersionStatus); i {
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
		file_tensorflow_serving_apis_get_model_status_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetModelStatusResponse); i {
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
			RawDescriptor: file_tensorflow_serving_apis_get_model_status_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_serving_apis_get_model_status_proto_goTypes,
		DependencyIndexes: file_tensorflow_serving_apis_get_model_status_proto_depIdxs,
		EnumInfos:         file_tensorflow_serving_apis_get_model_status_proto_enumTypes,
		MessageInfos:      file_tensorflow_serving_apis_get_model_status_proto_msgTypes,
	}.Build()
	File_tensorflow_serving_apis_get_model_status_proto = out.File
	file_tensorflow_serving_apis_get_model_status_proto_rawDesc = nil
	file_tensorflow_serving_apis_get_model_status_proto_goTypes = nil
	file_tensorflow_serving_apis_get_model_status_proto_depIdxs = nil
}
