// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	core "github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/core"
	mock "github.com/stretchr/testify/mock"

	time "time"

	v1 "k8s.io/api/core/v1"

	v1alpha1 "github.com/flyteorg/flyte/flytepropeller/pkg/apis/flyteworkflow/v1alpha1"
)

// ExecutableNode is an autogenerated mock type for the ExecutableNode type
type ExecutableNode struct {
	mock.Mock
}

type ExecutableNode_GetActiveDeadline struct {
	*mock.Call
}

func (_m ExecutableNode_GetActiveDeadline) Return(_a0 *time.Duration) *ExecutableNode_GetActiveDeadline {
	return &ExecutableNode_GetActiveDeadline{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableNode) OnGetActiveDeadline() *ExecutableNode_GetActiveDeadline {
	c_call := _m.On("GetActiveDeadline")
	return &ExecutableNode_GetActiveDeadline{Call: c_call}
}

func (_m *ExecutableNode) OnGetActiveDeadlineMatch(matchers ...interface{}) *ExecutableNode_GetActiveDeadline {
	c_call := _m.On("GetActiveDeadline", matchers...)
	return &ExecutableNode_GetActiveDeadline{Call: c_call}
}

// GetActiveDeadline provides a mock function with given fields:
func (_m *ExecutableNode) GetActiveDeadline() *time.Duration {
	ret := _m.Called()

	var r0 *time.Duration
	if rf, ok := ret.Get(0).(func() *time.Duration); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*time.Duration)
		}
	}

	return r0
}

type ExecutableNode_GetArrayNode struct {
	*mock.Call
}

func (_m ExecutableNode_GetArrayNode) Return(_a0 v1alpha1.ExecutableArrayNode) *ExecutableNode_GetArrayNode {
	return &ExecutableNode_GetArrayNode{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableNode) OnGetArrayNode() *ExecutableNode_GetArrayNode {
	c_call := _m.On("GetArrayNode")
	return &ExecutableNode_GetArrayNode{Call: c_call}
}

func (_m *ExecutableNode) OnGetArrayNodeMatch(matchers ...interface{}) *ExecutableNode_GetArrayNode {
	c_call := _m.On("GetArrayNode", matchers...)
	return &ExecutableNode_GetArrayNode{Call: c_call}
}

// GetArrayNode provides a mock function with given fields:
func (_m *ExecutableNode) GetArrayNode() v1alpha1.ExecutableArrayNode {
	ret := _m.Called()

	var r0 v1alpha1.ExecutableArrayNode
	if rf, ok := ret.Get(0).(func() v1alpha1.ExecutableArrayNode); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.ExecutableArrayNode)
		}
	}

	return r0
}

type ExecutableNode_GetBranchNode struct {
	*mock.Call
}

func (_m ExecutableNode_GetBranchNode) Return(_a0 v1alpha1.ExecutableBranchNode) *ExecutableNode_GetBranchNode {
	return &ExecutableNode_GetBranchNode{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableNode) OnGetBranchNode() *ExecutableNode_GetBranchNode {
	c_call := _m.On("GetBranchNode")
	return &ExecutableNode_GetBranchNode{Call: c_call}
}

func (_m *ExecutableNode) OnGetBranchNodeMatch(matchers ...interface{}) *ExecutableNode_GetBranchNode {
	c_call := _m.On("GetBranchNode", matchers...)
	return &ExecutableNode_GetBranchNode{Call: c_call}
}

// GetBranchNode provides a mock function with given fields:
func (_m *ExecutableNode) GetBranchNode() v1alpha1.ExecutableBranchNode {
	ret := _m.Called()

	var r0 v1alpha1.ExecutableBranchNode
	if rf, ok := ret.Get(0).(func() v1alpha1.ExecutableBranchNode); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.ExecutableBranchNode)
		}
	}

	return r0
}

type ExecutableNode_GetConfig struct {
	*mock.Call
}

