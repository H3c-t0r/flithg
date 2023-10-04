// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/core/dynamic_job.proto

package core

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Describes a set of tasks to execute and how the final outputs are produced.
type DynamicJobSpec struct {
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
	Subworkflows         []*WorkflowTemplate `protobuf:"bytes,5,rep,name=subworkflows,proto3" json:"subworkflows,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *DynamicJobSpec) Reset()         { *m = DynamicJobSpec{} }
func (m *DynamicJobSpec) String() string { return proto.CompactTextString(m) }
func (*DynamicJobSpec) ProtoMessage()    {}
func (*DynamicJobSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_87e0015f5a750c69, []int{0}
}

func (m *DynamicJobSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DynamicJobSpec.Unmarshal(m, b)
}
func (m *DynamicJobSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DynamicJobSpec.Marshal(b, m, deterministic)
}
func (m *DynamicJobSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DynamicJobSpec.Merge(m, src)
}
func (m *DynamicJobSpec) XXX_Size() int {
	return xxx_messageInfo_DynamicJobSpec.Size(m)
}
func (m *DynamicJobSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_DynamicJobSpec.DiscardUnknown(m)
}

var xxx_messageInfo_DynamicJobSpec proto.InternalMessageInfo

func (m *DynamicJobSpec) GetNodes() []*Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

func (m *DynamicJobSpec) GetMinSuccesses() int64 {
	if m != nil {
		return m.MinSuccesses
	}
	return 0
}

func (m *DynamicJobSpec) GetOutputs() []*Binding {
	if m != nil {
		return m.Outputs
	}
	return nil
}

func (m *DynamicJobSpec) GetTasks() []*TaskTemplate {
	if m != nil {
		return m.Tasks
	}
	return nil
}

func (m *DynamicJobSpec) GetSubworkflows() []*WorkflowTemplate {
	if m != nil {
		return m.Subworkflows
	}
	return nil
}

func init() {
	proto.RegisterType((*DynamicJobSpec)(nil), "flyteidl.core.DynamicJobSpec")
}

func init() { proto.RegisterFile("flyteidl/core/dynamic_job.proto", fileDescriptor_87e0015f5a750c69) }

var fileDescriptor_87e0015f5a750c69 = []byte{
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xcf, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0xd9, 0xe6, 0x14, 0xe2, 0xe6, 0x21, 0x82, 0xd4, 0x29, 0x6c, 0xe8, 0x65, 0x1e, 0x6c,
	0xfc, 0x71, 0x13, 0x4f, 0xd3, 0x93, 0x07, 0x0f, 0xdd, 0x40, 0xf0, 0x32, 0x9a, 0xf4, 0xad, 0xc6,
	0xa5, 0x79, 0xa5, 0x2f, 0x65, 0xec, 0x5f, 0xf0, 0xaf, 0x16, 0xd3, 0x75, 0xd0, 0xe2, 0x2d, 0xe4,
	0xfb, 0xf9, 0xbc, 0xe4, 0x7d, 0xd9, 0x78, 0x65, 0xb6, 0x0e, 0x74, 0x62, 0x84, 0xc2, 0x02, 0x44,
	0xb2, 0xb5, 0x71, 0xa6, 0xd5, 0xf2, 0x1b, 0x65, 0x98, 0x17, 0xe8, 0x90, 0x0f, 0x6b, 0x20, 0xfc,
	0x03, 0x46, 0xe7, 0x4d, 0xde, 0xc5, 0xb4, 0xa6, 0x8a, 0x1c, 0x5d, 0x36, 0xa3, 0x0d, 0x16, 0xeb,
	0x95, 0xc1, 0xcd, 0xff, 0xa9, 0xd1, 0x0e, 0x8a, 0xd8, 0xec, 0xdc, 0xab, 0x9f, 0x2e, 0x3b, 0x79,
	0xad, 0xde, 0x7e, 0x43, 0x39, 0xcf, 0x41, 0xf1, 0x1b, 0xd6, 0xb7, 0x98, 0x00, 0x05, 0x9d, 0x49,
	0x6f, 0x7a, 0xfc, 0x70, 0x1a, 0x36, 0x3e, 0x12, 0xbe, 0x63, 0x02, 0x51, 0x45, 0xf0, 0x6b, 0x36,
	0xcc, 0xb4, 0x5d, 0x52, 0xa9, 0x14, 0x10, 0x01, 0x05, 0xdd, 0x49, 0x67, 0xda, 0x8b, 0x06, 0x99,
	0xb6, 0xf3, 0xfa, 0x8e, 0xdf, 0xb1, 0x23, 0x2c, 0x5d, 0x5e, 0x3a, 0x0a, 0x7a, 0x7e, 0xe2, 0x59,
	0x6b, 0xe2, 0x4c, 0xdb, 0x44, 0xdb, 0x34, 0xaa, 0x31, 0x7e, 0xcf, 0xfa, 0x7e, 0xbf, 0xe0, 0xc0,
	0xf3, 0x17, 0x2d, 0x7e, 0x11, 0xd3, 0x7a, 0x01, 0x59, 0x6e, 0x62, 0x07, 0x51, 0x45, 0xf2, 0x17,
	0x36, 0xa0, 0x52, 0xd6, 0xab, 0x53, 0xd0, 0xf7, 0xe6, 0xb8, 0x65, 0x7e, 0xec, 0xf2, 0xbd, 0xdd,
	0x90, 0x66, 0xcf, 0x9f, 0x4f, 0xa9, 0x76, 0x5f, 0xa5, 0x0c, 0x15, 0x66, 0xc2, 0xab, 0x58, 0xa4,
	0xd5, 0x41, 0xec, 0x6b, 0x4c, 0xc1, 0x8a, 0x5c, 0xde, 0xa6, 0x28, 0x1a, 0xcd, 0xca, 0x43, 0xdf,
	0xe8, 0xe3, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7b, 0xfe, 0xc3, 0xc1, 0xda, 0x01, 0x00, 0x00,
}
