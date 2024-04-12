// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	core "github.com/flyteorg/flyte/flyteplugins/go/tasks/pluginmachinery/core"
	executors "github.com/flyteorg/flyte/flytepropeller/pkg/controller/executors"

	interfaces "github.com/flyteorg/flyte/flytepropeller/pkg/controller/nodes/interfaces"

	io "github.com/flyteorg/flyte/flyteplugins/go/tasks/pluginmachinery/io"

	ioutils "github.com/flyteorg/flyte/flyteplugins/go/tasks/pluginmachinery/ioutils"

	mock "github.com/stretchr/testify/mock"

	storage "github.com/flyteorg/flyte/flytestdlib/storage"

	v1alpha1 "github.com/flyteorg/flyte/flytepropeller/pkg/apis/flyteworkflow/v1alpha1"
)

// NodeExecutionContext is an autogenerated mock type for the NodeExecutionContext type
type NodeExecutionContext struct {
	mock.Mock
}

type NodeExecutionContext_ContextualNodeLookup struct {
	*mock.Call
}

func (_m NodeExecutionContext_ContextualNodeLookup) Return(_a0 executors.NodeLookup) *NodeExecutionContext_ContextualNodeLookup {
	return &NodeExecutionContext_ContextualNodeLookup{Call: _m.Call.Return(_a0)}
}

func (_m *NodeExecutionContext) OnContextualNodeLookup() *NodeExecutionContext_ContextualNodeLookup {
	c_call := _m.On("ContextualNodeLookup")
	return &NodeExecutionContext_ContextualNodeLookup{Call: c_call}
}

func (_m *NodeExecutionContext) OnContextualNodeLookupMatch(matchers ...interface{}) *NodeExecutionContext_ContextualNodeLookup {
	c_call := _m.On("ContextualNodeLookup", matchers...)
	return &NodeExecutionContext_ContextualNodeLookup{Call: c_call}
}

// ContextualNodeLookup provides a mock function with given fields:
func (_m *NodeExecutionContext) ContextualNodeLookup() executors.NodeLookup {
	ret := _m.Called()

	var r0 executors.NodeLookup
	if rf, ok := ret.Get(0).(func() executors.NodeLookup); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(executors.NodeLookup)
		}
	}

	return r0
}

type NodeExecutionContext_CurrentAttempt struct {
	*mock.Call
}

func (_m NodeExecutionContext_CurrentAttempt) Return(_a0 uint32) *NodeExecutionContext_CurrentAttempt {
	return &NodeExecutionContext_CurrentAttempt{Call: _m.Call.Return(_a0)}
}

func (_m *NodeExecutionContext) OnCurrentAttempt() *NodeExecutionContext_CurrentAttempt {
	c_call := _m.On("CurrentAttempt")
	return &NodeExecutionContext_CurrentAttempt{Call: c_call}
}

func (_m *NodeExecutionContext) OnCurrentAttemptMatch(matchers ...interface{}) *NodeExecutionContext_CurrentAttempt {
	c_call := _m.On("CurrentAttempt", matchers...)
	return &NodeExecutionContext_CurrentAttempt{Call: c_call}
}

// CurrentAttempt provides a mock function with given fields:
func (_m *NodeExecutionContext) CurrentAttempt() uint32 {
	ret := _m.Called()

	var r0 uint32
	if rf, ok := ret.Get(0).(func() uint32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint32)
	}

	return r0
}

type NodeExecutionContext_DataStore struct {
	*mock.Call
}

func (_m NodeExecutionContext_DataStore) Return(_a0 *storage.DataStore) *NodeExecutionContext_DataStore {
	return &NodeExecutionContext_DataStore{Call: _m.Call.Return(_a0)}
}

func (_m *NodeExecutionContext) OnDataStore() *NodeExecutionContext_DataStore {
	c_call := _m.On("DataStore")
	return &NodeExecutionContext_DataStore{Call: c_call}
}

func (_m *NodeExecutionContext) OnDataStoreMatch(matchers ...interface{}) *NodeExecutionContext_DataStore {
	c_call := _m.On("DataStore", matchers...)
	return &NodeExecutionContext_DataStore{Call: c_call}
}

// DataStore provides a mock function with given fields:
func (_m *NodeExecutionContext) DataStore() *storage.DataStore {
	ret := _m.Called()

	var r0 *storage.DataStore
	if rf, ok := ret.Get(0).(func() *storage.DataStore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*storage.DataStore)
		}
	}

	return r0
}

