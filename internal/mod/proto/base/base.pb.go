// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.24.2
// source: internal/mod/proto/base.proto

package base

import (
	reflect "reflect"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_internal_mod_proto_base_proto protoreflect.FileDescriptor

var file_internal_mod_proto_base_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6d, 0x6f, 0x64, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x04, 0x62, 0x61, 0x73, 0x65, 0x1a, 0x22, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x6d, 0x6f, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xae, 0x01, 0x0a, 0x04, 0x42, 0x61,
	0x73, 0x65, 0x12, 0x34, 0x0a, 0x0b, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x12, 0x0b, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16,
	0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x0b, 0x57, 0x6f, 0x72, 0x6b,
	0x65, 0x72, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x13, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x57,
	0x6f, 0x72, 0x6b, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x0b, 0x2e, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x11, 0x57,
	0x6f, 0x72, 0x6b, 0x65, 0x72, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x13, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x11, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x57, 0x6f, 0x72,
	0x6b, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x07, 0x5a, 0x05, 0x2f, 0x62,
	0x61, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_internal_mod_proto_base_proto_goTypes = []any{
	(*Empty)(nil),            // 0: base.Empty
	(*WorkerMessage)(nil),    // 1: base.WorkerMessage
	(*HealthCheckReply)(nil), // 2: base.HealthCheckReply
	(*WorkerReply)(nil),      // 3: base.WorkerReply
}
var file_internal_mod_proto_base_proto_depIdxs = []int32{
	0, // 0: base.Base.HealthCheck:input_type -> base.Empty
	1, // 1: base.Base.WorkerCheck:input_type -> base.WorkerMessage
	1, // 2: base.Base.WorkerCheckResult:input_type -> base.WorkerMessage
	2, // 3: base.Base.HealthCheck:output_type -> base.HealthCheckReply
	0, // 4: base.Base.WorkerCheck:output_type -> base.Empty
	3, // 5: base.Base.WorkerCheckResult:output_type -> base.WorkerReply
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_mod_proto_base_proto_init() }
func file_internal_mod_proto_base_proto_init() {
	if File_internal_mod_proto_base_proto != nil {
		return
	}
	file_internal_mod_proto_base_type_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_mod_proto_base_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_mod_proto_base_proto_goTypes,
		DependencyIndexes: file_internal_mod_proto_base_proto_depIdxs,
	}.Build()
	File_internal_mod_proto_base_proto = out.File
	file_internal_mod_proto_base_proto_rawDesc = nil
	file_internal_mod_proto_base_proto_goTypes = nil
	file_internal_mod_proto_base_proto_depIdxs = nil
}