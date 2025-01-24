// Code generated by mockery v2.43.2. DO NOT EDIT.

package logpoller

import (
	context "context"
	iter "iter"

	mock "github.com/stretchr/testify/mock"
)

// mockFilters is an autogenerated mock type for the filtersI type
type mockFilters struct {
	mock.Mock
}

type mockFilters_Expecter struct {
	mock *mock.Mock
}

func (_m *mockFilters) EXPECT() *mockFilters_Expecter {
	return &mockFilters_Expecter{mock: &_m.Mock}
}

// DecodeSubKey provides a mock function with given fields: ctx, raw, ID, subKeyPath
func (_m *mockFilters) DecodeSubKey(ctx context.Context, raw []byte, ID int64, subKeyPath []string) (interface{}, error) {
	ret := _m.Called(ctx, raw, ID, subKeyPath)

	if len(ret) == 0 {
		panic("no return value specified for DecodeSubKey")
	}

	var r0 interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []byte, int64, []string) (interface{}, error)); ok {
		return rf(ctx, raw, ID, subKeyPath)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []byte, int64, []string) interface{}); ok {
		r0 = rf(ctx, raw, ID, subKeyPath)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []byte, int64, []string) error); ok {
		r1 = rf(ctx, raw, ID, subKeyPath)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockFilters_DecodeSubKey_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DecodeSubKey'
type mockFilters_DecodeSubKey_Call struct {
	*mock.Call
}

// DecodeSubKey is a helper method to define mock.On call
//   - ctx context.Context
//   - raw []byte
//   - ID int64
//   - subKeyPath []string
func (_e *mockFilters_Expecter) DecodeSubKey(ctx interface{}, raw interface{}, ID interface{}, subKeyPath interface{}) *mockFilters_DecodeSubKey_Call {
	return &mockFilters_DecodeSubKey_Call{Call: _e.mock.On("DecodeSubKey", ctx, raw, ID, subKeyPath)}
}

func (_c *mockFilters_DecodeSubKey_Call) Run(run func(ctx context.Context, raw []byte, ID int64, subKeyPath []string)) *mockFilters_DecodeSubKey_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]byte), args[2].(int64), args[3].([]string))
	})
	return _c
}

func (_c *mockFilters_DecodeSubKey_Call) Return(_a0 interface{}, _a1 error) *mockFilters_DecodeSubKey_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockFilters_DecodeSubKey_Call) RunAndReturn(run func(context.Context, []byte, int64, []string) (interface{}, error)) *mockFilters_DecodeSubKey_Call {
	_c.Call.Return(run)
	return _c
}

// GetDistinctAddresses provides a mock function with given fields: ctx
func (_m *mockFilters) GetDistinctAddresses(ctx context.Context) ([]PublicKey, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetDistinctAddresses")
	}

	var r0 []PublicKey
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]PublicKey, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []PublicKey); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]PublicKey)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockFilters_GetDistinctAddresses_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetDistinctAddresses'
type mockFilters_GetDistinctAddresses_Call struct {
	*mock.Call
}

// GetDistinctAddresses is a helper method to define mock.On call
//   - ctx context.Context
func (_e *mockFilters_Expecter) GetDistinctAddresses(ctx interface{}) *mockFilters_GetDistinctAddresses_Call {
	return &mockFilters_GetDistinctAddresses_Call{Call: _e.mock.On("GetDistinctAddresses", ctx)}
}

func (_c *mockFilters_GetDistinctAddresses_Call) Run(run func(ctx context.Context)) *mockFilters_GetDistinctAddresses_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *mockFilters_GetDistinctAddresses_Call) Return(_a0 []PublicKey, _a1 error) *mockFilters_GetDistinctAddresses_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockFilters_GetDistinctAddresses_Call) RunAndReturn(run func(context.Context) ([]PublicKey, error)) *mockFilters_GetDistinctAddresses_Call {
	_c.Call.Return(run)
	return _c
}

// GetFiltersToBackfill provides a mock function with given fields:
func (_m *mockFilters) GetFiltersToBackfill() []Filter {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetFiltersToBackfill")
	}

	var r0 []Filter
	if rf, ok := ret.Get(0).(func() []Filter); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Filter)
		}
	}

	return r0
}

// mockFilters_GetFiltersToBackfill_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetFiltersToBackfill'
type mockFilters_GetFiltersToBackfill_Call struct {
	*mock.Call
}

// GetFiltersToBackfill is a helper method to define mock.On call
func (_e *mockFilters_Expecter) GetFiltersToBackfill() *mockFilters_GetFiltersToBackfill_Call {
	return &mockFilters_GetFiltersToBackfill_Call{Call: _e.mock.On("GetFiltersToBackfill")}
}

