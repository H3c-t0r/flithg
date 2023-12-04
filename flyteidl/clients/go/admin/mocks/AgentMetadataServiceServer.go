// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	context "context"

	admin "github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/admin"

	mock "github.com/stretchr/testify/mock"
)

// AgentMetadataServiceServer is an autogenerated mock type for the AgentMetadataServiceServer type
type AgentMetadataServiceServer struct {
	mock.Mock
}

type AgentMetadataServiceServer_GetAgent struct {
	*mock.Call
}

func (_m AgentMetadataServiceServer_GetAgent) Return(_a0 *admin.GetAgentResponse, _a1 error) *AgentMetadataServiceServer_GetAgent {
	return &AgentMetadataServiceServer_GetAgent{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *AgentMetadataServiceServer) OnGetAgent(_a0 context.Context, _a1 *admin.GetAgentRequest) *AgentMetadataServiceServer_GetAgent {
	c_call := _m.On("GetAgent", _a0, _a1)
	return &AgentMetadataServiceServer_GetAgent{Call: c_call}
}

func (_m *AgentMetadataServiceServer) OnGetAgentMatch(matchers ...interface{}) *AgentMetadataServiceServer_GetAgent {
	c_call := _m.On("GetAgent", matchers...)
	return &AgentMetadataServiceServer_GetAgent{Call: c_call}
}

// GetAgent provides a mock function with given fields: _a0, _a1
func (_m *AgentMetadataServiceServer) GetAgent(_a0 context.Context, _a1 *admin.GetAgentRequest) (*admin.GetAgentResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *admin.GetAgentResponse
	if rf, ok := ret.Get(0).(func(context.Context, *admin.GetAgentRequest) *admin.GetAgentResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.GetAgentResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *admin.GetAgentRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type AgentMetadataServiceServer_ListAgent struct {
	*mock.Call
}

func (_m AgentMetadataServiceServer_ListAgent) Return(_a0 *admin.ListAgentsResponse, _a1 error) *AgentMetadataServiceServer_ListAgent {
	return &AgentMetadataServiceServer_ListAgent{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *AgentMetadataServiceServer) OnListAgent(_a0 context.Context, _a1 *admin.ListAgentsRequest) *AgentMetadataServiceServer_ListAgent {
	c_call := _m.On("ListAgent", _a0, _a1)
	return &AgentMetadataServiceServer_ListAgent{Call: c_call}
}

func (_m *AgentMetadataServiceServer) OnListAgentMatch(matchers ...interface{}) *AgentMetadataServiceServer_ListAgent {
	c_call := _m.On("ListAgent", matchers...)
	return &AgentMetadataServiceServer_ListAgent{Call: c_call}
}

// ListAgent provides a mock function with given fields: _a0, _a1
func (_m *AgentMetadataServiceServer) ListAgent(_a0 context.Context, _a1 *admin.ListAgentsRequest) (*admin.ListAgentsResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *admin.ListAgentsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *admin.ListAgentsRequest) *admin.ListAgentsResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.ListAgentsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *admin.ListAgentsRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}