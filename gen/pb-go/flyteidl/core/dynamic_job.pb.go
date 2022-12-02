// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: flyteidl/core/dynamic_job.proto

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

// Describes a set of tasks to execute and how the final outputs are produced.
type DynamicJobSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A collection of nodes to execute.
	Nodes []*Node `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
	// An absolute number of successful completions of nodes required to mark this job as succeeded. As soon as this
	// criteria is met, the dynamic job will be marked as successful and outputs will be computed. If this number
	// becomes impossible to reach (e.g. number of currently running tasks + number of already succeeded tasks <
	// min_successes) the task will be aborted immediately and marked as failed. The default value of this field, if not
	// specified, is the count of nodes repeated field.
	MinSuccesses int64 `protobuf:"varint,2,opt,name=min_successes,json=minSuccesses,proto3" json:"min_successes,omitempty"`
	// Describes how to bind the final output of the dynamic job from the outputs of executed nodes. The referenced ids
	// in bindings should have the generated id for the subtask.
	Outputs []*Binding `protobuf:"bytes,3,rep,name=outputs,proto3" json:"outputs,omitempty"`
	// [Optional] A complete list of task specs referenced in nodes.
	Tasks []*TaskTemplate `protobuf:"bytes,4,rep,name=tasks,proto3" json:"tasks,omitempty"`
	// [Optional] A complete list of task specs referenced in nodes.
	Subworkflows []*WorkflowTemplate `protobuf:"bytes,5,rep,name=subworkflows,proto3" json:"subworkflows,omitempty"`
}

func (x *DynamicJobSpec) Reset() {
	*x = DynamicJobSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_core_dynamic_job_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DynamicJobSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DynamicJobSpec) ProtoMessage() {}

func (x *DynamicJobSpec) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_core_dynamic_job_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DynamicJobSpec.ProtoReflect.Descriptor instead.
func (*DynamicJobSpec) Descriptor() ([]byte, []int) {
	return file_flyteidl_core_dynamic_job_proto_rawDescGZIP(), []int{0}
}

func (x *DynamicJobSpec) GetNodes() []*Node {
	if x != nil {
		return x.Nodes
	}
	return nil
}

func (x *DynamicJobSpec) GetMinSuccesses() int64 {
	if x != nil {
		return x.MinSuccesses
	}
	return 0
}

func (x *DynamicJobSpec) GetOutputs() []*Binding {
	if x != nil {
		return x.Outputs
	}
	return nil
}

func (x *DynamicJobSpec) GetTasks() []*TaskTemplate {
	if x != nil {
		return x.Tasks
	}
	return nil
}

func (x *DynamicJobSpec) GetSubworkflows() []*WorkflowTemplate {
	if x != nil {
		return x.Subworkflows
	}
	return nil
}

var File_flyteidl_core_dynamic_job_proto protoreflect.FileDescriptor

var file_flyteidl_core_dynamic_job_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f,
	0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f, 0x6a, 0x6f, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0d, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x1a, 0x1c, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f,
	0x6c, 0x69, 0x74, 0x65, 0x72, 0x61, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19,
	0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x74, 0x61,
	0x73, 0x6b, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x66, 0x6c, 0x79, 0x74, 0x65,
	0x69, 0x64, 0x6c, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f,
	0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a, 0x02, 0x0a, 0x0e, 0x44, 0x79, 0x6e, 0x61,
	0x6d, 0x69, 0x63, 0x4a, 0x6f, 0x62, 0x53, 0x70, 0x65, 0x63, 0x12, 0x29, 0x0a, 0x05, 0x6e, 0x6f,
	0x64, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x66, 0x6c, 0x79, 0x74,
	0x65, 0x69, 0x64, 0x6c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x05,
	0x6e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x69, 0x6e, 0x5f, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x6d, 0x69,
	0x6e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x65, 0x73, 0x12, 0x30, 0x0a, 0x07, 0x6f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x66, 0x6c,
	0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x42, 0x69, 0x6e, 0x64,
	0x69, 0x6e, 0x67, 0x52, 0x07, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x12, 0x31, 0x0a, 0x05,
	0x74, 0x61, 0x73, 0x6b, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x66, 0x6c,
	0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x54, 0x61, 0x73, 0x6b,
	0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x12,
	0x43, 0x0a, 0x0c, 0x73, 0x75, 0x62, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x54, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x0c, 0x73, 0x75, 0x62, 0x77, 0x6f, 0x72, 0x6b, 0x66,
	0x6c, 0x6f, 0x77, 0x73, 0x42, 0xae, 0x01, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x2e, 0x66, 0x6c, 0x79,
	0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x42, 0x0f, 0x44, 0x79, 0x6e, 0x61,
	0x6d, 0x69, 0x63, 0x4a, 0x6f, 0x62, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x48, 0x02, 0x50, 0x01, 0x5a,
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
	file_flyteidl_core_dynamic_job_proto_rawDescOnce sync.Once
	file_flyteidl_core_dynamic_job_proto_rawDescData = file_flyteidl_core_dynamic_job_proto_rawDesc
)

func file_flyteidl_core_dynamic_job_proto_rawDescGZIP() []byte {
	file_flyteidl_core_dynamic_job_proto_rawDescOnce.Do(func() {
		file_flyteidl_core_dynamic_job_proto_rawDescData = protoimpl.X.CompressGZIP(file_flyteidl_core_dynamic_job_proto_rawDescData)
	})
	return file_flyteidl_core_dynamic_job_proto_rawDescData
}

var file_flyteidl_core_dynamic_job_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_flyteidl_core_dynamic_job_proto_goTypes = []interface{}{
	(*DynamicJobSpec)(nil),   // 0: flyteidl.core.DynamicJobSpec
	(*Node)(nil),             // 1: flyteidl.core.Node
	(*Binding)(nil),          // 2: flyteidl.core.Binding
	(*TaskTemplate)(nil),     // 3: flyteidl.core.TaskTemplate
	(*WorkflowTemplate)(nil), // 4: flyteidl.core.WorkflowTemplate
}
var file_flyteidl_core_dynamic_job_proto_depIdxs = []int32{
	1, // 0: flyteidl.core.DynamicJobSpec.nodes:type_name -> flyteidl.core.Node
	2, // 1: flyteidl.core.DynamicJobSpec.outputs:type_name -> flyteidl.core.Binding
	3, // 2: flyteidl.core.DynamicJobSpec.tasks:type_name -> flyteidl.core.TaskTemplate
	4, // 3: flyteidl.core.DynamicJobSpec.subworkflows:type_name -> flyteidl.core.WorkflowTemplate
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_flyteidl_core_dynamic_job_proto_init() }
func file_flyteidl_core_dynamic_job_proto_init() {
	if File_flyteidl_core_dynamic_job_proto != nil {
		return
	}
	file_flyteidl_core_literals_proto_init()
	file_flyteidl_core_tasks_proto_init()
	file_flyteidl_core_workflow_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_flyteidl_core_dynamic_job_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DynamicJobSpec); i {
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
			RawDescriptor: file_flyteidl_core_dynamic_job_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_flyteidl_core_dynamic_job_proto_goTypes,
		DependencyIndexes: file_flyteidl_core_dynamic_job_proto_depIdxs,
		MessageInfos:      file_flyteidl_core_dynamic_job_proto_msgTypes,
	}.Build()
	File_flyteidl_core_dynamic_job_proto = out.File
	file_flyteidl_core_dynamic_job_proto_rawDesc = nil
	file_flyteidl_core_dynamic_job_proto_goTypes = nil
	file_flyteidl_core_dynamic_job_proto_depIdxs = nil
}
