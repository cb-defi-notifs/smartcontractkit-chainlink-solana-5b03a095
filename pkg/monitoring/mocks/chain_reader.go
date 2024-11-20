// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	pkgsolana "github.com/smartcontractkit/chainlink-solana/pkg/solana"

	rpc "github.com/gagliardetto/solana-go/rpc"

	solana "github.com/gagliardetto/solana-go"
)

// ChainReader is an autogenerated mock type for the ChainReader type
type ChainReader struct {
	mock.Mock
}

type ChainReader_Expecter struct {
	mock *mock.Mock
}

func (_m *ChainReader) EXPECT() *ChainReader_Expecter {
	return &ChainReader_Expecter{mock: &_m.Mock}
}

// GetBalance provides a mock function with given fields: ctx, account, commitment
func (_m *ChainReader) GetBalance(ctx context.Context, account solana.PublicKey, commitment rpc.CommitmentType) (*rpc.GetBalanceResult, error) {
	ret := _m.Called(ctx, account, commitment)

	if len(ret) == 0 {
		panic("no return value specified for GetBalance")
	}

	var r0 *rpc.GetBalanceResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, solana.PublicKey, rpc.CommitmentType) (*rpc.GetBalanceResult, error)); ok {
		return rf(ctx, account, commitment)
	}
	if rf, ok := ret.Get(0).(func(context.Context, solana.PublicKey, rpc.CommitmentType) *rpc.GetBalanceResult); ok {
		r0 = rf(ctx, account, commitment)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rpc.GetBalanceResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, solana.PublicKey, rpc.CommitmentType) error); ok {
		r1 = rf(ctx, account, commitment)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ChainReader_GetBalance_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBalance'
type ChainReader_GetBalance_Call struct {
	*mock.Call
}

// GetBalance is a helper method to define mock.On call
//   - ctx context.Context
//   - account solana.PublicKey
//   - commitment rpc.CommitmentType
func (_e *ChainReader_Expecter) GetBalance(ctx interface{}, account interface{}, commitment interface{}) *ChainReader_GetBalance_Call {
	return &ChainReader_GetBalance_Call{Call: _e.mock.On("GetBalance", ctx, account, commitment)}
}

func (_c *ChainReader_GetBalance_Call) Run(run func(ctx context.Context, account solana.PublicKey, commitment rpc.CommitmentType)) *ChainReader_GetBalance_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(solana.PublicKey), args[2].(rpc.CommitmentType))
	})
	return _c
}

func (_c *ChainReader_GetBalance_Call) Return(out *rpc.GetBalanceResult, err error) *ChainReader_GetBalance_Call {
	_c.Call.Return(out, err)
	return _c
}

func (_c *ChainReader_GetBalance_Call) RunAndReturn(run func(context.Context, solana.PublicKey, rpc.CommitmentType) (*rpc.GetBalanceResult, error)) *ChainReader_GetBalance_Call {
	_c.Call.Return(run)
	return _c
}

// GetLatestBlock provides a mock function with given fields: ctx, commitment
func (_m *ChainReader) GetLatestBlock(ctx context.Context, commitment rpc.CommitmentType) (*rpc.GetBlockResult, error) {
	ret := _m.Called(ctx, commitment)

	if len(ret) == 0 {
		panic("no return value specified for GetLatestBlock")
	}

	var r0 *rpc.GetBlockResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, rpc.CommitmentType) (*rpc.GetBlockResult, error)); ok {
		return rf(ctx, commitment)
	}
	if rf, ok := ret.Get(0).(func(context.Context, rpc.CommitmentType) *rpc.GetBlockResult); ok {
		r0 = rf(ctx, commitment)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rpc.GetBlockResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, rpc.CommitmentType) error); ok {
		r1 = rf(ctx, commitment)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ChainReader_GetLatestBlock_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLatestBlock'
type ChainReader_GetLatestBlock_Call struct {
	*mock.Call
}

// GetLatestBlock is a helper method to define mock.On call
//   - ctx context.Context
//   - commitment rpc.CommitmentType
func (_e *ChainReader_Expecter) GetLatestBlock(ctx interface{}, commitment interface{}) *ChainReader_GetLatestBlock_Call {
	return &ChainReader_GetLatestBlock_Call{Call: _e.mock.On("GetLatestBlock", ctx, commitment)}
}

func (_c *ChainReader_GetLatestBlock_Call) Run(run func(ctx context.Context, commitment rpc.CommitmentType)) *ChainReader_GetLatestBlock_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(rpc.CommitmentType))
	})
	return _c
}

func (_c *ChainReader_GetLatestBlock_Call) Return(_a0 *rpc.GetBlockResult, _a1 error) *ChainReader_GetLatestBlock_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ChainReader_GetLatestBlock_Call) RunAndReturn(run func(context.Context, rpc.CommitmentType) (*rpc.GetBlockResult, error)) *ChainReader_GetLatestBlock_Call {
	_c.Call.Return(run)
	return _c
}

