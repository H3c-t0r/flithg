// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: flyteidl/core/artifact_id.proto

package core

import (
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

type ArtifactKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Project and domain and suffix needs to be unique across a given artifact store.
	Project string `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	Domain  string `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
	Name    string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Org     string `protobuf:"bytes,4,opt,name=org,proto3" json:"org,omitempty"`
}

func (x *ArtifactKey) Reset() {
	*x = ArtifactKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_core_artifact_id_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArtifactKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArtifactKey) ProtoMessage() {}

func (x *ArtifactKey) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_core_artifact_id_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArtifactKey.ProtoReflect.Descriptor instead.
func (*ArtifactKey) Descriptor() ([]byte, []int) {
	return file_flyteidl_core_artifact_id_proto_rawDescGZIP(), []int{0}
}

func (x *ArtifactKey) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *ArtifactKey) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *ArtifactKey) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ArtifactKey) GetOrg() string {
	if x != nil {
		return x.Org
	}
	return ""
}

// Only valid for triggers
type ArtifactBindingData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// These two fields are only relevant in the partition value case
	//
	// Types that are assignable to PartitionData:
	//
	//	*ArtifactBindingData_PartitionKey
	//	*ArtifactBindingData_BindToTimePartition
	PartitionData isArtifactBindingData_PartitionData `protobuf_oneof:"partition_data"`
	// This is only relevant in the time partition case
	Transform string `protobuf:"bytes,4,opt,name=transform,proto3" json:"transform,omitempty"`
}

func (x *ArtifactBindingData) Reset() {
	*x = ArtifactBindingData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_core_artifact_id_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArtifactBindingData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArtifactBindingData) ProtoMessage() {}

func (x *ArtifactBindingData) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_core_artifact_id_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArtifactBindingData.ProtoReflect.Descriptor instead.
func (*ArtifactBindingData) Descriptor() ([]byte, []int) {
	return file_flyteidl_core_artifact_id_proto_rawDescGZIP(), []int{1}
}

func (m *ArtifactBindingData) GetPartitionData() isArtifactBindingData_PartitionData {
	if m != nil {
		return m.PartitionData
	}
	return nil
}

func (x *ArtifactBindingData) GetPartitionKey() string {
	if x, ok := x.GetPartitionData().(*ArtifactBindingData_PartitionKey); ok {
		return x.PartitionKey
	}
	return ""
}

func (x *ArtifactBindingData) GetBindToTimePartition() bool {
	if x, ok := x.GetPartitionData().(*ArtifactBindingData_BindToTimePartition); ok {
		return x.BindToTimePartition
	}
	return false
}

func (x *ArtifactBindingData) GetTransform() string {
	if x != nil {
		return x.Transform
	}
	return ""
}

type isArtifactBindingData_PartitionData interface {
	isArtifactBindingData_PartitionData()
}

type ArtifactBindingData_PartitionKey struct {
	PartitionKey string `protobuf:"bytes,2,opt,name=partition_key,json=partitionKey,proto3,oneof"`
}

type ArtifactBindingData_BindToTimePartition struct {
	BindToTimePartition bool `protobuf:"varint,3,opt,name=bind_to_time_partition,json=bindToTimePartition,proto3,oneof"`
}

func (*ArtifactBindingData_PartitionKey) isArtifactBindingData_PartitionData() {}

func (*ArtifactBindingData_BindToTimePartition) isArtifactBindingData_PartitionData() {}

type InputBindingData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Var string `protobuf:"bytes,1,opt,name=var,proto3" json:"var,omitempty"`
}

func (x *InputBindingData) Reset() {
	*x = InputBindingData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_core_artifact_id_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InputBindingData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InputBindingData) ProtoMessage() {}

func (x *InputBindingData) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_core_artifact_id_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InputBindingData.ProtoReflect.Descriptor instead.
func (*InputBindingData) Descriptor() ([]byte, []int) {
	return file_flyteidl_core_artifact_id_proto_rawDescGZIP(), []int{2}
}

func (x *InputBindingData) GetVar() string {
	if x != nil {
		return x.Var
	}
	return ""
}

type LabelValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Value:
	//
	//	*LabelValue_StaticValue
	//	*LabelValue_TimeValue
	//	*LabelValue_TriggeredBinding
	//	*LabelValue_InputBinding
	Value isLabelValue_Value `protobuf_oneof:"value"`
}

func (x *LabelValue) Reset() {
	*x = LabelValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_core_artifact_id_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LabelValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LabelValue) ProtoMessage() {}

func (x *LabelValue) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_core_artifact_id_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LabelValue.ProtoReflect.Descriptor instead.
func (*LabelValue) Descriptor() ([]byte, []int) {
	return file_flyteidl_core_artifact_id_proto_rawDescGZIP(), []int{3}
}

func (m *LabelValue) GetValue() isLabelValue_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *LabelValue) GetStaticValue() string {
	if x, ok := x.GetValue().(*LabelValue_StaticValue); ok {
		return x.StaticValue
	}
	return ""
}

func (x *LabelValue) GetTimeValue() *timestamppb.Timestamp {
	if x, ok := x.GetValue().(*LabelValue_TimeValue); ok {
		return x.TimeValue
	}
	return nil
}

func (x *LabelValue) GetTriggeredBinding() *ArtifactBindingData {
	if x, ok := x.GetValue().(*LabelValue_TriggeredBinding); ok {
		return x.TriggeredBinding
	}
	return nil
}

func (x *LabelValue) GetInputBinding() *InputBindingData {
	if x, ok := x.GetValue().(*LabelValue_InputBinding); ok {
		return x.InputBinding
	}
	return nil
}

type isLabelValue_Value interface {
	isLabelValue_Value()
}

type LabelValue_StaticValue struct {
	// The string static value is for use in the Partitions object
	StaticValue string `protobuf:"bytes,1,opt,name=static_value,json=staticValue,proto3,oneof"`
}

type LabelValue_TimeValue struct {
	// The time value is for use in the TimePartition case
	TimeValue *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=time_value,json=timeValue,proto3,oneof"`
}

type LabelValue_TriggeredBinding struct {
	TriggeredBinding *ArtifactBindingData `protobuf:"bytes,3,opt,name=triggered_binding,json=triggeredBinding,proto3,oneof"`
}

type LabelValue_InputBinding struct {
	InputBinding *InputBindingData `protobuf:"bytes,4,opt,name=input_binding,json=inputBinding,proto3,oneof"`
}

func (*LabelValue_StaticValue) isLabelValue_Value() {}

func (*LabelValue_TimeValue) isLabelValue_Value() {}

func (*LabelValue_TriggeredBinding) isLabelValue_Value() {}

func (*LabelValue_InputBinding) isLabelValue_Value() {}

type Partitions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value map[string]*LabelValue `protobuf:"bytes,1,rep,name=value,proto3" json:"value,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Partitions) Reset() {
	*x = Partitions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_core_artifact_id_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Partitions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Partitions) ProtoMessage() {}

func (x *Partitions) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_core_artifact_id_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Partitions.ProtoReflect.Descriptor instead.
func (*Partitions) Descriptor() ([]byte, []int) {
	return file_flyteidl_core_artifact_id_proto_rawDescGZIP(), []int{4}
}

func (x *Partitions) GetValue() map[string]*LabelValue {
	if x != nil {
		return x.Value
	}
	return nil
}

type TimePartition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value *LabelValue `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *TimePartition) Reset() {
	*x = TimePartition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_core_artifact_id_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimePartition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimePartition) ProtoMessage() {}

func (x *TimePartition) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_core_artifact_id_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimePartition.ProtoReflect.Descriptor instead.
func (*TimePartition) Descriptor() ([]byte, []int) {
	return file_flyteidl_core_artifact_id_proto_rawDescGZIP(), []int{5}
}

func (x *TimePartition) GetValue() *LabelValue {
	if x != nil {
		return x.Value
	}
	return nil
}

