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
	// Secret to be passed to the agent.
	// Key is the name of the secret in secret manager.
	// Value is the actual secret value.
	Secrets              map[string]string `protobuf:"bytes,5,rep,name=secrets,proto3" json:"secrets,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
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

func (m *CreateTaskRequest) GetSecrets() map[string]string {
	if m != nil {
		return m.Secrets
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
	Resource *Resource `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	// log information for the task execution
	LogLinks             []*core.TaskLog `protobuf:"bytes,2,rep,name=log_links,json=logLinks,proto3" json:"log_links,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
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

func (m *GetTaskResponse) GetLogLinks() []*core.TaskLog {
	if m != nil {
		return m.LogLinks
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

// A message containing the agent metadata.
type Agent struct {
	// Name is the developer-assigned name of the agent.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// SupportedTaskTypes are the types of the tasks that the agent can handle.
	SupportedTaskType string `protobuf:"bytes,2,opt,name=supported_task_type,json=supportedTaskType,proto3" json:"supported_task_type,omitempty"`
	// IsSync indicates whether this agent is a sync agent. Sync agents are expected to return their results synchronously when called by propeller. Given that sync agents can affect the performance of the system, it's important to enforce strict timeout policies. An Async agent, on the other hand, is required to be able to identify jobs by an identifier and query for job statuses as jobs progress.
	IsSync bool `protobuf:"varint,3,opt,name=is_sync,json=isSync,proto3" json:"is_sync,omitempty"`
	// SecretNames is a list of secrets the agent requires to execute tasks.
	SecretNames          []string `protobuf:"bytes,4,rep,name=secret_names,json=secretNames,proto3" json:"secret_names,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Agent) Reset()         { *m = Agent{} }
func (m *Agent) String() string { return proto.CompactTextString(m) }
func (*Agent) ProtoMessage()    {}
func (*Agent) Descriptor() ([]byte, []int) {
	return fileDescriptor_c434e52bb0028071, []int{8}
}

func (m *Agent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Agent.Unmarshal(m, b)
}
func (m *Agent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Agent.Marshal(b, m, deterministic)
}
func (m *Agent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Agent.Merge(m, src)
}
func (m *Agent) XXX_Size() int {
	return xxx_messageInfo_Agent.Size(m)
}
func (m *Agent) XXX_DiscardUnknown() {
	xxx_messageInfo_Agent.DiscardUnknown(m)
}

var xxx_messageInfo_Agent proto.InternalMessageInfo

func (m *Agent) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Agent) GetSupportedTaskType() string {
	if m != nil {
		return m.SupportedTaskType
	}
	return ""
}

func (m *Agent) GetIsSync() bool {
	if m != nil {
		return m.IsSync
	}
	return false
}

func (m *Agent) GetSecretNames() []string {
	if m != nil {
		return m.SecretNames
	}
	return nil
}

