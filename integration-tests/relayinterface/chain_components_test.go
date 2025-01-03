/*
Package relayinterface contains the interface tests for chain components.
*/
package relayinterface

import (
	"context"
	"encoding/binary"
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"sync"
	"testing"
	"time"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/gagliardetto/solana-go/rpc/ws"
	"github.com/gagliardetto/solana-go/text"
	"github.com/stretchr/testify/require"

	commoncodec "github.com/smartcontractkit/chainlink-common/pkg/codec"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	commontestutils "github.com/smartcontractkit/chainlink-common/pkg/loop/testutils"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	. "github.com/smartcontractkit/chainlink-common/pkg/types/interfacetests" //nolint common practice to import test mods with .
	"github.com/smartcontractkit/chainlink-common/pkg/types/query/primitives"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"

	"github.com/smartcontractkit/chainlink-solana/pkg/solana/codec"

	contract "github.com/smartcontractkit/chainlink-solana/contracts/generated/contract_reader_interface"
	"github.com/smartcontractkit/chainlink-solana/integration-tests/solclient"
	"github.com/smartcontractkit/chainlink-solana/integration-tests/utils"
	"github.com/smartcontractkit/chainlink-solana/pkg/solana/chainreader"
	"github.com/smartcontractkit/chainlink-solana/pkg/solana/config"
	solanautils "github.com/smartcontractkit/chainlink-solana/pkg/solana/utils"
)

func TestChainComponents(t *testing.T) {
	t.Parallel()
	helper := &helper{}
	helper.Init(t)

	t.Run("RunChainComponentsSolanaTests", func(t *testing.T) {
		t.Parallel()
		it := &SolanaChainComponentsInterfaceTester[*testing.T]{Helper: helper}
		DisableTests(it)
		it.Setup(t)
		RunChainComponentsSolanaTests(t, it)
	})

	t.Run("RunChainComponentsInLoopSolanaTests", func(t *testing.T) {
		t.Parallel()
		it := &SolanaChainComponentsInterfaceTester[*testing.T]{Helper: helper}
		DisableTests(it)
		wrapped := commontestutils.WrapContractReaderTesterForLoop(it)
		wrapped.Setup(t)
		RunChainComponentsInLoopSolanaTests(t, wrapped)
	})
}

func DisableTests(it *SolanaChainComponentsInterfaceTester[*testing.T]) {
	it.DisableTests([]string{
		// disable tests that set values
		ContractReaderGetLatestValueBasedOnConfidenceLevel,
		// disable anything returning a struct or requiring input params for now
		ContractReaderGetLatestValueAsValuesDotValue,
		ContractReaderGetLatestValue,
		ContractReaderGetLatestValueWithModifiersUsingOwnMapstrctureOverrides,
		// events not yet supported
		ContractReaderGetLatestValueGetsLatestForEvent,
		ContractReaderGetLatestValueBasedOnConfidenceLevelForEvent,
		ContractReaderGetLatestValueReturnsNotFoundWhenNotTriggeredForEvent,
		ContractReaderGetLatestValueWithFilteringForEvent,
		// disable anything in batch relating to input params or structs for now
		ContractReaderBatchGetLatestValue,
		ContractReaderBatchGetLatestValueWithModifiersOwnMapstructureOverride,
		ContractReaderBatchGetLatestValueDifferentParamsResultsRetainOrder,
		ContractReaderBatchGetLatestValueDifferentParamsResultsRetainOrderMultipleContracts,
		ContractReaderBatchGetLatestValueSetsErrorsProperly,
		// query key not implemented yet
		ContractReaderQueryKeyNotFound,
		ContractReaderQueryKeyReturnsData,
		ContractReaderQueryKeyReturnsDataAsValuesDotValue,
		ContractReaderQueryKeyCanFilterWithValueComparator,
		ContractReaderQueryKeyCanLimitResultsWithCursor,
		ContractReaderQueryKeysReturnsDataTwoEventTypes,
		ContractReaderQueryKeysNotFound,
		ContractReaderQueryKeysReturnsData,
		ContractReaderQueryKeysReturnsDataAsValuesDotValue,
		ContractReaderQueryKeysCanFilterWithValueComparator,
		ContractReaderQueryKeysCanLimitResultsWithCursor,
	})
}

func RunChainComponentsSolanaTests[T TestingT[T]](t T, it *SolanaChainComponentsInterfaceTester[T]) {
	RunContractReaderSolanaTests(t, it)
	// Add ChainWriter tests here
}

func RunChainComponentsInLoopSolanaTests[T TestingT[T]](t T, it ChainComponentsInterfaceTester[T]) {
	RunContractReaderInLoopTests(t, it)
	// Add ChainWriter tests here
}

func RunContractReaderSolanaTests[T TestingT[T]](t T, it *SolanaChainComponentsInterfaceTester[T]) {
	RunContractReaderInterfaceTests(t, it, false, true)

	var testCases []Testcase[T]

	RunTests(t, it, testCases)
}

