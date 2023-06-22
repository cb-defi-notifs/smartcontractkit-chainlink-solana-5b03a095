package types

import (
	"strconv"
	"testing"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/smartcontractkit/chainlink-solana/pkg/internal/utils"
	"github.com/stretchr/testify/assert"
)

func configureBlockResult() rpc.GetBlockResult {
	result := rpc.GetBlockResult{
		Blockhash:         utils.NewHash(),
		PreviousBlockhash: utils.NewHash(),
		ParentSlot:        0,
		Transactions:      []rpc.TransactionWithMeta{},
		Signatures:        []solana.Signature{},
		Rewards:           []rpc.BlockReward{},
		BlockTime:         nil,
		BlockHeight:       nil,
	}
	return result
}

func TestHead_NewHead(t *testing.T) {
	emptyBlockResult := configureBlockResult()
	t.Parallel()

	tests := []struct {
		slot     int64
		block    rpc.GetBlockResult
		parent   *Head
		id       ChainID
		wantSlot int64
	}{
		// with no parent
		{10, emptyBlockResult, nil, Mainnet, 10},
		// with parent
		{20, emptyBlockResult,
			NewHead(10, emptyBlockResult, nil, Mainnet),
			Mainnet, 20},
		{30, emptyBlockResult,
			NewHead(20, emptyBlockResult,
				NewHead(10, emptyBlockResult, nil, Mainnet),
				Mainnet),
			Mainnet, 30},
	}

	for _, test := range tests {
		t.Run(
			strconv.FormatInt(test.wantSlot, 10), // convert to base 10
			func(t *testing.T) {
				head := NewHead(test.slot, test.block, test.parent, test.id)
				assert.Equal(t, test.wantSlot, head.Slot)
				assert.Equal(t, test.block, head.Block)
				assert.Equal(t, test.parent, head.Parent)
				assert.Equal(t, test.id, head.ID)
			})
	}
}

func TestHead_ChainLength(t *testing.T) {
	blockResult := configureBlockResult()
	id := Mainnet

	head := NewHead(0, blockResult, NewHead(0, blockResult, NewHead(0, blockResult, nil, id), id), id)

	assert.Equal(t, uint32(3), head.ChainLength())

	var head2 *Head
	assert.Equal(t, uint32(0), head2.ChainLength())
}

func TestHead_EarliestHeadInChain(t *testing.T) {
	blockResult := configureBlockResult()
	id := Mainnet

	head := NewHead(3, blockResult,
		NewHead(2, blockResult,
			NewHead(1, blockResult, nil, id), id), id)

	assert.Equal(t, int64(1), head.EarliestHeadInChain().BlockNumber())
}

func TestHead_GetParentHash(t *testing.T) {
	blockResult := configureBlockResult()
	id := Mainnet

	head := NewHead(3, blockResult,
		NewHead(2, blockResult,
			NewHead(1, blockResult, nil, id), id), id)

	assert.Equal(t, head.Parent.BlockHash(), head.GetParentHash())
}

func TestHead_GetParent(t *testing.T) {
	blockResult := configureBlockResult()
	id := Mainnet

	head := NewHead(3, blockResult,
		NewHead(2, blockResult,
			NewHead(1, blockResult, nil, id), id), id)

	assert.Equal(t, head.Parent, head.GetParent())
}

func TestHead_HasChainID(t *testing.T) {
	t.Parallel()
	blockResult := configureBlockResult() // Assuming this function creates a mock rpc.GetBlockResult

	tests := []struct {
		name    string
		chainID ChainID
		want    bool
	}{
		{
			"HasChainID returns true when ChainID is not 'unknown'",
			Devnet, // replace with correct initialization
			true,
		},
		{
			"HasChainID returns false when ChainID is 'unknown'",
			99,
			false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			head := NewHead(0, blockResult, nil, test.chainID)
			assert.Equal(t, test.want, head.HasChainID())
		})
	}

	t.Run("HasChainID returns false when Head is nil", func(t *testing.T) {
		var head *Head
		assert.False(t, head.HasChainID())
	})
}