type ArtifactID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ArtifactKey *ArtifactKey `protobuf:"bytes,1,opt,name=artifact_key,json=artifactKey,proto3" json:"artifact_key,omitempty"`
	Version     string       `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	// Think of a partition as a tag on an Artifact, except it's a key-value pair.
	// Different partitions naturally have different versions (execution ids).
	Partitions *Partitions `protobuf:"bytes,3,opt,name=partitions,proto3" json:"partitions,omitempty"`
	// There is no such thing as an empty time partition - if it's not set, then there is no time partition.
	TimePartition *TimePartition `protobuf:"bytes,4,opt,name=time_partition,json=timePartition,proto3" json:"time_partition,omitempty"`
}

func (x *ArtifactID) Reset() {
	*x = ArtifactID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_core_artifact_id_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArtifactID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArtifactID) ProtoMessage() {}

func (x *ArtifactID) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_core_artifact_id_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArtifactID.ProtoReflect.Descriptor instead.
func (*ArtifactID) Descriptor() ([]byte, []int) {
	return file_flyteidl_core_artifact_id_proto_rawDescGZIP(), []int{6}
}

func (x *ArtifactID) GetArtifactKey() *ArtifactKey {
	if x != nil {
		return x.ArtifactKey
	}
	return nil
}

func (x *ArtifactID) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *ArtifactID) GetPartitions() *Partitions {
	if x != nil {
		return x.Partitions
	}
	return nil
}

func (x *ArtifactID) GetTimePartition() *TimePartition {
	if x != nil {
		return x.TimePartition
	}
	return nil
}

type ArtifactTag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ArtifactKey *ArtifactKey `protobuf:"bytes,1,opt,name=artifact_key,json=artifactKey,proto3" json:"artifact_key,omitempty"`
	Value       *LabelValue  `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ArtifactTag) Reset() {
	*x = ArtifactTag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_core_artifact_id_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArtifactTag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArtifactTag) ProtoMessage() {}

func (x *ArtifactTag) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_core_artifact_id_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArtifactTag.ProtoReflect.Descriptor instead.
func (*ArtifactTag) Descriptor() ([]byte, []int) {
	return file_flyteidl_core_artifact_id_proto_rawDescGZIP(), []int{7}
}

func (x *ArtifactTag) GetArtifactKey() *ArtifactKey {
	if x != nil {
		return x.ArtifactKey
	}
	return nil
}

func (x *ArtifactTag) GetValue() *LabelValue {
	if x != nil {
		return x.Value
	}
	return nil
}

// Uniqueness constraints for Artifacts
//   - project, domain, name, version, partitions
//
// Option 2 (tags are standalone, point to an individual artifact id):
//   - project, domain, name, alias (points to one partition if partitioned)
//   - project, domain, name, partition key, partition value
type ArtifactQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Identifier:
	//
	//	*ArtifactQuery_ArtifactId
	//	*ArtifactQuery_ArtifactTag
	//	*ArtifactQuery_Uri
	//	*ArtifactQuery_Binding
	Identifier isArtifactQuery_Identifier `protobuf_oneof:"identifier"`
}

func (x *ArtifactQuery) Reset() {
	*x = ArtifactQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_core_artifact_id_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArtifactQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArtifactQuery) ProtoMessage() {}

func (x *ArtifactQuery) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_core_artifact_id_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArtifactQuery.ProtoReflect.Descriptor instead.
func (*ArtifactQuery) Descriptor() ([]byte, []int) {
	return file_flyteidl_core_artifact_id_proto_rawDescGZIP(), []int{8}
}

func (m *ArtifactQuery) GetIdentifier() isArtifactQuery_Identifier {
	if m != nil {
		return m.Identifier
	}
	return nil
}

func (x *ArtifactQuery) GetArtifactId() *ArtifactID {
	if x, ok := x.GetIdentifier().(*ArtifactQuery_ArtifactId); ok {
		return x.ArtifactId
	}
	return nil
}

func (x *ArtifactQuery) GetArtifactTag() *ArtifactTag {
	if x, ok := x.GetIdentifier().(*ArtifactQuery_ArtifactTag); ok {
		return x.ArtifactTag
	}
	return nil
}