type NodeExecutionContext_EnqueueOwnerFunc struct {
	*mock.Call
}

func (_m NodeExecutionContext_EnqueueOwnerFunc) Return(_a0 func() error) *NodeExecutionContext_EnqueueOwnerFunc {
	return &NodeExecutionContext_EnqueueOwnerFunc{Call: _m.Call.Return(_a0)}
}

func (_m *NodeExecutionContext) OnEnqueueOwnerFunc() *NodeExecutionContext_EnqueueOwnerFunc {
	c_call := _m.On("EnqueueOwnerFunc")
	return &NodeExecutionContext_EnqueueOwnerFunc{Call: c_call}
}

func (_m *NodeExecutionContext) OnEnqueueOwnerFuncMatch(matchers ...interface{}) *NodeExecutionContext_EnqueueOwnerFunc {
	c_call := _m.On("EnqueueOwnerFunc", matchers...)
	return &NodeExecutionContext_EnqueueOwnerFunc{Call: c_call}
}

// EnqueueOwnerFunc provides a mock function with given fields:
func (_m *NodeExecutionContext) EnqueueOwnerFunc() func() error {
	ret := _m.Called()

	var r0 func() error
	if rf, ok := ret.Get(0).(func() func() error); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(func() error)
		}
	}

	return r0
}

type NodeExecutionContext_EventsRecorder struct {
	*mock.Call
}

func (_m NodeExecutionContext_EventsRecorder) Return(_a0 interfaces.EventRecorder) *NodeExecutionContext_EventsRecorder {
	return &NodeExecutionContext_EventsRecorder{Call: _m.Call.Return(_a0)}
}

func (_m *NodeExecutionContext) OnEventsRecorder() *NodeExecutionContext_EventsRecorder {
	c_call := _m.On("EventsRecorder")
	return &NodeExecutionContext_EventsRecorder{Call: c_call}
}

func (_m *NodeExecutionContext) OnEventsRecorderMatch(matchers ...interface{}) *NodeExecutionContext_EventsRecorder {
	c_call := _m.On("EventsRecorder", matchers...)
	return &NodeExecutionContext_EventsRecorder{Call: c_call}
}

// EventsRecorder provides a mock function with given fields:
func (_m *NodeExecutionContext) EventsRecorder() interfaces.EventRecorder {
	ret := _m.Called()

	var r0 interfaces.EventRecorder
	if rf, ok := ret.Get(0).(func() interfaces.EventRecorder); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interfaces.EventRecorder)
		}
	}

	return r0
}

type NodeExecutionContext_ExecutionContext struct {
	*mock.Call
}

func (_m NodeExecutionContext_ExecutionContext) Return(_a0 executors.ExecutionContext) *NodeExecutionContext_ExecutionContext {
	return &NodeExecutionContext_ExecutionContext{Call: _m.Call.Return(_a0)}
}

func (_m *NodeExecutionContext) OnExecutionContext() *NodeExecutionContext_ExecutionContext {
	c_call := _m.On("ExecutionContext")
	return &NodeExecutionContext_ExecutionContext{Call: c_call}
}

func (_m *NodeExecutionContext) OnExecutionContextMatch(matchers ...interface{}) *NodeExecutionContext_ExecutionContext {
	c_call := _m.On("ExecutionContext", matchers...)
	return &NodeExecutionContext_ExecutionContext{Call: c_call}
}

// ExecutionContext provides a mock function with given fields:
func (_m *NodeExecutionContext) ExecutionContext() executors.ExecutionContext {
	ret := _m.Called()

	var r0 executors.ExecutionContext
	if rf, ok := ret.Get(0).(func() executors.ExecutionContext); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(executors.ExecutionContext)
		}
	}

	return r0
}

type NodeExecutionContext_GetExecutionEnvClient struct {
	*mock.Call
}

func (_m NodeExecutionContext_GetExecutionEnvClient) Return(_a0 core.ExecutionEnvClient) *NodeExecutionContext_GetExecutionEnvClient {
	return &NodeExecutionContext_GetExecutionEnvClient{Call: _m.Call.Return(_a0)}
}

func (_m *NodeExecutionContext) OnGetExecutionEnvClient() *NodeExecutionContext_GetExecutionEnvClient {
	c_call := _m.On("GetExecutionEnvClient")
	return &NodeExecutionContext_GetExecutionEnvClient{Call: c_call}
}

