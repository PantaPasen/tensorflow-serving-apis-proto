// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: tensorflow_serving/apis/model_service.proto

package apis

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_tensorflow_serving_apis_model_service_proto protoreflect.FileDescriptor

var file_tensorflow_serving_apis_model_service_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e,
	0x67, 0x1a, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x67, 0x65, 0x74, 0x5f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x32, 0xe7, 0x01, 0x0a, 0x0c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x67, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x29, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f,
	0x77, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2a, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6e, 0x0a, 0x19, 0x48,
	0x61, 0x6e, 0x64, 0x6c, 0x65, 0x52, 0x65, 0x6c, 0x6f, 0x61, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x65,
	0x6c, 0x6f, 0x61, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x28, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x65, 0x6c, 0x6f, 0x61, 0x64, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xe7, 0x01, 0x0a, 0x16,
	0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x42, 0x11, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x61, 0x6e, 0x74, 0x61, 0x70, 0x61, 0x73,
	0x65, 0x6e, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2d, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x6e, 0x67, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0xf8, 0x01, 0x01, 0xa2, 0x02,
	0x03, 0x54, 0x53, 0x58, 0xaa, 0x02, 0x12, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f,
	0x77, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0xca, 0x02, 0x12, 0x54, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0xe2, 0x02,
	0x1e, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5c, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x6e, 0x67, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x13, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x3a, 0x3a, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_tensorflow_serving_apis_model_service_proto_goTypes = []interface{}{
	(*GetModelStatusRequest)(nil),  // 0: tensorflow.serving.GetModelStatusRequest
	(*ReloadConfigRequest)(nil),    // 1: tensorflow.serving.ReloadConfigRequest
	(*GetModelStatusResponse)(nil), // 2: tensorflow.serving.GetModelStatusResponse
	(*ReloadConfigResponse)(nil),   // 3: tensorflow.serving.ReloadConfigResponse
}
var file_tensorflow_serving_apis_model_service_proto_depIdxs = []int32{
	0, // 0: tensorflow.serving.ModelService.GetModelStatus:input_type -> tensorflow.serving.GetModelStatusRequest
	1, // 1: tensorflow.serving.ModelService.HandleReloadConfigRequest:input_type -> tensorflow.serving.ReloadConfigRequest
	2, // 2: tensorflow.serving.ModelService.GetModelStatus:output_type -> tensorflow.serving.GetModelStatusResponse
	3, // 3: tensorflow.serving.ModelService.HandleReloadConfigRequest:output_type -> tensorflow.serving.ReloadConfigResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tensorflow_serving_apis_model_service_proto_init() }
func file_tensorflow_serving_apis_model_service_proto_init() {
	if File_tensorflow_serving_apis_model_service_proto != nil {
		return
	}
	file_tensorflow_serving_apis_get_model_status_proto_init()
	file_tensorflow_serving_apis_model_management_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tensorflow_serving_apis_model_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tensorflow_serving_apis_model_service_proto_goTypes,
		DependencyIndexes: file_tensorflow_serving_apis_model_service_proto_depIdxs,
	}.Build()
	File_tensorflow_serving_apis_model_service_proto = out.File
	file_tensorflow_serving_apis_model_service_proto_rawDesc = nil
	file_tensorflow_serving_apis_model_service_proto_goTypes = nil
	file_tensorflow_serving_apis_model_service_proto_depIdxs = nil
}
