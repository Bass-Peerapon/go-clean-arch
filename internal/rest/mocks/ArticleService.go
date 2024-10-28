// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/bxcodec/go-clean-arch/domain"
	mock "github.com/stretchr/testify/mock"
)

// ArticleService is an autogenerated mock type for the ArticleService type
type ArticleService struct {
	mock.Mock
}

// Delete provides a mock function with given fields: ctx, id
func (_m *ArticleService) Delete(ctx context.Context, id int64) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Fetch provides a mock function with given fields: ctx, cursor, num
func (_m *ArticleService) Fetch(ctx context.Context, cursor string, num int64) ([]domain.Article, string, error) {
	ret := _m.Called(ctx, cursor, num)

	var r0 []domain.Article
	var r1 string
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int64) ([]domain.Article, string, error)); ok {
		return rf(ctx, cursor, num)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, int64) []domain.Article); ok {
		r0 = rf(ctx, cursor, num)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Article)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, int64) string); ok {
		r1 = rf(ctx, cursor, num)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(context.Context, string, int64) error); ok {
		r2 = rf(ctx, cursor, num)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *ArticleService) GetByID(ctx context.Context, id int64) (domain.Article, error) {
	ret := _m.Called(ctx, id)

	var r0 domain.Article
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (domain.Article, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) domain.Article); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(domain.Article)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByTitle provides a mock function with given fields: ctx, title
func (_m *ArticleService) GetByTitle(ctx context.Context, title string) (domain.Article, error) {
	ret := _m.Called(ctx, title)

	var r0 domain.Article
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (domain.Article, error)); ok {
		return rf(ctx, title)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) domain.Article); ok {
		r0 = rf(ctx, title)
	} else {
		r0 = ret.Get(0).(domain.Article)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, title)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: _a0, _a1
func (_m *ArticleService) Store(_a0 context.Context, _a1 *domain.Article) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Article) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: ctx, ar
func (_m *ArticleService) Update(ctx context.Context, ar *domain.Article) error {
	ret := _m.Called(ctx, ar)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Article) error); ok {
		r0 = rf(ctx, ar)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewArticleService creates a new instance of ArticleService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewArticleService(t interface {
	mock.TestingT
	Cleanup(func())
}) *ArticleService {
	mock := &ArticleService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
