// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: flyteidl/admin/workflow.proto

package admin

import (
	core "github.com/flyteorg/flyte/gen/pb-go/flyteidl/core"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Represents a request structure to create a revision of a workflow.
// See :ref:`ref_flyteidl.admin.Workflow` for more details
type WorkflowCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id represents the unique identifier of the workflow.
	// +required
	Id *core.Identifier `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Represents the specification for workflow.
	// +required
	Spec *WorkflowSpec `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
}

func (x *WorkflowCreateRequest) Reset() {
	*x = WorkflowCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_admin_workflow_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkflowCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkflowCreateRequest) ProtoMessage() {}

func (x *WorkflowCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_admin_workflow_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkflowCreateRequest.ProtoReflect.Descriptor instead.
func (*WorkflowCreateRequest) Descriptor() ([]byte, []int) {
	return file_flyteidl_admin_workflow_proto_rawDescGZIP(), []int{0}
}

func (x *WorkflowCreateRequest) GetId() *core.Identifier {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *WorkflowCreateRequest) GetSpec() *WorkflowSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

type WorkflowCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *WorkflowCreateResponse) Reset() {
	*x = WorkflowCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_admin_workflow_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkflowCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkflowCreateResponse) ProtoMessage() {}

func (x *WorkflowCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_admin_workflow_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkflowCreateResponse.ProtoReflect.Descriptor instead.
func (*WorkflowCreateResponse) Descriptor() ([]byte, []int) {
	return file_flyteidl_admin_workflow_proto_rawDescGZIP(), []int{1}
}

// Represents the workflow structure stored in the Admin
// A workflow is created by ordering tasks and associating outputs to inputs
// in order to produce a directed-acyclic execution graph.
type Workflow struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id represents the unique identifier of the workflow.
	Id *core.Identifier `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// closure encapsulates all the fields that maps to a compiled version of the workflow.
	Closure *WorkflowClosure `protobuf:"bytes,2,opt,name=closure,proto3" json:"closure,omitempty"`
	// One-liner overview of the entity.
	ShortDescription string `protobuf:"bytes,3,opt,name=short_description,json=shortDescription,proto3" json:"short_description,omitempty"`
}

func (x *Workflow) Reset() {
	*x = Workflow{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_admin_workflow_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Workflow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Workflow) ProtoMessage() {}

func (x *Workflow) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_admin_workflow_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Workflow.ProtoReflect.Descriptor instead.
func (*Workflow) Descriptor() ([]byte, []int) {
	return file_flyteidl_admin_workflow_proto_rawDescGZIP(), []int{2}
}

func (x *Workflow) GetId() *core.Identifier {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Workflow) GetClosure() *WorkflowClosure {
	if x != nil {
		return x.Closure
	}
	return nil
}

func (x *Workflow) GetShortDescription() string {
	if x != nil {
		return x.ShortDescription
	}
	return ""
}

// Represents a list of workflows returned from the admin.
// See :ref:`ref_flyteidl.admin.Workflow` for more details
type WorkflowList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of workflows returned based on the request.
	Workflows []*Workflow `protobuf:"bytes,1,rep,name=workflows,proto3" json:"workflows,omitempty"`
	// In the case of multiple pages of results, the server-provided token can be used to fetch the next page
	// in a query. If there are no more results, this value will be empty.
	Token string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *WorkflowList) Reset() {
	*x = WorkflowList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_admin_workflow_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkflowList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkflowList) ProtoMessage() {}

func (x *WorkflowList) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_admin_workflow_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkflowList.ProtoReflect.Descriptor instead.
func (*WorkflowList) Descriptor() ([]byte, []int) {
	return file_flyteidl_admin_workflow_proto_rawDescGZIP(), []int{3}
}

func (x *WorkflowList) GetWorkflows() []*Workflow {
	if x != nil {
		return x.Workflows
	}
	return nil
}

func (x *WorkflowList) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

// Represents a structure that encapsulates the specification of the workflow.
type WorkflowSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Template of the task that encapsulates all the metadata of the workflow.
	Template *core.WorkflowTemplate `protobuf:"bytes,1,opt,name=template,proto3" json:"template,omitempty"`
	// Workflows that are embedded into other workflows need to be passed alongside the parent workflow to the
	// propeller compiler (since the compiler doesn't have any knowledge of other workflows - ie, it doesn't reach out
	// to Admin to see other registered workflows).  In fact, subworkflows do not even need to be registered.
	SubWorkflows []*core.WorkflowTemplate `protobuf:"bytes,2,rep,name=sub_workflows,json=subWorkflows,proto3" json:"sub_workflows,omitempty"`
	// Represents the specification for description entity.
	Description *DescriptionEntity `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *WorkflowSpec) Reset() {
	*x = WorkflowSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_admin_workflow_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkflowSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkflowSpec) ProtoMessage() {}

