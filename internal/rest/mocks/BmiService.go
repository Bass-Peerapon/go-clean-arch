// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// BmiService is an autogenerated mock type for the BmiService type
type BmiService struct {
	mock.Mock
}

// Calculate provides a mock function with given fields: ctx, weight, height
func (_m *BmiService) Calculate(ctx context.Context, weight float64, height float64) (float64, error) {
	ret := _m.Called(ctx, weight, height)

	var r0 float64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, float64, float64) (float64, error)); ok {
		return rf(ctx, weight, height)
	}
	if rf, ok := ret.Get(0).(func(context.Context, float64, float64) float64); ok {
		r0 = rf(ctx, weight, height)
	} else {
		r0 = ret.Get(0).(float64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, float64, float64) error); ok {
		r1 = rf(ctx, weight, height)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewBmiService creates a new instance of BmiService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBmiService(t interface {
	mock.TestingT
	Cleanup(func())
}) *BmiService {
	mock := &BmiService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