func (_c *mockFilters_GetFiltersToBackfill_Call) Run(run func()) *mockFilters_GetFiltersToBackfill_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *mockFilters_GetFiltersToBackfill_Call) Return(_a0 []Filter) *mockFilters_GetFiltersToBackfill_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mockFilters_GetFiltersToBackfill_Call) RunAndReturn(run func() []Filter) *mockFilters_GetFiltersToBackfill_Call {
	_c.Call.Return(run)
	return _c
}

// IncrementSeqNum provides a mock function with given fields: filterID
func (_m *mockFilters) IncrementSeqNum(filterID int64) int64 {
	ret := _m.Called(filterID)

	if len(ret) == 0 {
		panic("no return value specified for IncrementSeqNum")
	}

	var r0 int64
	if rf, ok := ret.Get(0).(func(int64) int64); ok {
		r0 = rf(filterID)
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// mockFilters_IncrementSeqNum_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IncrementSeqNum'
type mockFilters_IncrementSeqNum_Call struct {
	*mock.Call
}

// IncrementSeqNum is a helper method to define mock.On call
//   - filterID int64
func (_e *mockFilters_Expecter) IncrementSeqNum(filterID interface{}) *mockFilters_IncrementSeqNum_Call {
	return &mockFilters_IncrementSeqNum_Call{Call: _e.mock.On("IncrementSeqNum", filterID)}
}

func (_c *mockFilters_IncrementSeqNum_Call) Run(run func(filterID int64)) *mockFilters_IncrementSeqNum_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64))
	})
	return _c
}

func (_c *mockFilters_IncrementSeqNum_Call) Return(_a0 int64) *mockFilters_IncrementSeqNum_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mockFilters_IncrementSeqNum_Call) RunAndReturn(run func(int64) int64) *mockFilters_IncrementSeqNum_Call {
	_c.Call.Return(run)
	return _c
}

// LoadFilters provides a mock function with given fields: ctx
func (_m *mockFilters) LoadFilters(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for LoadFilters")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// mockFilters_LoadFilters_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LoadFilters'
type mockFilters_LoadFilters_Call struct {
	*mock.Call
}

// LoadFilters is a helper method to define mock.On call
//   - ctx context.Context
func (_e *mockFilters_Expecter) LoadFilters(ctx interface{}) *mockFilters_LoadFilters_Call {
	return &mockFilters_LoadFilters_Call{Call: _e.mock.On("LoadFilters", ctx)}
}

func (_c *mockFilters_LoadFilters_Call) Run(run func(ctx context.Context)) *mockFilters_LoadFilters_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *mockFilters_LoadFilters_Call) Return(_a0 error) *mockFilters_LoadFilters_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mockFilters_LoadFilters_Call) RunAndReturn(run func(context.Context) error) *mockFilters_LoadFilters_Call {
	_c.Call.Return(run)
	return _c
}

// MarkFilterBackfilled provides a mock function with given fields: ctx, filterID
func (_m *mockFilters) MarkFilterBackfilled(ctx context.Context, filterID int64) error {
	ret := _m.Called(ctx, filterID)

	if len(ret) == 0 {
		panic("no return value specified for MarkFilterBackfilled")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, filterID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// mockFilters_MarkFilterBackfilled_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MarkFilterBackfilled'
type mockFilters_MarkFilterBackfilled_Call struct {
	*mock.Call
}

// MarkFilterBackfilled is a helper method to define mock.On call
//   - ctx context.Context
//   - filterID int64
func (_e *mockFilters_Expecter) MarkFilterBackfilled(ctx interface{}, filterID interface{}) *mockFilters_MarkFilterBackfilled_Call {
	return &mockFilters_MarkFilterBackfilled_Call{Call: _e.mock.On("MarkFilterBackfilled", ctx, filterID)}
}

func (_c *mockFilters_MarkFilterBackfilled_Call) Run(run func(ctx context.Context, filterID int64)) *mockFilters_MarkFilterBackfilled_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *mockFilters_MarkFilterBackfilled_Call) Return(_a0 error) *mockFilters_MarkFilterBackfilled_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mockFilters_MarkFilterBackfilled_Call) RunAndReturn(run func(context.Context, int64) error) *mockFilters_MarkFilterBackfilled_Call {
	_c.Call.Return(run)
	return _c
}

// MatchingFiltersForEncodedEvent provides a mock function with given fields: event
func (_m *mockFilters) MatchingFiltersForEncodedEvent(event ProgramEvent) iter.Seq[Filter] {
	ret := _m.Called(event)

	if len(ret) == 0 {
		panic("no return value specified for MatchingFiltersForEncodedEvent")
	}

	var r0 iter.Seq[Filter]
	if rf, ok := ret.Get(0).(func(ProgramEvent) iter.Seq[Filter]); ok {
		r0 = rf(event)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(iter.Seq[Filter])
		}
	}

	return r0
}

// mockFilters_MatchingFiltersForEncodedEvent_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MatchingFiltersForEncodedEvent'
type mockFilters_MatchingFiltersForEncodedEvent_Call struct {
	*mock.Call
}

