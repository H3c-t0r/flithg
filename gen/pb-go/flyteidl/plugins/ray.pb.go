// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: flyteidl/plugins/ray.proto

package plugins

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

// RayJobSpec defines the desired state of RayJob
type RayJob struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// RayClusterSpec is the cluster template to run the job
	RayCluster *RayCluster `protobuf:"bytes,1,opt,name=ray_cluster,json=rayCluster,proto3" json:"ray_cluster,omitempty"`
	// runtime_env is base64 encoded.
	// Ray runtime environments: https://docs.ray.io/en/latest/ray-core/handling-dependencies.html#runtime-environments
	RuntimeEnv string `protobuf:"bytes,2,opt,name=runtime_env,json=runtimeEnv,proto3" json:"runtime_env,omitempty"`
}

func (x *RayJob) Reset() {
	*x = RayJob{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_plugins_ray_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RayJob) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RayJob) ProtoMessage() {}

func (x *RayJob) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_plugins_ray_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RayJob.ProtoReflect.Descriptor instead.
func (*RayJob) Descriptor() ([]byte, []int) {
	return file_flyteidl_plugins_ray_proto_rawDescGZIP(), []int{0}
}

func (x *RayJob) GetRayCluster() *RayCluster {
	if x != nil {
		return x.RayCluster
	}
	return nil
}

func (x *RayJob) GetRuntimeEnv() string {
	if x != nil {
		return x.RuntimeEnv
	}
	return ""
}

// Define Ray cluster defines the desired state of RayCluster
type RayCluster struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// HeadGroupSpecs are the spec for the head pod
	HeadGroupSpec *HeadGroupSpec `protobuf:"bytes,1,opt,name=head_group_spec,json=headGroupSpec,proto3" json:"head_group_spec,omitempty"`
	// WorkerGroupSpecs are the specs for the worker pods
	WorkerGroupSpec []*WorkerGroupSpec `protobuf:"bytes,2,rep,name=worker_group_spec,json=workerGroupSpec,proto3" json:"worker_group_spec,omitempty"`
}

func (x *RayCluster) Reset() {
	*x = RayCluster{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_plugins_ray_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RayCluster) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RayCluster) ProtoMessage() {}

func (x *RayCluster) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_plugins_ray_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RayCluster.ProtoReflect.Descriptor instead.
func (*RayCluster) Descriptor() ([]byte, []int) {
	return file_flyteidl_plugins_ray_proto_rawDescGZIP(), []int{1}
}

func (x *RayCluster) GetHeadGroupSpec() *HeadGroupSpec {
	if x != nil {
		return x.HeadGroupSpec
	}
	return nil
}

func (x *RayCluster) GetWorkerGroupSpec() []*WorkerGroupSpec {
	if x != nil {
		return x.WorkerGroupSpec
	}
	return nil
}

// HeadGroupSpec are the spec for the head pod
type HeadGroupSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. RayStartParams are the params of the start command: address, object-store-memory.
	// Refer to https://docs.ray.io/en/latest/ray-core/package-ref.html#ray-start
	RayStartParams map[string]string `protobuf:"bytes,1,rep,name=ray_start_params,json=rayStartParams,proto3" json:"ray_start_params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *HeadGroupSpec) Reset() {
	*x = HeadGroupSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_plugins_ray_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeadGroupSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeadGroupSpec) ProtoMessage() {}

func (x *HeadGroupSpec) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_plugins_ray_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeadGroupSpec.ProtoReflect.Descriptor instead.
func (*HeadGroupSpec) Descriptor() ([]byte, []int) {
	return file_flyteidl_plugins_ray_proto_rawDescGZIP(), []int{2}
}

func (x *HeadGroupSpec) GetRayStartParams() map[string]string {
	if x != nil {
		return x.RayStartParams
	}
	return nil
}

// WorkerGroupSpec are the specs for the worker pods
type WorkerGroupSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. RayCluster can have multiple worker groups, and it distinguishes them by name
	GroupName string `protobuf:"bytes,1,opt,name=group_name,json=groupName,proto3" json:"group_name,omitempty"`
	// Required. Desired replicas of the worker group. Defaults to 1.
	Replicas int32 `protobuf:"varint,2,opt,name=replicas,proto3" json:"replicas,omitempty"`
	// Optional. Min replicas of the worker group. MinReplicas defaults to 1.
	MinReplicas int32 `protobuf:"varint,3,opt,name=min_replicas,json=minReplicas,proto3" json:"min_replicas,omitempty"`
	// Optional. Max replicas of the worker group. MaxReplicas defaults to maxInt32
	MaxReplicas int32 `protobuf:"varint,4,opt,name=max_replicas,json=maxReplicas,proto3" json:"max_replicas,omitempty"`
	// Optional. RayStartParams are the params of the start command: address, object-store-memory.
	// Refer to https://docs.ray.io/en/latest/ray-core/package-ref.html#ray-start
	RayStartParams map[string]string `protobuf:"bytes,5,rep,name=ray_start_params,json=rayStartParams,proto3" json:"ray_start_params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *WorkerGroupSpec) Reset() {
	*x = WorkerGroupSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_plugins_ray_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkerGroupSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkerGroupSpec) ProtoMessage() {}