// A request to get an agent.
type GetAgentRequest struct {
	// The name of the agent.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAgentRequest) Reset()         { *m = GetAgentRequest{} }
func (m *GetAgentRequest) String() string { return proto.CompactTextString(m) }
func (*GetAgentRequest) ProtoMessage()    {}
func (*GetAgentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c434e52bb0028071, []int{9}
}

func (m *GetAgentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAgentRequest.Unmarshal(m, b)
}
func (m *GetAgentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAgentRequest.Marshal(b, m, deterministic)
}
func (m *GetAgentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAgentRequest.Merge(m, src)
}
func (m *GetAgentRequest) XXX_Size() int {
	return xxx_messageInfo_GetAgentRequest.Size(m)
}
func (m *GetAgentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAgentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAgentRequest proto.InternalMessageInfo

func (m *GetAgentRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// A response containing an agent.
type GetAgentResponse struct {
	Agent                *Agent   `protobuf:"bytes,1,opt,name=agent,proto3" json:"agent,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAgentResponse) Reset()         { *m = GetAgentResponse{} }
func (m *GetAgentResponse) String() string { return proto.CompactTextString(m) }
func (*GetAgentResponse) ProtoMessage()    {}
func (*GetAgentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c434e52bb0028071, []int{10}
}

func (m *GetAgentResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAgentResponse.Unmarshal(m, b)
}
func (m *GetAgentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAgentResponse.Marshal(b, m, deterministic)
}
func (m *GetAgentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAgentResponse.Merge(m, src)
}
func (m *GetAgentResponse) XXX_Size() int {
	return xxx_messageInfo_GetAgentResponse.Size(m)
}
func (m *GetAgentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAgentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAgentResponse proto.InternalMessageInfo

func (m *GetAgentResponse) GetAgent() *Agent {
	if m != nil {
		return m.Agent
	}
	return nil
}

// A request to list all agents.
type ListAgentsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListAgentsRequest) Reset()         { *m = ListAgentsRequest{} }
func (m *ListAgentsRequest) String() string { return proto.CompactTextString(m) }
func (*ListAgentsRequest) ProtoMessage()    {}
func (*ListAgentsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c434e52bb0028071, []int{11}
}

func (m *ListAgentsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListAgentsRequest.Unmarshal(m, b)
}
func (m *ListAgentsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListAgentsRequest.Marshal(b, m, deterministic)
}
func (m *ListAgentsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAgentsRequest.Merge(m, src)
}
func (m *ListAgentsRequest) XXX_Size() int {
	return xxx_messageInfo_ListAgentsRequest.Size(m)
}
func (m *ListAgentsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAgentsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListAgentsRequest proto.InternalMessageInfo

// A response containing a list of agents.
type ListAgentsResponse struct {
	Agents               []*Agent `protobuf:"bytes,1,rep,name=agents,proto3" json:"agents,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListAgentsResponse) Reset()         { *m = ListAgentsResponse{} }
func (m *ListAgentsResponse) String() string { return proto.CompactTextString(m) }
func (*ListAgentsResponse) ProtoMessage()    {}
func (*ListAgentsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c434e52bb0028071, []int{12}
}

func (m *ListAgentsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListAgentsResponse.Unmarshal(m, b)
}
func (m *ListAgentsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListAgentsResponse.Marshal(b, m, deterministic)
}
func (m *ListAgentsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAgentsResponse.Merge(m, src)
}
func (m *ListAgentsResponse) XXX_Size() int {
	return xxx_messageInfo_ListAgentsResponse.Size(m)
}
func (m *ListAgentsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAgentsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListAgentsResponse proto.InternalMessageInfo

func (m *ListAgentsResponse) GetAgents() []*Agent {
	if m != nil {
		return m.Agents
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
	proto.RegisterMapType((map[string]string)(nil), "flyteidl.admin.CreateTaskRequest.SecretsEntry")
	proto.RegisterType((*CreateTaskResponse)(nil), "flyteidl.admin.CreateTaskResponse")
	proto.RegisterType((*GetTaskRequest)(nil), "flyteidl.admin.GetTaskRequest")
	proto.RegisterType((*GetTaskResponse)(nil), "flyteidl.admin.GetTaskResponse")
	proto.RegisterType((*Resource)(nil), "flyteidl.admin.Resource")
	proto.RegisterType((*DeleteTaskRequest)(nil), "flyteidl.admin.DeleteTaskRequest")
	proto.RegisterType((*DeleteTaskResponse)(nil), "flyteidl.admin.DeleteTaskResponse")
	proto.RegisterType((*Agent)(nil), "flyteidl.admin.Agent")
	proto.RegisterType((*GetAgentRequest)(nil), "flyteidl.admin.GetAgentRequest")
	proto.RegisterType((*GetAgentResponse)(nil), "flyteidl.admin.GetAgentResponse")
	proto.RegisterType((*ListAgentsRequest)(nil), "flyteidl.admin.ListAgentsRequest")
	proto.RegisterType((*ListAgentsResponse)(nil), "flyteidl.admin.ListAgentsResponse")
}

func init() { proto.RegisterFile("flyteidl/admin/agent.proto", fileDescriptor_c434e52bb0028071) }

var fileDescriptor_c434e52bb0028071 = []byte{
	// 937 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0x7f, 0x6f, 0xdb, 0x54,
	0x14, 0x25, 0x4d, 0x93, 0x26, 0x37, 0xd9, 0x96, 0xbc, 0x35, 0xcc, 0xcd, 0x06, 0x2a, 0x46, 0x43,
	0x15, 0xd3, 0x1c, 0xad, 0x45, 0xd0, 0x0d, 0xc1, 0x94, 0xb5, 0xa6, 0x54, 0x4a, 0xa3, 0xea, 0x25,
	0x45, 0x80, 0x04, 0xd6, 0x8b, 0x73, 0x1b, 0xac, 0x38, 0xb6, 0xf1, 0x7b, 0xae, 0x16, 0x89, 0xbf,
	0x41, 0x7c, 0x03, 0x3e, 0x2e, 0x7a, 0xef, 0xd9, 0xce, 0x8f, 0x65, 0xa8, 0x15, 0xff, 0xd9, 0xf7,
	0x9c, 0x7b, 0xee, 0x7d, 0xf7, 0x1e, 0xff, 0x80, 0xf6, 0xb5, 0x3f, 0x17, 0xe8, 0x8d, 0xfd, 0x0e,
	0x1b, 0xcf, 0xbc, 0xa0, 0xc3, 0x26, 0x18, 0x08, 0x2b, 0x8a, 0x43, 0x11, 0x92, 0xfb, 0x19, 0x66,
	0x29, 0xac, 0xfd, 0x24, 0xe7, 0xba, 0x61, 0x8c, 0x1d, 0xdf, 0x13, 0x18, 0x33, 0x9f, 0x6b, 0x76,
	0x7b, 0x6f, 0x15, 0x15, 0x8c, 0x4f, 0x33, 0xe8, 0xa3, 0x55, 0xc8, 0x0b, 0x04, 0xc6, 0xd7, 0xcc,
	0xc5, 0x14, 0xfe, 0x78, 0x0d, 0x1e, 0x63, 0x20, 0xbc, 0x6b, 0x0f, 0xe3, 0xcd, 0xe9, 0xf8, 0x16,
	0xdd, 0x44, 0x78, 0x61, 0xa0, 0x61, 0xf3, 0x9f, 0x12, 0xb4, 0x86, 0x8c, 0x4f, 0xed, 0x2c, 0x7e,
	0x81, 0x82, 0x8d, 0x99, 0x60, 0x84, 0x42, 0x53, 0xb6, 0xe1, 0xe4, 0x19, 0x8e, 0x37, 0x36, 0x0a,
	0xfb, 0x85, 0x83, 0xda, 0xe1, 0x67, 0x56, 0x7e, 0x38, 0x29, 0x6a, 0xad, 0x08, 0x9c, 0xe7, 0x1d,
	0xd0, 0x07, 0x62, 0x15, 0x20, 0x4f, 0xa0, 0x1a, 0xb0, 0x19, 0xf2, 0x88, 0xb9, 0x68, 0x6c, 0xed,
	0x17, 0x0e, 0xaa, 0x74, 0x11, 0x20, 0xe7, 0x50, 0xf6, 0xd9, 0x08, 0x7d, 0x6e, 0x14, 0xf7, 0x8b,
	0x07, 0xb5, 0xc3, 0x17, 0xd6, 0xea, 0x0c, 0xad, 0x8d, 0x8d, 0x5a, 0x3d, 0x95, 0x63, 0x07, 0x22,
	0x9e, 0xd3, 0x54, 0x80, 0xfc, 0x08, 0x35, 0x16, 0x04, 0xa1, 0x60, 0x92, 0xc9, 0x8d, 0x6d, 0xa5,
	0xf7, 0xe5, 0xed, 0xf4, 0xba, 0x8b, 0x44, 0x2d, 0xba, 0x2c, 0x45, 0x2c, 0x78, 0x38, 0x3d, 0xe6,
	0x0e, 0xc7, 0xf8, 0xc6, 0x73, 0xd1, 0x61, 0xae, 0x1b, 0x26, 0x81, 0x30, 0x4a, 0xea, 0x30, 0xcd,
	0xe9, 0x31, 0x1f, 0x68, 0xa4, 0xab, 0x01, 0x22, 0xa0, 0x85, 0xc1, 0x8d, 0x17, 0x87, 0xc1, 0x0c,
	0x03, 0xe1, 0xdc, 0xb0, 0xd8, 0x63, 0x23, 0x1f, 0xb9, 0x51, 0x56, 0x3d, 0xbd, 0xbe, 0x5d, 0x4f,
	0xf6, 0x42, 0xe2, 0x87, 0x4c, 0x41, 0x37, 0xb7, 0x8b, 0x1b, 0xa0, 0xf6, 0x4b, 0xa8, 0x2d, 0x8d,
	0x85, 0x34, 0xa0, 0x38, 0xc5, 0xb9, 0xda, 0x5e, 0x95, 0xca, 0x4b, 0xb2, 0x0b, 0xa5, 0x1b, 0xe6,
	0x27, 0xd9, 0x16, 0xf4, 0xcd, 0xab, 0xad, 0xe3, 0x42, 0xfb, 0x5b, 0x68, 0xac, 0x4f, 0xe0, 0x4e,
	0xf9, 0x67, 0xb0, 0xf7, 0xde, 0x6e, 0xef, 0x22, 0x64, 0xfe, 0x5d, 0x84, 0xe6, 0x49, 0x8c, 0x4c,
	0xa0, 0x9c, 0x09, 0xc5, 0xdf, 0x13, 0xe4, 0x82, 0xbc, 0x80, 0xb2, 0x17, 0x44, 0x89, 0xe0, 0xa9,
	0x17, 0xf7, 0xd6, 0xbc, 0xd8, 0xd3, 0x0f, 0xd6, 0x05, 0x8b, 0x68, 0x4a, 0x24, 0x5f, 0x41, 0x45,
	0xe0, 0x2c, 0xf2, 0x99, 0xd0, 0x55, 0x6a, 0x87, 0x8f, 0x37, 0x18, 0x78, 0x98, 0x52, 0x68, 0x4e,
	0x26, 0x9f, 0xc2, 0xbd, 0x30, 0x11, 0x51, 0x22, 0x9c, 0x28, 0xc6, 0x6b, 0xef, 0xad, 0x51, 0x54,
	0x3d, 0xd6, 0x75, 0xf0, 0x52, 0xc5, 0xc8, 0x2f, 0xf0, 0x68, 0xed, 0x39, 0x99, 0xa5, 0x5b, 0x33,
	0xb6, 0x55, 0xb1, 0xa7, 0xb7, 0x5a, 0x31, 0x6d, 0x89, 0x8d, 0x8f, 0xe1, 0xf7, 0xb0, 0xc3, 0xd1,
	0x8d, 0x51, 0x70, 0xa3, 0xa4, 0x1c, 0x63, 0xad, 0xcb, 0xbd, 0x33, 0x23, 0x6b, 0xa0, 0x13, 0xb4,
	0x41, 0xb2, 0xf4, 0xf6, 0x2b, 0xa8, 0x2f, 0x03, 0x77, 0xda, 0xc5, 0x4b, 0x20, 0xcb, 0x65, 0x78,
	0x14, 0x06, 0x5c, 0xcd, 0x27, 0x46, 0x1e, 0x26, 0xb1, 0x8b, 0xea, 0xd0, 0x4a, 0xab, 0x4e, 0xeb,
	0x59, 0x50, 0x1e, 0xc2, 0xa4, 0x70, 0xff, 0x0c, 0xc5, 0xf2, 0x0a, 0x1f, 0x43, 0x55, 0x4d, 0x4c,
	0xcc, 0x23, 0x4c, 0xcb, 0x57, 0x64, 0x60, 0x38, 0x8f, 0x36, 0x68, 0x6e, 0x6d, 0xd0, 0xfc, 0x03,
	0x1e, 0xe4, 0x9a, 0x69, 0x2f, 0x5f, 0x40, 0x25, 0xa3, 0xa4, 0xce, 0x30, 0xd6, 0x07, 0x45, 0x53,
	0x9c, 0xe6, 0x4c, 0x72, 0x04, 0x55, 0x3f, 0x9c, 0x38, 0xbe, 0x17, 0x4c, 0xb9, 0xb1, 0xa5, 0xe6,
	0xfb, 0xe1, 0x06, 0x6f, 0xf4, 0xc2, 0x09, 0xad, 0xf8, 0xe1, 0xa4, 0x27, 0x79, 0xe6, 0x9f, 0x05,
	0xa8, 0x64, 0x5a, 0xe4, 0x19, 0x94, 0xb8, 0x90, 0xce, 0x92, 0x45, 0xef, 0x1f, 0xb6, 0xd6, 0x8b,
	0x0e, 0x24, 0x48, 0x35, 0x87, 0x1c, 0xc1, 0x8e, 0xf6, 0x0e, 0x4f, 0x8d, 0xf8, 0x1f, 0xee, 0xcd,
	0x98, 0xc4, 0x80, 0x9d, 0x19, 0x72, 0xce, 0x26, 0x98, 0xfa, 0x2f, 0xbb, 0x35, 0xaf, 0xa0, 0x79,
	0x8a, 0x3e, 0xae, 0x3e, 0x20, 0xff, 0x7f, 0xba, 0xbb, 0x40, 0x96, 0x65, 0xf5, 0x80, 0xcd, 0xbf,
	0x0a, 0x50, 0xea, 0xca, 0x0f, 0x1c, 0x21, 0xb0, 0x2d, 0x5f, 0xda, 0xa9, 0xb8, 0xba, 0x96, 0xaf,
	0x45, 0x9e, 0x44, 0x51, 0x18, 0x0b, 0x1c, 0x3b, 0x8b, 0xfa, 0xda, 0x48, 0xcd, 0x1c, 0x1a, 0x66,
	0x8d, 0x3c, 0x82, 0x1d, 0x8f, 0x3b, 0x7c, 0x1e, 0xb8, 0xea, 0x50, 0x15, 0x5a, 0xf6, 0xf8, 0x60,
	0x1e, 0xb8, 0xe4, 0x13, 0xa8, 0x6b, 0xc3, 0x3a, 0xea, 0xc3, 0xa0, 0x5e, 0xdd, 0x55, 0x5a, 0xd3,
	0xb1, 0xbe, 0x0c, 0x99, 0x4f, 0xd5, 0xf6, 0x55, 0x2f, 0xd9, 0xa1, 0x37, 0xb4, 0x64, 0xbe, 0x86,
	0xc6, 0x82, 0x96, 0xba, 0xe4, 0x19, 0x94, 0xd4, 0x47, 0x3a, 0xb5, 0xc8, 0x3b, 0xdb, 0xd2, 0x6c,
	0xcd, 0x31, 0x1f, 0x42, 0xb3, 0xe7, 0x71, 0xad, 0xc0, 0xd3, 0x4a, 0xe6, 0x09, 0x90, 0xe5, 0x60,
	0xaa, 0xfb, 0x1c, 0xca, 0x2a, 0x47, 0xbe, 0x95, 0x8a, 0xef, 0x17, 0x4e, 0x49, 0x9f, 0xff, 0x0a,
	0x25, 0xe5, 0x0b, 0xd2, 0x82, 0x26, 0xb5, 0x87, 0xf4, 0xa7, 0xee, 0x9b, 0x9e, 0xed, 0x7c, 0xd7,
	0x3d, 0xef, 0x5d, 0x51, 0xbb, 0xf1, 0x81, 0x0c, 0x5f, 0xda, 0xf4, 0xa2, 0xdb, 0xb7, 0xfb, 0xc3,
	0x3c, 0x5c, 0x20, 0x35, 0xd8, 0xb9, 0xb4, 0xfb, 0xa7, 0xe7, 0xfd, 0xb3, 0xc6, 0x96, 0xbc, 0xa1,
	0x57, 0xfd, 0xbe, 0xbc, 0x29, 0x92, 0x7b, 0x50, 0x1d, 0x5c, 0x9d, 0x9c, 0xd8, 0xf6, 0xa9, 0x7d,
	0xda, 0xd8, 0x7e, 0xf3, 0xcd, 0xcf, 0x5f, 0x4f, 0x3c, 0xf1, 0x5b, 0x32, 0xb2, 0xdc, 0x70, 0xd6,
	0x51, 0xad, 0x84, 0xf1, 0x44, 0x5f, 0x74, 0xf2, 0x1f, 0x82, 0x09, 0x06, 0x9d, 0x68, 0xf4, 0x7c,
	0x12, 0x76, 0x56, 0xff, 0x63, 0x46, 0x65, 0xf5, 0x6f, 0x70, 0xf4, 0x6f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x86, 0xf5, 0x06, 0x75, 0xe0, 0x08, 0x00, 0x00,
}