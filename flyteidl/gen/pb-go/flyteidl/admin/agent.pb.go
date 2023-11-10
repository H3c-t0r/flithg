// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/admin/agent.proto

package admin

import (
	fmt "fmt"
	core "github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/core"
	proto "github.com/golang/protobuf/proto"
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

// The state of the execution is used to control its visibility in the UI/CLI.
type State int32

const (
	State_RETRYABLE_FAILURE State = 0
	State_PERMANENT_FAILURE State = 1
	State_PENDING           State = 2
	State_RUNNING           State = 3
	State_SUCCEEDED         State = 4
)

var State_name = map[int32]string{
	0: "RETRYABLE_FAILURE",
	1: "PERMANENT_FAILURE",
	2: "PENDING",
	3: "RUNNING",
	4: "SUCCEEDED",
}

var State_value = map[string]int32{
	"RETRYABLE_FAILURE": 0,
	"PERMANENT_FAILURE": 1,
	"PENDING":           2,
	"RUNNING":           3,
	"SUCCEEDED":         4,
}

func (x State) String() string {
	return proto.EnumName(State_name, int32(x))
}

func (State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c434e52bb0028071, []int{0}
}

// Represents a subset of runtime task execution metadata that are relevant to external plugins.
type TaskExecutionMetadata struct {
	// ID of the task execution
	TaskExecutionId *core.TaskExecutionIdentifier `protobuf:"bytes,1,opt,name=task_execution_id,json=taskExecutionId,proto3" json:"task_execution_id,omitempty"`
	// k8s namespace where the task is executed in
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// Labels attached to the task execution
	Labels map[string]string `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Annotations attached to the task execution
	Annotations map[string]string `protobuf:"bytes,4,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// k8s service account associated with the task execution
	K8SServiceAccount string `protobuf:"bytes,5,opt,name=k8s_service_account,json=k8sServiceAccount,proto3" json:"k8s_service_account,omitempty"`
	// Environment variables attached to the task execution
	EnvironmentVariables map[string]string `protobuf:"bytes,6,rep,name=environment_variables,json=environmentVariables,proto3" json:"environment_variables,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *TaskExecutionMetadata) Reset()         { *m = TaskExecutionMetadata{} }
func (m *TaskExecutionMetadata) String() string { return proto.CompactTextString(m) }
func (*TaskExecutionMetadata) ProtoMessage()    {}
func (*TaskExecutionMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_c434e52bb0028071, []int{0}
}

func (m *TaskExecutionMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskExecutionMetadata.Unmarshal(m, b)
}
func (m *TaskExecutionMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskExecutionMetadata.Marshal(b, m, deterministic)
}
func (m *TaskExecutionMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskExecutionMetadata.Merge(m, src)
}
func (m *TaskExecutionMetadata) XXX_Size() int {
	return xxx_messageInfo_TaskExecutionMetadata.Size(m)
}
func (m *TaskExecutionMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskExecutionMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_TaskExecutionMetadata proto.InternalMessageInfo

func (m *TaskExecutionMetadata) GetTaskExecutionId() *core.TaskExecutionIdentifier {
	if m != nil {
		return m.TaskExecutionId
	}
	return nil
}

func (m *TaskExecutionMetadata) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *TaskExecutionMetadata) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *TaskExecutionMetadata) GetAnnotations() map[string]string {
	if m != nil {
		return m.Annotations
	}
	return nil
}

func (m *TaskExecutionMetadata) GetK8SServiceAccount() string {
	if m != nil {
		return m.K8SServiceAccount
	}
	return ""
}

func (m *TaskExecutionMetadata) GetEnvironmentVariables() map[string]string {
	if m != nil {
		return m.EnvironmentVariables
	}
	return nil
}

// Represents a request structure to create task.
type CreateTaskRequest struct {
	// The inputs required to start the execution. All required inputs must be
	// included in this map. If not required and not provided, defaults apply.
	// +optional
	Inputs *core.LiteralMap `protobuf:"bytes,1,opt,name=inputs,proto3" json:"inputs,omitempty"`
	// Template of the task that encapsulates all the metadata of the task.
	Template *core.TaskTemplate `protobuf:"bytes,2,opt,name=template,proto3" json:"template,omitempty"`
	// Prefix for where task output data will be written. (e.g. s3://my-bucket/randomstring)
	OutputPrefix string `protobuf:"bytes,3,opt,name=output_prefix,json=outputPrefix,proto3" json:"output_prefix,omitempty"`
	// subset of runtime task execution metadata.
	TaskExecutionMetadata *TaskExecutionMetadata `protobuf:"bytes,4,opt,name=task_execution_metadata,json=taskExecutionMetadata,proto3" json:"task_execution_metadata,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}               `json:"-"`
	XXX_unrecognized      []byte                 `json:"-"`
	XXX_sizecache         int32                  `json:"-"`
}

func (m *CreateTaskRequest) Reset()         { *m = CreateTaskRequest{} }
func (m *CreateTaskRequest) String() string { return proto.CompactTextString(m) }
func (*CreateTaskRequest) ProtoMessage()    {}
func (*CreateTaskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c434e52bb0028071, []int{1}
}

func (m *CreateTaskRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTaskRequest.Unmarshal(m, b)
}
func (m *CreateTaskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTaskRequest.Marshal(b, m, deterministic)
}
func (m *CreateTaskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTaskRequest.Merge(m, src)
}
func (m *CreateTaskRequest) XXX_Size() int {
	return xxx_messageInfo_CreateTaskRequest.Size(m)
}
func (m *CreateTaskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTaskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTaskRequest proto.InternalMessageInfo

func (m *CreateTaskRequest) GetInputs() *core.LiteralMap {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *CreateTaskRequest) GetTemplate() *core.TaskTemplate {
	if m != nil {
		return m.Template
	}
	return nil
}

func (m *CreateTaskRequest) GetOutputPrefix() string {
	if m != nil {
		return m.OutputPrefix
	}
	return ""
}

func (m *CreateTaskRequest) GetTaskExecutionMetadata() *TaskExecutionMetadata {
	if m != nil {
		return m.TaskExecutionMetadata
	}
	return nil
}

// Represents a create response structure.
type CreateTaskResponse struct {
	// Metadata is created by the agent. It could be a string (jobId) or a dict (more complex metadata).
	ResourceMeta         []byte   `protobuf:"bytes,1,opt,name=resource_meta,json=resourceMeta,proto3" json:"resource_meta,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTaskResponse) Reset()         { *m = CreateTaskResponse{} }
func (m *CreateTaskResponse) String() string { return proto.CompactTextString(m) }
func (*CreateTaskResponse) ProtoMessage()    {}
func (*CreateTaskResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c434e52bb0028071, []int{2}
}

func (m *CreateTaskResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTaskResponse.Unmarshal(m, b)
}
func (m *CreateTaskResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTaskResponse.Marshal(b, m, deterministic)
}
func (m *CreateTaskResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTaskResponse.Merge(m, src)
}
func (m *CreateTaskResponse) XXX_Size() int {
	return xxx_messageInfo_CreateTaskResponse.Size(m)
}
func (m *CreateTaskResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTaskResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTaskResponse proto.InternalMessageInfo

func (m *CreateTaskResponse) GetResourceMeta() []byte {
	if m != nil {
		return m.ResourceMeta
	}
	return nil
}

// A message used to fetch a job resource from flyte agent server.
type GetTaskRequest struct {
	// A predefined yet extensible Task type identifier.
	TaskType string `protobuf:"bytes,1,opt,name=task_type,json=taskType,proto3" json:"task_type,omitempty"`
	// Metadata about the resource to be pass to the agent.
	ResourceMeta         []byte   `protobuf:"bytes,2,opt,name=resource_meta,json=resourceMeta,proto3" json:"resource_meta,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTaskRequest) Reset()         { *m = GetTaskRequest{} }
func (m *GetTaskRequest) String() string { return proto.CompactTextString(m) }
func (*GetTaskRequest) ProtoMessage()    {}
func (*GetTaskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c434e52bb0028071, []int{3}
}

func (m *GetTaskRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTaskRequest.Unmarshal(m, b)
}
func (m *GetTaskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTaskRequest.Marshal(b, m, deterministic)
}
func (m *GetTaskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTaskRequest.Merge(m, src)
}
func (m *GetTaskRequest) XXX_Size() int {
	return xxx_messageInfo_GetTaskRequest.Size(m)
}
func (m *GetTaskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTaskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTaskRequest proto.InternalMessageInfo

func (m *GetTaskRequest) GetTaskType() string {
	if m != nil {
		return m.TaskType
	}
	return ""
}

func (m *GetTaskRequest) GetResourceMeta() []byte {
	if m != nil {
		return m.ResourceMeta
	}
	return nil
}

// Response to get an individual task resource.
type GetTaskResponse struct {
	Resource             *Resource `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GetTaskResponse) Reset()         { *m = GetTaskResponse{} }
func (m *GetTaskResponse) String() string { return proto.CompactTextString(m) }
func (*GetTaskResponse) ProtoMessage()    {}
func (*GetTaskResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c434e52bb0028071, []int{4}
}

func (m *GetTaskResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTaskResponse.Unmarshal(m, b)
}
func (m *GetTaskResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTaskResponse.Marshal(b, m, deterministic)
}
func (m *GetTaskResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTaskResponse.Merge(m, src)
}
func (m *GetTaskResponse) XXX_Size() int {
	return xxx_messageInfo_GetTaskResponse.Size(m)
}
func (m *GetTaskResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTaskResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetTaskResponse proto.InternalMessageInfo

func (m *GetTaskResponse) GetResource() *Resource {
	if m != nil {
		return m.Resource
	}
	return nil
}

type Resource struct {
	// The state of the execution is used to control its visibility in the UI/CLI.
	State State `protobuf:"varint,1,opt,name=state,proto3,enum=flyteidl.admin.State" json:"state,omitempty"`
	// The outputs of the execution. It's typically used by sql task. Agent service will create a
	// Structured dataset pointing to the query result table.
	// +optional
	Outputs *core.LiteralMap `protobuf:"bytes,2,opt,name=outputs,proto3" json:"outputs,omitempty"`
	// A descriptive message for the current state. e.g. waiting for cluster.
	Message              string   `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Resource) Reset()         { *m = Resource{} }
func (m *Resource) String() string { return proto.CompactTextString(m) }
func (*Resource) ProtoMessage()    {}
func (*Resource) Descriptor() ([]byte, []int) {
	return fileDescriptor_c434e52bb0028071, []int{5}
}

func (m *Resource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resource.Unmarshal(m, b)
}
func (m *Resource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resource.Marshal(b, m, deterministic)
}
func (m *Resource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resource.Merge(m, src)
}
func (m *Resource) XXX_Size() int {
	return xxx_messageInfo_Resource.Size(m)
}
func (m *Resource) XXX_DiscardUnknown() {
	xxx_messageInfo_Resource.DiscardUnknown(m)
}

var xxx_messageInfo_Resource proto.InternalMessageInfo

func (m *Resource) GetState() State {
	if m != nil {
		return m.State
	}
	return State_RETRYABLE_FAILURE
}

func (m *Resource) GetOutputs() *core.LiteralMap {
	if m != nil {
		return m.Outputs
	}
	return nil
}

func (m *Resource) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// A message used to delete a task.
type DeleteTaskRequest struct {
	// A predefined yet extensible Task type identifier.
	TaskType string `protobuf:"bytes,1,opt,name=task_type,json=taskType,proto3" json:"task_type,omitempty"`
	// Metadata about the resource to be pass to the agent.
	ResourceMeta         []byte   `protobuf:"bytes,2,opt,name=resource_meta,json=resourceMeta,proto3" json:"resource_meta,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteTaskRequest) Reset()         { *m = DeleteTaskRequest{} }
func (m *DeleteTaskRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteTaskRequest) ProtoMessage()    {}
func (*DeleteTaskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c434e52bb0028071, []int{6}
}

func (m *DeleteTaskRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTaskRequest.Unmarshal(m, b)
}
func (m *DeleteTaskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTaskRequest.Marshal(b, m, deterministic)
}
func (m *DeleteTaskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTaskRequest.Merge(m, src)
}
func (m *DeleteTaskRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteTaskRequest.Size(m)
}
func (m *DeleteTaskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTaskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTaskRequest proto.InternalMessageInfo

func (m *DeleteTaskRequest) GetTaskType() string {
	if m != nil {
		return m.TaskType
	}
	return ""
}

func (m *DeleteTaskRequest) GetResourceMeta() []byte {
	if m != nil {
		return m.ResourceMeta
	}
	return nil
}

// Response to delete a task.
type DeleteTaskResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteTaskResponse) Reset()         { *m = DeleteTaskResponse{} }
func (m *DeleteTaskResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteTaskResponse) ProtoMessage()    {}
func (*DeleteTaskResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c434e52bb0028071, []int{7}
}

func (m *DeleteTaskResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTaskResponse.Unmarshal(m, b)
}
func (m *DeleteTaskResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTaskResponse.Marshal(b, m, deterministic)
}
func (m *DeleteTaskResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTaskResponse.Merge(m, src)
}
func (m *DeleteTaskResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteTaskResponse.Size(m)
}
func (m *DeleteTaskResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTaskResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTaskResponse proto.InternalMessageInfo

// A message used to do a task
type DoTaskRequest struct {
	// The inputs required to start the execution. All required inputs must be
	// included in this map. If not required and not provided, defaults apply.
	// +optional
	Inputs *core.LiteralMap `protobuf:"bytes,1,opt,name=inputs,proto3" json:"inputs,omitempty"`
	// Template of the task that encapsulates all the metadata of the task.
	Template *core.TaskTemplate `protobuf:"bytes,2,opt,name=template,proto3" json:"template,omitempty"`
	// Prefix for where task output data will be written. (e.g. s3://my-bucket/randomstring)
	OutputPrefix string `protobuf:"bytes,3,opt,name=output_prefix,json=outputPrefix,proto3" json:"output_prefix,omitempty"`
	// subset of runtime task execution metadata.
	TaskExecutionMetadata *TaskExecutionMetadata `protobuf:"bytes,4,opt,name=task_execution_metadata,json=taskExecutionMetadata,proto3" json:"task_execution_metadata,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}               `json:"-"`
	XXX_unrecognized      []byte                 `json:"-"`
	XXX_sizecache         int32                  `json:"-"`
}

func (m *DoTaskRequest) Reset()         { *m = DoTaskRequest{} }
func (m *DoTaskRequest) String() string { return proto.CompactTextString(m) }
func (*DoTaskRequest) ProtoMessage()    {}
func (*DoTaskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c434e52bb0028071, []int{8}
}

func (m *DoTaskRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DoTaskRequest.Unmarshal(m, b)
}
func (m *DoTaskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DoTaskRequest.Marshal(b, m, deterministic)
}
func (m *DoTaskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoTaskRequest.Merge(m, src)
}
func (m *DoTaskRequest) XXX_Size() int {
	return xxx_messageInfo_DoTaskRequest.Size(m)
}
func (m *DoTaskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DoTaskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DoTaskRequest proto.InternalMessageInfo

func (m *DoTaskRequest) GetInputs() *core.LiteralMap {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *DoTaskRequest) GetTemplate() *core.TaskTemplate {
	if m != nil {
		return m.Template
	}
	return nil
}

func (m *DoTaskRequest) GetOutputPrefix() string {
	if m != nil {
		return m.OutputPrefix
	}
	return ""
}

func (m *DoTaskRequest) GetTaskExecutionMetadata() *TaskExecutionMetadata {
	if m != nil {
		return m.TaskExecutionMetadata
	}
	return nil
}

type DoTaskResponse struct {
	Resource             *Resource `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *DoTaskResponse) Reset()         { *m = DoTaskResponse{} }
func (m *DoTaskResponse) String() string { return proto.CompactTextString(m) }
func (*DoTaskResponse) ProtoMessage()    {}
func (*DoTaskResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c434e52bb0028071, []int{9}
}

func (m *DoTaskResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DoTaskResponse.Unmarshal(m, b)
}
func (m *DoTaskResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DoTaskResponse.Marshal(b, m, deterministic)
}
func (m *DoTaskResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoTaskResponse.Merge(m, src)
}
func (m *DoTaskResponse) XXX_Size() int {
	return xxx_messageInfo_DoTaskResponse.Size(m)
}
func (m *DoTaskResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DoTaskResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DoTaskResponse proto.InternalMessageInfo

func (m *DoTaskResponse) GetResource() *Resource {
	if m != nil {
		return m.Resource
	}
	return nil
}

func init() {
	proto.RegisterEnum("flyteidl.admin.State", State_name, State_value)
	proto.RegisterType((*TaskExecutionMetadata)(nil), "flyteidl.admin.TaskExecutionMetadata")
	proto.RegisterMapType((map[string]string)(nil), "flyteidl.admin.TaskExecutionMetadata.AnnotationsEntry")
	proto.RegisterMapType((map[string]string)(nil), "flyteidl.admin.TaskExecutionMetadata.EnvironmentVariablesEntry")
	proto.RegisterMapType((map[string]string)(nil), "flyteidl.admin.TaskExecutionMetadata.LabelsEntry")
	proto.RegisterType((*CreateTaskRequest)(nil), "flyteidl.admin.CreateTaskRequest")
	proto.RegisterType((*CreateTaskResponse)(nil), "flyteidl.admin.CreateTaskResponse")
	proto.RegisterType((*GetTaskRequest)(nil), "flyteidl.admin.GetTaskRequest")
	proto.RegisterType((*GetTaskResponse)(nil), "flyteidl.admin.GetTaskResponse")
	proto.RegisterType((*Resource)(nil), "flyteidl.admin.Resource")
	proto.RegisterType((*DeleteTaskRequest)(nil), "flyteidl.admin.DeleteTaskRequest")
	proto.RegisterType((*DeleteTaskResponse)(nil), "flyteidl.admin.DeleteTaskResponse")
	proto.RegisterType((*DoTaskRequest)(nil), "flyteidl.admin.DoTaskRequest")
	proto.RegisterType((*DoTaskResponse)(nil), "flyteidl.admin.DoTaskResponse")
}

func init() { proto.RegisterFile("flyteidl/admin/agent.proto", fileDescriptor_c434e52bb0028071) }

var fileDescriptor_c434e52bb0028071 = []byte{
	// 757 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe4, 0x55, 0xdf, 0x6f, 0x32, 0x45,
	0x14, 0x15, 0x28, 0x14, 0x2e, 0x2d, 0x1f, 0x8c, 0x25, 0x6e, 0xf9, 0xaa, 0x69, 0x30, 0x9a, 0x46,
	0xe3, 0x92, 0x52, 0xa3, 0xad, 0x46, 0x0d, 0x2d, 0x5b, 0x42, 0x42, 0x49, 0x33, 0x05, 0xa3, 0x26,
	0x4a, 0x86, 0xe5, 0x82, 0x1b, 0x96, 0xd9, 0x75, 0x67, 0x96, 0x94, 0x67, 0x13, 0xff, 0x06, 0xff,
	0x5b, 0xcd, 0xce, 0xfe, 0x28, 0x10, 0x34, 0x6d, 0xfa, 0xf8, 0xbd, 0xed, 0xdc, 0x73, 0xef, 0x99,
	0xb3, 0xe7, 0xdc, 0xcd, 0x42, 0x6d, 0x6a, 0xaf, 0x24, 0x5a, 0x13, 0xbb, 0xc1, 0x26, 0x0b, 0x8b,
	0x37, 0xd8, 0x0c, 0xb9, 0xd4, 0x5d, 0xcf, 0x91, 0x0e, 0x29, 0xc5, 0x98, 0xae, 0xb0, 0xda, 0x49,
	0xd2, 0x6b, 0x3a, 0x1e, 0x36, 0x6c, 0x4b, 0xa2, 0xc7, 0x6c, 0x11, 0x76, 0xd7, 0x8e, 0x37, 0x51,
	0xc9, 0xc4, 0x3c, 0x86, 0x3e, 0xdc, 0x84, 0x2c, 0x2e, 0xd1, 0x9b, 0x32, 0x13, 0x23, 0xf8, 0xa3,
	0x2d, 0x78, 0x82, 0x5c, 0x5a, 0x53, 0x0b, 0xbd, 0x10, 0xaf, 0xff, 0x9d, 0x85, 0xea, 0x80, 0x89,
	0xb9, 0xf1, 0x88, 0xa6, 0x2f, 0x2d, 0x87, 0xdf, 0xa1, 0x64, 0x13, 0x26, 0x19, 0xa1, 0x50, 0x09,
	0xee, 0x19, 0x61, 0x8c, 0x8c, 0xac, 0x89, 0x96, 0x3a, 0x4d, 0x9d, 0x15, 0x9b, 0x9f, 0xea, 0x89,
	0xfa, 0x80, 0x55, 0xdf, 0x20, 0xe8, 0x26, 0x57, 0xd0, 0x37, 0x72, 0x13, 0x20, 0x27, 0x50, 0xe0,
	0x6c, 0x81, 0xc2, 0x65, 0x26, 0x6a, 0xe9, 0xd3, 0xd4, 0x59, 0x81, 0x3e, 0x15, 0x48, 0x17, 0x72,
	0x36, 0x1b, 0xa3, 0x2d, 0xb4, 0xcc, 0x69, 0xe6, 0xac, 0xd8, 0x3c, 0xd7, 0x37, 0x4d, 0xd2, 0x77,
	0x0a, 0xd5, 0x7b, 0x6a, 0xc6, 0xe0, 0xd2, 0x5b, 0xd1, 0x88, 0x80, 0xfc, 0x04, 0x45, 0xc6, 0xb9,
	0x23, 0x59, 0xd0, 0x29, 0xb4, 0x3d, 0xc5, 0xf7, 0xd5, 0xf3, 0xf8, 0x5a, 0x4f, 0x83, 0x21, 0xe9,
	0x3a, 0x15, 0xd1, 0xe1, 0xfd, 0xf9, 0xa5, 0x18, 0x09, 0xf4, 0x96, 0x96, 0x89, 0x23, 0x66, 0x9a,
	0x8e, 0xcf, 0xa5, 0x96, 0x55, 0x2f, 0x53, 0x99, 0x5f, 0x8a, 0x87, 0x10, 0x69, 0x85, 0x00, 0x91,
	0x50, 0x45, 0xbe, 0xb4, 0x3c, 0x87, 0x2f, 0x90, 0xcb, 0xd1, 0x92, 0x79, 0x16, 0x1b, 0xdb, 0x28,
	0xb4, 0x9c, 0xd2, 0xf4, 0xc3, 0xf3, 0x34, 0x19, 0x4f, 0x14, 0x3f, 0xc6, 0x0c, 0xa1, 0xb8, 0x23,
	0xdc, 0x01, 0xd5, 0xae, 0xa0, 0xb8, 0x66, 0x0b, 0x29, 0x43, 0x66, 0x8e, 0x2b, 0x95, 0x5e, 0x81,
	0x06, 0x8f, 0xe4, 0x08, 0xb2, 0x4b, 0x66, 0xfb, 0x71, 0x0a, 0xe1, 0xe1, 0x9b, 0xf4, 0x65, 0xaa,
	0xf6, 0x3d, 0x94, 0xb7, 0x1d, 0x78, 0xd1, 0x7c, 0x07, 0x8e, 0xff, 0x53, 0xed, 0x4b, 0x88, 0xea,
	0x7f, 0xa6, 0xa1, 0x72, 0xe3, 0x21, 0x93, 0x18, 0x78, 0x42, 0xf1, 0x0f, 0x1f, 0x85, 0x24, 0xe7,
	0x90, 0xb3, 0xb8, 0xeb, 0x4b, 0x11, 0xed, 0xe2, 0xf1, 0xd6, 0x2e, 0xf6, 0xc2, 0x2f, 0xe7, 0x8e,
	0xb9, 0x34, 0x6a, 0x24, 0x5f, 0x43, 0x5e, 0xe2, 0xc2, 0xb5, 0x99, 0x0c, 0x6f, 0x29, 0x36, 0xdf,
	0xee, 0x58, 0xe0, 0x41, 0xd4, 0x42, 0x93, 0x66, 0xf2, 0x31, 0x1c, 0x3a, 0xbe, 0x74, 0x7d, 0x39,
	0x72, 0x3d, 0x9c, 0x5a, 0x8f, 0x5a, 0x46, 0x69, 0x3c, 0x08, 0x8b, 0xf7, 0xaa, 0x46, 0x7e, 0x85,
	0x0f, 0xb6, 0xbe, 0x93, 0x45, 0x94, 0x9a, 0xb6, 0xa7, 0x2e, 0xfb, 0xe4, 0x59, 0x11, 0xd3, 0xaa,
	0xdc, 0x55, 0xae, 0x5f, 0x01, 0x59, 0x37, 0x41, 0xb8, 0x0e, 0x17, 0x4a, 0x99, 0x87, 0xc2, 0xf1,
	0x3d, 0x13, 0xd5, 0x75, 0xca, 0x8c, 0x03, 0x7a, 0x10, 0x17, 0x83, 0xf1, 0x3a, 0x85, 0x52, 0x07,
	0xe5, 0xba, 0x79, 0x6f, 0xa1, 0xa0, 0xb4, 0xca, 0x95, 0x8b, 0x51, 0x08, 0xf9, 0xa0, 0x30, 0x58,
	0xb9, 0x3b, 0x38, 0xd3, 0x3b, 0x38, 0x3b, 0xf0, 0x26, 0xe1, 0x8c, 0xb4, 0x7c, 0x09, 0xf9, 0xb8,
	0x25, 0xca, 0x44, 0xdb, 0x7e, 0x63, 0x1a, 0xe1, 0x34, 0xe9, 0xac, 0xff, 0x95, 0x82, 0x7c, 0x5c,
	0x26, 0x9f, 0x43, 0x56, 0xc8, 0x20, 0x9e, 0x60, 0xbe, 0xd4, 0xac, 0x6e, 0xcf, 0x3f, 0x04, 0x20,
	0x0d, 0x7b, 0xc8, 0x05, 0xec, 0x87, 0x01, 0x88, 0x28, 0xcd, 0xff, 0x59, 0x81, 0xb8, 0x93, 0x68,
	0xb0, 0xbf, 0x40, 0x21, 0xd8, 0x0c, 0xa3, 0x10, 0xe3, 0x63, 0x7d, 0x08, 0x95, 0x36, 0xda, 0xb8,
	0xb9, 0x65, 0xaf, 0x37, 0xea, 0x08, 0xc8, 0x3a, 0x6d, 0xe8, 0x55, 0xfd, 0x9f, 0x14, 0x1c, 0xb6,
	0x9d, 0x77, 0x79, 0x9f, 0x6f, 0xa1, 0x14, 0x1b, 0xf0, 0x9a, 0xfd, 0xf9, 0xec, 0x37, 0xc8, 0xaa,
	0xad, 0x20, 0x55, 0xa8, 0x50, 0x63, 0x40, 0x7f, 0x6e, 0x5d, 0xf7, 0x8c, 0xd1, 0x6d, 0xab, 0xdb,
	0x1b, 0x52, 0xa3, 0xfc, 0x5e, 0x50, 0xbe, 0x37, 0xe8, 0x5d, 0xab, 0x6f, 0xf4, 0x07, 0x49, 0x39,
	0x45, 0x8a, 0xb0, 0x7f, 0x6f, 0xf4, 0xdb, 0xdd, 0x7e, 0xa7, 0x9c, 0x0e, 0x0e, 0x74, 0xd8, 0xef,
	0x07, 0x87, 0x0c, 0x39, 0x84, 0xc2, 0xc3, 0xf0, 0xe6, 0xc6, 0x30, 0xda, 0x46, 0xbb, 0xbc, 0x77,
	0xfd, 0xdd, 0x2f, 0xdf, 0xce, 0x2c, 0xf9, 0xbb, 0x3f, 0xd6, 0x4d, 0x67, 0xd1, 0x50, 0x7a, 0x1c,
	0x6f, 0x16, 0x3e, 0x34, 0x92, 0x9f, 0xea, 0x0c, 0x79, 0xc3, 0x1d, 0x7f, 0x31, 0x73, 0x1a, 0x9b,
	0xff, 0xfa, 0x71, 0x4e, 0xfd, 0x5e, 0x2f, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x15, 0xe0, 0x43,
	0xe4, 0x04, 0x08, 0x00, 0x00,
}