// GetLatestTransmission provides a mock function with given fields: ctx, account, commitment
func (_m *ChainReader) GetLatestTransmission(ctx context.Context, account solana.PublicKey, commitment rpc.CommitmentType) (pkgsolana.Answer, uint64, error) {
	ret := _m.Called(ctx, account, commitment)

	if len(ret) == 0 {
		panic("no return value specified for GetLatestTransmission")
	}

	var r0 pkgsolana.Answer
	var r1 uint64
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, solana.PublicKey, rpc.CommitmentType) (pkgsolana.Answer, uint64, error)); ok {
		return rf(ctx, account, commitment)
	}
	if rf, ok := ret.Get(0).(func(context.Context, solana.PublicKey, rpc.CommitmentType) pkgsolana.Answer); ok {
		r0 = rf(ctx, account, commitment)
	} else {
		r0 = ret.Get(0).(pkgsolana.Answer)
	}

	if rf, ok := ret.Get(1).(func(context.Context, solana.PublicKey, rpc.CommitmentType) uint64); ok {
		r1 = rf(ctx, account, commitment)
	} else {
		r1 = ret.Get(1).(uint64)
	}

	if rf, ok := ret.Get(2).(func(context.Context, solana.PublicKey, rpc.CommitmentType) error); ok {
		r2 = rf(ctx, account, commitment)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ChainReader_GetLatestTransmission_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLatestTransmission'
type ChainReader_GetLatestTransmission_Call struct {
	*mock.Call
}

// GetLatestTransmission is a helper method to define mock.On call
//   - ctx context.Context
//   - account solana.PublicKey
//   - commitment rpc.CommitmentType
func (_e *ChainReader_Expecter) GetLatestTransmission(ctx interface{}, account interface{}, commitment interface{}) *ChainReader_GetLatestTransmission_Call {
	return &ChainReader_GetLatestTransmission_Call{Call: _e.mock.On("GetLatestTransmission", ctx, account, commitment)}
}

func (_c *ChainReader_GetLatestTransmission_Call) Run(run func(ctx context.Context, account solana.PublicKey, commitment rpc.CommitmentType)) *ChainReader_GetLatestTransmission_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(solana.PublicKey), args[2].(rpc.CommitmentType))
	})
	return _c
}

func (_c *ChainReader_GetLatestTransmission_Call) Return(answer pkgsolana.Answer, blockHeight uint64, err error) *ChainReader_GetLatestTransmission_Call {
	_c.Call.Return(answer, blockHeight, err)
	return _c
}

func (_c *ChainReader_GetLatestTransmission_Call) RunAndReturn(run func(context.Context, solana.PublicKey, rpc.CommitmentType) (pkgsolana.Answer, uint64, error)) *ChainReader_GetLatestTransmission_Call {
	_c.Call.Return(run)
	return _c
}

// GetSignaturesForAddressWithOpts provides a mock function with given fields: ctx, account, opts
func (_m *ChainReader) GetSignaturesForAddressWithOpts(ctx context.Context, account solana.PublicKey, opts *rpc.GetSignaturesForAddressOpts) ([]*rpc.TransactionSignature, error) {
	ret := _m.Called(ctx, account, opts)

	if len(ret) == 0 {
		panic("no return value specified for GetSignaturesForAddressWithOpts")
	}

	var r0 []*rpc.TransactionSignature
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, solana.PublicKey, *rpc.GetSignaturesForAddressOpts) ([]*rpc.TransactionSignature, error)); ok {
		return rf(ctx, account, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, solana.PublicKey, *rpc.GetSignaturesForAddressOpts) []*rpc.TransactionSignature); ok {
		r0 = rf(ctx, account, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*rpc.TransactionSignature)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, solana.PublicKey, *rpc.GetSignaturesForAddressOpts) error); ok {
		r1 = rf(ctx, account, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ChainReader_GetSignaturesForAddressWithOpts_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSignaturesForAddressWithOpts'
type ChainReader_GetSignaturesForAddressWithOpts_Call struct {
	*mock.Call
}

// GetSignaturesForAddressWithOpts is a helper method to define mock.On call
//   - ctx context.Context
//   - account solana.PublicKey
//   - opts *rpc.GetSignaturesForAddressOpts
func (_e *ChainReader_Expecter) GetSignaturesForAddressWithOpts(ctx interface{}, account interface{}, opts interface{}) *ChainReader_GetSignaturesForAddressWithOpts_Call {
	return &ChainReader_GetSignaturesForAddressWithOpts_Call{Call: _e.mock.On("GetSignaturesForAddressWithOpts", ctx, account, opts)}
}

func (_c *ChainReader_GetSignaturesForAddressWithOpts_Call) Run(run func(ctx context.Context, account solana.PublicKey, opts *rpc.GetSignaturesForAddressOpts)) *ChainReader_GetSignaturesForAddressWithOpts_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(solana.PublicKey), args[2].(*rpc.GetSignaturesForAddressOpts))
	})
	return _c
}

