// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	metrics "github.com/smartcontractkit/chainlink-solana/pkg/monitoring/metrics"
	mock "github.com/stretchr/testify/mock"
)

// NodeSuccess is an autogenerated mock type for the NodeSuccess type
type NodeSuccess struct {
	mock.Mock
}

type NodeSuccess_Expecter struct {
	mock *mock.Mock
}

func (_m *NodeSuccess) EXPECT() *NodeSuccess_Expecter {
	return &NodeSuccess_Expecter{mock: &_m.Mock}
}

// Add provides a mock function with given fields: count, i
func (_m *NodeSuccess) Add(count int, i metrics.NodeFeedInput) {
	_m.Called(count, i)
}

// NodeSuccess_Add_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Add'
type NodeSuccess_Add_Call struct {
	*mock.Call
}

// Add is a helper method to define mock.On call
//   - count int
//   - i metrics.NodeFeedInput
func (_e *NodeSuccess_Expecter) Add(count interface{}, i interface{}) *NodeSuccess_Add_Call {
	return &NodeSuccess_Add_Call{Call: _e.mock.On("Add", count, i)}
}

func (_c *NodeSuccess_Add_Call) Run(run func(count int, i metrics.NodeFeedInput)) *NodeSuccess_Add_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int), args[1].(metrics.NodeFeedInput))
	})
	return _c
}

func (_c *NodeSuccess_Add_Call) Return() *NodeSuccess_Add_Call {
	_c.Call.Return()
	return _c
}

func (_c *NodeSuccess_Add_Call) RunAndReturn(run func(int, metrics.NodeFeedInput)) *NodeSuccess_Add_Call {
	_c.Call.Return(run)
	return _c
}

// Cleanup provides a mock function with given fields: i
func (_m *NodeSuccess) Cleanup(i metrics.NodeFeedInput) {
	_m.Called(i)
}

// NodeSuccess_Cleanup_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Cleanup'
type NodeSuccess_Cleanup_Call struct {
	*mock.Call
}

// Cleanup is a helper method to define mock.On call
//   - i metrics.NodeFeedInput
func (_e *NodeSuccess_Expecter) Cleanup(i interface{}) *NodeSuccess_Cleanup_Call {
	return &NodeSuccess_Cleanup_Call{Call: _e.mock.On("Cleanup", i)}
}

func (_c *NodeSuccess_Cleanup_Call) Run(run func(i metrics.NodeFeedInput)) *NodeSuccess_Cleanup_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(metrics.NodeFeedInput))
	})
	return _c
}

func (_c *NodeSuccess_Cleanup_Call) Return() *NodeSuccess_Cleanup_Call {
	_c.Call.Return()
	return _c
}

func (_c *NodeSuccess_Cleanup_Call) RunAndReturn(run func(metrics.NodeFeedInput)) *NodeSuccess_Cleanup_Call {
	_c.Call.Return(run)
	return _c
}

// NewNodeSuccess creates a new instance of NodeSuccess. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewNodeSuccess(t interface {
	mock.TestingT
	Cleanup(func())
}) *NodeSuccess {
	mock := &NodeSuccess{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