func (_m *NodeExecutionContext) OnGetExecutionEnvClientMatch(matchers ...interface{}) *NodeExecutionContext_GetExecutionEnvClient {
	c_call := _m.On("GetExecutionEnvClient", matchers...)
	return &NodeExecutionContext_GetExecutionEnvClient{Call: c_call}
}

// GetExecutionEnvClient provides a mock function with given fields:
func (_m *NodeExecutionContext) GetExecutionEnvClient() core.ExecutionEnvClient {
	ret := _m.Called()

	var r0 core.ExecutionEnvClient
	if rf, ok := ret.Get(0).(func() core.ExecutionEnvClient); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(core.ExecutionEnvClient)
		}
	}

	return r0
}

type NodeExecutionContext_InputReader struct {
	*mock.Call
}

func (_m NodeExecutionContext_InputReader) Return(_a0 io.InputReader) *NodeExecutionContext_InputReader {
	return &NodeExecutionContext_InputReader{Call: _m.Call.Return(_a0)}
}

func (_m *NodeExecutionContext) OnInputReader() *NodeExecutionContext_InputReader {
	c_call := _m.On("InputReader")
	return &NodeExecutionContext_InputReader{Call: c_call}
}

func (_m *NodeExecutionContext) OnInputReaderMatch(matchers ...interface{}) *NodeExecutionContext_InputReader {
	c_call := _m.On("InputReader", matchers...)
	return &NodeExecutionContext_InputReader{Call: c_call}
}

// InputReader provides a mock function with given fields:
func (_m *NodeExecutionContext) InputReader() io.InputReader {
	ret := _m.Called()

	var r0 io.InputReader
	if rf, ok := ret.Get(0).(func() io.InputReader); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.InputReader)
		}
	}

	return r0
}

type NodeExecutionContext_MaxDatasetSizeBytes struct {
	*mock.Call
}

func (_m NodeExecutionContext_MaxDatasetSizeBytes) Return(_a0 int64) *NodeExecutionContext_MaxDatasetSizeBytes {
	return &NodeExecutionContext_MaxDatasetSizeBytes{Call: _m.Call.Return(_a0)}
}

func (_m *NodeExecutionContext) OnMaxDatasetSizeBytes() *NodeExecutionContext_MaxDatasetSizeBytes {
	c_call := _m.On("MaxDatasetSizeBytes")
	return &NodeExecutionContext_MaxDatasetSizeBytes{Call: c_call}
}

func (_m *NodeExecutionContext) OnMaxDatasetSizeBytesMatch(matchers ...interface{}) *NodeExecutionContext_MaxDatasetSizeBytes {
	c_call := _m.On("MaxDatasetSizeBytes", matchers...)
	return &NodeExecutionContext_MaxDatasetSizeBytes{Call: c_call}
}

// MaxDatasetSizeBytes provides a mock function with given fields:
func (_m *NodeExecutionContext) MaxDatasetSizeBytes() int64 {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

type NodeExecutionContext_Node struct {
	*mock.Call
}

func (_m NodeExecutionContext_Node) Return(_a0 v1alpha1.ExecutableNode) *NodeExecutionContext_Node {
	return &NodeExecutionContext_Node{Call: _m.Call.Return(_a0)}
}

func (_m *NodeExecutionContext) OnNode() *NodeExecutionContext_Node {
	c_call := _m.On("Node")
	return &NodeExecutionContext_Node{Call: c_call}
}

func (_m *NodeExecutionContext) OnNodeMatch(matchers ...interface{}) *NodeExecutionContext_Node {
	c_call := _m.On("Node", matchers...)
	return &NodeExecutionContext_Node{Call: c_call}
}

// Node provides a mock function with given fields:
func (_m *NodeExecutionContext) Node() v1alpha1.ExecutableNode {
	ret := _m.Called()

	var r0 v1alpha1.ExecutableNode
	if rf, ok := ret.Get(0).(func() v1alpha1.ExecutableNode); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.ExecutableNode)
		}
	}

	return r0
}

type NodeExecutionContext_NodeExecutionMetadata struct {
	*mock.Call
}

func (_m NodeExecutionContext_NodeExecutionMetadata) Return(_a0 interfaces.NodeExecutionMetadata) *NodeExecutionContext_NodeExecutionMetadata {
	return &NodeExecutionContext_NodeExecutionMetadata{Call: _m.Call.Return(_a0)}
}

func (_m *NodeExecutionContext) OnNodeExecutionMetadata() *NodeExecutionContext_NodeExecutionMetadata {
	c_call := _m.On("NodeExecutionMetadata")
	return &NodeExecutionContext_NodeExecutionMetadata{Call: c_call}
}

