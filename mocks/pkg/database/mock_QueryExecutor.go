// Code generated by mockery v2.52.4. DO NOT EDIT.

package database

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	sql "database/sql"

	sqlx "github.com/jmoiron/sqlx"
)

// MockQueryExecutor is an autogenerated mock type for the QueryExecutor type
type MockQueryExecutor struct {
	mock.Mock
}

type MockQueryExecutor_Expecter struct {
	mock *mock.Mock
}

func (_m *MockQueryExecutor) EXPECT() *MockQueryExecutor_Expecter {
	return &MockQueryExecutor_Expecter{mock: &_m.Mock}
}

// NamedExecContext provides a mock function with given fields: ctx, query, arg
func (_m *MockQueryExecutor) NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error) {
	ret := _m.Called(ctx, query, arg)

	if len(ret) == 0 {
		panic("no return value specified for NamedExecContext")
	}

	var r0 sql.Result
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, interface{}) (sql.Result, error)); ok {
		return rf(ctx, query, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, interface{}) sql.Result); ok {
		r0 = rf(ctx, query, arg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(sql.Result)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, interface{}) error); ok {
		r1 = rf(ctx, query, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQueryExecutor_NamedExecContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'NamedExecContext'
type MockQueryExecutor_NamedExecContext_Call struct {
	*mock.Call
}

// NamedExecContext is a helper method to define mock.On call
//   - ctx context.Context
//   - query string
//   - arg interface{}
func (_e *MockQueryExecutor_Expecter) NamedExecContext(ctx interface{}, query interface{}, arg interface{}) *MockQueryExecutor_NamedExecContext_Call {
	return &MockQueryExecutor_NamedExecContext_Call{Call: _e.mock.On("NamedExecContext", ctx, query, arg)}
}

func (_c *MockQueryExecutor_NamedExecContext_Call) Run(run func(ctx context.Context, query string, arg interface{})) *MockQueryExecutor_NamedExecContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(interface{}))
	})
	return _c
}

func (_c *MockQueryExecutor_NamedExecContext_Call) Return(_a0 sql.Result, _a1 error) *MockQueryExecutor_NamedExecContext_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockQueryExecutor_NamedExecContext_Call) RunAndReturn(run func(context.Context, string, interface{}) (sql.Result, error)) *MockQueryExecutor_NamedExecContext_Call {
	_c.Call.Return(run)
	return _c
}

// QueryRowxContext provides a mock function with given fields: ctx, query, args
func (_m *MockQueryExecutor) QueryRowxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row {
	var _ca []interface{}
	_ca = append(_ca, ctx, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for QueryRowxContext")
	}

	var r0 *sqlx.Row
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) *sqlx.Row); ok {
		r0 = rf(ctx, query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sqlx.Row)
		}
	}

	return r0
}

// MockQueryExecutor_QueryRowxContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'QueryRowxContext'
type MockQueryExecutor_QueryRowxContext_Call struct {
	*mock.Call
}

// QueryRowxContext is a helper method to define mock.On call
//   - ctx context.Context
//   - query string
//   - args ...interface{}
func (_e *MockQueryExecutor_Expecter) QueryRowxContext(ctx interface{}, query interface{}, args ...interface{}) *MockQueryExecutor_QueryRowxContext_Call {
	return &MockQueryExecutor_QueryRowxContext_Call{Call: _e.mock.On("QueryRowxContext",
		append([]interface{}{ctx, query}, args...)...)}
}

func (_c *MockQueryExecutor_QueryRowxContext_Call) Run(run func(ctx context.Context, query string, args ...interface{})) *MockQueryExecutor_QueryRowxContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(context.Context), args[1].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockQueryExecutor_QueryRowxContext_Call) Return(_a0 *sqlx.Row) *MockQueryExecutor_QueryRowxContext_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockQueryExecutor_QueryRowxContext_Call) RunAndReturn(run func(context.Context, string, ...interface{}) *sqlx.Row) *MockQueryExecutor_QueryRowxContext_Call {
	_c.Call.Return(run)
	return _c
}

// QueryxContext provides a mock function with given fields: ctx, query, args
func (_m *MockQueryExecutor) QueryxContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error) {
	var _ca []interface{}
	_ca = append(_ca, ctx, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for QueryxContext")
	}

	var r0 *sqlx.Rows
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) (*sqlx.Rows, error)); ok {
		return rf(ctx, query, args...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) *sqlx.Rows); ok {
		r0 = rf(ctx, query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sqlx.Rows)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, ...interface{}) error); ok {
		r1 = rf(ctx, query, args...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQueryExecutor_QueryxContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'QueryxContext'
type MockQueryExecutor_QueryxContext_Call struct {
	*mock.Call
}

// QueryxContext is a helper method to define mock.On call
//   - ctx context.Context
//   - query string
//   - args ...interface{}
func (_e *MockQueryExecutor_Expecter) QueryxContext(ctx interface{}, query interface{}, args ...interface{}) *MockQueryExecutor_QueryxContext_Call {
	return &MockQueryExecutor_QueryxContext_Call{Call: _e.mock.On("QueryxContext",
		append([]interface{}{ctx, query}, args...)...)}
}

func (_c *MockQueryExecutor_QueryxContext_Call) Run(run func(ctx context.Context, query string, args ...interface{})) *MockQueryExecutor_QueryxContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(context.Context), args[1].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockQueryExecutor_QueryxContext_Call) Return(_a0 *sqlx.Rows, _a1 error) *MockQueryExecutor_QueryxContext_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockQueryExecutor_QueryxContext_Call) RunAndReturn(run func(context.Context, string, ...interface{}) (*sqlx.Rows, error)) *MockQueryExecutor_QueryxContext_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockQueryExecutor creates a new instance of MockQueryExecutor. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockQueryExecutor(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockQueryExecutor {
	mock := &MockQueryExecutor{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