func (x *ArtifactQuery) GetUri() string {
	if x, ok := x.GetIdentifier().(*ArtifactQuery_Uri); ok {
		return x.Uri
	}
	return ""
}

func (x *ArtifactQuery) GetBinding() *ArtifactBindingData {
	if x, ok := x.GetIdentifier().(*ArtifactQuery_Binding); ok {
		return x.Binding
	}
	return nil
}

type isArtifactQuery_Identifier interface {
	isArtifactQuery_Identifier()
}

type ArtifactQuery_ArtifactId struct {
	ArtifactId *ArtifactID `protobuf:"bytes,1,opt,name=artifact_id,json=artifactId,proto3,oneof"`
}

type ArtifactQuery_ArtifactTag struct {
	ArtifactTag *ArtifactTag `protobuf:"bytes,2,opt,name=artifact_tag,json=artifactTag,proto3,oneof"`
}

type ArtifactQuery_Uri struct {
	Uri string `protobuf:"bytes,3,opt,name=uri,proto3,oneof"`
}

type ArtifactQuery_Binding struct {
	// This is used in the trigger case, where a user specifies a value for an input that is one of the triggering
	// artifacts, or a partition value derived from a triggering artifact.
	Binding *ArtifactBindingData `protobuf:"bytes,4,opt,name=binding,proto3,oneof"`
}

func (*ArtifactQuery_ArtifactId) isArtifactQuery_Identifier() {}

func (*ArtifactQuery_ArtifactTag) isArtifactQuery_Identifier() {}

func (*ArtifactQuery_Uri) isArtifactQuery_Identifier() {}

func (*ArtifactQuery_Binding) isArtifactQuery_Identifier() {}

var File_flyteidl_core_artifact_id_proto protoreflect.FileDescriptor

