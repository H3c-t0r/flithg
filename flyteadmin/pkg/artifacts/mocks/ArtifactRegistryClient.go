// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	context "context"

	artifacts "github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/artifacts"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"
)

// ArtifactRegistryClient is an autogenerated mock type for the ArtifactRegistryClient type
type ArtifactRegistryClient struct {
	mock.Mock
}

type ArtifactRegistryClient_ActivateTrigger struct {
	*mock.Call
}

func (_m ArtifactRegistryClient_ActivateTrigger) Return(_a0 *artifacts.ActivateTriggerResponse, _a1 error) *ArtifactRegistryClient_ActivateTrigger {
	return &ArtifactRegistryClient_ActivateTrigger{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *ArtifactRegistryClient) OnActivateTrigger(ctx context.Context, in *artifacts.ActivateTriggerRequest, opts ...grpc.CallOption) *ArtifactRegistryClient_ActivateTrigger {
	c_call := _m.On("ActivateTrigger", ctx, in, opts)
	return &ArtifactRegistryClient_ActivateTrigger{Call: c_call}
}

func (_m *ArtifactRegistryClient) OnActivateTriggerMatch(matchers ...interface{}) *ArtifactRegistryClient_ActivateTrigger {
	c_call := _m.On("ActivateTrigger", matchers...)
	return &ArtifactRegistryClient_ActivateTrigger{Call: c_call}
}

// ActivateTrigger provides a mock function with given fields: ctx, in, opts
func (_m *ArtifactRegistryClient) ActivateTrigger(ctx context.Context, in *artifacts.ActivateTriggerRequest, opts ...grpc.CallOption) (*artifacts.ActivateTriggerResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *artifacts.ActivateTriggerResponse
	if rf, ok := ret.Get(0).(func(context.Context, *artifacts.ActivateTriggerRequest, ...grpc.CallOption) *artifacts.ActivateTriggerResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*artifacts.ActivateTriggerResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *artifacts.ActivateTriggerRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type ArtifactRegistryClient_AddTag struct {
	*mock.Call
}

func (_m ArtifactRegistryClient_AddTag) Return(_a0 *artifacts.AddTagResponse, _a1 error) *ArtifactRegistryClient_AddTag {
	return &ArtifactRegistryClient_AddTag{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *ArtifactRegistryClient) OnAddTag(ctx context.Context, in *artifacts.AddTagRequest, opts ...grpc.CallOption) *ArtifactRegistryClient_AddTag {
	c_call := _m.On("AddTag", ctx, in, opts)
	return &ArtifactRegistryClient_AddTag{Call: c_call}
}

func (_m *ArtifactRegistryClient) OnAddTagMatch(matchers ...interface{}) *ArtifactRegistryClient_AddTag {
	c_call := _m.On("AddTag", matchers...)
	return &ArtifactRegistryClient_AddTag{Call: c_call}
}

// AddTag provides a mock function with given fields: ctx, in, opts
func (_m *ArtifactRegistryClient) AddTag(ctx context.Context, in *artifacts.AddTagRequest, opts ...grpc.CallOption) (*artifacts.AddTagResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *artifacts.AddTagResponse
	if rf, ok := ret.Get(0).(func(context.Context, *artifacts.AddTagRequest, ...grpc.CallOption) *artifacts.AddTagResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*artifacts.AddTagResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *artifacts.AddTagRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type ArtifactRegistryClient_CreateArtifact struct {
	*mock.Call
}

func (_m ArtifactRegistryClient_CreateArtifact) Return(_a0 *artifacts.CreateArtifactResponse, _a1 error) *ArtifactRegistryClient_CreateArtifact {
	return &ArtifactRegistryClient_CreateArtifact{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *ArtifactRegistryClient) OnCreateArtifact(ctx context.Context, in *artifacts.CreateArtifactRequest, opts ...grpc.CallOption) *ArtifactRegistryClient_CreateArtifact {
	c_call := _m.On("CreateArtifact", ctx, in, opts)
	return &ArtifactRegistryClient_CreateArtifact{Call: c_call}
}

func (_m *ArtifactRegistryClient) OnCreateArtifactMatch(matchers ...interface{}) *ArtifactRegistryClient_CreateArtifact {
	c_call := _m.On("CreateArtifact", matchers...)
	return &ArtifactRegistryClient_CreateArtifact{Call: c_call}
}

// CreateArtifact provides a mock function with given fields: ctx, in, opts
func (_m *ArtifactRegistryClient) CreateArtifact(ctx context.Context, in *artifacts.CreateArtifactRequest, opts ...grpc.CallOption) (*artifacts.CreateArtifactResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *artifacts.CreateArtifactResponse
	if rf, ok := ret.Get(0).(func(context.Context, *artifacts.CreateArtifactRequest, ...grpc.CallOption) *artifacts.CreateArtifactResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*artifacts.CreateArtifactResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *artifacts.CreateArtifactRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type ArtifactRegistryClient_CreateTrigger struct {
	*mock.Call
}

func (_m ArtifactRegistryClient_CreateTrigger) Return(_a0 *artifacts.CreateTriggerResponse, _a1 error) *ArtifactRegistryClient_CreateTrigger {
	return &ArtifactRegistryClient_CreateTrigger{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *ArtifactRegistryClient) OnCreateTrigger(ctx context.Context, in *artifacts.CreateTriggerRequest, opts ...grpc.CallOption) *ArtifactRegistryClient_CreateTrigger {
	c_call := _m.On("CreateTrigger", ctx, in, opts)
	return &ArtifactRegistryClient_CreateTrigger{Call: c_call}
}

func (_m *ArtifactRegistryClient) OnCreateTriggerMatch(matchers ...interface{}) *ArtifactRegistryClient_CreateTrigger {
	c_call := _m.On("CreateTrigger", matchers...)
	return &ArtifactRegistryClient_CreateTrigger{Call: c_call}
}

// CreateTrigger provides a mock function with given fields: ctx, in, opts
func (_m *ArtifactRegistryClient) CreateTrigger(ctx context.Context, in *artifacts.CreateTriggerRequest, opts ...grpc.CallOption) (*artifacts.CreateTriggerResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *artifacts.CreateTriggerResponse
	if rf, ok := ret.Get(0).(func(context.Context, *artifacts.CreateTriggerRequest, ...grpc.CallOption) *artifacts.CreateTriggerResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*artifacts.CreateTriggerResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *artifacts.CreateTriggerRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type ArtifactRegistryClient_DeactivateAllTriggers struct {
	*mock.Call
}

func (_m ArtifactRegistryClient_DeactivateAllTriggers) Return(_a0 *artifacts.DeactivateAllTriggersResponse, _a1 error) *ArtifactRegistryClient_DeactivateAllTriggers {
	return &ArtifactRegistryClient_DeactivateAllTriggers{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *ArtifactRegistryClient) OnDeactivateAllTriggers(ctx context.Context, in *artifacts.DeactivateAllTriggersRequest, opts ...grpc.CallOption) *ArtifactRegistryClient_DeactivateAllTriggers {
	c_call := _m.On("DeactivateAllTriggers", ctx, in, opts)
	return &ArtifactRegistryClient_DeactivateAllTriggers{Call: c_call}
}

func (_m *ArtifactRegistryClient) OnDeactivateAllTriggersMatch(matchers ...interface{}) *ArtifactRegistryClient_DeactivateAllTriggers {
	c_call := _m.On("DeactivateAllTriggers", matchers...)
	return &ArtifactRegistryClient_DeactivateAllTriggers{Call: c_call}
}

// DeactivateAllTriggers provides a mock function with given fields: ctx, in, opts
func (_m *ArtifactRegistryClient) DeactivateAllTriggers(ctx context.Context, in *artifacts.DeactivateAllTriggersRequest, opts ...grpc.CallOption) (*artifacts.DeactivateAllTriggersResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *artifacts.DeactivateAllTriggersResponse
	if rf, ok := ret.Get(0).(func(context.Context, *artifacts.DeactivateAllTriggersRequest, ...grpc.CallOption) *artifacts.DeactivateAllTriggersResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*artifacts.DeactivateAllTriggersResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *artifacts.DeactivateAllTriggersRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type ArtifactRegistryClient_DeactivateTrigger struct {
	*mock.Call
}

func (_m ArtifactRegistryClient_DeactivateTrigger) Return(_a0 *artifacts.DeactivateTriggerResponse, _a1 error) *ArtifactRegistryClient_DeactivateTrigger {
	return &ArtifactRegistryClient_DeactivateTrigger{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *ArtifactRegistryClient) OnDeactivateTrigger(ctx context.Context, in *artifacts.DeactivateTriggerRequest, opts ...grpc.CallOption) *ArtifactRegistryClient_DeactivateTrigger {
	c_call := _m.On("DeactivateTrigger", ctx, in, opts)
	return &ArtifactRegistryClient_DeactivateTrigger{Call: c_call}
}

func (_m *ArtifactRegistryClient) OnDeactivateTriggerMatch(matchers ...interface{}) *ArtifactRegistryClient_DeactivateTrigger {
	c_call := _m.On("DeactivateTrigger", matchers...)
	return &ArtifactRegistryClient_DeactivateTrigger{Call: c_call}
}

// DeactivateTrigger provides a mock function with given fields: ctx, in, opts
func (_m *ArtifactRegistryClient) DeactivateTrigger(ctx context.Context, in *artifacts.DeactivateTriggerRequest, opts ...grpc.CallOption) (*artifacts.DeactivateTriggerResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *artifacts.DeactivateTriggerResponse
	if rf, ok := ret.Get(0).(func(context.Context, *artifacts.DeactivateTriggerRequest, ...grpc.CallOption) *artifacts.DeactivateTriggerResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*artifacts.DeactivateTriggerResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *artifacts.DeactivateTriggerRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type ArtifactRegistryClient_FindByWorkflowExec struct {
	*mock.Call
}

func (_m ArtifactRegistryClient_FindByWorkflowExec) Return(_a0 *artifacts.SearchArtifactsResponse, _a1 error) *ArtifactRegistryClient_FindByWorkflowExec {
	return &ArtifactRegistryClient_FindByWorkflowExec{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *ArtifactRegistryClient) OnFindByWorkflowExec(ctx context.Context, in *artifacts.FindByWorkflowExecRequest, opts ...grpc.CallOption) *ArtifactRegistryClient_FindByWorkflowExec {
	c_call := _m.On("FindByWorkflowExec", ctx, in, opts)
	return &ArtifactRegistryClient_FindByWorkflowExec{Call: c_call}
}

func (_m *ArtifactRegistryClient) OnFindByWorkflowExecMatch(matchers ...interface{}) *ArtifactRegistryClient_FindByWorkflowExec {
	c_call := _m.On("FindByWorkflowExec", matchers...)
	return &ArtifactRegistryClient_FindByWorkflowExec{Call: c_call}
}

// FindByWorkflowExec provides a mock function with given fields: ctx, in, opts
func (_m *ArtifactRegistryClient) FindByWorkflowExec(ctx context.Context, in *artifacts.FindByWorkflowExecRequest, opts ...grpc.CallOption) (*artifacts.SearchArtifactsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *artifacts.SearchArtifactsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *artifacts.FindByWorkflowExecRequest, ...grpc.CallOption) *artifacts.SearchArtifactsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*artifacts.SearchArtifactsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *artifacts.FindByWorkflowExecRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type ArtifactRegistryClient_GetArtifact struct {
	*mock.Call
}

func (_m ArtifactRegistryClient_GetArtifact) Return(_a0 *artifacts.GetArtifactResponse, _a1 error) *ArtifactRegistryClient_GetArtifact {
	return &ArtifactRegistryClient_GetArtifact{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *ArtifactRegistryClient) OnGetArtifact(ctx context.Context, in *artifacts.GetArtifactRequest, opts ...grpc.CallOption) *ArtifactRegistryClient_GetArtifact {
	c_call := _m.On("GetArtifact", ctx, in, opts)
	return &ArtifactRegistryClient_GetArtifact{Call: c_call}
}

func (_m *ArtifactRegistryClient) OnGetArtifactMatch(matchers ...interface{}) *ArtifactRegistryClient_GetArtifact {
	c_call := _m.On("GetArtifact", matchers...)
	return &ArtifactRegistryClient_GetArtifact{Call: c_call}
}

// GetArtifact provides a mock function with given fields: ctx, in, opts
func (_m *ArtifactRegistryClient) GetArtifact(ctx context.Context, in *artifacts.GetArtifactRequest, opts ...grpc.CallOption) (*artifacts.GetArtifactResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *artifacts.GetArtifactResponse
	if rf, ok := ret.Get(0).(func(context.Context, *artifacts.GetArtifactRequest, ...grpc.CallOption) *artifacts.GetArtifactResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*artifacts.GetArtifactResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *artifacts.GetArtifactRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type ArtifactRegistryClient_GetCard struct {
	*mock.Call
}

func (_m ArtifactRegistryClient_GetCard) Return(_a0 *artifacts.GetCardResponse, _a1 error) *ArtifactRegistryClient_GetCard {
	return &ArtifactRegistryClient_GetCard{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *ArtifactRegistryClient) OnGetCard(ctx context.Context, in *artifacts.GetCardRequest, opts ...grpc.CallOption) *ArtifactRegistryClient_GetCard {
	c_call := _m.On("GetCard", ctx, in, opts)
	return &ArtifactRegistryClient_GetCard{Call: c_call}
}

func (_m *ArtifactRegistryClient) OnGetCardMatch(matchers ...interface{}) *ArtifactRegistryClient_GetCard {
	c_call := _m.On("GetCard", matchers...)
	return &ArtifactRegistryClient_GetCard{Call: c_call}
}

// GetCard provides a mock function with given fields: ctx, in, opts
func (_m *ArtifactRegistryClient) GetCard(ctx context.Context, in *artifacts.GetCardRequest, opts ...grpc.CallOption) (*artifacts.GetCardResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *artifacts.GetCardResponse
	if rf, ok := ret.Get(0).(func(context.Context, *artifacts.GetCardRequest, ...grpc.CallOption) *artifacts.GetCardResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*artifacts.GetCardResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *artifacts.GetCardRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type ArtifactRegistryClient_ListUsage struct {
	*mock.Call
}

func (_m ArtifactRegistryClient_ListUsage) Return(_a0 *artifacts.ListUsageResponse, _a1 error) *ArtifactRegistryClient_ListUsage {
	return &ArtifactRegistryClient_ListUsage{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *ArtifactRegistryClient) OnListUsage(ctx context.Context, in *artifacts.ListUsageRequest, opts ...grpc.CallOption) *ArtifactRegistryClient_ListUsage {
	c_call := _m.On("ListUsage", ctx, in, opts)
	return &ArtifactRegistryClient_ListUsage{Call: c_call}
}

func (_m *ArtifactRegistryClient) OnListUsageMatch(matchers ...interface{}) *ArtifactRegistryClient_ListUsage {
	c_call := _m.On("ListUsage", matchers...)
	return &ArtifactRegistryClient_ListUsage{Call: c_call}
}

// ListUsage provides a mock function with given fields: ctx, in, opts
func (_m *ArtifactRegistryClient) ListUsage(ctx context.Context, in *artifacts.ListUsageRequest, opts ...grpc.CallOption) (*artifacts.ListUsageResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *artifacts.ListUsageResponse
	if rf, ok := ret.Get(0).(func(context.Context, *artifacts.ListUsageRequest, ...grpc.CallOption) *artifacts.ListUsageResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*artifacts.ListUsageResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *artifacts.ListUsageRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type ArtifactRegistryClient_RegisterConsumer struct {
	*mock.Call
}

func (_m ArtifactRegistryClient_RegisterConsumer) Return(_a0 *artifacts.RegisterResponse, _a1 error) *ArtifactRegistryClient_RegisterConsumer {
	return &ArtifactRegistryClient_RegisterConsumer{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *ArtifactRegistryClient) OnRegisterConsumer(ctx context.Context, in *artifacts.RegisterConsumerRequest, opts ...grpc.CallOption) *ArtifactRegistryClient_RegisterConsumer {
	c_call := _m.On("RegisterConsumer", ctx, in, opts)
	return &ArtifactRegistryClient_RegisterConsumer{Call: c_call}
}

func (_m *ArtifactRegistryClient) OnRegisterConsumerMatch(matchers ...interface{}) *ArtifactRegistryClient_RegisterConsumer {
	c_call := _m.On("RegisterConsumer", matchers...)
	return &ArtifactRegistryClient_RegisterConsumer{Call: c_call}
}

// RegisterConsumer provides a mock function with given fields: ctx, in, opts
func (_m *ArtifactRegistryClient) RegisterConsumer(ctx context.Context, in *artifacts.RegisterConsumerRequest, opts ...grpc.CallOption) (*artifacts.RegisterResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *artifacts.RegisterResponse
	if rf, ok := ret.Get(0).(func(context.Context, *artifacts.RegisterConsumerRequest, ...grpc.CallOption) *artifacts.RegisterResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*artifacts.RegisterResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *artifacts.RegisterConsumerRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type ArtifactRegistryClient_RegisterProducer struct {
	*mock.Call
}

func (_m ArtifactRegistryClient_RegisterProducer) Return(_a0 *artifacts.RegisterResponse, _a1 error) *ArtifactRegistryClient_RegisterProducer {
	return &ArtifactRegistryClient_RegisterProducer{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *ArtifactRegistryClient) OnRegisterProducer(ctx context.Context, in *artifacts.RegisterProducerRequest, opts ...grpc.CallOption) *ArtifactRegistryClient_RegisterProducer {
	c_call := _m.On("RegisterProducer", ctx, in, opts)
	return &ArtifactRegistryClient_RegisterProducer{Call: c_call}
}

func (_m *ArtifactRegistryClient) OnRegisterProducerMatch(matchers ...interface{}) *ArtifactRegistryClient_RegisterProducer {
	c_call := _m.On("RegisterProducer", matchers...)
	return &ArtifactRegistryClient_RegisterProducer{Call: c_call}
}

// RegisterProducer provides a mock function with given fields: ctx, in, opts
func (_m *ArtifactRegistryClient) RegisterProducer(ctx context.Context, in *artifacts.RegisterProducerRequest, opts ...grpc.CallOption) (*artifacts.RegisterResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *artifacts.RegisterResponse
	if rf, ok := ret.Get(0).(func(context.Context, *artifacts.RegisterProducerRequest, ...grpc.CallOption) *artifacts.RegisterResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*artifacts.RegisterResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *artifacts.RegisterProducerRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type ArtifactRegistryClient_SearchArtifacts struct {
	*mock.Call
}

func (_m ArtifactRegistryClient_SearchArtifacts) Return(_a0 *artifacts.SearchArtifactsResponse, _a1 error) *ArtifactRegistryClient_SearchArtifacts {
	return &ArtifactRegistryClient_SearchArtifacts{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *ArtifactRegistryClient) OnSearchArtifacts(ctx context.Context, in *artifacts.SearchArtifactsRequest, opts ...grpc.CallOption) *ArtifactRegistryClient_SearchArtifacts {
	c_call := _m.On("SearchArtifacts", ctx, in, opts)
	return &ArtifactRegistryClient_SearchArtifacts{Call: c_call}
}

func (_m *ArtifactRegistryClient) OnSearchArtifactsMatch(matchers ...interface{}) *ArtifactRegistryClient_SearchArtifacts {
	c_call := _m.On("SearchArtifacts", matchers...)
	return &ArtifactRegistryClient_SearchArtifacts{Call: c_call}
}

// SearchArtifacts provides a mock function with given fields: ctx, in, opts
func (_m *ArtifactRegistryClient) SearchArtifacts(ctx context.Context, in *artifacts.SearchArtifactsRequest, opts ...grpc.CallOption) (*artifacts.SearchArtifactsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *artifacts.SearchArtifactsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *artifacts.SearchArtifactsRequest, ...grpc.CallOption) *artifacts.SearchArtifactsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*artifacts.SearchArtifactsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *artifacts.SearchArtifactsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type ArtifactRegistryClient_SetExecutionInputs struct {
	*mock.Call
}

func (_m ArtifactRegistryClient_SetExecutionInputs) Return(_a0 *artifacts.ExecutionInputsResponse, _a1 error) *ArtifactRegistryClient_SetExecutionInputs {
	return &ArtifactRegistryClient_SetExecutionInputs{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *ArtifactRegistryClient) OnSetExecutionInputs(ctx context.Context, in *artifacts.ExecutionInputsRequest, opts ...grpc.CallOption) *ArtifactRegistryClient_SetExecutionInputs {
	c_call := _m.On("SetExecutionInputs", ctx, in, opts)
	return &ArtifactRegistryClient_SetExecutionInputs{Call: c_call}
}

func (_m *ArtifactRegistryClient) OnSetExecutionInputsMatch(matchers ...interface{}) *ArtifactRegistryClient_SetExecutionInputs {
	c_call := _m.On("SetExecutionInputs", matchers...)
	return &ArtifactRegistryClient_SetExecutionInputs{Call: c_call}
}

// SetExecutionInputs provides a mock function with given fields: ctx, in, opts
func (_m *ArtifactRegistryClient) SetExecutionInputs(ctx context.Context, in *artifacts.ExecutionInputsRequest, opts ...grpc.CallOption) (*artifacts.ExecutionInputsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *artifacts.ExecutionInputsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *artifacts.ExecutionInputsRequest, ...grpc.CallOption) *artifacts.ExecutionInputsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*artifacts.ExecutionInputsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *artifacts.ExecutionInputsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