// MatchingFiltersForEncodedEvent is a helper method to define mock.On call
//   - event ProgramEvent
func (_e *mockFilters_Expecter) MatchingFiltersForEncodedEvent(event interface{}) *mockFilters_MatchingFiltersForEncodedEvent_Call {
	return &mockFilters_MatchingFiltersForEncodedEvent_Call{Call: _e.mock.On("MatchingFiltersForEncodedEvent", event)}
}

func (_c *mockFilters_MatchingFiltersForEncodedEvent_Call) Run(run func(event ProgramEvent)) *mockFilters_MatchingFiltersForEncodedEvent_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(ProgramEvent))
	})
	return _c
}

func (_c *mockFilters_MatchingFiltersForEncodedEvent_Call) Return(_a0 iter.Seq[Filter]) *mockFilters_MatchingFiltersForEncodedEvent_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mockFilters_MatchingFiltersForEncodedEvent_Call) RunAndReturn(run func(ProgramEvent) iter.Seq[Filter]) *mockFilters_MatchingFiltersForEncodedEvent_Call {
	_c.Call.Return(run)
	return _c
}

// PruneFilters provides a mock function with given fields: ctx
func (_m *mockFilters) PruneFilters(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for PruneFilters")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// mockFilters_PruneFilters_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PruneFilters'
type mockFilters_PruneFilters_Call struct {
	*mock.Call
}

// PruneFilters is a helper method to define mock.On call
//   - ctx context.Context
func (_e *mockFilters_Expecter) PruneFilters(ctx interface{}) *mockFilters_PruneFilters_Call {
	return &mockFilters_PruneFilters_Call{Call: _e.mock.On("PruneFilters", ctx)}
}

func (_c *mockFilters_PruneFilters_Call) Run(run func(ctx context.Context)) *mockFilters_PruneFilters_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *mockFilters_PruneFilters_Call) Return(_a0 error) *mockFilters_PruneFilters_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mockFilters_PruneFilters_Call) RunAndReturn(run func(context.Context) error) *mockFilters_PruneFilters_Call {
	_c.Call.Return(run)
	return _c
}

// RegisterFilter provides a mock function with given fields: ctx, filter
func (_m *mockFilters) RegisterFilter(ctx context.Context, filter Filter) error {
	ret := _m.Called(ctx, filter)

	if len(ret) == 0 {
		panic("no return value specified for RegisterFilter")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, Filter) error); ok {
		r0 = rf(ctx, filter)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// mockFilters_RegisterFilter_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RegisterFilter'
type mockFilters_RegisterFilter_Call struct {
	*mock.Call
}

// RegisterFilter is a helper method to define mock.On call
//   - ctx context.Context
//   - filter Filter
func (_e *mockFilters_Expecter) RegisterFilter(ctx interface{}, filter interface{}) *mockFilters_RegisterFilter_Call {
	return &mockFilters_RegisterFilter_Call{Call: _e.mock.On("RegisterFilter", ctx, filter)}
}

func (_c *mockFilters_RegisterFilter_Call) Run(run func(ctx context.Context, filter Filter)) *mockFilters_RegisterFilter_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(Filter))
	})
	return _c
}

func (_c *mockFilters_RegisterFilter_Call) Return(_a0 error) *mockFilters_RegisterFilter_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mockFilters_RegisterFilter_Call) RunAndReturn(run func(context.Context, Filter) error) *mockFilters_RegisterFilter_Call {
	_c.Call.Return(run)
	return _c
}

// UnregisterFilter provides a mock function with given fields: ctx, name
func (_m *mockFilters) UnregisterFilter(ctx context.Context, name string) error {
	ret := _m.Called(ctx, name)

	if len(ret) == 0 {
		panic("no return value specified for UnregisterFilter")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// mockFilters_UnregisterFilter_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UnregisterFilter'
type mockFilters_UnregisterFilter_Call struct {
	*mock.Call
}

// UnregisterFilter is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
func (_e *mockFilters_Expecter) UnregisterFilter(ctx interface{}, name interface{}) *mockFilters_UnregisterFilter_Call {
	return &mockFilters_UnregisterFilter_Call{Call: _e.mock.On("UnregisterFilter", ctx, name)}
}

func (_c *mockFilters_UnregisterFilter_Call) Run(run func(ctx context.Context, name string)) *mockFilters_UnregisterFilter_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *mockFilters_UnregisterFilter_Call) Return(_a0 error) *mockFilters_UnregisterFilter_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mockFilters_UnregisterFilter_Call) RunAndReturn(run func(context.Context, string) error) *mockFilters_UnregisterFilter_Call {
	_c.Call.Return(run)
	return _c
}

// newMockFilters creates a new instance of mockFilters. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockFilters(t interface {
	mock.TestingT
	Cleanup(func())
}) *mockFilters {
	mock := &mockFilters{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
