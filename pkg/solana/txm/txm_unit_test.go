package txm_test

import (
	"errors"
	"math/big"
	"testing"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/system"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	solanaClient "github.com/smartcontractkit/chainlink-solana/pkg/solana/client"
	clientmocks "github.com/smartcontractkit/chainlink-solana/pkg/solana/client/mocks"
	"github.com/smartcontractkit/chainlink-solana/pkg/solana/config"
	"github.com/smartcontractkit/chainlink-solana/pkg/solana/fees"
	solanatxm "github.com/smartcontractkit/chainlink-solana/pkg/solana/txm"
	keyMocks "github.com/smartcontractkit/chainlink-solana/pkg/solana/txm/mocks"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/utils"
	bigmath "github.com/smartcontractkit/chainlink-common/pkg/utils/big_math"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"
)

func TestTxm_EstimateComputeUnitLimit(t *testing.T) {
	t.Parallel()
	ctx := tests.Context(t)

	// setup mock keystore
	mkey := keyMocks.NewSimpleKeystore(t)
	mkey.On("Sign", mock.Anything, mock.Anything, mock.Anything).Return([]byte{}, nil)

	// setup key
	key, err := solana.NewRandomPrivateKey()
	require.NoError(t, err)
	pubKey := key.PublicKey()

	// setup receiver key
	privKeyReceiver, err := solana.NewRandomPrivateKey()
	require.NoError(t, err)
	pubKeyReceiver := privKeyReceiver.PublicKey()

	// set up txm
	lggr := logger.Test(t)
	cfg := config.NewDefault()
	client := clientmocks.NewReaderWriter(t)
	require.NoError(t, err)
	loader := utils.NewLazyLoad(func() (solanaClient.ReaderWriter, error) { return client, nil })
	txm := solanatxm.NewTxm("localnet", loader, nil, cfg, mkey, lggr)

	t.Run("successfully sets estimated compute unit limit", func(t *testing.T) {
		usedCompute := uint64(100)
		client.On("LatestBlockhash", mock.Anything).Return(&rpc.GetLatestBlockhashResult{
			Value: &rpc.LatestBlockhashResult{
				LastValidBlockHeight: 100,
				Blockhash:            solana.Hash{},
			},
		}, nil).Once()
		client.On("SimulateTx", mock.Anything, mock.IsType(&solana.Transaction{}), mock.IsType(&rpc.SimulateTransactionOpts{})).Run(func(args mock.Arguments) {
			// Validate max compute unit limit is set in transaction
			tx := args.Get(1).(*solana.Transaction)
			limit, parseErr := fees.ParseComputeUnitLimit(tx.Message.Instructions[len(tx.Message.Instructions)-1].Data)
			require.NoError(t, parseErr)
			require.Equal(t, fees.ComputeUnitLimit(solanatxm.MaxComputeUnitLimit), limit)

			// Validate signature verification is enabled
			opts := args.Get(2).(*rpc.SimulateTransactionOpts)
			require.True(t, opts.SigVerify)
		}).Return(&rpc.SimulateTransactionResult{
			Err:           nil,
			UnitsConsumed: &usedCompute,
		}, nil).Once()
		tx := createTx(t, client, pubKey, pubKey, pubKeyReceiver, solana.LAMPORTS_PER_SOL)
		computeUnitLimit, estimateErr := txm.EstimateComputeUnitLimit(ctx, tx, "")
		require.NoError(t, estimateErr)
		usedComputeWithBuffer := bigmath.AddPercentage(new(big.Int).SetUint64(usedCompute), solanatxm.EstimateComputeUnitLimitBuffer).Uint64()
		require.Equal(t, usedComputeWithBuffer, uint64(computeUnitLimit))
	})

	t.Run("failed to simulate tx", func(t *testing.T) {
		client.On("LatestBlockhash", mock.Anything).Return(&rpc.GetLatestBlockhashResult{
			Value: &rpc.LatestBlockhashResult{
				LastValidBlockHeight: 100,
				Blockhash:            solana.Hash{},
			},
		}, nil).Once()
		client.On("SimulateTx", mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("failed to simulate")).Once()
		tx := createTx(t, client, pubKey, pubKey, pubKeyReceiver, solana.LAMPORTS_PER_SOL)
		_, estimateErr := txm.EstimateComputeUnitLimit(ctx, tx, "")
		require.Error(t, estimateErr)
	})

	t.Run("simulation returns error for tx", func(t *testing.T) {
		client.On("LatestBlockhash", mock.Anything).Return(&rpc.GetLatestBlockhashResult{
			Value: &rpc.LatestBlockhashResult{
				LastValidBlockHeight: 100,
				Blockhash:            solana.Hash{},
			},
		}, nil).Once()
		client.On("SimulateTx", mock.Anything, mock.Anything, mock.Anything).Return(&rpc.SimulateTransactionResult{
			Err: errors.New("InstructionError"),
		}, nil).Once()
		tx := createTx(t, client, pubKey, pubKey, pubKeyReceiver, solana.LAMPORTS_PER_SOL)
		_, err = txm.EstimateComputeUnitLimit(ctx, tx, "")
		require.Error(t, err)
	})

	t.Run("simulation returns nil err with 0 compute unit limit", func(t *testing.T) {
		client.On("LatestBlockhash", mock.Anything).Return(&rpc.GetLatestBlockhashResult{
			Value: &rpc.LatestBlockhashResult{
				LastValidBlockHeight: 100,
				Blockhash:            solana.Hash{},
			},
		}, nil).Once()
		client.On("SimulateTx", mock.Anything, mock.Anything, mock.Anything).Return(&rpc.SimulateTransactionResult{
			Err: nil,
		}, nil).Once()
		tx := createTx(t, client, pubKey, pubKey, pubKeyReceiver, solana.LAMPORTS_PER_SOL)
		computeUnitLimit, err := txm.EstimateComputeUnitLimit(ctx, tx, "")
		require.NoError(t, err)
		require.Equal(t, uint32(0), computeUnitLimit)
	})

	t.Run("simulation returns max compute unit limit if adding buffer exceeds it", func(t *testing.T) {
		usedCompute := uint64(1_400_000)
		client.On("LatestBlockhash", mock.Anything).Return(&rpc.GetLatestBlockhashResult{
			Value: &rpc.LatestBlockhashResult{
				LastValidBlockHeight: 100,
				Blockhash:            solana.Hash{},
			},
		}, nil).Once()
		client.On("SimulateTx", mock.Anything, mock.IsType(&solana.Transaction{}), mock.IsType(&rpc.SimulateTransactionOpts{})).Run(func(args mock.Arguments) {
			// Validate max compute unit limit is set in transaction
			tx := args.Get(1).(*solana.Transaction)
			limit, err := fees.ParseComputeUnitLimit(tx.Message.Instructions[len(tx.Message.Instructions)-1].Data)
			require.NoError(t, err)
			require.Equal(t, fees.ComputeUnitLimit(solanatxm.MaxComputeUnitLimit), limit)

			// Validate signature verification is enabled
			opts := args.Get(2).(*rpc.SimulateTransactionOpts)
			require.True(t, opts.SigVerify)
		}).Return(&rpc.SimulateTransactionResult{
			Err:           nil,
			UnitsConsumed: &usedCompute,
		}, nil).Once()
		tx := createTx(t, client, pubKey, pubKey, pubKeyReceiver, solana.LAMPORTS_PER_SOL)
		computeUnitLimit, err := txm.EstimateComputeUnitLimit(ctx, tx, "")
		require.NoError(t, err)
		require.Equal(t, uint32(1_400_000), computeUnitLimit)
	})
}

func createTx(t *testing.T, client solanaClient.ReaderWriter, signer solana.PublicKey, sender solana.PublicKey, receiver solana.PublicKey, amt uint64) *solana.Transaction {
	// create transfer tx
	hash, err := client.LatestBlockhash(tests.Context(t))
	require.NoError(t, err)
	tx, err := solana.NewTransaction(
		[]solana.Instruction{
			system.NewTransferInstruction(
				amt,
				sender,
				receiver,
			).Build(),
		},
		hash.Value.Blockhash,
		solana.TransactionPayer(signer),
	)
	require.NoError(t, err)
	return tx
}