func RunContractReaderInLoopTests[T TestingT[T]](t T, it ChainComponentsInterfaceTester[T]) {
	RunContractReaderInterfaceTests(t, it, false, true)

	var testCases []Testcase[T]

	RunTests(t, it, testCases)
}

type SolanaChainComponentsInterfaceTesterHelper[T TestingT[T]] interface {
	Init(t T)
	RPCClient() *chainreader.RPCClientWrapper
	Context(t T) context.Context
	Logger(t T) logger.Logger
	GetJSONEncodedIDL(t T) []byte
	CreateAccount(t T, value uint64) solana.PublicKey
}

type SolanaChainComponentsInterfaceTester[T TestingT[T]] struct {
	TestSelectionSupport
	Helper               SolanaChainComponentsInterfaceTesterHelper[T]
	cr                   *chainreader.SolanaChainReaderService
	contractReaderConfig config.ContractReader
}

func (it *SolanaChainComponentsInterfaceTester[T]) Setup(t T) {
	t.Cleanup(func() {})

	it.contractReaderConfig = config.ContractReader{
		Namespaces: map[string]config.ChainContractReader{
			AnyContractName: {
				IDL: mustUnmarshalIDL(t, string(it.Helper.GetJSONEncodedIDL(t))),
				Reads: map[string]config.ReadDefinition{
					MethodReturningUint64: {
						ChainSpecificName: "DataAccount",
						ReadType:          config.Account,
						OutputModifications: commoncodec.ModifiersConfig{
							&commoncodec.PropertyExtractorConfig{FieldName: "U64Value"},
						},
					},
					MethodReturningUint64Slice: {
						ChainSpecificName: "DataAccount",
						OutputModifications: commoncodec.ModifiersConfig{
							&commoncodec.PropertyExtractorConfig{FieldName: "U64Slice"},
						},
					},
				},
			},
			AnySecondContractName: {
				IDL: mustUnmarshalIDL(t, string(it.Helper.GetJSONEncodedIDL(t))),
				Reads: map[string]config.ReadDefinition{
					MethodReturningUint64: {
						ChainSpecificName: "DataAccount",
						OutputModifications: commoncodec.ModifiersConfig{
							&commoncodec.PropertyExtractorConfig{FieldName: "U64Value"},
						},
					},
				},
			},
		},
	}
}

func (it *SolanaChainComponentsInterfaceTester[T]) Name() string {
	return ""
}

func (it *SolanaChainComponentsInterfaceTester[T]) GetAccountBytes(i int) []byte {
	return nil
}

func (it *SolanaChainComponentsInterfaceTester[T]) GetAccountString(i int) string {
	return ""
}

func (it *SolanaChainComponentsInterfaceTester[T]) GetContractReader(t T) types.ContractReader {
	ctx := it.Helper.Context(t)
	if it.cr != nil {
		return it.cr
	}

	svc, err := chainreader.NewChainReaderService(it.Helper.Logger(t), it.Helper.RPCClient(), it.contractReaderConfig)

	require.NoError(t, err)
	require.NoError(t, svc.Start(ctx))

	it.cr = svc

	return svc
}

func (it *SolanaChainComponentsInterfaceTester[T]) GetContractWriter(t T) types.ContractWriter {
	return nil
}

func (it *SolanaChainComponentsInterfaceTester[T]) GetBindings(t T) []types.BoundContract {
	// Create a new account with fresh state for each test
	return []types.BoundContract{
		{Name: AnyContractName, Address: it.Helper.CreateAccount(t, AnyValueToReadWithoutAnArgument).String()},
		{Name: AnySecondContractName, Address: it.Helper.CreateAccount(t, AnyDifferentValueToReadWithoutAnArgument).String()},
	}
}

func (it *SolanaChainComponentsInterfaceTester[T]) DirtyContracts() {}

func (it *SolanaChainComponentsInterfaceTester[T]) MaxWaitTimeForEvents() time.Duration {
	return time.Second
}

func (it *SolanaChainComponentsInterfaceTester[T]) GenerateBlocksTillConfidenceLevel(t T, contractName, readName string, confidenceLevel primitives.ConfidenceLevel) {

}

type helper struct {
	programID solana.PublicKey
	rpcURL    string
	wsURL     string
	rpcClient *rpc.Client
	wsClient  *ws.Client
	idlBts    []byte
	nonce     uint64
	nonceMu   sync.Mutex
}

