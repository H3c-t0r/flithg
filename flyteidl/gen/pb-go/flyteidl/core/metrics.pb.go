// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/core/metrics.proto

package core

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
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

// Span represents a duration trace of Flyte execution. The id field denotes a Flyte execution entity or an operation
// which uniquely identifies the Span. The spans attribute allows this Span to be further broken down into more
// precise definitions.
type Span struct {
	// start_time defines the instance this span began.
	StartTime *timestamp.Timestamp `protobuf:"bytes,1,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// end_time defines the instance this span completed.
	EndTime *timestamp.Timestamp `protobuf:"bytes,2,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	// Types that are valid to be assigned to Id:
	//	*Span_WorkflowId
	//	*Span_NodeId
	//	*Span_TaskId
	//	*Span_OperationId
	Id isSpan_Id `protobuf_oneof:"id"`
	// spans defines a collection of Spans that breakdown this execution.
	Spans                []*Span  `protobuf:"bytes,7,rep,name=spans,proto3" json:"spans,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Span) Reset()         { *m = Span{} }
func (m *Span) String() string { return proto.CompactTextString(m) }
func (*Span) ProtoMessage()    {}
func (*Span) Descriptor() ([]byte, []int) {
	return fileDescriptor_756935f796ae3119, []int{0}
}

func (m *Span) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Span.Unmarshal(m, b)
}
func (m *Span) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Span.Marshal(b, m, deterministic)
}
func (m *Span) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Span.Merge(m, src)
}
func (m *Span) XXX_Size() int {
	return xxx_messageInfo_Span.Size(m)
}
func (m *Span) XXX_DiscardUnknown() {
	xxx_messageInfo_Span.DiscardUnknown(m)
}

var xxx_messageInfo_Span proto.InternalMessageInfo

func (m *Span) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *Span) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

type isSpan_Id interface {
	isSpan_Id()
}

type Span_WorkflowId struct {
	WorkflowId *WorkflowExecutionIdentifier `protobuf:"bytes,3,opt,name=workflow_id,json=workflowId,proto3,oneof"`
}

type Span_NodeId struct {
	NodeId *NodeExecutionIdentifier `protobuf:"bytes,4,opt,name=node_id,json=nodeId,proto3,oneof"`
}

type Span_TaskId struct {
	TaskId *TaskExecutionIdentifier `protobuf:"bytes,5,opt,name=task_id,json=taskId,proto3,oneof"`
}

type Span_OperationId struct {
	OperationId string `protobuf:"bytes,6,opt,name=operation_id,json=operationId,proto3,oneof"`
}

func (*Span_WorkflowId) isSpan_Id() {}

func (*Span_NodeId) isSpan_Id() {}

func (*Span_TaskId) isSpan_Id() {}

func (*Span_OperationId) isSpan_Id() {}

func (m *Span) GetId() isSpan_Id {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Span) GetWorkflowId() *WorkflowExecutionIdentifier {
	if x, ok := m.GetId().(*Span_WorkflowId); ok {
		return x.WorkflowId
	}
	return nil
}

func (m *Span) GetNodeId() *NodeExecutionIdentifier {
	if x, ok := m.GetId().(*Span_NodeId); ok {
		return x.NodeId
	}
	return nil
}

func (m *Span) GetTaskId() *TaskExecutionIdentifier {
	if x, ok := m.GetId().(*Span_TaskId); ok {
		return x.TaskId
	}
	return nil
}

func (m *Span) GetOperationId() string {
	if x, ok := m.GetId().(*Span_OperationId); ok {
		return x.OperationId
	}
	return ""
}

func (m *Span) GetSpans() []*Span {
	if m != nil {
		return m.Spans
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Span) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Span_WorkflowId)(nil),
		(*Span_NodeId)(nil),
		(*Span_TaskId)(nil),
		(*Span_OperationId)(nil),
	}
}

// ExecutionMetrics is a collection of metrics that are collected during the execution of a Flyte task.
type ExecutionMetricResult struct {
	// The metric this data represents. e.g. EXECUTION_METRIC_USED_CPU_AVG or EXECUTION_METRIC_USED_MEMORY_BYTES_AVG.
	Metric string `protobuf:"bytes,1,opt,name=metric,proto3" json:"metric,omitempty"`
	// The result data in prometheus range query result format
	// https://prometheus.io/docs/prometheus/latest/querying/api/#expression-query-result-formats.
	// This may include multiple time series, differentiated by their metric labels.
	// Start time is greater of (execution attempt start, 48h ago)
	// End time is lesser of (execution attempt end, now)
	Data                 *_struct.Struct `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ExecutionMetricResult) Reset()         { *m = ExecutionMetricResult{} }
func (m *ExecutionMetricResult) String() string { return proto.CompactTextString(m) }
func (*ExecutionMetricResult) ProtoMessage()    {}
func (*ExecutionMetricResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_756935f796ae3119, []int{1}
}

func (m *ExecutionMetricResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecutionMetricResult.Unmarshal(m, b)
}
func (m *ExecutionMetricResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecutionMetricResult.Marshal(b, m, deterministic)
}
func (m *ExecutionMetricResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecutionMetricResult.Merge(m, src)
}
func (m *ExecutionMetricResult) XXX_Size() int {
	return xxx_messageInfo_ExecutionMetricResult.Size(m)
}
func (m *ExecutionMetricResult) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecutionMetricResult.DiscardUnknown(m)
}

var xxx_messageInfo_ExecutionMetricResult proto.InternalMessageInfo

func (m *ExecutionMetricResult) GetMetric() string {
	if m != nil {
		return m.Metric
	}
	return ""
}

