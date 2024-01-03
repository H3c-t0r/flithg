// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	context "context"

	admin "github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/admin"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"
)

// AgentMetadataServiceClient is an autogenerated mock type for the AgentMetadataServiceClient type
type AgentMetadataServiceClient struct {
	mock.Mock
}

type AgentMetadataServiceClient_GetAgent struct {
	*mock.Call
}

func (_m AgentMetadataServiceClient_GetAgent) Return(_a0 *admin.GetAgentResponse, _a1 error) *AgentMetadataServiceClient_GetAgent {
	return &AgentMetadataServiceClient_GetAgent{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *AgentMetadataServiceClient) OnGetAgent(ctx context.Context, in *admin.GetAgentRequest, opts ...grpc.CallOption) *AgentMetadataServiceClient_GetAgent {
	c_call := _m.On("GetAgent", ctx, in, opts)
	return &AgentMetadataServiceClient_GetAgent{Call: c_call}
}

func (_m *AgentMetadataServiceClient) OnGetAgentMatch(matchers ...interface{}) *AgentMetadataServiceClient_GetAgent {
	c_call := _m.On("GetAgent", matchers...)
	return &AgentMetadataServiceClient_GetAgent{Call: c_call}
}

// GetAgent provides a mock function with given fields: ctx, in, opts
func (_m *AgentMetadataServiceClient) GetAgent(ctx context.Context, in *admin.GetAgentRequest, opts ...grpc.CallOption) (*admin.GetAgentResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *admin.GetAgentResponse
	if rf, ok := ret.Get(0).(func(context.Context, *admin.GetAgentRequest, ...grpc.CallOption) *admin.GetAgentResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.GetAgentResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *admin.GetAgentRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type AgentMetadataServiceClient_ListAgents struct {
	*mock.Call
}

func (_m AgentMetadataServiceClient_ListAgents) Return(_a0 *admin.ListAgentsResponse, _a1 error) *AgentMetadataServiceClient_ListAgents {
	return &AgentMetadataServiceClient_ListAgents{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *AgentMetadataServiceClient) OnListAgents(ctx context.Context, in *admin.ListAgentsRequest, opts ...grpc.CallOption) *AgentMetadataServiceClient_ListAgents {
	c_call := _m.On("ListAgents", ctx, in, opts)
	return &AgentMetadataServiceClient_ListAgents{Call: c_call}
}

func (_m *AgentMetadataServiceClient) OnListAgentsMatch(matchers ...interface{}) *AgentMetadataServiceClient_ListAgents {
	c_call := _m.On("ListAgents", matchers...)
	return &AgentMetadataServiceClient_ListAgents{Call: c_call}
}

// ListAgents provides a mock function with given fields: ctx, in, opts
func (_m *AgentMetadataServiceClient) ListAgents(ctx context.Context, in *admin.ListAgentsRequest, opts ...grpc.CallOption) (*admin.ListAgentsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *admin.ListAgentsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *admin.ListAgentsRequest, ...grpc.CallOption) *admin.ListAgentsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.ListAgentsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *admin.ListAgentsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