func (_m ExecutableNode_GetConfig) Return(_a0 *v1.ConfigMap) *ExecutableNode_GetConfig {
	return &ExecutableNode_GetConfig{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableNode) OnGetConfig() *ExecutableNode_GetConfig {
	c_call := _m.On("GetConfig")
	return &ExecutableNode_GetConfig{Call: c_call}
}

func (_m *ExecutableNode) OnGetConfigMatch(matchers ...interface{}) *ExecutableNode_GetConfig {
	c_call := _m.On("GetConfig", matchers...)
	return &ExecutableNode_GetConfig{Call: c_call}
}

// GetConfig provides a mock function with given fields:
func (_m *ExecutableNode) GetConfig() *v1.ConfigMap {
	ret := _m.Called()

	var r0 *v1.ConfigMap
	if rf, ok := ret.Get(0).(func() *v1.ConfigMap); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.ConfigMap)
		}
	}

	return r0
}

type ExecutableNode_GetExecutionDeadline struct {
	*mock.Call
}

func (_m ExecutableNode_GetExecutionDeadline) Return(_a0 *time.Duration) *ExecutableNode_GetExecutionDeadline {
	return &ExecutableNode_GetExecutionDeadline{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableNode) OnGetExecutionDeadline() *ExecutableNode_GetExecutionDeadline {
	c_call := _m.On("GetExecutionDeadline")
	return &ExecutableNode_GetExecutionDeadline{Call: c_call}
}

func (_m *ExecutableNode) OnGetExecutionDeadlineMatch(matchers ...interface{}) *ExecutableNode_GetExecutionDeadline {
	c_call := _m.On("GetExecutionDeadline", matchers...)
	return &ExecutableNode_GetExecutionDeadline{Call: c_call}
}

// GetExecutionDeadline provides a mock function with given fields:
func (_m *ExecutableNode) GetExecutionDeadline() *time.Duration {
	ret := _m.Called()

	var r0 *time.Duration
	if rf, ok := ret.Get(0).(func() *time.Duration); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*time.Duration)
		}
	}

	return r0
}

type ExecutableNode_GetExtendedResources struct {
	*mock.Call
}

func (_m ExecutableNode_GetExtendedResources) Return(_a0 *core.ExtendedResources) *ExecutableNode_GetExtendedResources {
	return &ExecutableNode_GetExtendedResources{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableNode) OnGetExtendedResources() *ExecutableNode_GetExtendedResources {
	c_call := _m.On("GetExtendedResources")
	return &ExecutableNode_GetExtendedResources{Call: c_call}
}

func (_m *ExecutableNode) OnGetExtendedResourcesMatch(matchers ...interface{}) *ExecutableNode_GetExtendedResources {
	c_call := _m.On("GetExtendedResources", matchers...)
	return &ExecutableNode_GetExtendedResources{Call: c_call}
}

// GetExtendedResources provides a mock function with given fields:
func (_m *ExecutableNode) GetExtendedResources() *core.ExtendedResources {
	ret := _m.Called()

	var r0 *core.ExtendedResources
	if rf, ok := ret.Get(0).(func() *core.ExtendedResources); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.ExtendedResources)
		}
	}

	return r0
}

type ExecutableNode_GetGateNode struct {
	*mock.Call
}

func (_m ExecutableNode_GetGateNode) Return(_a0 v1alpha1.ExecutableGateNode) *ExecutableNode_GetGateNode {
	return &ExecutableNode_GetGateNode{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableNode) OnGetGateNode() *ExecutableNode_GetGateNode {
	c_call := _m.On("GetGateNode")
	return &ExecutableNode_GetGateNode{Call: c_call}
}

func (_m *ExecutableNode) OnGetGateNodeMatch(matchers ...interface{}) *ExecutableNode_GetGateNode {
	c_call := _m.On("GetGateNode", matchers...)
	return &ExecutableNode_GetGateNode{Call: c_call}
}

// GetGateNode provides a mock function with given fields:
func (_m *ExecutableNode) GetGateNode() v1alpha1.ExecutableGateNode {
	ret := _m.Called()

	var r0 v1alpha1.ExecutableGateNode
	if rf, ok := ret.Get(0).(func() v1alpha1.ExecutableGateNode); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.ExecutableGateNode)
		}
	}

	return r0
}

