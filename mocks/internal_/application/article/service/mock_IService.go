// Code generated by mockery v2.52.4. DO NOT EDIT.

package service

import (
	context "context"

	model "github.com/radityacandra/go-cms/internal/application/article/model"
	mock "github.com/stretchr/testify/mock"

	types "github.com/radityacandra/go-cms/internal/application/article/types"
)

// MockIService is an autogenerated mock type for the IService type
type MockIService struct {
	mock.Mock
}

type MockIService_Expecter struct {
	mock *mock.Mock
}

func (_m *MockIService) EXPECT() *MockIService_Expecter {
	return &MockIService_Expecter{mock: &_m.Mock}
}

// CalculateTagAssociations provides a mock function with given fields: ctx, article
func (_m *MockIService) CalculateTagAssociations(ctx context.Context, article model.Article) error {
	ret := _m.Called(ctx, article)

	if len(ret) == 0 {
		panic("no return value specified for CalculateTagAssociations")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.Article) error); ok {
		r0 = rf(ctx, article)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockIService_CalculateTagAssociations_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CalculateTagAssociations'
type MockIService_CalculateTagAssociations_Call struct {
	*mock.Call
}

// CalculateTagAssociations is a helper method to define mock.On call
//   - ctx context.Context
//   - article model.Article
func (_e *MockIService_Expecter) CalculateTagAssociations(ctx interface{}, article interface{}) *MockIService_CalculateTagAssociations_Call {
	return &MockIService_CalculateTagAssociations_Call{Call: _e.mock.On("CalculateTagAssociations", ctx, article)}
}

func (_c *MockIService_CalculateTagAssociations_Call) Run(run func(ctx context.Context, article model.Article)) *MockIService_CalculateTagAssociations_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.Article))
	})
	return _c
}

func (_c *MockIService_CalculateTagAssociations_Call) Return(_a0 error) *MockIService_CalculateTagAssociations_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockIService_CalculateTagAssociations_Call) RunAndReturn(run func(context.Context, model.Article) error) *MockIService_CalculateTagAssociations_Call {
	_c.Call.Return(run)
	return _c
}

// CreateArticle provides a mock function with given fields: ctx, input
func (_m *MockIService) CreateArticle(ctx context.Context, input types.CreateArticleInput) (string, error) {
	ret := _m.Called(ctx, input)

	if len(ret) == 0 {
		panic("no return value specified for CreateArticle")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, types.CreateArticleInput) (string, error)); ok {
		return rf(ctx, input)
	}
	if rf, ok := ret.Get(0).(func(context.Context, types.CreateArticleInput) string); ok {
		r0 = rf(ctx, input)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, types.CreateArticleInput) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockIService_CreateArticle_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateArticle'
type MockIService_CreateArticle_Call struct {
	*mock.Call
}

// CreateArticle is a helper method to define mock.On call
//   - ctx context.Context
//   - input types.CreateArticleInput
func (_e *MockIService_Expecter) CreateArticle(ctx interface{}, input interface{}) *MockIService_CreateArticle_Call {
	return &MockIService_CreateArticle_Call{Call: _e.mock.On("CreateArticle", ctx, input)}
}

func (_c *MockIService_CreateArticle_Call) Run(run func(ctx context.Context, input types.CreateArticleInput)) *MockIService_CreateArticle_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(types.CreateArticleInput))
	})
	return _c
}

func (_c *MockIService_CreateArticle_Call) Return(_a0 string, _a1 error) *MockIService_CreateArticle_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockIService_CreateArticle_Call) RunAndReturn(run func(context.Context, types.CreateArticleInput) (string, error)) *MockIService_CreateArticle_Call {
	_c.Call.Return(run)
	return _c
}

// CreateArticleRevision provides a mock function with given fields: ctx, input
func (_m *MockIService) CreateArticleRevision(ctx context.Context, input types.CreateArticleRevisionInput) (string, error) {
	ret := _m.Called(ctx, input)

	if len(ret) == 0 {
		panic("no return value specified for CreateArticleRevision")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, types.CreateArticleRevisionInput) (string, error)); ok {
		return rf(ctx, input)
	}
	if rf, ok := ret.Get(0).(func(context.Context, types.CreateArticleRevisionInput) string); ok {
		r0 = rf(ctx, input)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, types.CreateArticleRevisionInput) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockIService_CreateArticleRevision_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateArticleRevision'
type MockIService_CreateArticleRevision_Call struct {
	*mock.Call
}

// CreateArticleRevision is a helper method to define mock.On call
//   - ctx context.Context
//   - input types.CreateArticleRevisionInput
func (_e *MockIService_Expecter) CreateArticleRevision(ctx interface{}, input interface{}) *MockIService_CreateArticleRevision_Call {
	return &MockIService_CreateArticleRevision_Call{Call: _e.mock.On("CreateArticleRevision", ctx, input)}
}

func (_c *MockIService_CreateArticleRevision_Call) Run(run func(ctx context.Context, input types.CreateArticleRevisionInput)) *MockIService_CreateArticleRevision_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(types.CreateArticleRevisionInput))
	})
	return _c
}

func (_c *MockIService_CreateArticleRevision_Call) Return(_a0 string, _a1 error) *MockIService_CreateArticleRevision_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockIService_CreateArticleRevision_Call) RunAndReturn(run func(context.Context, types.CreateArticleRevisionInput) (string, error)) *MockIService_CreateArticleRevision_Call {
	_c.Call.Return(run)
	return _c
}

