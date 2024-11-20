// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	metrics "github.com/smartcontractkit/chainlink-solana/pkg/monitoring/metrics"
	mock "github.com/stretchr/testify/mock"
)

// NetworkFees is an autogenerated mock type for the NetworkFees type
type NetworkFees struct {
	mock.Mock
}

type NetworkFees_Expecter struct {
	mock *mock.Mock
}

func (_m *NetworkFees) EXPECT() *NetworkFees_Expecter {
	return &NetworkFees_Expecter{mock: &_m.Mock}
}

// Cleanup provides a mock function with given fields:
func (_m *NetworkFees) Cleanup() {
	_m.Called()
}

// NetworkFees_Cleanup_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Cleanup'
type NetworkFees_Cleanup_Call struct {
	*mock.Call
}

// Cleanup is a helper method to define mock.On call
func (_e *NetworkFees_Expecter) Cleanup() *NetworkFees_Cleanup_Call {
	return &NetworkFees_Cleanup_Call{Call: _e.mock.On("Cleanup")}
}

func (_c *NetworkFees_Cleanup_Call) Run(run func()) *NetworkFees_Cleanup_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *NetworkFees_Cleanup_Call) Return() *NetworkFees_Cleanup_Call {
	_c.Call.Return()
	return _c
}

func (_c *NetworkFees_Cleanup_Call) RunAndReturn(run func()) *NetworkFees_Cleanup_Call {
	_c.Call.Return(run)
	return _c
}

// Set provides a mock function with given fields: slot, chain
func (_m *NetworkFees) Set(slot metrics.NetworkFeesInput, chain string) {
	_m.Called(slot, chain)
}

// NetworkFees_Set_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Set'
type NetworkFees_Set_Call struct {
	*mock.Call
}

// Set is a helper method to define mock.On call
//   - slot metrics.NetworkFeesInput
//   - chain string
func (_e *NetworkFees_Expecter) Set(slot interface{}, chain interface{}) *NetworkFees_Set_Call {
	return &NetworkFees_Set_Call{Call: _e.mock.On("Set", slot, chain)}
}

func (_c *NetworkFees_Set_Call) Run(run func(slot metrics.NetworkFeesInput, chain string)) *NetworkFees_Set_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(metrics.NetworkFeesInput), args[1].(string))
	})
	return _c
}

func (_c *NetworkFees_Set_Call) Return() *NetworkFees_Set_Call {
	_c.Call.Return()
	return _c
}

func (_c *NetworkFees_Set_Call) RunAndReturn(run func(metrics.NetworkFeesInput, string)) *NetworkFees_Set_Call {
	_c.Call.Return(run)
	return _c
}

// NewNetworkFees creates a new instance of NetworkFees. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewNetworkFees(t interface {
	mock.TestingT
	Cleanup(func())
}) *NetworkFees {
	mock := &NetworkFees{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
