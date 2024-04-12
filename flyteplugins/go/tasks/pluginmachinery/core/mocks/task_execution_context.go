// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	catalog "github.com/flyteorg/flyte/flyteplugins/go/tasks/pluginmachinery/catalog"
	core "github.com/flyteorg/flyte/flyteplugins/go/tasks/pluginmachinery/core"

	io "github.com/flyteorg/flyte/flyteplugins/go/tasks/pluginmachinery/io"

	mock "github.com/stretchr/testify/mock"

	storage "github.com/flyteorg/flyte/flytestdlib/storage"
)

// TaskExecutionContext is an autogenerated mock type for the TaskExecutionContext type
type TaskExecutionContext struct {
	mock.Mock
}

type TaskExecutionContext_Catalog struct {
	*mock.Call
}

func (_m TaskExecutionContext_Catalog) Return(_a0 catalog.AsyncClient) *TaskExecutionContext_Catalog {
	return &TaskExecutionContext_Catalog{Call: _m.Call.Return(_a0)}
}

func (_m *TaskExecutionContext) OnCatalog() *TaskExecutionContext_Catalog {
	c_call := _m.On("Catalog")
	return &TaskExecutionContext_Catalog{Call: c_call}
}

func (_m *TaskExecutionContext) OnCatalogMatch(matchers ...interface{}) *TaskExecutionContext_Catalog {
	c_call := _m.On("Catalog", matchers...)
	return &TaskExecutionContext_Catalog{Call: c_call}
}

// Catalog provides a mock function with given fields:
func (_m *TaskExecutionContext) Catalog() catalog.AsyncClient {
	ret := _m.Called()

	var r0 catalog.AsyncClient
	if rf, ok := ret.Get(0).(func() catalog.AsyncClient); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(catalog.AsyncClient)
		}
	}

	return r0
}

type TaskExecutionContext_DataStore struct {
	*mock.Call
}

func (_m TaskExecutionContext_DataStore) Return(_a0 *storage.DataStore) *TaskExecutionContext_DataStore {
	return &TaskExecutionContext_DataStore{Call: _m.Call.Return(_a0)}
}

func (_m *TaskExecutionContext) OnDataStore() *TaskExecutionContext_DataStore {
	c_call := _m.On("DataStore")
	return &TaskExecutionContext_DataStore{Call: c_call}
}

func (_m *TaskExecutionContext) OnDataStoreMatch(matchers ...interface{}) *TaskExecutionContext_DataStore {
	c_call := _m.On("DataStore", matchers...)
	return &TaskExecutionContext_DataStore{Call: c_call}
}

// DataStore provides a mock function with given fields:
func (_m *TaskExecutionContext) DataStore() *storage.DataStore {
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

type TaskExecutionContext_EventsRecorder struct {
	*mock.Call
}

func (_m TaskExecutionContext_EventsRecorder) Return(_a0 core.EventsRecorder) *TaskExecutionContext_EventsRecorder {
	return &TaskExecutionContext_EventsRecorder{Call: _m.Call.Return(_a0)}
}

func (_m *TaskExecutionContext) OnEventsRecorder() *TaskExecutionContext_EventsRecorder {
	c_call := _m.On("EventsRecorder")
	return &TaskExecutionContext_EventsRecorder{Call: c_call}
}

func (_m *TaskExecutionContext) OnEventsRecorderMatch(matchers ...interface{}) *TaskExecutionContext_EventsRecorder {
	c_call := _m.On("EventsRecorder", matchers...)
	return &TaskExecutionContext_EventsRecorder{Call: c_call}
}

// EventsRecorder provides a mock function with given fields:
func (_m *TaskExecutionContext) EventsRecorder() core.EventsRecorder {
	ret := _m.Called()

	var r0 core.EventsRecorder
	if rf, ok := ret.Get(0).(func() core.EventsRecorder); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(core.EventsRecorder)
		}
	}

	return r0
}

type TaskExecutionContext_GetExecutionEnvClient struct {
	*mock.Call
}

func (_m TaskExecutionContext_GetExecutionEnvClient) Return(_a0 core.ExecutionEnvClient) *TaskExecutionContext_GetExecutionEnvClient {
	return &TaskExecutionContext_GetExecutionEnvClient{Call: _m.Call.Return(_a0)}
}