// DetailArticle provides a mock function with given fields: ctx, articleId, userId
func (_m *MockIService) DetailArticle(ctx context.Context, articleId string, userId string) (types.DetailArticleOutput, error) {
	ret := _m.Called(ctx, articleId, userId)

	if len(ret) == 0 {
		panic("no return value specified for DetailArticle")
	}

	var r0 types.DetailArticleOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (types.DetailArticleOutput, error)); ok {
		return rf(ctx, articleId, userId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) types.DetailArticleOutput); ok {
		r0 = rf(ctx, articleId, userId)
	} else {
		r0 = ret.Get(0).(types.DetailArticleOutput)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, articleId, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockIService_DetailArticle_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DetailArticle'
type MockIService_DetailArticle_Call struct {
	*mock.Call
}

// DetailArticle is a helper method to define mock.On call
//   - ctx context.Context
//   - articleId string
//   - userId string
func (_e *MockIService_Expecter) DetailArticle(ctx interface{}, articleId interface{}, userId interface{}) *MockIService_DetailArticle_Call {
	return &MockIService_DetailArticle_Call{Call: _e.mock.On("DetailArticle", ctx, articleId, userId)}
}

func (_c *MockIService_DetailArticle_Call) Run(run func(ctx context.Context, articleId string, userId string)) *MockIService_DetailArticle_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *MockIService_DetailArticle_Call) Return(_a0 types.DetailArticleOutput, _a1 error) *MockIService_DetailArticle_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockIService_DetailArticle_Call) RunAndReturn(run func(context.Context, string, string) (types.DetailArticleOutput, error)) *MockIService_DetailArticle_Call {
	_c.Call.Return(run)
	return _c
}

// DetailArticleRevision provides a mock function with given fields: ctx, articleId, revisionId
func (_m *MockIService) DetailArticleRevision(ctx context.Context, articleId string, revisionId string) (types.DetailArticleRevisionOutput, error) {
	ret := _m.Called(ctx, articleId, revisionId)

	if len(ret) == 0 {
		panic("no return value specified for DetailArticleRevision")
	}

	var r0 types.DetailArticleRevisionOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (types.DetailArticleRevisionOutput, error)); ok {
		return rf(ctx, articleId, revisionId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) types.DetailArticleRevisionOutput); ok {
		r0 = rf(ctx, articleId, revisionId)
	} else {
		r0 = ret.Get(0).(types.DetailArticleRevisionOutput)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, articleId, revisionId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockIService_DetailArticleRevision_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DetailArticleRevision'
type MockIService_DetailArticleRevision_Call struct {
	*mock.Call
}

// DetailArticleRevision is a helper method to define mock.On call
//   - ctx context.Context
//   - articleId string
//   - revisionId string
func (_e *MockIService_Expecter) DetailArticleRevision(ctx interface{}, articleId interface{}, revisionId interface{}) *MockIService_DetailArticleRevision_Call {
	return &MockIService_DetailArticleRevision_Call{Call: _e.mock.On("DetailArticleRevision", ctx, articleId, revisionId)}
}

func (_c *MockIService_DetailArticleRevision_Call) Run(run func(ctx context.Context, articleId string, revisionId string)) *MockIService_DetailArticleRevision_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *MockIService_DetailArticleRevision_Call) Return(_a0 types.DetailArticleRevisionOutput, _a1 error) *MockIService_DetailArticleRevision_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockIService_DetailArticleRevision_Call) RunAndReturn(run func(context.Context, string, string) (types.DetailArticleRevisionOutput, error)) *MockIService_DetailArticleRevision_Call {
	_c.Call.Return(run)
	return _c
}

// ListArticle provides a mock function with given fields: ctx, input
func (_m *MockIService) ListArticle(ctx context.Context, input types.ListArticleInput) (types.ListArticleOutput, error) {
	ret := _m.Called(ctx, input)

	if len(ret) == 0 {
		panic("no return value specified for ListArticle")
	}

	var r0 types.ListArticleOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, types.ListArticleInput) (types.ListArticleOutput, error)); ok {
		return rf(ctx, input)
	}
	if rf, ok := ret.Get(0).(func(context.Context, types.ListArticleInput) types.ListArticleOutput); ok {
		r0 = rf(ctx, input)
	} else {
		r0 = ret.Get(0).(types.ListArticleOutput)
	}

	if rf, ok := ret.Get(1).(func(context.Context, types.ListArticleInput) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockIService_ListArticle_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListArticle'
type MockIService_ListArticle_Call struct {
	*mock.Call
}

// ListArticle is a helper method to define mock.On call
//   - ctx context.Context
//   - input types.ListArticleInput
func (_e *MockIService_Expecter) ListArticle(ctx interface{}, input interface{}) *MockIService_ListArticle_Call {
	return &MockIService_ListArticle_Call{Call: _e.mock.On("ListArticle", ctx, input)}
}

func (_c *MockIService_ListArticle_Call) Run(run func(ctx context.Context, input types.ListArticleInput)) *MockIService_ListArticle_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(types.ListArticleInput))
	})
	return _c
}

func (_c *MockIService_ListArticle_Call) Return(_a0 types.ListArticleOutput, _a1 error) *MockIService_ListArticle_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockIService_ListArticle_Call) RunAndReturn(run func(context.Context, types.ListArticleInput) (types.ListArticleOutput, error)) *MockIService_ListArticle_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockIService creates a new instance of MockIService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockIService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockIService {
	mock := &MockIService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