func (_c *ChainReader_GetSignaturesForAddressWithOpts_Call) Return(out []*rpc.TransactionSignature, err error) *ChainReader_GetSignaturesForAddressWithOpts_Call {
	_c.Call.Return(out, err)
	return _c
}

func (_c *ChainReader_GetSignaturesForAddressWithOpts_Call) RunAndReturn(run func(context.Context, solana.PublicKey, *rpc.GetSignaturesForAddressOpts) ([]*rpc.TransactionSignature, error)) *ChainReader_GetSignaturesForAddressWithOpts_Call {
	_c.Call.Return(run)
	return _c
}

// GetSlot provides a mock function with given fields: ctx
func (_m *ChainReader) GetSlot(ctx context.Context) (uint64, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetSlot")
	}

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (uint64, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) uint64); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ChainReader_GetSlot_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSlot'
type ChainReader_GetSlot_Call struct {
	*mock.Call
}

// GetSlot is a helper method to define mock.On call
//   - ctx context.Context
func (_e *ChainReader_Expecter) GetSlot(ctx interface{}) *ChainReader_GetSlot_Call {
	return &ChainReader_GetSlot_Call{Call: _e.mock.On("GetSlot", ctx)}
}

func (_c *ChainReader_GetSlot_Call) Run(run func(ctx context.Context)) *ChainReader_GetSlot_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *ChainReader_GetSlot_Call) Return(slot uint64, err error) *ChainReader_GetSlot_Call {
	_c.Call.Return(slot, err)
	return _c
}

func (_c *ChainReader_GetSlot_Call) RunAndReturn(run func(context.Context) (uint64, error)) *ChainReader_GetSlot_Call {
	_c.Call.Return(run)
	return _c
}

// GetState provides a mock function with given fields: ctx, account, commitment
func (_m *ChainReader) GetState(ctx context.Context, account solana.PublicKey, commitment rpc.CommitmentType) (pkgsolana.State, uint64, error) {
	ret := _m.Called(ctx, account, commitment)

	if len(ret) == 0 {
		panic("no return value specified for GetState")
	}

	var r0 pkgsolana.State
	var r1 uint64
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, solana.PublicKey, rpc.CommitmentType) (pkgsolana.State, uint64, error)); ok {
		return rf(ctx, account, commitment)
	}
	if rf, ok := ret.Get(0).(func(context.Context, solana.PublicKey, rpc.CommitmentType) pkgsolana.State); ok {
		r0 = rf(ctx, account, commitment)
	} else {
		r0 = ret.Get(0).(pkgsolana.State)
	}

	if rf, ok := ret.Get(1).(func(context.Context, solana.PublicKey, rpc.CommitmentType) uint64); ok {
		r1 = rf(ctx, account, commitment)
	} else {
		r1 = ret.Get(1).(uint64)
	}

	if rf, ok := ret.Get(2).(func(context.Context, solana.PublicKey, rpc.CommitmentType) error); ok {
		r2 = rf(ctx, account, commitment)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ChainReader_GetState_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetState'
type ChainReader_GetState_Call struct {
	*mock.Call
}

// GetState is a helper method to define mock.On call
//   - ctx context.Context
//   - account solana.PublicKey
//   - commitment rpc.CommitmentType
func (_e *ChainReader_Expecter) GetState(ctx interface{}, account interface{}, commitment interface{}) *ChainReader_GetState_Call {
	return &ChainReader_GetState_Call{Call: _e.mock.On("GetState", ctx, account, commitment)}
}

func (_c *ChainReader_GetState_Call) Run(run func(ctx context.Context, account solana.PublicKey, commitment rpc.CommitmentType)) *ChainReader_GetState_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(solana.PublicKey), args[2].(rpc.CommitmentType))
	})
	return _c
}

func (_c *ChainReader_GetState_Call) Return(state pkgsolana.State, blockHeight uint64, err error) *ChainReader_GetState_Call {
	_c.Call.Return(state, blockHeight, err)
	return _c
}

func (_c *ChainReader_GetState_Call) RunAndReturn(run func(context.Context, solana.PublicKey, rpc.CommitmentType) (pkgsolana.State, uint64, error)) *ChainReader_GetState_Call {
	_c.Call.Return(run)
	return _c
}

