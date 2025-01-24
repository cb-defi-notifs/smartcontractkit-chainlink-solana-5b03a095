// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	rpc "github.com/gagliardetto/solana-go/rpc"

	solana "github.com/gagliardetto/solana-go"
)

// RPCClient is an autogenerated mock type for the RPCClient type
type RPCClient struct {
	mock.Mock
}

type RPCClient_Expecter struct {
	mock *mock.Mock
}

func (_m *RPCClient) EXPECT() *RPCClient_Expecter {
	return &RPCClient_Expecter{mock: &_m.Mock}
}

// GetBlockWithOpts provides a mock function with given fields: _a0, _a1, _a2
func (_m *RPCClient) GetBlockWithOpts(_a0 context.Context, _a1 uint64, _a2 *rpc.GetBlockOpts) (*rpc.GetBlockResult, error) {
	ret := _m.Called(_a0, _a1, _a2)

	if len(ret) == 0 {
		panic("no return value specified for GetBlockWithOpts")
	}

	var r0 *rpc.GetBlockResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, *rpc.GetBlockOpts) (*rpc.GetBlockResult, error)); ok {
		return rf(_a0, _a1, _a2)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, *rpc.GetBlockOpts) *rpc.GetBlockResult); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rpc.GetBlockResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, *rpc.GetBlockOpts) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RPCClient_GetBlockWithOpts_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBlockWithOpts'
type RPCClient_GetBlockWithOpts_Call struct {
	*mock.Call
}

// GetBlockWithOpts is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 uint64
//   - _a2 *rpc.GetBlockOpts
func (_e *RPCClient_Expecter) GetBlockWithOpts(_a0 interface{}, _a1 interface{}, _a2 interface{}) *RPCClient_GetBlockWithOpts_Call {
	return &RPCClient_GetBlockWithOpts_Call{Call: _e.mock.On("GetBlockWithOpts", _a0, _a1, _a2)}
}

func (_c *RPCClient_GetBlockWithOpts_Call) Run(run func(_a0 context.Context, _a1 uint64, _a2 *rpc.GetBlockOpts)) *RPCClient_GetBlockWithOpts_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint64), args[2].(*rpc.GetBlockOpts))
	})
	return _c
}

func (_c *RPCClient_GetBlockWithOpts_Call) Return(_a0 *rpc.GetBlockResult, _a1 error) *RPCClient_GetBlockWithOpts_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *RPCClient_GetBlockWithOpts_Call) RunAndReturn(run func(context.Context, uint64, *rpc.GetBlockOpts) (*rpc.GetBlockResult, error)) *RPCClient_GetBlockWithOpts_Call {
	_c.Call.Return(run)
	return _c
}

// GetSignaturesForAddressWithOpts provides a mock function with given fields: _a0, _a1, _a2
func (_m *RPCClient) GetSignaturesForAddressWithOpts(_a0 context.Context, _a1 solana.PublicKey, _a2 *rpc.GetSignaturesForAddressOpts) ([]*rpc.TransactionSignature, error) {
	ret := _m.Called(_a0, _a1, _a2)

	if len(ret) == 0 {
		panic("no return value specified for GetSignaturesForAddressWithOpts")
	}

	var r0 []*rpc.TransactionSignature
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, solana.PublicKey, *rpc.GetSignaturesForAddressOpts) ([]*rpc.TransactionSignature, error)); ok {
		return rf(_a0, _a1, _a2)
	}
	if rf, ok := ret.Get(0).(func(context.Context, solana.PublicKey, *rpc.GetSignaturesForAddressOpts) []*rpc.TransactionSignature); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*rpc.TransactionSignature)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, solana.PublicKey, *rpc.GetSignaturesForAddressOpts) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RPCClient_GetSignaturesForAddressWithOpts_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSignaturesForAddressWithOpts'
type RPCClient_GetSignaturesForAddressWithOpts_Call struct {
	*mock.Call
}

// GetSignaturesForAddressWithOpts is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 solana.PublicKey
//   - _a2 *rpc.GetSignaturesForAddressOpts
func (_e *RPCClient_Expecter) GetSignaturesForAddressWithOpts(_a0 interface{}, _a1 interface{}, _a2 interface{}) *RPCClient_GetSignaturesForAddressWithOpts_Call {
	return &RPCClient_GetSignaturesForAddressWithOpts_Call{Call: _e.mock.On("GetSignaturesForAddressWithOpts", _a0, _a1, _a2)}
}

func (_c *RPCClient_GetSignaturesForAddressWithOpts_Call) Run(run func(_a0 context.Context, _a1 solana.PublicKey, _a2 *rpc.GetSignaturesForAddressOpts)) *RPCClient_GetSignaturesForAddressWithOpts_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(solana.PublicKey), args[2].(*rpc.GetSignaturesForAddressOpts))
	})
	return _c
}

func (_c *RPCClient_GetSignaturesForAddressWithOpts_Call) Return(_a0 []*rpc.TransactionSignature, _a1 error) *RPCClient_GetSignaturesForAddressWithOpts_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *RPCClient_GetSignaturesForAddressWithOpts_Call) RunAndReturn(run func(context.Context, solana.PublicKey, *rpc.GetSignaturesForAddressOpts) ([]*rpc.TransactionSignature, error)) *RPCClient_GetSignaturesForAddressWithOpts_Call {
	_c.Call.Return(run)
	return _c
}

// SlotHeightWithCommitment provides a mock function with given fields: ctx, commitment
func (_m *RPCClient) SlotHeightWithCommitment(ctx context.Context, commitment rpc.CommitmentType) (uint64, error) {
	ret := _m.Called(ctx, commitment)

	if len(ret) == 0 {
		panic("no return value specified for SlotHeightWithCommitment")
	}

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, rpc.CommitmentType) (uint64, error)); ok {
		return rf(ctx, commitment)
	}
	if rf, ok := ret.Get(0).(func(context.Context, rpc.CommitmentType) uint64); ok {
		r0 = rf(ctx, commitment)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, rpc.CommitmentType) error); ok {
		r1 = rf(ctx, commitment)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RPCClient_SlotHeightWithCommitment_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SlotHeightWithCommitment'
type RPCClient_SlotHeightWithCommitment_Call struct {
	*mock.Call
}

// SlotHeightWithCommitment is a helper method to define mock.On call
//   - ctx context.Context
//   - commitment rpc.CommitmentType
func (_e *RPCClient_Expecter) SlotHeightWithCommitment(ctx interface{}, commitment interface{}) *RPCClient_SlotHeightWithCommitment_Call {
	return &RPCClient_SlotHeightWithCommitment_Call{Call: _e.mock.On("SlotHeightWithCommitment", ctx, commitment)}
}

func (_c *RPCClient_SlotHeightWithCommitment_Call) Run(run func(ctx context.Context, commitment rpc.CommitmentType)) *RPCClient_SlotHeightWithCommitment_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(rpc.CommitmentType))
	})
	return _c
}

func (_c *RPCClient_SlotHeightWithCommitment_Call) Return(_a0 uint64, _a1 error) *RPCClient_SlotHeightWithCommitment_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *RPCClient_SlotHeightWithCommitment_Call) RunAndReturn(run func(context.Context, rpc.CommitmentType) (uint64, error)) *RPCClient_SlotHeightWithCommitment_Call {
	_c.Call.Return(run)
	return _c
}

// NewRPCClient creates a new instance of RPCClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRPCClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *RPCClient {
	mock := &RPCClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