var file_flyteidl_core_artifact_id_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f,
	0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0d, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2f, 0x63, 0x6f, 0x72, 0x65,
	0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x65, 0x0a, 0x0b, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x4b, 0x65, 0x79,
	0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6f, 0x72, 0x67, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6f, 0x72, 0x67, 0x22, 0xa3, 0x01, 0x0a, 0x13, 0x41, 0x72, 0x74,
	0x69, 0x66, 0x61, 0x63, 0x74, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x25, 0x0a, 0x0d, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6b, 0x65,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0c, 0x70, 0x61, 0x72, 0x74, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x12, 0x35, 0x0a, 0x16, 0x62, 0x69, 0x6e, 0x64, 0x5f,
	0x74, 0x6f, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x13, 0x62, 0x69, 0x6e, 0x64, 0x54,
	0x6f, 0x54, 0x69, 0x6d, 0x65, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c,
	0x0a, 0x09, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x42, 0x10, 0x0a, 0x0e,
	0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x22, 0x24,
	0x0a, 0x10, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x76, 0x61, 0x72, 0x22, 0x92, 0x02, 0x0a, 0x0a, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x23, 0x0a, 0x0c, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x5f, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x73, 0x74, 0x61,
	0x74, 0x69, 0x63, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x3b, 0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65,
	0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x00, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x51, 0x0a, 0x11, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72,
	0x65, 0x64, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x22, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x10, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x65,
	0x64, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x46, 0x0a, 0x0d, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1f, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x74, 0x61,
	0x48, 0x00, 0x52, 0x0c, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x42, 0x07, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x9d, 0x01, 0x0a, 0x0a, 0x50, 0x61,
	0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x3a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69,
	0x64, 0x6c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x1a, 0x53, 0x0a, 0x0a, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x2f, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x40, 0x0a, 0x0d, 0x54, 0x69, 0x6d,
	0x65, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2f, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x66, 0x6c, 0x79, 0x74,
	0x65, 0x69, 0x64, 0x6c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xe5, 0x01, 0x0a, 0x0a,
	0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x49, 0x44, 0x12, 0x3d, 0x0a, 0x0c, 0x61, 0x72,
	0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x4b, 0x65, 0x79, 0x52, 0x0b, 0x61, 0x72,
	0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69,
	0x64, 0x6c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x43,
	0x0a, 0x0e, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64,
	0x6c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x50, 0x61, 0x72, 0x74, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x74, 0x69, 0x6d, 0x65, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x7d, 0x0a, 0x0b, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x54,
	0x61, 0x67, 0x12, 0x3d, 0x0a, 0x0c, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x5f, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65,
	0x69, 0x64, 0x6c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63,
	0x74, 0x4b, 0x65, 0x79, 0x52, 0x0b, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x4b, 0x65,
	0x79, 0x12, 0x2f, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x22, 0xf0, 0x01, 0x0a, 0x0d, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x12, 0x3c, 0x0a, 0x0b, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x66, 0x6c, 0x79, 0x74,
	0x65, 0x69, 0x64, 0x6c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61,
	0x63, 0x74, 0x49, 0x44, 0x48, 0x00, 0x52, 0x0a, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74,
	0x49, 0x64, 0x12, 0x3f, 0x0a, 0x0c, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x5f, 0x74,
	0x61, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65,
	0x69, 0x64, 0x6c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63,
	0x74, 0x54, 0x61, 0x67, 0x48, 0x00, 0x52, 0x0b, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74,
	0x54, 0x61, 0x67, 0x12, 0x12, 0x0a, 0x03, 0x75, 0x72, 0x69, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x03, 0x75, 0x72, 0x69, 0x12, 0x3e, 0x0a, 0x07, 0x62, 0x69, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65,
	0x69, 0x64, 0x6c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63,
	0x74, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x07,
	0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x42, 0x0c, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x42, 0xb5, 0x01, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x2e, 0x66, 0x6c,
	0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x42, 0x0f, 0x41, 0x72, 0x74,
	0x69, 0x66, 0x61, 0x63, 0x74, 0x49, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3a,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x6c, 0x79, 0x74, 0x65,
	0x6f, 0x72, 0x67, 0x2f, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x2f, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69,
	0x64, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x62, 0x2d, 0x67, 0x6f, 0x2f, 0x66, 0x6c, 0x79,
	0x74, 0x65, 0x69, 0x64, 0x6c, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0xa2, 0x02, 0x03, 0x46, 0x43, 0x58,
	0xaa, 0x02, 0x0d, 0x46, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x43, 0x6f, 0x72, 0x65,
	0xca, 0x02, 0x0d, 0x46, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x5c, 0x43, 0x6f, 0x72, 0x65,
	0xe2, 0x02, 0x19, 0x46, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x5c, 0x43, 0x6f, 0x72, 0x65,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0e, 0x46,
	0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x3a, 0x3a, 0x43, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_flyteidl_core_artifact_id_proto_rawDescOnce sync.Once
	file_flyteidl_core_artifact_id_proto_rawDescData = file_flyteidl_core_artifact_id_proto_rawDesc
)

func file_flyteidl_core_artifact_id_proto_rawDescGZIP() []byte {
	file_flyteidl_core_artifact_id_proto_rawDescOnce.Do(func() {
		file_flyteidl_core_artifact_id_proto_rawDescData = protoimpl.X.CompressGZIP(file_flyteidl_core_artifact_id_proto_rawDescData)
	})
	return file_flyteidl_core_artifact_id_proto_rawDescData
}