func (m *ExecutionMetricResult) GetData() *_struct.Struct {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*Span)(nil), "flyteidl.core.Span")
	proto.RegisterType((*ExecutionMetricResult)(nil), "flyteidl.core.ExecutionMetricResult")
}

func init() { proto.RegisterFile("flyteidl/core/metrics.proto", fileDescriptor_756935f796ae3119) }

var fileDescriptor_756935f796ae3119 = []byte{
	// 393 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x51, 0xcb, 0xd3, 0x30,
	0x14, 0x86, 0xbf, 0x7d, 0xeb, 0x3a, 0x97, 0xea, 0x4d, 0x44, 0x1d, 0x53, 0x74, 0x4c, 0x90, 0xa9,
	0x98, 0xc0, 0xc4, 0x0b, 0xc5, 0x1b, 0x07, 0xc2, 0x7a, 0x31, 0x2f, 0xba, 0x81, 0x20, 0xc2, 0x48,
	0x9b, 0xb4, 0x86, 0xb5, 0x49, 0x49, 0x52, 0xa6, 0xbf, 0xc7, 0x3f, 0x2a, 0x49, 0xda, 0xc2, 0xe6,
	0xc4, 0xbb, 0x9e, 0x9e, 0xf7, 0x79, 0xd2, 0xe6, 0x1c, 0xf0, 0x38, 0x2f, 0x7f, 0x19, 0xc6, 0x69,
	0x89, 0x33, 0xa9, 0x18, 0xae, 0x98, 0x51, 0x3c, 0xd3, 0xa8, 0x56, 0xd2, 0x48, 0x78, 0xaf, 0x6b,
	0x22, 0xdb, 0x9c, 0x3d, 0x3d, 0xcf, 0x72, 0xca, 0x84, 0xe1, 0x39, 0x67, 0xca, 0xc7, 0x67, 0xcf,
	0x0a, 0x29, 0x8b, 0x92, 0x61, 0x57, 0xa5, 0x4d, 0x8e, 0x0d, 0xaf, 0x98, 0x36, 0xa4, 0xaa, 0xdb,
	0xc0, 0x93, 0xcb, 0x80, 0x36, 0xaa, 0xc9, 0x8c, 0xef, 0x2e, 0x7e, 0x0f, 0x41, 0xb0, 0xab, 0x89,
	0x80, 0xef, 0x01, 0xd0, 0x86, 0x28, 0x73, 0xb0, 0xfc, 0x74, 0x30, 0x1f, 0x2c, 0xa3, 0xd5, 0x0c,
	0x79, 0x16, 0x75, 0x2c, 0xda, 0x77, 0xf2, 0x64, 0xe2, 0xd2, 0xb6, 0x86, 0xef, 0xc0, 0x1d, 0x26,
	0xa8, 0x07, 0x6f, 0xff, 0x0b, 0x8e, 0x99, 0xa0, 0x0e, 0xdb, 0x82, 0xe8, 0x24, 0xd5, 0x31, 0x2f,
	0xe5, 0xe9, 0xc0, 0xe9, 0x74, 0xe8, 0xc8, 0x57, 0xe8, 0xec, 0xf7, 0xd1, 0xd7, 0x36, 0xf1, 0xf9,
	0x27, 0xcb, 0x1a, 0xc3, 0xa5, 0x88, 0xfb, 0x0b, 0xd8, 0xdc, 0x24, 0xa0, 0x13, 0xc4, 0x14, 0x7e,
	0x02, 0x63, 0x21, 0x29, 0xb3, 0xaa, 0xc0, 0xa9, 0x5e, 0x5c, 0xa8, 0xbe, 0x48, 0xca, 0xae, 0x6b,
	0x42, 0x0b, 0x7a, 0x85, 0x21, 0xfa, 0x68, 0x15, 0xa3, 0xab, 0x8a, 0x3d, 0xd1, 0xc7, 0x7f, 0x28,
	0x2c, 0x18, 0x53, 0xf8, 0x1c, 0xdc, 0x95, 0x35, 0x53, 0xc4, 0x06, 0xac, 0x27, 0x9c, 0x0f, 0x96,
	0x93, 0xcd, 0x4d, 0x12, 0xf5, 0x6f, 0x63, 0x0a, 0x5f, 0x82, 0x91, 0xae, 0x89, 0xd0, 0xd3, 0xf1,
	0x7c, 0xb8, 0x8c, 0x56, 0xf7, 0x2f, 0x4e, 0xb1, 0xf3, 0x48, 0x7c, 0x62, 0x1d, 0x80, 0x5b, 0x4e,
	0x17, 0xdf, 0xc1, 0x83, 0xfe, 0xd8, 0xad, 0xdb, 0x96, 0x84, 0xe9, 0xa6, 0x34, 0xf0, 0x21, 0x08,
	0xfd, 0xf6, 0xb8, 0x89, 0x4d, 0x92, 0xb6, 0x82, 0xaf, 0x41, 0x40, 0x89, 0x21, 0xed, 0x38, 0x1e,
	0xfd, 0x35, 0x8e, 0x9d, 0xdb, 0x81, 0xc4, 0x85, 0xd6, 0x1f, 0xbf, 0x7d, 0x28, 0xb8, 0xf9, 0xd1,
	0xa4, 0x28, 0x93, 0x15, 0x76, 0xdf, 0x22, 0x55, 0xe1, 0x1f, 0x70, 0xbf, 0x7e, 0x05, 0x13, 0xb8,
	0x4e, 0xdf, 0x14, 0x12, 0x9f, 0x6d, 0x64, 0x1a, 0x3a, 0xe9, 0xdb, 0x3f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0xd7, 0xa7, 0xf1, 0xb0, 0xd5, 0x02, 0x00, 0x00,
}