func (x *WorkflowSpec) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_admin_workflow_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkflowSpec.ProtoReflect.Descriptor instead.
func (*WorkflowSpec) Descriptor() ([]byte, []int) {
	return file_flyteidl_admin_workflow_proto_rawDescGZIP(), []int{4}
}

func (x *WorkflowSpec) GetTemplate() *core.WorkflowTemplate {
	if x != nil {
		return x.Template
	}
	return nil
}

func (x *WorkflowSpec) GetSubWorkflows() []*core.WorkflowTemplate {
	if x != nil {
		return x.SubWorkflows
	}
	return nil
}

func (x *WorkflowSpec) GetDescription() *DescriptionEntity {
	if x != nil {
		return x.Description
	}
	return nil
}

// A container holding the compiled workflow produced from the WorkflowSpec and additional metadata.
type WorkflowClosure struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Represents the compiled representation of the workflow from the specification provided.
	CompiledWorkflow *core.CompiledWorkflowClosure `protobuf:"bytes,1,opt,name=compiled_workflow,json=compiledWorkflow,proto3" json:"compiled_workflow,omitempty"`
	// Time at which the workflow was created.
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *WorkflowClosure) Reset() {
	*x = WorkflowClosure{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_admin_workflow_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkflowClosure) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkflowClosure) ProtoMessage() {}

func (x *WorkflowClosure) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_admin_workflow_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkflowClosure.ProtoReflect.Descriptor instead.
func (*WorkflowClosure) Descriptor() ([]byte, []int) {
	return file_flyteidl_admin_workflow_proto_rawDescGZIP(), []int{5}
}

func (x *WorkflowClosure) GetCompiledWorkflow() *core.CompiledWorkflowClosure {
	if x != nil {
		return x.CompiledWorkflow
	}
	return nil
}

func (x *WorkflowClosure) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

// The workflow id is already used and the structure is different
type WorkflowErrorExistsDifferentStructure struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id *core.Identifier `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *WorkflowErrorExistsDifferentStructure) Reset() {
	*x = WorkflowErrorExistsDifferentStructure{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_admin_workflow_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkflowErrorExistsDifferentStructure) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkflowErrorExistsDifferentStructure) ProtoMessage() {}

func (x *WorkflowErrorExistsDifferentStructure) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_admin_workflow_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkflowErrorExistsDifferentStructure.ProtoReflect.Descriptor instead.
func (*WorkflowErrorExistsDifferentStructure) Descriptor() ([]byte, []int) {
	return file_flyteidl_admin_workflow_proto_rawDescGZIP(), []int{6}
}

func (x *WorkflowErrorExistsDifferentStructure) GetId() *core.Identifier {
	if x != nil {
		return x.Id
	}
	return nil
}

// The workflow id is already used with an identical sctructure
type WorkflowErrorExistsIdenticalStructure struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id *core.Identifier `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *WorkflowErrorExistsIdenticalStructure) Reset() {
	*x = WorkflowErrorExistsIdenticalStructure{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_admin_workflow_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkflowErrorExistsIdenticalStructure) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkflowErrorExistsIdenticalStructure) ProtoMessage() {}

func (x *WorkflowErrorExistsIdenticalStructure) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_admin_workflow_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkflowErrorExistsIdenticalStructure.ProtoReflect.Descriptor instead.
func (*WorkflowErrorExistsIdenticalStructure) Descriptor() ([]byte, []int) {
	return file_flyteidl_admin_workflow_proto_rawDescGZIP(), []int{7}
}

func (x *WorkflowErrorExistsIdenticalStructure) GetId() *core.Identifier {
	if x != nil {
		return x.Id
	}
	return nil
}