func (_m *TaskExecutionContext) OnGetExecutionEnvClient() *TaskExecutionContext_GetExecutionEnvClient {
	c_call := _m.On("GetExecutionEnvClient")
	return &TaskExecutionContext_GetExecutionEnvClient{Call: c_call}
}

func (_m *TaskExecutionContext) OnGetExecutionEnvClientMatch(matchers ...interface{}) *TaskExecutionContext_GetExecutionEnvClient {
	c_call := _m.On("GetExecutionEnvClient", matchers...)
	return &TaskExecutionContext_GetExecutionEnvClient{Call: c_call}
}

// GetExecutionEnvClient provides a mock function with given fields:
func (_m *TaskExecutionContext) GetExecutionEnvClient() core.ExecutionEnvClient {
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

type TaskExecutionContext_InputReader struct {
	*mock.Call
}

func (_m TaskExecutionContext_InputReader) Return(_a0 io.InputReader) *TaskExecutionContext_InputReader {
	return &TaskExecutionContext_InputReader{Call: _m.Call.Return(_a0)}
}

func (_m *TaskExecutionContext) OnInputReader() *TaskExecutionContext_InputReader {
	c_call := _m.On("InputReader")
	return &TaskExecutionContext_InputReader{Call: c_call}
}

func (_m *TaskExecutionContext) OnInputReaderMatch(matchers ...interface{}) *TaskExecutionContext_InputReader {
	c_call := _m.On("InputReader", matchers...)
	return &TaskExecutionContext_InputReader{Call: c_call}
}

// InputReader provides a mock function with given fields:
func (_m *TaskExecutionContext) InputReader() io.InputReader {
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

type TaskExecutionContext_MaxDatasetSizeBytes struct {
	*mock.Call
}

func (_m TaskExecutionContext_MaxDatasetSizeBytes) Return(_a0 int64) *TaskExecutionContext_MaxDatasetSizeBytes {
	return &TaskExecutionContext_MaxDatasetSizeBytes{Call: _m.Call.Return(_a0)}
}

func (_m *TaskExecutionContext) OnMaxDatasetSizeBytes() *TaskExecutionContext_MaxDatasetSizeBytes {
	c_call := _m.On("MaxDatasetSizeBytes")
	return &TaskExecutionContext_MaxDatasetSizeBytes{Call: c_call}
}

func (_m *TaskExecutionContext) OnMaxDatasetSizeBytesMatch(matchers ...interface{}) *TaskExecutionContext_MaxDatasetSizeBytes {
	c_call := _m.On("MaxDatasetSizeBytes", matchers...)
	return &TaskExecutionContext_MaxDatasetSizeBytes{Call: c_call}
}

// MaxDatasetSizeBytes provides a mock function with given fields:
func (_m *TaskExecutionContext) MaxDatasetSizeBytes() int64 {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

type TaskExecutionContext_OutputWriter struct {
	*mock.Call
}

func (_m TaskExecutionContext_OutputWriter) Return(_a0 io.OutputWriter) *TaskExecutionContext_OutputWriter {
	return &TaskExecutionContext_OutputWriter{Call: _m.Call.Return(_a0)}
}

func (_m *TaskExecutionContext) OnOutputWriter() *TaskExecutionContext_OutputWriter {
	c_call := _m.On("OutputWriter")
	return &TaskExecutionContext_OutputWriter{Call: c_call}
}

func (_m *TaskExecutionContext) OnOutputWriterMatch(matchers ...interface{}) *TaskExecutionContext_OutputWriter {
	c_call := _m.On("OutputWriter", matchers...)
	return &TaskExecutionContext_OutputWriter{Call: c_call}
}

// OutputWriter provides a mock function with given fields:
func (_m *TaskExecutionContext) OutputWriter() io.OutputWriter {
	ret := _m.Called()

	var r0 io.OutputWriter
	if rf, ok := ret.Get(0).(func() io.OutputWriter); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.OutputWriter)
		}
	}

	return r0
}

type TaskExecutionContext_PluginStateReader struct {
	*mock.Call
}