func (x *WorkerGroupSpec) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_plugins_ray_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkerGroupSpec.ProtoReflect.Descriptor instead.
func (*WorkerGroupSpec) Descriptor() ([]byte, []int) {
	return file_flyteidl_plugins_ray_proto_rawDescGZIP(), []int{3}
}

func (x *WorkerGroupSpec) GetGroupName() string {
	if x != nil {
		return x.GroupName
	}
	return ""
}

func (x *WorkerGroupSpec) GetReplicas() int32 {
	if x != nil {
		return x.Replicas
	}
	return 0
}

func (x *WorkerGroupSpec) GetMinReplicas() int32 {
	if x != nil {
		return x.MinReplicas
	}
	return 0
}

func (x *WorkerGroupSpec) GetMaxReplicas() int32 {
	if x != nil {
		return x.MaxReplicas
	}
	return 0
}

func (x *WorkerGroupSpec) GetRayStartParams() map[string]string {
	if x != nil {
		return x.RayStartParams
	}
	return nil
}

var File_flyteidl_plugins_ray_proto protoreflect.FileDescriptor

var file_flyteidl_plugins_ray_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x73, 0x2f, 0x72, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x66, 0x6c,
	0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x22, 0x68,
	0x0a, 0x06, 0x52, 0x61, 0x79, 0x4a, 0x6f, 0x62, 0x12, 0x3d, 0x0a, 0x0b, 0x72, 0x61, 0x79, 0x5f,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73,
	0x2e, 0x52, 0x61, 0x79, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x0a, 0x72, 0x61, 0x79,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x75, 0x6e, 0x74, 0x69,
	0x6d, 0x65, 0x5f, 0x65, 0x6e, 0x76, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x75,
	0x6e, 0x74, 0x69, 0x6d, 0x65, 0x45, 0x6e, 0x76, 0x22, 0xa4, 0x01, 0x0a, 0x0a, 0x52, 0x61, 0x79,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x47, 0x0a, 0x0f, 0x68, 0x65, 0x61, 0x64, 0x5f,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1f, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x70, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x73, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x70, 0x65,
	0x63, 0x52, 0x0d, 0x68, 0x65, 0x61, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x70, 0x65, 0x63,
	0x12, 0x4d, 0x0a, 0x11, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x5f, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x66, 0x6c,
	0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2e, 0x57,
	0x6f, 0x72, 0x6b, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x70, 0x65, 0x63, 0x52, 0x0f,
	0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x70, 0x65, 0x63, 0x22,
	0xb1, 0x01, 0x0a, 0x0d, 0x48, 0x65, 0x61, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x70, 0x65,
	0x63, 0x12, 0x5d, 0x0a, 0x10, 0x72, 0x61, 0x79, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x70,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x66, 0x6c,
	0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2e, 0x48,
	0x65, 0x61, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x52, 0x61, 0x79,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x0e, 0x72, 0x61, 0x79, 0x53, 0x74, 0x61, 0x72, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x1a, 0x41, 0x0a, 0x13, 0x52, 0x61, 0x79, 0x53, 0x74, 0x61, 0x72, 0x74, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0xb6, 0x02, 0x0a, 0x0f, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x53, 0x70, 0x65, 0x63, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x69, 0x6e, 0x5f, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x61, 0x78, 0x5f, 0x72, 0x65, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6d, 0x61, 0x78,
	0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x12, 0x5f, 0x0a, 0x10, 0x72, 0x61, 0x79, 0x5f,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x35, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x70, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x73, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x52, 0x61, 0x79, 0x53, 0x74, 0x61, 0x72, 0x74, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0e, 0x72, 0x61, 0x79, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x41, 0x0a, 0x13, 0x52, 0x61, 0x79,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0xb9, 0x01, 0x0a,
	0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x70, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x73, 0x42, 0x08, 0x52, 0x61, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x48,
	0x02, 0x50, 0x01, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x66, 0x6c, 0x79, 0x74, 0x65, 0x6f, 0x72, 0x67, 0x2f, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x70, 0x62, 0x2d, 0x67, 0x6f, 0x2f, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64,
	0x6c, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0xa2, 0x02, 0x03, 0x46, 0x50, 0x58, 0xaa,
	0x02, 0x10, 0x46, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x73, 0xca, 0x02, 0x10, 0x46, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x5c, 0x50, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x73, 0xe2, 0x02, 0x1c, 0x46, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c,
	0x5c, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x11, 0x46, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x3a,
	0x3a, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_flyteidl_plugins_ray_proto_rawDescOnce sync.Once
	file_flyteidl_plugins_ray_proto_rawDescData = file_flyteidl_plugins_ray_proto_rawDesc
)