// When a CreateWorkflowRequest failes due to matching id
type CreateWorkflowFailureReason struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Reason:
	//
	//	*CreateWorkflowFailureReason_ExistsDifferentStructure
	//	*CreateWorkflowFailureReason_ExistsIdenticalStructure
	Reason isCreateWorkflowFailureReason_Reason `protobuf_oneof:"reason"`
}

func (x *CreateWorkflowFailureReason) Reset() {
	*x = CreateWorkflowFailureReason{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_admin_workflow_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateWorkflowFailureReason) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWorkflowFailureReason) ProtoMessage() {}

func (x *CreateWorkflowFailureReason) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_admin_workflow_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWorkflowFailureReason.ProtoReflect.Descriptor instead.
func (*CreateWorkflowFailureReason) Descriptor() ([]byte, []int) {
	return file_flyteidl_admin_workflow_proto_rawDescGZIP(), []int{8}
}

func (m *CreateWorkflowFailureReason) GetReason() isCreateWorkflowFailureReason_Reason {
	if m != nil {
		return m.Reason
	}
	return nil
}

func (x *CreateWorkflowFailureReason) GetExistsDifferentStructure() *WorkflowErrorExistsDifferentStructure {
	if x, ok := x.GetReason().(*CreateWorkflowFailureReason_ExistsDifferentStructure); ok {
		return x.ExistsDifferentStructure
	}
	return nil
}

func (x *CreateWorkflowFailureReason) GetExistsIdenticalStructure() *WorkflowErrorExistsIdenticalStructure {
	if x, ok := x.GetReason().(*CreateWorkflowFailureReason_ExistsIdenticalStructure); ok {
		return x.ExistsIdenticalStructure
	}
	return nil
}

type isCreateWorkflowFailureReason_Reason interface {
	isCreateWorkflowFailureReason_Reason()
}

type CreateWorkflowFailureReason_ExistsDifferentStructure struct {
	ExistsDifferentStructure *WorkflowErrorExistsDifferentStructure `protobuf:"bytes,1,opt,name=exists_different_structure,json=existsDifferentStructure,proto3,oneof"`
}

type CreateWorkflowFailureReason_ExistsIdenticalStructure struct {
	ExistsIdenticalStructure *WorkflowErrorExistsIdenticalStructure `protobuf:"bytes,2,opt,name=exists_identical_structure,json=existsIdenticalStructure,proto3,oneof"`
}

func (*CreateWorkflowFailureReason_ExistsDifferentStructure) isCreateWorkflowFailureReason_Reason() {}

func (*CreateWorkflowFailureReason_ExistsIdenticalStructure) isCreateWorkflowFailureReason_Reason() {}

var File_flyteidl_admin_workflow_proto protoreflect.FileDescriptor

