// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.3
// source: package/v1/stage_question_type.proto

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

type StageQuestionType int32

const (
	StageQuestionType_STAGE_QUESTION_TYPE_UNSPECIFIED StageQuestionType = 0
	StageQuestionType_STANDARD                        StageQuestionType = 1
	StageQuestionType_SAFE                            StageQuestionType = 2
	StageQuestionType_SECRET                          StageQuestionType = 3
	StageQuestionType_SUPER_SECRET                    StageQuestionType = 4
	StageQuestionType_AUCTION                         StageQuestionType = 5
)

// Enum value maps for StageQuestionType.
var (
	StageQuestionType_name = map[int32]string{
		0: "STAGE_QUESTION_TYPE_UNSPECIFIED",
		1: "STANDARD",
		2: "SAFE",
		3: "SECRET",
		4: "SUPER_SECRET",
		5: "AUCTION",
	}
	StageQuestionType_value = map[string]int32{
		"STAGE_QUESTION_TYPE_UNSPECIFIED": 0,
		"STANDARD":                        1,
		"SAFE":                            2,
		"SECRET":                          3,
		"SUPER_SECRET":                    4,
		"AUCTION":                         5,
	}
)

func (x StageQuestionType) Enum() *StageQuestionType {
	p := new(StageQuestionType)
	*p = x
	return p
}

func (x StageQuestionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StageQuestionType) Descriptor() protoreflect.EnumDescriptor {
	return file_package_v1_stage_question_type_proto_enumTypes[0].Descriptor()
}

func (StageQuestionType) Type() protoreflect.EnumType {
	return &file_package_v1_stage_question_type_proto_enumTypes[0]
}

func (x StageQuestionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StageQuestionType.Descriptor instead.
func (StageQuestionType) EnumDescriptor() ([]byte, []int) {
	return file_package_v1_stage_question_type_proto_rawDescGZIP(), []int{0}
}

var File_package_v1_stage_question_type_proto protoreflect.FileDescriptor

var file_package_v1_stage_question_type_proto_rawDesc = []byte{
	0x0a, 0x24, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61,
	0x67, 0x65, 0x5f, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e,
	0x76, 0x31, 0x2a, 0x7b, 0x0a, 0x11, 0x53, 0x74, 0x61, 0x67, 0x65, 0x51, 0x75, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x1f, 0x53, 0x54, 0x41, 0x47, 0x45,
	0x5f, 0x51, 0x55, 0x45, 0x53, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08,
	0x53, 0x54, 0x41, 0x4e, 0x44, 0x41, 0x52, 0x44, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x41,
	0x46, 0x45, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x45, 0x43, 0x52, 0x45, 0x54, 0x10, 0x03,
	0x12, 0x10, 0x0a, 0x0c, 0x53, 0x55, 0x50, 0x45, 0x52, 0x5f, 0x53, 0x45, 0x43, 0x52, 0x45, 0x54,
	0x10, 0x04, 0x12, 0x0b, 0x0a, 0x07, 0x41, 0x55, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x05, 0x42,
	0x16, 0x5a, 0x14, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_package_v1_stage_question_type_proto_rawDescOnce sync.Once
	file_package_v1_stage_question_type_proto_rawDescData = file_package_v1_stage_question_type_proto_rawDesc
)

func file_package_v1_stage_question_type_proto_rawDescGZIP() []byte {
	file_package_v1_stage_question_type_proto_rawDescOnce.Do(func() {
		file_package_v1_stage_question_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_package_v1_stage_question_type_proto_rawDescData)
	})
	return file_package_v1_stage_question_type_proto_rawDescData
}

var file_package_v1_stage_question_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_package_v1_stage_question_type_proto_goTypes = []interface{}{
	(StageQuestionType)(0), // 0: package.v1.StageQuestionType
}
var file_package_v1_stage_question_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_package_v1_stage_question_type_proto_init() }
func file_package_v1_stage_question_type_proto_init() {
	if File_package_v1_stage_question_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_package_v1_stage_question_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_package_v1_stage_question_type_proto_goTypes,
		DependencyIndexes: file_package_v1_stage_question_type_proto_depIdxs,
		EnumInfos:         file_package_v1_stage_question_type_proto_enumTypes,
	}.Build()
	File_package_v1_stage_question_type_proto = out.File
	file_package_v1_stage_question_type_proto_rawDesc = nil
	file_package_v1_stage_question_type_proto_goTypes = nil
	file_package_v1_stage_question_type_proto_depIdxs = nil
}
