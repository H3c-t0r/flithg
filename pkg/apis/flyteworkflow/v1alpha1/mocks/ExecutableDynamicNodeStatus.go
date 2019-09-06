// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import v1alpha1 "github.com/lyft/flytepropeller/pkg/apis/flyteworkflow/v1alpha1"

// ExecutableDynamicNodeStatus is an autogenerated mock type for the ExecutableDynamicNodeStatus type
type ExecutableDynamicNodeStatus struct {
	mock.Mock
}

// GetDynamicNodePhase provides a mock function with given fields:
func (_m *ExecutableDynamicNodeStatus) GetDynamicNodePhase() v1alpha1.DynamicNodePhase {
	ret := _m.Called()

	var r0 v1alpha1.DynamicNodePhase
	if rf, ok := ret.Get(0).(func() v1alpha1.DynamicNodePhase); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(v1alpha1.DynamicNodePhase)
	}

	return r0
}