var file_flyteidl_admin_workflow_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x1a,
	0x27, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69,
	0x64, 0x6c, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c,
	0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c,
	0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x74, 0x0a, 0x15, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f,
	0x77, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x66, 0x6c, 0x79,
	0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x02, 0x69, 0x64, 0x12, 0x30, 0x0a, 0x04, 0x73, 0x70, 0x65,
	0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69,
	0x64, 0x6c, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f,
	0x77, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x22, 0x18, 0x0a, 0x16, 0x57,
	0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x9d, 0x01, 0x0a, 0x08, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c,
	0x6f, 0x77, 0x12, 0x29, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x02, 0x69, 0x64, 0x12, 0x39, 0x0a,
	0x07, 0x63, 0x6c, 0x6f, 0x73, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f,
	0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e,
	0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x43, 0x6c, 0x6f, 0x73, 0x75, 0x72, 0x65, 0x52,
	0x07, 0x63, 0x6c, 0x6f, 0x73, 0x75, 0x72, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x10, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x5c, 0x0a, 0x0c, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f,
	0x77, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f,
	0x77, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65,
	0x69, 0x64, 0x6c, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c,
	0x6f, 0x77, 0x52, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x22, 0xd6, 0x01, 0x0a, 0x0c, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x53, 0x70, 0x65, 0x63, 0x12, 0x3b, 0x0a, 0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64,
	0x6c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x54,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x12, 0x44, 0x0a, 0x0d, 0x73, 0x75, 0x62, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f,
	0x77, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65,
	0x69, 0x64, 0x6c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f,
	0x77, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x0c, 0x73, 0x75, 0x62, 0x57, 0x6f,
	0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x12, 0x43, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x66,
	0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xa1, 0x01, 0x0a,
	0x0f, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x43, 0x6c, 0x6f, 0x73, 0x75, 0x72, 0x65,
	0x12, 0x53, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x64, 0x5f, 0x77, 0x6f, 0x72,
	0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x66, 0x6c,
	0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x70,
	0x69, 0x6c, 0x65, 0x64, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x43, 0x6c, 0x6f, 0x73,
	0x75, 0x72, 0x65, 0x52, 0x10, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x64, 0x57, 0x6f, 0x72,
	0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x22, 0x52, 0x0a, 0x25, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x44, 0x69, 0x66, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x74,
	0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x29, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x52, 0x0a, 0x25, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x63, 0x61, 0x6c, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x29, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x66, 0x6c, 0x79, 0x74,
	0x65, 0x69, 0x64, 0x6c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x52, 0x02, 0x69, 0x64, 0x22, 0x95, 0x02, 0x0a, 0x1b, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x46, 0x61, 0x69, 0x6c, 0x75,
	0x72, 0x65, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x75, 0x0a, 0x1a, 0x65, 0x78, 0x69, 0x73,
	0x74, 0x73, 0x5f, 0x64, 0x69, 0x66, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x66,
	0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x57, 0x6f,
	0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x45, 0x78, 0x69, 0x73, 0x74,
	0x73, 0x44, 0x69, 0x66, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x75, 0x72, 0x65, 0x48, 0x00, 0x52, 0x18, 0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x44, 0x69, 0x66,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12,
	0x75, 0x0a, 0x1a, 0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x63, 0x61, 0x6c, 0x5f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61,
	0x6c, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x48, 0x00, 0x52, 0x18, 0x65, 0x78,
	0x69, 0x73, 0x74, 0x73, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x53, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x42, 0xb2, 0x01, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64,
	0x6c, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x42, 0x0d, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f,
	0x77, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x48, 0x02, 0x50, 0x01, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x6f, 0x72, 0x67, 0x2f,
	0x66, 0x6c, 0x79, 0x74, 0x65, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x62, 0x2d, 0x67, 0x6f, 0x2f,
	0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0xa2, 0x02,
	0x03, 0x46, 0x41, 0x58, 0xaa, 0x02, 0x0e, 0x46, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0xca, 0x02, 0x0e, 0x46, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c,
	0x5c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0xe2, 0x02, 0x1a, 0x46, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64,
	0x6c, 0x5c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x0f, 0x46, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x3a, 0x3a,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_flyteidl_admin_workflow_proto_rawDescOnce sync.Once
	file_flyteidl_admin_workflow_proto_rawDescData = file_flyteidl_admin_workflow_proto_rawDesc
)

func file_flyteidl_admin_workflow_proto_rawDescGZIP() []byte {
	file_flyteidl_admin_workflow_proto_rawDescOnce.Do(func() {
		file_flyteidl_admin_workflow_proto_rawDescData = protoimpl.X.CompressGZIP(file_flyteidl_admin_workflow_proto_rawDescData)
	})
	return file_flyteidl_admin_workflow_proto_rawDescData
}

