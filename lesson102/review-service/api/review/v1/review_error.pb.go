// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.20.1
// source: api/review/v1/review_error.proto

package v1

import (
	_ "github.com/go-kratos/kratos/v2/errors"
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

type ErrorReason int32

const (
	// 为某个枚举单独设置错误码
	ErrorReason_NEED_LOGIN     ErrorReason = 0
	ErrorReason_DB_FAILED      ErrorReason = 1
	ErrorReason_ORDER_REVIEWED ErrorReason = 100
)

// Enum value maps for ErrorReason.
var (
	ErrorReason_name = map[int32]string{
		0:   "NEED_LOGIN",
		1:   "DB_FAILED",
		100: "ORDER_REVIEWED",
	}
	ErrorReason_value = map[string]int32{
		"NEED_LOGIN":     0,
		"DB_FAILED":      1,
		"ORDER_REVIEWED": 100,
	}
)

func (x ErrorReason) Enum() *ErrorReason {
	p := new(ErrorReason)
	*p = x
	return p
}

func (x ErrorReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorReason) Descriptor() protoreflect.EnumDescriptor {
	return file_api_review_v1_review_error_proto_enumTypes[0].Descriptor()
}

func (ErrorReason) Type() protoreflect.EnumType {
	return &file_api_review_v1_review_error_proto_enumTypes[0]
}

func (x ErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorReason.Descriptor instead.
func (ErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_api_review_v1_review_error_proto_rawDescGZIP(), []int{0}
}

var File_api_review_v1_review_error_proto protoreflect.FileDescriptor

var file_api_review_v1_review_error_proto_rawDesc = []byte{
	0x0a, 0x20, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2f, 0x76, 0x31, 0x2f,
	0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0d, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x76,
	0x31, 0x1a, 0x13, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x58, 0x0a, 0x0b, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x0a, 0x4e, 0x45, 0x45, 0x44, 0x5f, 0x4c, 0x4f,
	0x47, 0x49, 0x4e, 0x10, 0x00, 0x1a, 0x04, 0xa8, 0x45, 0x91, 0x03, 0x12, 0x13, 0x0a, 0x09, 0x44,
	0x42, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x1a, 0x04, 0xa8, 0x45, 0xf4, 0x03,
	0x12, 0x18, 0x0a, 0x0e, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x52, 0x45, 0x56, 0x49, 0x45, 0x57,
	0x45, 0x44, 0x10, 0x64, 0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03, 0x1a, 0x04, 0xa0, 0x45, 0xf4, 0x03,
	0x42, 0x32, 0x0a, 0x0d, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x76,
	0x31, 0x50, 0x01, 0x5a, 0x1f, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2d, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2f, 0x76,
	0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_review_v1_review_error_proto_rawDescOnce sync.Once
	file_api_review_v1_review_error_proto_rawDescData = file_api_review_v1_review_error_proto_rawDesc
)

func file_api_review_v1_review_error_proto_rawDescGZIP() []byte {
	file_api_review_v1_review_error_proto_rawDescOnce.Do(func() {
		file_api_review_v1_review_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_review_v1_review_error_proto_rawDescData)
	})
	return file_api_review_v1_review_error_proto_rawDescData
}

var file_api_review_v1_review_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_review_v1_review_error_proto_goTypes = []interface{}{
	(ErrorReason)(0), // 0: api.review.v1.ErrorReason
}
var file_api_review_v1_review_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_review_v1_review_error_proto_init() }
func file_api_review_v1_review_error_proto_init() {
	if File_api_review_v1_review_error_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_review_v1_review_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_review_v1_review_error_proto_goTypes,
		DependencyIndexes: file_api_review_v1_review_error_proto_depIdxs,
		EnumInfos:         file_api_review_v1_review_error_proto_enumTypes,
	}.Build()
	File_api_review_v1_review_error_proto = out.File
	file_api_review_v1_review_error_proto_rawDesc = nil
	file_api_review_v1_review_error_proto_goTypes = nil
	file_api_review_v1_review_error_proto_depIdxs = nil
}