func (h *helper) Init(t *testing.T) {
	t.Helper()

	privateKey, err := solana.PrivateKeyFromBase58(solclient.DefaultPrivateKeysSolValidator[1])
	require.NoError(t, err)

	h.rpcURL, h.wsURL = solanautils.SetupTestValidatorWithAnchorPrograms(t, privateKey.PublicKey().String(), []string{"contract-reader-interface"})
	h.wsClient, err = ws.Connect(tests.Context(t), h.wsURL)
	h.rpcClient = rpc.New(h.rpcURL)

	require.NoError(t, err)

	solanautils.FundAccounts(t, []solana.PrivateKey{privateKey}, h.rpcClient)

	pubkey, err := solana.PublicKeyFromBase58(programPubKey)
	require.NoError(t, err)

	contract.SetProgramID(pubkey)
	h.programID = pubkey
}

func (h *helper) RPCClient() *chainreader.RPCClientWrapper {
	return &chainreader.RPCClientWrapper{Client: h.rpcClient}
}

func (h *helper) Context(t *testing.T) context.Context {
	return tests.Context(t)
}

func (h *helper) Logger(t *testing.T) logger.Logger {
	return logger.Test(t)
}

func (h *helper) GetJSONEncodedIDL(t *testing.T) []byte {
	t.Helper()

	if h.idlBts != nil {
		return h.idlBts
	}

	soPath := filepath.Join(utils.IDLDir, "contract_reader_interface.json")

	_, err := os.Stat(soPath)
	if err != nil {
		t.Log(err.Error())
		t.FailNow()
	}

	bts, err := os.ReadFile(soPath)
	require.NoError(t, err)

	h.idlBts = bts

	return h.idlBts
}

func (h *helper) CreateAccount(t *testing.T, value uint64) solana.PublicKey {
	t.Helper()

	// avoid collisions in parallel tests
	h.nonceMu.Lock()
	h.nonce++
	nonce := h.nonce
	h.nonceMu.Unlock()

	bts := make([]byte, 8)
	binary.LittleEndian.PutUint64(bts, nonce*value)

	pubKey, _, err := solana.FindProgramAddress([][]byte{[]byte("data"), bts}, h.programID)
	require.NoError(t, err)

	// Getting the default localnet private key
	privateKey, err := solana.PrivateKeyFromBase58(solclient.DefaultPrivateKeysSolValidator[1])
	require.NoError(t, err)

	h.runInitialize(t, nonce, value, pubKey, func(key solana.PublicKey) *solana.PrivateKey {
		return &privateKey
	}, privateKey.PublicKey())

	return pubKey
}

func (h *helper) runInitialize(
	t *testing.T,
	nonce uint64,
	value uint64,
	data solana.PublicKey,
	signerFunc func(key solana.PublicKey) *solana.PrivateKey,
	payer solana.PublicKey,
) {
	t.Helper()

	inst, err := contract.NewInitializeInstruction(nonce*value, value, data, payer, solana.SystemProgramID).ValidateAndBuild()
	require.NoError(t, err)

	h.sendInstruction(t, inst, signerFunc, payer)
}

func (h *helper) sendInstruction(
	t *testing.T,
	inst *contract.Instruction,
	signerFunc func(key solana.PublicKey) *solana.PrivateKey,
	payer solana.PublicKey,
) {
	t.Helper()

	ctx := tests.Context(t)

	recent, err := h.rpcClient.GetLatestBlockhash(ctx, rpc.CommitmentFinalized)
	require.NoError(t, err)

	tx, err := solana.NewTransaction(
		[]solana.Instruction{
			inst,
		},
		recent.Value.Blockhash,
		solana.TransactionPayer(payer),
	)
	require.NoError(t, err)

	_, err = tx.EncodeTree(text.NewTreeEncoder(io.Discard, "Initialize"))
	require.NoError(t, err)

	_, err = tx.Sign(signerFunc)
	require.NoError(t, err)

	sig, err := h.rpcClient.SendTransactionWithOpts(
		ctx, tx,
		rpc.TransactionOpts{
			PreflightCommitment: rpc.CommitmentConfirmed,
		},
	)
	require.NoError(t, err)

	h.waitForTX(t, sig, rpc.CommitmentFinalized)
}

func (h *helper) waitForTX(t *testing.T, sig solana.Signature, commitment rpc.CommitmentType) {
	t.Helper()

	sub, err := h.wsClient.SignatureSubscribe(
		sig,
		commitment,
	)
	require.NoError(t, err)

	defer sub.Unsubscribe()

	res, err := sub.Recv()
	require.NoError(t, err)

	if res.Value.Err != nil {
		t.Logf("transaction confirmation failed: %v", res.Value.Err)
		t.FailNow()
	}
}

func mustUnmarshalIDL[T TestingT[T]](t T, rawIDL string) codec.IDL {
	var idl codec.IDL
	if err := json.Unmarshal([]byte(rawIDL), &idl); err != nil {
		t.Errorf("failed to unmarshal test IDL", err)
		t.FailNow()
	}

	return idl
}

const programPubKey = "6AfuXF6HapDUhQfE4nQG9C1SGtA1YjP3icaJyRfU4RyE"
