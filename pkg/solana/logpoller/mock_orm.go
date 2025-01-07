// Code generated by mockery v2.43.2. DO NOT EDIT.

package logpoller

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// mockORM is an autogenerated mock type for the ORM type
type mockORM struct {
	mock.Mock
}

type mockORM_Expecter struct {
	mock *mock.Mock
}

func (_m *mockORM) EXPECT() *mockORM_Expecter {
	return &mockORM_Expecter{mock: &_m.Mock}
}

// DeleteFilters provides a mock function with given fields: ctx, filters
func (_m *mockORM) DeleteFilters(ctx context.Context, filters map[int64]Filter) error {
	ret := _m.Called(ctx, filters)

	if len(ret) == 0 {
		panic("no return value specified for DeleteFilters")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, map[int64]Filter) error); ok {
		r0 = rf(ctx, filters)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// mockORM_DeleteFilters_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteFilters'
type mockORM_DeleteFilters_Call struct {
	*mock.Call
}

// DeleteFilters is a helper method to define mock.On call
//   - ctx context.Context
//   - filters map[int64]Filter
func (_e *mockORM_Expecter) DeleteFilters(ctx interface{}, filters interface{}) *mockORM_DeleteFilters_Call {
	return &mockORM_DeleteFilters_Call{Call: _e.mock.On("DeleteFilters", ctx, filters)}
}

func (_c *mockORM_DeleteFilters_Call) Run(run func(ctx context.Context, filters map[int64]Filter)) *mockORM_DeleteFilters_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(map[int64]Filter))
	})
	return _c
}

func (_c *mockORM_DeleteFilters_Call) Return(_a0 error) *mockORM_DeleteFilters_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mockORM_DeleteFilters_Call) RunAndReturn(run func(context.Context, map[int64]Filter) error) *mockORM_DeleteFilters_Call {
	_c.Call.Return(run)
	return _c
}

// InsertFilter provides a mock function with given fields: ctx, filter
func (_m *mockORM) InsertFilter(ctx context.Context, filter Filter) (int64, error) {
	ret := _m.Called(ctx, filter)

	if len(ret) == 0 {
		panic("no return value specified for InsertFilter")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, Filter) (int64, error)); ok {
		return rf(ctx, filter)
	}
	if rf, ok := ret.Get(0).(func(context.Context, Filter) int64); ok {
		r0 = rf(ctx, filter)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, Filter) error); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockORM_InsertFilter_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'InsertFilter'
type mockORM_InsertFilter_Call struct {
	*mock.Call
}

// InsertFilter is a helper method to define mock.On call
//   - ctx context.Context
//   - filter Filter
func (_e *mockORM_Expecter) InsertFilter(ctx interface{}, filter interface{}) *mockORM_InsertFilter_Call {
	return &mockORM_InsertFilter_Call{Call: _e.mock.On("InsertFilter", ctx, filter)}
}

func (_c *mockORM_InsertFilter_Call) Run(run func(ctx context.Context, filter Filter)) *mockORM_InsertFilter_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(Filter))
	})
	return _c
}

func (_c *mockORM_InsertFilter_Call) Return(id int64, err error) *mockORM_InsertFilter_Call {
	_c.Call.Return(id, err)
	return _c
}

func (_c *mockORM_InsertFilter_Call) RunAndReturn(run func(context.Context, Filter) (int64, error)) *mockORM_InsertFilter_Call {
	_c.Call.Return(run)
	return _c
}

// MarkFilterBackfilled provides a mock function with given fields: ctx, id
func (_m *mockORM) MarkFilterBackfilled(ctx context.Context, id int64) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for MarkFilterBackfilled")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// mockORM_MarkFilterBackfilled_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MarkFilterBackfilled'
type mockORM_MarkFilterBackfilled_Call struct {
	*mock.Call
}

// MarkFilterBackfilled is a helper method to define mock.On call
//   - ctx context.Context
//   - id int64
func (_e *mockORM_Expecter) MarkFilterBackfilled(ctx interface{}, id interface{}) *mockORM_MarkFilterBackfilled_Call {
	return &mockORM_MarkFilterBackfilled_Call{Call: _e.mock.On("MarkFilterBackfilled", ctx, id)}
}

func (_c *mockORM_MarkFilterBackfilled_Call) Run(run func(ctx context.Context, id int64)) *mockORM_MarkFilterBackfilled_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *mockORM_MarkFilterBackfilled_Call) Return(err error) *mockORM_MarkFilterBackfilled_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *mockORM_MarkFilterBackfilled_Call) RunAndReturn(run func(context.Context, int64) error) *mockORM_MarkFilterBackfilled_Call {
	_c.Call.Return(run)
	return _c
}

// MarkFilterDeleted provides a mock function with given fields: ctx, id
func (_m *mockORM) MarkFilterDeleted(ctx context.Context, id int64) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for MarkFilterDeleted")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// mockORM_MarkFilterDeleted_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MarkFilterDeleted'
type mockORM_MarkFilterDeleted_Call struct {
	*mock.Call
}

// MarkFilterDeleted is a helper method to define mock.On call
//   - ctx context.Context
//   - id int64
func (_e *mockORM_Expecter) MarkFilterDeleted(ctx interface{}, id interface{}) *mockORM_MarkFilterDeleted_Call {
	return &mockORM_MarkFilterDeleted_Call{Call: _e.mock.On("MarkFilterDeleted", ctx, id)}
}

func (_c *mockORM_MarkFilterDeleted_Call) Run(run func(ctx context.Context, id int64)) *mockORM_MarkFilterDeleted_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *mockORM_MarkFilterDeleted_Call) Return(err error) *mockORM_MarkFilterDeleted_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *mockORM_MarkFilterDeleted_Call) RunAndReturn(run func(context.Context, int64) error) *mockORM_MarkFilterDeleted_Call {
	_c.Call.Return(run)
	return _c
}

// SelectFilters provides a mock function with given fields: ctx
func (_m *mockORM) SelectFilters(ctx context.Context) ([]Filter, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for SelectFilters")
	}

	var r0 []Filter
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]Filter, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []Filter); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Filter)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockORM_SelectFilters_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SelectFilters'
type mockORM_SelectFilters_Call struct {
	*mock.Call
}

// SelectFilters is a helper method to define mock.On call
//   - ctx context.Context
func (_e *mockORM_Expecter) SelectFilters(ctx interface{}) *mockORM_SelectFilters_Call {
	return &mockORM_SelectFilters_Call{Call: _e.mock.On("SelectFilters", ctx)}
}

func (_c *mockORM_SelectFilters_Call) Run(run func(ctx context.Context)) *mockORM_SelectFilters_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *mockORM_SelectFilters_Call) Return(_a0 []Filter, _a1 error) *mockORM_SelectFilters_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockORM_SelectFilters_Call) RunAndReturn(run func(context.Context) ([]Filter, error)) *mockORM_SelectFilters_Call {
	_c.Call.Return(run)
	return _c
}

// newMockORM creates a new instance of mockORM. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockORM(t interface {
	mock.TestingT
	Cleanup(func())
}) *mockORM {
	mock := &mockORM{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