// GetTokenAccountBalance provides a mock function with given fields: ctx, account, commitment
func (_m *ChainReader) GetTokenAccountBalance(ctx context.Context, account solana.PublicKey, commitment rpc.CommitmentType) (*rpc.GetTokenAccountBalanceResult, error) {
	ret := _m.Called(ctx, account, commitment)

	if len(ret) == 0 {
		panic("no return value specified for GetTokenAccountBalance")
	}

	var r0 *rpc.GetTokenAccountBalanceResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, solana.PublicKey, rpc.CommitmentType) (*rpc.GetTokenAccountBalanceResult, error)); ok {
		return rf(ctx, account, commitment)
	}
	if rf, ok := ret.Get(0).(func(context.Context, solana.PublicKey, rpc.CommitmentType) *rpc.GetTokenAccountBalanceResult); ok {
		r0 = rf(ctx, account, commitment)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rpc.GetTokenAccountBalanceResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, solana.PublicKey, rpc.CommitmentType) error); ok {
		r1 = rf(ctx, account, commitment)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ChainReader_GetTokenAccountBalance_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTokenAccountBalance'
type ChainReader_GetTokenAccountBalance_Call struct {
	*mock.Call
}

// GetTokenAccountBalance is a helper method to define mock.On call
//   - ctx context.Context
//   - account solana.PublicKey
//   - commitment rpc.CommitmentType
func (_e *ChainReader_Expecter) GetTokenAccountBalance(ctx interface{}, account interface{}, commitment interface{}) *ChainReader_GetTokenAccountBalance_Call {
	return &ChainReader_GetTokenAccountBalance_Call{Call: _e.mock.On("GetTokenAccountBalance", ctx, account, commitment)}
}

func (_c *ChainReader_GetTokenAccountBalance_Call) Run(run func(ctx context.Context, account solana.PublicKey, commitment rpc.CommitmentType)) *ChainReader_GetTokenAccountBalance_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(solana.PublicKey), args[2].(rpc.CommitmentType))
	})
	return _c
}

func (_c *ChainReader_GetTokenAccountBalance_Call) Return(out *rpc.GetTokenAccountBalanceResult, err error) *ChainReader_GetTokenAccountBalance_Call {
	_c.Call.Return(out, err)
	return _c
}

func (_c *ChainReader_GetTokenAccountBalance_Call) RunAndReturn(run func(context.Context, solana.PublicKey, rpc.CommitmentType) (*rpc.GetTokenAccountBalanceResult, error)) *ChainReader_GetTokenAccountBalance_Call {
	_c.Call.Return(run)
	return _c
}

// GetTransaction provides a mock function with given fields: ctx, txSig, opts
func (_m *ChainReader) GetTransaction(ctx context.Context, txSig solana.Signature, opts *rpc.GetTransactionOpts) (*rpc.GetTransactionResult, error) {
	ret := _m.Called(ctx, txSig, opts)

	if len(ret) == 0 {
		panic("no return value specified for GetTransaction")
	}

	var r0 *rpc.GetTransactionResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, solana.Signature, *rpc.GetTransactionOpts) (*rpc.GetTransactionResult, error)); ok {
		return rf(ctx, txSig, opts)
	}
	if rf, ok := ret.Get(0).(func(context.Context, solana.Signature, *rpc.GetTransactionOpts) *rpc.GetTransactionResult); ok {
		r0 = rf(ctx, txSig, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rpc.GetTransactionResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, solana.Signature, *rpc.GetTransactionOpts) error); ok {
		r1 = rf(ctx, txSig, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ChainReader_GetTransaction_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTransaction'
type ChainReader_GetTransaction_Call struct {
	*mock.Call
}

// GetTransaction is a helper method to define mock.On call
//   - ctx context.Context
//   - txSig solana.Signature
//   - opts *rpc.GetTransactionOpts
func (_e *ChainReader_Expecter) GetTransaction(ctx interface{}, txSig interface{}, opts interface{}) *ChainReader_GetTransaction_Call {
	return &ChainReader_GetTransaction_Call{Call: _e.mock.On("GetTransaction", ctx, txSig, opts)}
}

func (_c *ChainReader_GetTransaction_Call) Run(run func(ctx context.Context, txSig solana.Signature, opts *rpc.GetTransactionOpts)) *ChainReader_GetTransaction_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(solana.Signature), args[2].(*rpc.GetTransactionOpts))
	})
	return _c
}

func (_c *ChainReader_GetTransaction_Call) Return(out *rpc.GetTransactionResult, err error) *ChainReader_GetTransaction_Call {
	_c.Call.Return(out, err)
	return _c
}

func (_c *ChainReader_GetTransaction_Call) RunAndReturn(run func(context.Context, solana.Signature, *rpc.GetTransactionOpts) (*rpc.GetTransactionResult, error)) *ChainReader_GetTransaction_Call {
	_c.Call.Return(run)
	return _c
}

// NewChainReader creates a new instance of ChainReader. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewChainReader(t interface {
	mock.TestingT
	Cleanup(func())
}) *ChainReader {
	mock := &ChainReader{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