type ExecutableNode_GetID struct {
	*mock.Call
}

func (_m ExecutableNode_GetID) Return(_a0 string) *ExecutableNode_GetID {
	return &ExecutableNode_GetID{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableNode) OnGetID() *ExecutableNode_GetID {
	c_call := _m.On("GetID")
	return &ExecutableNode_GetID{Call: c_call}
}

func (_m *ExecutableNode) OnGetIDMatch(matchers ...interface{}) *ExecutableNode_GetID {
	c_call := _m.On("GetID", matchers...)
	return &ExecutableNode_GetID{Call: c_call}
}

// GetID provides a mock function with given fields:
func (_m *ExecutableNode) GetID() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type ExecutableNode_GetInputBindings struct {
	*mock.Call
}

func (_m ExecutableNode_GetInputBindings) Return(_a0 []*v1alpha1.Binding) *ExecutableNode_GetInputBindings {
	return &ExecutableNode_GetInputBindings{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableNode) OnGetInputBindings() *ExecutableNode_GetInputBindings {
	c_call := _m.On("GetInputBindings")
	return &ExecutableNode_GetInputBindings{Call: c_call}
}

func (_m *ExecutableNode) OnGetInputBindingsMatch(matchers ...interface{}) *ExecutableNode_GetInputBindings {
	c_call := _m.On("GetInputBindings", matchers...)
	return &ExecutableNode_GetInputBindings{Call: c_call}
}

// GetInputBindings provides a mock function with given fields:
func (_m *ExecutableNode) GetInputBindings() []*v1alpha1.Binding {
	ret := _m.Called()

	var r0 []*v1alpha1.Binding
	if rf, ok := ret.Get(0).(func() []*v1alpha1.Binding); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v1alpha1.Binding)
		}
	}

	return r0
}

type ExecutableNode_GetKind struct {
	*mock.Call
}

func (_m ExecutableNode_GetKind) Return(_a0 v1alpha1.NodeKind) *ExecutableNode_GetKind {
	return &ExecutableNode_GetKind{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableNode) OnGetKind() *ExecutableNode_GetKind {
	c_call := _m.On("GetKind")
	return &ExecutableNode_GetKind{Call: c_call}
}

func (_m *ExecutableNode) OnGetKindMatch(matchers ...interface{}) *ExecutableNode_GetKind {
	c_call := _m.On("GetKind", matchers...)
	return &ExecutableNode_GetKind{Call: c_call}
}

// GetKind provides a mock function with given fields:
func (_m *ExecutableNode) GetKind() v1alpha1.NodeKind {
	ret := _m.Called()

	var r0 v1alpha1.NodeKind
	if rf, ok := ret.Get(0).(func() v1alpha1.NodeKind); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(v1alpha1.NodeKind)
	}

	return r0
}

type ExecutableNode_GetName struct {
	*mock.Call
}

