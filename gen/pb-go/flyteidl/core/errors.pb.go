// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: flyteidl/core/errors.proto

package core

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

// Defines a generic error type that dictates the behavior of the retry strategy.
type ContainerError_Kind int32

const (
	ContainerError_NON_RECOVERABLE ContainerError_Kind = 0
	ContainerError_RECOVERABLE     ContainerError_Kind = 1
)

// Enum value maps for ContainerError_Kind.
var (
	ContainerError_Kind_name = map[int32]string{
		0: "NON_RECOVERABLE",
		1: "RECOVERABLE",
	}
	ContainerError_Kind_value = map[string]int32{
		"NON_RECOVERABLE": 0,
		"RECOVERABLE":     1,
	}
)

func (x ContainerError_Kind) Enum() *ContainerError_Kind {
	p := new(ContainerError_Kind)
	*p = x
	return p
}

func (x ContainerError_Kind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ContainerError_Kind) Descriptor() protoreflect.EnumDescriptor {
	return file_flyteidl_core_errors_proto_enumTypes[0].Descriptor()
}

func (ContainerError_Kind) Type() protoreflect.EnumType {
	return &file_flyteidl_core_errors_proto_enumTypes[0]
}

func (x ContainerError_Kind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ContainerError_Kind.Descriptor instead.
func (ContainerError_Kind) EnumDescriptor() ([]byte, []int) {
	return file_flyteidl_core_errors_proto_rawDescGZIP(), []int{0, 0}
}

// Error message to propagate detailed errors from container executions to the execution
// engine.
type ContainerError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A simplified code for errors, so that we can provide a glossary of all possible errors.
	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	// A detailed error message.
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// An abstract error kind for this error. Defaults to Non_Recoverable if not specified.
	Kind ContainerError_Kind `protobuf:"varint,3,opt,name=kind,proto3,enum=flyteidl.core.ContainerError_Kind" json:"kind,omitempty"`
	// Defines the origin of the error (system, user, unknown).
	Origin ExecutionError_ErrorKind `protobuf:"varint,4,opt,name=origin,proto3,enum=flyteidl.core.ExecutionError_ErrorKind" json:"origin,omitempty"`
}

func (x *ContainerError) Reset() {
	*x = ContainerError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_core_errors_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContainerError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContainerError) ProtoMessage() {}

func (x *ContainerError) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_core_errors_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContainerError.ProtoReflect.Descriptor instead.
func (*ContainerError) Descriptor() ([]byte, []int) {
	return file_flyteidl_core_errors_proto_rawDescGZIP(), []int{0}
}

func (x *ContainerError) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *ContainerError) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ContainerError) GetKind() ContainerError_Kind {
	if x != nil {
		return x.Kind
	}
	return ContainerError_NON_RECOVERABLE
}

func (x *ContainerError) GetOrigin() ExecutionError_ErrorKind {
	if x != nil {
		return x.Origin
	}
	return ExecutionError_UNKNOWN
}

