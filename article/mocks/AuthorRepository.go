// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/bxcodec/go-clean-arch/domain"
	mock "github.com/stretchr/testify/mock"
)

// AuthorRepository is an autogenerated mock type for the AuthorRepository type
type AuthorRepository struct {
	mock.Mock
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *AuthorRepository) GetByID(ctx context.Context, id int64) (domain.Author, error) {
	ret := _m.Called(ctx, id)

	var r0 domain.Author
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (domain.Author, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) domain.Author); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(domain.Author)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewAuthorRepository creates a new instance of AuthorRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAuthorRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *AuthorRepository {
	mock := &AuthorRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