func (_m *NodeExecutionContext) OnNodeExecutionMetadataMatch(matchers ...interface{}) *NodeExecutionContext_NodeExecutionMetadata {
	c_call := _m.On("NodeExecutionMetadata", matchers...)
	return &NodeExecutionContext_NodeExecutionMetadata{Call: c_call}
}

// NodeExecutionMetadata provides a mock function with given fields:
func (_m *NodeExecutionContext) NodeExecutionMetadata() interfaces.NodeExecutionMetadata {
	ret := _m.Called()

	var r0 interfaces.NodeExecutionMetadata
	if rf, ok := ret.Get(0).(func() interfaces.NodeExecutionMetadata); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interfaces.NodeExecutionMetadata)
		}
	}

	return r0
}

type NodeExecutionContext_NodeID struct {
	*mock.Call
}

func (_m NodeExecutionContext_NodeID) Return(_a0 string) *NodeExecutionContext_NodeID {
	return &NodeExecutionContext_NodeID{Call: _m.Call.Return(_a0)}
}

func (_m *NodeExecutionContext) OnNodeID() *NodeExecutionContext_NodeID {
	c_call := _m.On("NodeID")
	return &NodeExecutionContext_NodeID{Call: c_call}
}

func (_m *NodeExecutionContext) OnNodeIDMatch(matchers ...interface{}) *NodeExecutionContext_NodeID {
	c_call := _m.On("NodeID", matchers...)
	return &NodeExecutionContext_NodeID{Call: c_call}
}

// NodeID provides a mock function with given fields:
func (_m *NodeExecutionContext) NodeID() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type NodeExecutionContext_NodeStateReader struct {
	*mock.Call
}

func (_m NodeExecutionContext_NodeStateReader) Return(_a0 interfaces.NodeStateReader) *NodeExecutionContext_NodeStateReader {
	return &NodeExecutionContext_NodeStateReader{Call: _m.Call.Return(_a0)}
}

func (_m *NodeExecutionContext) OnNodeStateReader() *NodeExecutionContext_NodeStateReader {
	c_call := _m.On("NodeStateReader")
	return &NodeExecutionContext_NodeStateReader{Call: c_call}
}

func (_m *NodeExecutionContext) OnNodeStateReaderMatch(matchers ...interface{}) *NodeExecutionContext_NodeStateReader {
	c_call := _m.On("NodeStateReader", matchers...)
	return &NodeExecutionContext_NodeStateReader{Call: c_call}
}

// NodeStateReader provides a mock function with given fields:
func (_m *NodeExecutionContext) NodeStateReader() interfaces.NodeStateReader {
	ret := _m.Called()

	var r0 interfaces.NodeStateReader
	if rf, ok := ret.Get(0).(func() interfaces.NodeStateReader); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interfaces.NodeStateReader)
		}
	}

	return r0
}

type NodeExecutionContext_NodeStateWriter struct {
	*mock.Call
}

func (_m NodeExecutionContext_NodeStateWriter) Return(_a0 interfaces.NodeStateWriter) *NodeExecutionContext_NodeStateWriter {
	return &NodeExecutionContext_NodeStateWriter{Call: _m.Call.Return(_a0)}
}

func (_m *NodeExecutionContext) OnNodeStateWriter() *NodeExecutionContext_NodeStateWriter {
	c_call := _m.On("NodeStateWriter")
	return &NodeExecutionContext_NodeStateWriter{Call: c_call}
}

func (_m *NodeExecutionContext) OnNodeStateWriterMatch(matchers ...interface{}) *NodeExecutionContext_NodeStateWriter {
	c_call := _m.On("NodeStateWriter", matchers...)
	return &NodeExecutionContext_NodeStateWriter{Call: c_call}
}

// NodeStateWriter provides a mock function with given fields:
func (_m *NodeExecutionContext) NodeStateWriter() interfaces.NodeStateWriter {
	ret := _m.Called()

	var r0 interfaces.NodeStateWriter
	if rf, ok := ret.Get(0).(func() interfaces.NodeStateWriter); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interfaces.NodeStateWriter)
		}
	}

	return r0
}

type NodeExecutionContext_NodeStatus struct {
	*mock.Call
}

func (_m NodeExecutionContext_NodeStatus) Return(_a0 v1alpha1.ExecutableNodeStatus) *NodeExecutionContext_NodeStatus {
	return &NodeExecutionContext_NodeStatus{Call: _m.Call.Return(_a0)}
}