var file_flyteidl_core_artifact_id_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_flyteidl_core_artifact_id_proto_goTypes = []interface{}{
	(*ArtifactKey)(nil),           // 0: flyteidl.core.ArtifactKey
	(*ArtifactBindingData)(nil),   // 1: flyteidl.core.ArtifactBindingData
	(*InputBindingData)(nil),      // 2: flyteidl.core.InputBindingData
	(*LabelValue)(nil),            // 3: flyteidl.core.LabelValue
	(*Partitions)(nil),            // 4: flyteidl.core.Partitions
	(*TimePartition)(nil),         // 5: flyteidl.core.TimePartition
	(*ArtifactID)(nil),            // 6: flyteidl.core.ArtifactID
	(*ArtifactTag)(nil),           // 7: flyteidl.core.ArtifactTag
	(*ArtifactQuery)(nil),         // 8: flyteidl.core.ArtifactQuery
	nil,                           // 9: flyteidl.core.Partitions.ValueEntry
	(*timestamppb.Timestamp)(nil), // 10: google.protobuf.Timestamp
}
var file_flyteidl_core_artifact_id_proto_depIdxs = []int32{
	10, // 0: flyteidl.core.LabelValue.time_value:type_name -> google.protobuf.Timestamp
	1,  // 1: flyteidl.core.LabelValue.triggered_binding:type_name -> flyteidl.core.ArtifactBindingData
	2,  // 2: flyteidl.core.LabelValue.input_binding:type_name -> flyteidl.core.InputBindingData
	9,  // 3: flyteidl.core.Partitions.value:type_name -> flyteidl.core.Partitions.ValueEntry
	3,  // 4: flyteidl.core.TimePartition.value:type_name -> flyteidl.core.LabelValue
	0,  // 5: flyteidl.core.ArtifactID.artifact_key:type_name -> flyteidl.core.ArtifactKey
	4,  // 6: flyteidl.core.ArtifactID.partitions:type_name -> flyteidl.core.Partitions
	5,  // 7: flyteidl.core.ArtifactID.time_partition:type_name -> flyteidl.core.TimePartition
	0,  // 8: flyteidl.core.ArtifactTag.artifact_key:type_name -> flyteidl.core.ArtifactKey
	3,  // 9: flyteidl.core.ArtifactTag.value:type_name -> flyteidl.core.LabelValue
	6,  // 10: flyteidl.core.ArtifactQuery.artifact_id:type_name -> flyteidl.core.ArtifactID
	7,  // 11: flyteidl.core.ArtifactQuery.artifact_tag:type_name -> flyteidl.core.ArtifactTag
	1,  // 12: flyteidl.core.ArtifactQuery.binding:type_name -> flyteidl.core.ArtifactBindingData
	3,  // 13: flyteidl.core.Partitions.ValueEntry.value:type_name -> flyteidl.core.LabelValue
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_flyteidl_core_artifact_id_proto_init() }
func file_flyteidl_core_artifact_id_proto_init() {
	if File_flyteidl_core_artifact_id_proto != nil {
		return
	}
	file_flyteidl_core_identifier_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_flyteidl_core_artifact_id_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArtifactKey); i {
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
		file_flyteidl_core_artifact_id_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArtifactBindingData); i {
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
		file_flyteidl_core_artifact_id_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InputBindingData); i {
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
		file_flyteidl_core_artifact_id_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LabelValue); i {
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
		file_flyteidl_core_artifact_id_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Partitions); i {
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
		file_flyteidl_core_artifact_id_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimePartition); i {
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
		file_flyteidl_core_artifact_id_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArtifactID); i {
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
		file_flyteidl_core_artifact_id_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArtifactTag); i {
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
		file_flyteidl_core_artifact_id_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArtifactQuery); i {
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
	file_flyteidl_core_artifact_id_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*ArtifactBindingData_PartitionKey)(nil),
		(*ArtifactBindingData_BindToTimePartition)(nil),
	}
	file_flyteidl_core_artifact_id_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*LabelValue_StaticValue)(nil),
		(*LabelValue_TimeValue)(nil),
		(*LabelValue_TriggeredBinding)(nil),
		(*LabelValue_InputBinding)(nil),
	}
	file_flyteidl_core_artifact_id_proto_msgTypes[8].OneofWrappers = []interface{}{
		(*ArtifactQuery_ArtifactId)(nil),
		(*ArtifactQuery_ArtifactTag)(nil),
		(*ArtifactQuery_Uri)(nil),
		(*ArtifactQuery_Binding)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_flyteidl_core_artifact_id_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_flyteidl_core_artifact_id_proto_goTypes,
		DependencyIndexes: file_flyteidl_core_artifact_id_proto_depIdxs,
		MessageInfos:      file_flyteidl_core_artifact_id_proto_msgTypes,
	}.Build()
	File_flyteidl_core_artifact_id_proto = out.File
	file_flyteidl_core_artifact_id_proto_rawDesc = nil
	file_flyteidl_core_artifact_id_proto_goTypes = nil
	file_flyteidl_core_artifact_id_proto_depIdxs = nil
}
