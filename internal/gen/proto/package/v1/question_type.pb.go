// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.3
// source: package/v1/question_type.proto

package packagev1

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

type QuestionType int32

const (
	QuestionType_QUESTION_TYPE_UNSPECIFIED  QuestionType = 0
	QuestionType_QUESTION_TYPE_STANDARD     QuestionType = 1
	QuestionType_QUESTION_TYPE_SAFE         QuestionType = 2
	QuestionType_QUESTION_TYPE_SECRET       QuestionType = 3
	QuestionType_QUESTION_TYPE_SUPER_SECRET QuestionType = 4
	QuestionType_QUESTION_TYPE_AUCTION      QuestionType = 5
)

// Enum value maps for QuestionType.
var (
	QuestionType_name = map[int32]string{
		0: "QUESTION_TYPE_UNSPECIFIED",
		1: "QUESTION_TYPE_STANDARD",
		2: "QUESTION_TYPE_SAFE",
		3: "QUESTION_TYPE_SECRET",
		4: "QUESTION_TYPE_SUPER_SECRET",
		5: "QUESTION_TYPE_AUCTION",
	}
	QuestionType_value = map[string]int32{
		"QUESTION_TYPE_UNSPECIFIED":  0,
		"QUESTION_TYPE_STANDARD":     1,
		"QUESTION_TYPE_SAFE":         2,
		"QUESTION_TYPE_SECRET":       3,
		"QUESTION_TYPE_SUPER_SECRET": 4,
		"QUESTION_TYPE_AUCTION":      5,
	}
)

func (x QuestionType) Enum() *QuestionType {
	p := new(QuestionType)
	*p = x
	return p
}

func (x QuestionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (QuestionType) Descriptor() protoreflect.EnumDescriptor {
	return file_package_v1_question_type_proto_enumTypes[0].Descriptor()
}

func (QuestionType) Type() protoreflect.EnumType {
	return &file_package_v1_question_type_proto_enumTypes[0]
}

func (x QuestionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use QuestionType.Descriptor instead.
func (QuestionType) EnumDescriptor() ([]byte, []int) {
	return file_package_v1_question_type_proto_rawDescGZIP(), []int{0}
}

var File_package_v1_question_type_proto protoreflect.FileDescriptor

var file_package_v1_question_type_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0a, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2a, 0xb6, 0x01, 0x0a,
	0x0c, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a,
	0x19, 0x51, 0x55, 0x45, 0x53, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1a, 0x0a, 0x16,
	0x51, 0x55, 0x45, 0x53, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x54,
	0x41, 0x4e, 0x44, 0x41, 0x52, 0x44, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x51, 0x55, 0x45, 0x53,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x41, 0x46, 0x45, 0x10, 0x02,
	0x12, 0x18, 0x0a, 0x14, 0x51, 0x55, 0x45, 0x53, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x53, 0x45, 0x43, 0x52, 0x45, 0x54, 0x10, 0x03, 0x12, 0x1e, 0x0a, 0x1a, 0x51, 0x55,
	0x45, 0x53, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x55, 0x50, 0x45,
	0x52, 0x5f, 0x53, 0x45, 0x43, 0x52, 0x45, 0x54, 0x10, 0x04, 0x12, 0x19, 0x0a, 0x15, 0x51, 0x55,
	0x45, 0x53, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x55, 0x43, 0x54,
	0x49, 0x4f, 0x4e, 0x10, 0x05, 0x42, 0x16, 0x5a, 0x14, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x2f, 0x76, 0x31, 0x3b, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_package_v1_question_type_proto_rawDescOnce sync.Once
	file_package_v1_question_type_proto_rawDescData = file_package_v1_question_type_proto_rawDesc
)

func file_package_v1_question_type_proto_rawDescGZIP() []byte {
	file_package_v1_question_type_proto_rawDescOnce.Do(func() {
		file_package_v1_question_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_package_v1_question_type_proto_rawDescData)
	})
	return file_package_v1_question_type_proto_rawDescData
}

var file_package_v1_question_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_package_v1_question_type_proto_goTypes = []interface{}{
	(QuestionType)(0), // 0: package.v1.QuestionType
}
var file_package_v1_question_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_package_v1_question_type_proto_init() }
func file_package_v1_question_type_proto_init() {
	if File_package_v1_question_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_package_v1_question_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_package_v1_question_type_proto_goTypes,
		DependencyIndexes: file_package_v1_question_type_proto_depIdxs,
		EnumInfos:         file_package_v1_question_type_proto_enumTypes,
	}.Build()
	File_package_v1_question_type_proto = out.File
	file_package_v1_question_type_proto_rawDesc = nil
	file_package_v1_question_type_proto_goTypes = nil
	file_package_v1_question_type_proto_depIdxs = nil
}
