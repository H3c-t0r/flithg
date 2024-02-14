// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// SecretFetcher is an autogenerated mock type for the SecretFetcher type
type SecretFetcher struct {
	mock.Mock
}

type SecretFetcher_GetSecretValue struct {
	*mock.Call
}

func (_m SecretFetcher_GetSecretValue) Return(_a0 string, _a1 error) *SecretFetcher_GetSecretValue {
	return &SecretFetcher_GetSecretValue{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *SecretFetcher) OnGetSecretValue(ctx context.Context, secretID string) *SecretFetcher_GetSecretValue {
	c_call := _m.On("GetSecretValue", ctx, secretID)
	return &SecretFetcher_GetSecretValue{Call: c_call}
}

func (_m *SecretFetcher) OnGetSecretValueMatch(matchers ...interface{}) *SecretFetcher_GetSecretValue {
	c_call := _m.On("GetSecretValue", matchers...)
	return &SecretFetcher_GetSecretValue{Call: c_call}
}

// GetSecretValue provides a mock function with given fields: ctx, secretID
func (_m *SecretFetcher) GetSecretValue(ctx context.Context, secretID string) (string, error) {
	ret := _m.Called(ctx, secretID)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, secretID)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, secretID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}