func (_m *NodeExecutionContext) OnNodeStatus() *NodeExecutionContext_NodeStatus {
	c_call := _m.On("NodeStatus")
	return &NodeExecutionContext_NodeStatus{Call: c_call}
}

func (_m *NodeExecutionContext) OnNodeStatusMatch(matchers ...interface{}) *NodeExecutionContext_NodeStatus {
	c_call := _m.On("NodeStatus", matchers...)
	return &NodeExecutionContext_NodeStatus{Call: c_call}
}

// NodeStatus provides a mock function with given fields:
func (_m *NodeExecutionContext) NodeStatus() v1alpha1.ExecutableNodeStatus {
	ret := _m.Called()

	var r0 v1alpha1.ExecutableNodeStatus
	if rf, ok := ret.Get(0).(func() v1alpha1.ExecutableNodeStatus); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.ExecutableNodeStatus)
		}
	}

	return r0
}

type NodeExecutionContext_OutputShardSelector struct {
	*mock.Call
}

func (_m NodeExecutionContext_OutputShardSelector) Return(_a0 ioutils.ShardSelector) *NodeExecutionContext_OutputShardSelector {
	return &NodeExecutionContext_OutputShardSelector{Call: _m.Call.Return(_a0)}
}

func (_m *NodeExecutionContext) OnOutputShardSelector() *NodeExecutionContext_OutputShardSelector {
	c_call := _m.On("OutputShardSelector")
	return &NodeExecutionContext_OutputShardSelector{Call: c_call}
}

func (_m *NodeExecutionContext) OnOutputShardSelectorMatch(matchers ...interface{}) *NodeExecutionContext_OutputShardSelector {
	c_call := _m.On("OutputShardSelector", matchers...)
	return &NodeExecutionContext_OutputShardSelector{Call: c_call}
}

// OutputShardSelector provides a mock function with given fields:
func (_m *NodeExecutionContext) OutputShardSelector() ioutils.ShardSelector {
	ret := _m.Called()

	var r0 ioutils.ShardSelector
	if rf, ok := ret.Get(0).(func() ioutils.ShardSelector); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ioutils.ShardSelector)
		}
	}

	return r0
}

type NodeExecutionContext_RawOutputPrefix struct {
	*mock.Call
}

func (_m NodeExecutionContext_RawOutputPrefix) Return(_a0 storage.DataReference) *NodeExecutionContext_RawOutputPrefix {
	return &NodeExecutionContext_RawOutputPrefix{Call: _m.Call.Return(_a0)}
}

func (_m *NodeExecutionContext) OnRawOutputPrefix() *NodeExecutionContext_RawOutputPrefix {
	c_call := _m.On("RawOutputPrefix")
	return &NodeExecutionContext_RawOutputPrefix{Call: c_call}
}

func (_m *NodeExecutionContext) OnRawOutputPrefixMatch(matchers ...interface{}) *NodeExecutionContext_RawOutputPrefix {
	c_call := _m.On("RawOutputPrefix", matchers...)
	return &NodeExecutionContext_RawOutputPrefix{Call: c_call}
}

// RawOutputPrefix provides a mock function with given fields:
func (_m *NodeExecutionContext) RawOutputPrefix() storage.DataReference {
	ret := _m.Called()

	var r0 storage.DataReference
	if rf, ok := ret.Get(0).(func() storage.DataReference); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(storage.DataReference)
	}

	return r0
}

type NodeExecutionContext_TaskReader struct {
	*mock.Call
}

func (_m NodeExecutionContext_TaskReader) Return(_a0 interfaces.TaskReader) *NodeExecutionContext_TaskReader {
	return &NodeExecutionContext_TaskReader{Call: _m.Call.Return(_a0)}
}

func (_m *NodeExecutionContext) OnTaskReader() *NodeExecutionContext_TaskReader {
	c_call := _m.On("TaskReader")
	return &NodeExecutionContext_TaskReader{Call: c_call}
}

func (_m *NodeExecutionContext) OnTaskReaderMatch(matchers ...interface{}) *NodeExecutionContext_TaskReader {
	c_call := _m.On("TaskReader", matchers...)
	return &NodeExecutionContext_TaskReader{Call: c_call}
}

// TaskReader provides a mock function with given fields:
func (_m *NodeExecutionContext) TaskReader() interfaces.TaskReader {
	ret := _m.Called()

	var r0 interfaces.TaskReader
	if rf, ok := ret.Get(0).(func() interfaces.TaskReader); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interfaces.TaskReader)
		}
	}

	return r0
}