func (_m TaskExecutionContext_PluginStateReader) Return(_a0 core.PluginStateReader) *TaskExecutionContext_PluginStateReader {
	return &TaskExecutionContext_PluginStateReader{Call: _m.Call.Return(_a0)}
}

func (_m *TaskExecutionContext) OnPluginStateReader() *TaskExecutionContext_PluginStateReader {
	c_call := _m.On("PluginStateReader")
	return &TaskExecutionContext_PluginStateReader{Call: c_call}
}

func (_m *TaskExecutionContext) OnPluginStateReaderMatch(matchers ...interface{}) *TaskExecutionContext_PluginStateReader {
	c_call := _m.On("PluginStateReader", matchers...)
	return &TaskExecutionContext_PluginStateReader{Call: c_call}
}

// PluginStateReader provides a mock function with given fields:
func (_m *TaskExecutionContext) PluginStateReader() core.PluginStateReader {
	ret := _m.Called()

	var r0 core.PluginStateReader
	if rf, ok := ret.Get(0).(func() core.PluginStateReader); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(core.PluginStateReader)
		}
	}

	return r0
}

type TaskExecutionContext_PluginStateWriter struct {
	*mock.Call
}

func (_m TaskExecutionContext_PluginStateWriter) Return(_a0 core.PluginStateWriter) *TaskExecutionContext_PluginStateWriter {
	return &TaskExecutionContext_PluginStateWriter{Call: _m.Call.Return(_a0)}
}

func (_m *TaskExecutionContext) OnPluginStateWriter() *TaskExecutionContext_PluginStateWriter {
	c_call := _m.On("PluginStateWriter")
	return &TaskExecutionContext_PluginStateWriter{Call: c_call}
}

func (_m *TaskExecutionContext) OnPluginStateWriterMatch(matchers ...interface{}) *TaskExecutionContext_PluginStateWriter {
	c_call := _m.On("PluginStateWriter", matchers...)
	return &TaskExecutionContext_PluginStateWriter{Call: c_call}
}

// PluginStateWriter provides a mock function with given fields:
func (_m *TaskExecutionContext) PluginStateWriter() core.PluginStateWriter {
	ret := _m.Called()

	var r0 core.PluginStateWriter
	if rf, ok := ret.Get(0).(func() core.PluginStateWriter); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(core.PluginStateWriter)
		}
	}

	return r0
}

type TaskExecutionContext_ResourceManager struct {
	*mock.Call
}

func (_m TaskExecutionContext_ResourceManager) Return(_a0 core.ResourceManager) *TaskExecutionContext_ResourceManager {
	return &TaskExecutionContext_ResourceManager{Call: _m.Call.Return(_a0)}
}

func (_m *TaskExecutionContext) OnResourceManager() *TaskExecutionContext_ResourceManager {
	c_call := _m.On("ResourceManager")
	return &TaskExecutionContext_ResourceManager{Call: c_call}
}

func (_m *TaskExecutionContext) OnResourceManagerMatch(matchers ...interface{}) *TaskExecutionContext_ResourceManager {
	c_call := _m.On("ResourceManager", matchers...)
	return &TaskExecutionContext_ResourceManager{Call: c_call}
}

// ResourceManager provides a mock function with given fields:
func (_m *TaskExecutionContext) ResourceManager() core.ResourceManager {
	ret := _m.Called()

	var r0 core.ResourceManager
	if rf, ok := ret.Get(0).(func() core.ResourceManager); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(core.ResourceManager)
		}
	}

	return r0
}

type TaskExecutionContext_SecretManager struct {
	*mock.Call
}

func (_m TaskExecutionContext_SecretManager) Return(_a0 core.SecretManager) *TaskExecutionContext_SecretManager {
	return &TaskExecutionContext_SecretManager{Call: _m.Call.Return(_a0)}
}

func (_m *TaskExecutionContext) OnSecretManager() *TaskExecutionContext_SecretManager {
	c_call := _m.On("SecretManager")
	return &TaskExecutionContext_SecretManager{Call: c_call}
}

func (_m *TaskExecutionContext) OnSecretManagerMatch(matchers ...interface{}) *TaskExecutionContext_SecretManager {
	c_call := _m.On("SecretManager", matchers...)
	return &TaskExecutionContext_SecretManager{Call: c_call}
}

