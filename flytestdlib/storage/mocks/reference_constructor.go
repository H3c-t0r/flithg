// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	context "context"

	storage "github.com/flyteorg/flyte/flytestdlib/storage"
	mock "github.com/stretchr/testify/mock"
)

// ReferenceConstructor is an autogenerated mock type for the ReferenceConstructor type
type ReferenceConstructor struct {
	mock.Mock
}

type ReferenceConstructor_ConstructReference struct {
	*mock.Call
}

func (_m ReferenceConstructor_ConstructReference) Return(_a0 storage.DataReference, _a1 error) *ReferenceConstructor_ConstructReference {
	return &ReferenceConstructor_ConstructReference{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *ReferenceConstructor) OnConstructReference(ctx context.Context, reference storage.DataReference, nestedKeys ...string) *ReferenceConstructor_ConstructReference {
	c_call := _m.On("ConstructReference", ctx, reference, nestedKeys)
	return &ReferenceConstructor_ConstructReference{Call: c_call}
}

func (_m *ReferenceConstructor) OnConstructReferenceMatch(matchers ...interface{}) *ReferenceConstructor_ConstructReference {
	c_call := _m.On("ConstructReference", matchers...)
	return &ReferenceConstructor_ConstructReference{Call: c_call}
}

// ConstructReference provides a mock function with given fields: ctx, reference, nestedKeys
func (_m *ReferenceConstructor) ConstructReference(ctx context.Context, reference storage.DataReference, nestedKeys ...string) (storage.DataReference, error) {
	_va := make([]interface{}, len(nestedKeys))
	for _i := range nestedKeys {
		_va[_i] = nestedKeys[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, reference)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 storage.DataReference
	if rf, ok := ret.Get(0).(func(context.Context, storage.DataReference, ...string) storage.DataReference); ok {
		r0 = rf(ctx, reference, nestedKeys...)
	} else {
		r0 = ret.Get(0).(storage.DataReference)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, storage.DataReference, ...string) error); ok {
		r1 = rf(ctx, reference, nestedKeys...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