func (_m ExecutableNode_GetName) Return(_a0 string) *ExecutableNode_GetName {
	return &ExecutableNode_GetName{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableNode) OnGetName() *ExecutableNode_GetName {
	c_call := _m.On("GetName")
	return &ExecutableNode_GetName{Call: c_call}
}

func (_m *ExecutableNode) OnGetNameMatch(matchers ...interface{}) *ExecutableNode_GetName {
	c_call := _m.On("GetName", matchers...)
	return &ExecutableNode_GetName{Call: c_call}
}

// GetName provides a mock function with given fields:
func (_m *ExecutableNode) GetName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type ExecutableNode_GetOutputAlias struct {
	*mock.Call
}

func (_m ExecutableNode_GetOutputAlias) Return(_a0 []v1alpha1.Alias) *ExecutableNode_GetOutputAlias {
	return &ExecutableNode_GetOutputAlias{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableNode) OnGetOutputAlias() *ExecutableNode_GetOutputAlias {
	c_call := _m.On("GetOutputAlias")
	return &ExecutableNode_GetOutputAlias{Call: c_call}
}

func (_m *ExecutableNode) OnGetOutputAliasMatch(matchers ...interface{}) *ExecutableNode_GetOutputAlias {
	c_call := _m.On("GetOutputAlias", matchers...)
	return &ExecutableNode_GetOutputAlias{Call: c_call}
}

// GetOutputAlias provides a mock function with given fields:
func (_m *ExecutableNode) GetOutputAlias() []v1alpha1.Alias {
	ret := _m.Called()

	var r0 []v1alpha1.Alias
	if rf, ok := ret.Get(0).(func() []v1alpha1.Alias); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]v1alpha1.Alias)
		}
	}

	return r0
}

type ExecutableNode_GetResources struct {
	*mock.Call
}

func (_m ExecutableNode_GetResources) Return(_a0 *v1.ResourceRequirements) *ExecutableNode_GetResources {
	return &ExecutableNode_GetResources{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableNode) OnGetResources() *ExecutableNode_GetResources {
	c_call := _m.On("GetResources")
	return &ExecutableNode_GetResources{Call: c_call}
}

func (_m *ExecutableNode) OnGetResourcesMatch(matchers ...interface{}) *ExecutableNode_GetResources {
	c_call := _m.On("GetResources", matchers...)
	return &ExecutableNode_GetResources{Call: c_call}
}

// GetResources provides a mock function with given fields:
func (_m *ExecutableNode) GetResources() *v1.ResourceRequirements {
	ret := _m.Called()

	var r0 *v1.ResourceRequirements
	if rf, ok := ret.Get(0).(func() *v1.ResourceRequirements); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.ResourceRequirements)
		}
	}

	return r0
}

type ExecutableNode_GetRetryStrategy struct {
	*mock.Call
}

func (_m ExecutableNode_GetRetryStrategy) Return(_a0 *v1alpha1.RetryStrategy) *ExecutableNode_GetRetryStrategy {
	return &ExecutableNode_GetRetryStrategy{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableNode) OnGetRetryStrategy() *ExecutableNode_GetRetryStrategy {
	c_call := _m.On("GetRetryStrategy")
	return &ExecutableNode_GetRetryStrategy{Call: c_call}
}

func (_m *ExecutableNode) OnGetRetryStrategyMatch(matchers ...interface{}) *ExecutableNode_GetRetryStrategy {
	c_call := _m.On("GetRetryStrategy", matchers...)
	return &ExecutableNode_GetRetryStrategy{Call: c_call}
}

// GetRetryStrategy provides a mock function with given fields:
func (_m *ExecutableNode) GetRetryStrategy() *v1alpha1.RetryStrategy {
	ret := _m.Called()

	var r0 *v1alpha1.RetryStrategy
	if rf, ok := ret.Get(0).(func() *v1alpha1.RetryStrategy); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.RetryStrategy)
		}
	}

	return r0
}

type ExecutableNode_GetTaskID struct {
	*mock.Call
}

func (_m ExecutableNode_GetTaskID) Return(_a0 *string) *ExecutableNode_GetTaskID {
	return &ExecutableNode_GetTaskID{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableNode) OnGetTaskID() *ExecutableNode_GetTaskID {
	c_call := _m.On("GetTaskID")
	return &ExecutableNode_GetTaskID{Call: c_call}
}

func (_m *ExecutableNode) OnGetTaskIDMatch(matchers ...interface{}) *ExecutableNode_GetTaskID {
	c_call := _m.On("GetTaskID", matchers...)
	return &ExecutableNode_GetTaskID{Call: c_call}
}

// GetTaskID provides a mock function with given fields:
func (_m *ExecutableNode) GetTaskID() *string {
	ret := _m.Called()

	var r0 *string
	if rf, ok := ret.Get(0).(func() *string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*string)
		}
	}

	return r0
}