// SecretManager provides a mock function with given fields:
func (_m *TaskExecutionContext) SecretManager() core.SecretManager {
	ret := _m.Called()

	var r0 core.SecretManager
	if rf, ok := ret.Get(0).(func() core.SecretManager); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(core.SecretManager)
		}
	}

	return r0
}

type TaskExecutionContext_TaskExecutionMetadata struct {
	*mock.Call
}

func (_m TaskExecutionContext_TaskExecutionMetadata) Return(_a0 core.TaskExecutionMetadata) *TaskExecutionContext_TaskExecutionMetadata {
	return &TaskExecutionContext_TaskExecutionMetadata{Call: _m.Call.Return(_a0)}
}

func (_m *TaskExecutionContext) OnTaskExecutionMetadata() *TaskExecutionContext_TaskExecutionMetadata {
	c_call := _m.On("TaskExecutionMetadata")
	return &TaskExecutionContext_TaskExecutionMetadata{Call: c_call}
}

func (_m *TaskExecutionContext) OnTaskExecutionMetadataMatch(matchers ...interface{}) *TaskExecutionContext_TaskExecutionMetadata {
	c_call := _m.On("TaskExecutionMetadata", matchers...)
	return &TaskExecutionContext_TaskExecutionMetadata{Call: c_call}
}

// TaskExecutionMetadata provides a mock function with given fields:
func (_m *TaskExecutionContext) TaskExecutionMetadata() core.TaskExecutionMetadata {
	ret := _m.Called()

	var r0 core.TaskExecutionMetadata
	if rf, ok := ret.Get(0).(func() core.TaskExecutionMetadata); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(core.TaskExecutionMetadata)
		}
	}

	return r0
}

type TaskExecutionContext_TaskReader struct {
	*mock.Call
}

func (_m TaskExecutionContext_TaskReader) Return(_a0 core.TaskReader) *TaskExecutionContext_TaskReader {
	return &TaskExecutionContext_TaskReader{Call: _m.Call.Return(_a0)}
}

func (_m *TaskExecutionContext) OnTaskReader() *TaskExecutionContext_TaskReader {
	c_call := _m.On("TaskReader")
	return &TaskExecutionContext_TaskReader{Call: c_call}
}

func (_m *TaskExecutionContext) OnTaskReaderMatch(matchers ...interface{}) *TaskExecutionContext_TaskReader {
	c_call := _m.On("TaskReader", matchers...)
	return &TaskExecutionContext_TaskReader{Call: c_call}
}

// TaskReader provides a mock function with given fields:
func (_m *TaskExecutionContext) TaskReader() core.TaskReader {
	ret := _m.Called()

	var r0 core.TaskReader
	if rf, ok := ret.Get(0).(func() core.TaskReader); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(core.TaskReader)
		}
	}

	return r0
}

type TaskExecutionContext_TaskRefreshIndicator struct {
	*mock.Call
}

func (_m TaskExecutionContext_TaskRefreshIndicator) Return(_a0 core.SignalAsync) *TaskExecutionContext_TaskRefreshIndicator {
	return &TaskExecutionContext_TaskRefreshIndicator{Call: _m.Call.Return(_a0)}
}

func (_m *TaskExecutionContext) OnTaskRefreshIndicator() *TaskExecutionContext_TaskRefreshIndicator {
	c_call := _m.On("TaskRefreshIndicator")
	return &TaskExecutionContext_TaskRefreshIndicator{Call: c_call}
}

func (_m *TaskExecutionContext) OnTaskRefreshIndicatorMatch(matchers ...interface{}) *TaskExecutionContext_TaskRefreshIndicator {
	c_call := _m.On("TaskRefreshIndicator", matchers...)
	return &TaskExecutionContext_TaskRefreshIndicator{Call: c_call}
}

// TaskRefreshIndicator provides a mock function with given fields:
func (_m *TaskExecutionContext) TaskRefreshIndicator() core.SignalAsync {
	ret := _m.Called()

	var r0 core.SignalAsync
	if rf, ok := ret.Get(0).(func() core.SignalAsync); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(core.SignalAsync)
		}
	}

	return r0
}