// Defines the errors.pb file format the container can produce to communicate
// failure reasons to the execution engine.
type ErrorDocument struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The error raised during execution.
	Error *ContainerError `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ErrorDocument) Reset() {
	*x = ErrorDocument{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_core_errors_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorDocument) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorDocument) ProtoMessage() {}

func (x *ErrorDocument) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_core_errors_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorDocument.ProtoReflect.Descriptor instead.
func (*ErrorDocument) Descriptor() ([]byte, []int) {
	return file_flyteidl_core_errors_proto_rawDescGZIP(), []int{1}
}

func (x *ErrorDocument) GetError() *ContainerError {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_flyteidl_core_errors_proto protoreflect.FileDescriptor

var file_flyteidl_core_errors_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x66, 0x6c,
	0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x1a, 0x1d, 0x66, 0x6c, 0x79,
	0x74, 0x65, 0x69, 0x64, 0x6c, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x65, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe5, 0x01, 0x0a, 0x0e, 0x43,
	0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x36, 0x0a, 0x04, 0x6b,
	0x69, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x66, 0x6c, 0x79, 0x74,
	0x65, 0x69, 0x64, 0x6c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x65, 0x72, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x4b, 0x69, 0x6e, 0x64, 0x52, 0x04, 0x6b,
	0x69, 0x6e, 0x64, 0x12, 0x3f, 0x0a, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4b, 0x69, 0x6e, 0x64, 0x52, 0x06, 0x6f, 0x72,
	0x69, 0x67, 0x69, 0x6e, 0x22, 0x2c, 0x0a, 0x04, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x13, 0x0a, 0x0f,
	0x4e, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x56, 0x45, 0x52, 0x41, 0x42, 0x4c, 0x45, 0x10,
	0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x52, 0x45, 0x43, 0x4f, 0x56, 0x45, 0x52, 0x41, 0x42, 0x4c, 0x45,
	0x10, 0x01, 0x22, 0x44, 0x0a, 0x0d, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x44, 0x6f, 0x63, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x33, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x42, 0xaa, 0x01, 0x0a, 0x11, 0x63, 0x6f, 0x6d,
	0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x42, 0x0b,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x48, 0x02, 0x50, 0x01, 0x5a,
	0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x6c, 0x79, 0x74,
	0x65, 0x6f, 0x72, 0x67, 0x2f, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70,
	0x62, 0x2d, 0x67, 0x6f, 0x2f, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2f, 0x63, 0x6f,
	0x72, 0x65, 0xa2, 0x02, 0x03, 0x46, 0x43, 0x58, 0xaa, 0x02, 0x0d, 0x46, 0x6c, 0x79, 0x74, 0x65,
	0x69, 0x64, 0x6c, 0x2e, 0x43, 0x6f, 0x72, 0x65, 0xca, 0x02, 0x0d, 0x46, 0x6c, 0x79, 0x74, 0x65,
	0x69, 0x64, 0x6c, 0x5c, 0x43, 0x6f, 0x72, 0x65, 0xe2, 0x02, 0x19, 0x46, 0x6c, 0x79, 0x74, 0x65,
	0x69, 0x64, 0x6c, 0x5c, 0x43, 0x6f, 0x72, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0e, 0x46, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x3a,
	0x3a, 0x43, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_flyteidl_core_errors_proto_rawDescOnce sync.Once
	file_flyteidl_core_errors_proto_rawDescData = file_flyteidl_core_errors_proto_rawDesc
)

func file_flyteidl_core_errors_proto_rawDescGZIP() []byte {
	file_flyteidl_core_errors_proto_rawDescOnce.Do(func() {
		file_flyteidl_core_errors_proto_rawDescData = protoimpl.X.CompressGZIP(file_flyteidl_core_errors_proto_rawDescData)
	})
	return file_flyteidl_core_errors_proto_rawDescData
}

var file_flyteidl_core_errors_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_flyteidl_core_errors_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_flyteidl_core_errors_proto_goTypes = []interface{}{
	(ContainerError_Kind)(0),      // 0: flyteidl.core.ContainerError.Kind
	(*ContainerError)(nil),        // 1: flyteidl.core.ContainerError
	(*ErrorDocument)(nil),         // 2: flyteidl.core.ErrorDocument
	(ExecutionError_ErrorKind)(0), // 3: flyteidl.core.ExecutionError.ErrorKind
}
var file_flyteidl_core_errors_proto_depIdxs = []int32{
	0, // 0: flyteidl.core.ContainerError.kind:type_name -> flyteidl.core.ContainerError.Kind
	3, // 1: flyteidl.core.ContainerError.origin:type_name -> flyteidl.core.ExecutionError.ErrorKind
	1, // 2: flyteidl.core.ErrorDocument.error:type_name -> flyteidl.core.ContainerError
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_flyteidl_core_errors_proto_init() }
func file_flyteidl_core_errors_proto_init() {
	if File_flyteidl_core_errors_proto != nil {
		return
	}
	file_flyteidl_core_execution_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_flyteidl_core_errors_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContainerError); i {
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
		file_flyteidl_core_errors_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorDocument); i {
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
			RawDescriptor: file_flyteidl_core_errors_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_flyteidl_core_errors_proto_goTypes,
		DependencyIndexes: file_flyteidl_core_errors_proto_depIdxs,
		EnumInfos:         file_flyteidl_core_errors_proto_enumTypes,
		MessageInfos:      file_flyteidl_core_errors_proto_msgTypes,
	}.Build()
	File_flyteidl_core_errors_proto = out.File
	file_flyteidl_core_errors_proto_rawDesc = nil
	file_flyteidl_core_errors_proto_goTypes = nil
	file_flyteidl_core_errors_proto_depIdxs = nil
}