func file_flyteidl_plugins_ray_proto_rawDescGZIP() []byte {
	file_flyteidl_plugins_ray_proto_rawDescOnce.Do(func() {
		file_flyteidl_plugins_ray_proto_rawDescData = protoimpl.X.CompressGZIP(file_flyteidl_plugins_ray_proto_rawDescData)
	})
	return file_flyteidl_plugins_ray_proto_rawDescData
}

var file_flyteidl_plugins_ray_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_flyteidl_plugins_ray_proto_goTypes = []interface{}{
	(*RayJob)(nil),          // 0: flyteidl.plugins.RayJob
	(*RayCluster)(nil),      // 1: flyteidl.plugins.RayCluster
	(*HeadGroupSpec)(nil),   // 2: flyteidl.plugins.HeadGroupSpec
	(*WorkerGroupSpec)(nil), // 3: flyteidl.plugins.WorkerGroupSpec
	nil,                     // 4: flyteidl.plugins.HeadGroupSpec.RayStartParamsEntry
	nil,                     // 5: flyteidl.plugins.WorkerGroupSpec.RayStartParamsEntry
}
var file_flyteidl_plugins_ray_proto_depIdxs = []int32{
	1, // 0: flyteidl.plugins.RayJob.ray_cluster:type_name -> flyteidl.plugins.RayCluster
	2, // 1: flyteidl.plugins.RayCluster.head_group_spec:type_name -> flyteidl.plugins.HeadGroupSpec
	3, // 2: flyteidl.plugins.RayCluster.worker_group_spec:type_name -> flyteidl.plugins.WorkerGroupSpec
	4, // 3: flyteidl.plugins.HeadGroupSpec.ray_start_params:type_name -> flyteidl.plugins.HeadGroupSpec.RayStartParamsEntry
	5, // 4: flyteidl.plugins.WorkerGroupSpec.ray_start_params:type_name -> flyteidl.plugins.WorkerGroupSpec.RayStartParamsEntry
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_flyteidl_plugins_ray_proto_init() }
func file_flyteidl_plugins_ray_proto_init() {
	if File_flyteidl_plugins_ray_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_flyteidl_plugins_ray_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RayJob); i {
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
		file_flyteidl_plugins_ray_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RayCluster); i {
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
		file_flyteidl_plugins_ray_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeadGroupSpec); i {
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
		file_flyteidl_plugins_ray_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkerGroupSpec); i {
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
			RawDescriptor: file_flyteidl_plugins_ray_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_flyteidl_plugins_ray_proto_goTypes,
		DependencyIndexes: file_flyteidl_plugins_ray_proto_depIdxs,
		MessageInfos:      file_flyteidl_plugins_ray_proto_msgTypes,
	}.Build()
	File_flyteidl_plugins_ray_proto = out.File
	file_flyteidl_plugins_ray_proto_rawDesc = nil
	file_flyteidl_plugins_ray_proto_goTypes = nil
	file_flyteidl_plugins_ray_proto_depIdxs = nil
}