var file_flyteidl_admin_workflow_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_flyteidl_admin_workflow_proto_goTypes = []interface{}{
	(*WorkflowCreateRequest)(nil),                 // 0: flyteidl.admin.WorkflowCreateRequest
	(*WorkflowCreateResponse)(nil),                // 1: flyteidl.admin.WorkflowCreateResponse
	(*Workflow)(nil),                              // 2: flyteidl.admin.Workflow
	(*WorkflowList)(nil),                          // 3: flyteidl.admin.WorkflowList
	(*WorkflowSpec)(nil),                          // 4: flyteidl.admin.WorkflowSpec
	(*WorkflowClosure)(nil),                       // 5: flyteidl.admin.WorkflowClosure
	(*WorkflowErrorExistsDifferentStructure)(nil), // 6: flyteidl.admin.WorkflowErrorExistsDifferentStructure
	(*WorkflowErrorExistsIdenticalStructure)(nil), // 7: flyteidl.admin.WorkflowErrorExistsIdenticalStructure
	(*CreateWorkflowFailureReason)(nil),           // 8: flyteidl.admin.CreateWorkflowFailureReason
	(*core.Identifier)(nil),                       // 9: flyteidl.core.Identifier
	(*core.WorkflowTemplate)(nil),                 // 10: flyteidl.core.WorkflowTemplate
	(*DescriptionEntity)(nil),                     // 11: flyteidl.admin.DescriptionEntity
	(*core.CompiledWorkflowClosure)(nil),          // 12: flyteidl.core.CompiledWorkflowClosure
	(*timestamppb.Timestamp)(nil),                 // 13: google.protobuf.Timestamp
}
var file_flyteidl_admin_workflow_proto_depIdxs = []int32{
	9,  // 0: flyteidl.admin.WorkflowCreateRequest.id:type_name -> flyteidl.core.Identifier
	4,  // 1: flyteidl.admin.WorkflowCreateRequest.spec:type_name -> flyteidl.admin.WorkflowSpec
	9,  // 2: flyteidl.admin.Workflow.id:type_name -> flyteidl.core.Identifier
	5,  // 3: flyteidl.admin.Workflow.closure:type_name -> flyteidl.admin.WorkflowClosure
	2,  // 4: flyteidl.admin.WorkflowList.workflows:type_name -> flyteidl.admin.Workflow
	10, // 5: flyteidl.admin.WorkflowSpec.template:type_name -> flyteidl.core.WorkflowTemplate
	10, // 6: flyteidl.admin.WorkflowSpec.sub_workflows:type_name -> flyteidl.core.WorkflowTemplate
	11, // 7: flyteidl.admin.WorkflowSpec.description:type_name -> flyteidl.admin.DescriptionEntity
	12, // 8: flyteidl.admin.WorkflowClosure.compiled_workflow:type_name -> flyteidl.core.CompiledWorkflowClosure
	13, // 9: flyteidl.admin.WorkflowClosure.created_at:type_name -> google.protobuf.Timestamp
	9,  // 10: flyteidl.admin.WorkflowErrorExistsDifferentStructure.id:type_name -> flyteidl.core.Identifier
	9,  // 11: flyteidl.admin.WorkflowErrorExistsIdenticalStructure.id:type_name -> flyteidl.core.Identifier
	6,  // 12: flyteidl.admin.CreateWorkflowFailureReason.exists_different_structure:type_name -> flyteidl.admin.WorkflowErrorExistsDifferentStructure
	7,  // 13: flyteidl.admin.CreateWorkflowFailureReason.exists_identical_structure:type_name -> flyteidl.admin.WorkflowErrorExistsIdenticalStructure
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_flyteidl_admin_workflow_proto_init() }
func file_flyteidl_admin_workflow_proto_init() {
	if File_flyteidl_admin_workflow_proto != nil {
		return
	}
	file_flyteidl_admin_description_entity_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_flyteidl_admin_workflow_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkflowCreateRequest); i {
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
		file_flyteidl_admin_workflow_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkflowCreateResponse); i {
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
		file_flyteidl_admin_workflow_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Workflow); i {
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
		file_flyteidl_admin_workflow_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkflowList); i {
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
		file_flyteidl_admin_workflow_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkflowSpec); i {
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
		file_flyteidl_admin_workflow_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkflowClosure); i {
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
		file_flyteidl_admin_workflow_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkflowErrorExistsDifferentStructure); i {
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
		file_flyteidl_admin_workflow_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkflowErrorExistsIdenticalStructure); i {
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
		file_flyteidl_admin_workflow_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateWorkflowFailureReason); i {
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
	file_flyteidl_admin_workflow_proto_msgTypes[8].OneofWrappers = []interface{}{
		(*CreateWorkflowFailureReason_ExistsDifferentStructure)(nil),
		(*CreateWorkflowFailureReason_ExistsIdenticalStructure)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_flyteidl_admin_workflow_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_flyteidl_admin_workflow_proto_goTypes,
		DependencyIndexes: file_flyteidl_admin_workflow_proto_depIdxs,
		MessageInfos:      file_flyteidl_admin_workflow_proto_msgTypes,
	}.Build()
	File_flyteidl_admin_workflow_proto = out.File
	file_flyteidl_admin_workflow_proto_rawDesc = nil
	file_flyteidl_admin_workflow_proto_goTypes = nil
	file_flyteidl_admin_workflow_proto_depIdxs = nil
}