type ExecutableNode_GetWorkflowNode struct {
	*mock.Call
}

func (_m ExecutableNode_GetWorkflowNode) Return(_a0 v1alpha1.ExecutableWorkflowNode) *ExecutableNode_GetWorkflowNode {
	return &ExecutableNode_GetWorkflowNode{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableNode) OnGetWorkflowNode() *ExecutableNode_GetWorkflowNode {
	c_call := _m.On("GetWorkflowNode")
	return &ExecutableNode_GetWorkflowNode{Call: c_call}
}

func (_m *ExecutableNode) OnGetWorkflowNodeMatch(matchers ...interface{}) *ExecutableNode_GetWorkflowNode {
	c_call := _m.On("GetWorkflowNode", matchers...)
	return &ExecutableNode_GetWorkflowNode{Call: c_call}
}

// GetWorkflowNode provides a mock function with given fields:
func (_m *ExecutableNode) GetWorkflowNode() v1alpha1.ExecutableWorkflowNode {
	ret := _m.Called()

	var r0 v1alpha1.ExecutableWorkflowNode
	if rf, ok := ret.Get(0).(func() v1alpha1.ExecutableWorkflowNode); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.ExecutableWorkflowNode)
		}
	}

	return r0
}

type ExecutableNode_IsEndNode struct {
	*mock.Call
}

func (_m ExecutableNode_IsEndNode) Return(_a0 bool) *ExecutableNode_IsEndNode {
	return &ExecutableNode_IsEndNode{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableNode) OnIsEndNode() *ExecutableNode_IsEndNode {
	c_call := _m.On("IsEndNode")
	return &ExecutableNode_IsEndNode{Call: c_call}
}

func (_m *ExecutableNode) OnIsEndNodeMatch(matchers ...interface{}) *ExecutableNode_IsEndNode {
	c_call := _m.On("IsEndNode", matchers...)
	return &ExecutableNode_IsEndNode{Call: c_call}
}

// IsEndNode provides a mock function with given fields:
func (_m *ExecutableNode) IsEndNode() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

type ExecutableNode_IsInterruptible struct {
	*mock.Call
}

func (_m ExecutableNode_IsInterruptible) Return(_a0 *bool) *ExecutableNode_IsInterruptible {
	return &ExecutableNode_IsInterruptible{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableNode) OnIsInterruptible() *ExecutableNode_IsInterruptible {
	c_call := _m.On("IsInterruptible")
	return &ExecutableNode_IsInterruptible{Call: c_call}
}

func (_m *ExecutableNode) OnIsInterruptibleMatch(matchers ...interface{}) *ExecutableNode_IsInterruptible {
	c_call := _m.On("IsInterruptible", matchers...)
	return &ExecutableNode_IsInterruptible{Call: c_call}
}

// IsInterruptible provides a mock function with given fields:
func (_m *ExecutableNode) IsInterruptible() *bool {
	ret := _m.Called()

	var r0 *bool
	if rf, ok := ret.Get(0).(func() *bool); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bool)
		}
	}

	return r0
}

type ExecutableNode_IsStartNode struct {
	*mock.Call
}

func (_m ExecutableNode_IsStartNode) Return(_a0 bool) *ExecutableNode_IsStartNode {
	return &ExecutableNode_IsStartNode{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableNode) OnIsStartNode() *ExecutableNode_IsStartNode {
	c_call := _m.On("IsStartNode")
	return &ExecutableNode_IsStartNode{Call: c_call}
}

func (_m *ExecutableNode) OnIsStartNodeMatch(matchers ...interface{}) *ExecutableNode_IsStartNode {
	c_call := _m.On("IsStartNode", matchers...)
	return &ExecutableNode_IsStartNode{Call: c_call}
}

// IsStartNode provides a mock function with given fields:
func (_m *ExecutableNode) IsStartNode() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}
