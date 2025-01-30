package chainwriter

import (
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"

	"github.com/smartcontractkit/chainlink-solana/pkg/solana/client"
	"github.com/smartcontractkit/chainlink-solana/pkg/solana/codec"
)

// Lookup is an interface that defines a method to resolve an address (or multiple addresses) from a given definition.
type Lookup interface {
	Resolve(ctx context.Context, args any, derivedTableMap map[string]map[string][]*solana.AccountMeta, reader client.Reader, idl string) ([]*solana.AccountMeta, error)
}

// AccountConstant represents a fixed address, provided in Base58 format, converted into a `solana.PublicKey`.
type AccountConstant struct {
	Name       string
	Address    string
	IsSigner   bool
	IsWritable bool
}

// AccountLookup dynamically derives an account address from args using a specified location path.
type AccountLookup struct {
	Name     string
	Location string
	// IsSigner and IsWritable can either be a constant bool or a location to a bitmap which decides the bools
	IsSigner   MetaBool
	IsWritable MetaBool
}

type MetaBool struct {
	Value          bool
	BitmapLocation string
}

type Seed struct {
	Static  []byte // Static seed value
	Dynamic Lookup // Dynamic lookup for seed
}

// PDALookups generates Program Derived Addresses (PDA) by combining a derived public key with one or more seeds.
type PDALookups struct {
	Name string
	// The public key of the PDA to be combined with seeds. If there are multiple PublicKeys
	// there will be multiple PDAs generated by combining each PublicKey with the seeds.
	PublicKey Lookup
	// Seeds to be derived from an additional lookup
	Seeds      []Seed
	IsSigner   bool
	IsWritable bool
	// OPTIONAL: On-chain location and type of desired data from PDA (e.g. a sub-account of the data account)
	InternalField InternalField
}

type InternalField struct {
	// must map directly to IDL type
	TypeName string
	Location string
}

// LookupTables represents a list of lookup tables that are used to derive addresses for a program.
type LookupTables struct {
	DerivedLookupTables []DerivedLookupTable
	StaticLookupTables  []solana.PublicKey
}

// DerivedLookupTable represents a lookup table that is used to derive addresses for a program.
type DerivedLookupTable struct {
	Name     string
	Accounts Lookup
}

// AccountsFromLookupTable extracts accounts from a lookup table that was previously read and stored in memory.
type AccountsFromLookupTable struct {
	LookupTableName string
	IncludeIndexes  []int
}

func (ac AccountConstant) Resolve(_ context.Context, _ any, _ map[string]map[string][]*solana.AccountMeta, _ client.Reader, _ string) ([]*solana.AccountMeta, error) {
	address, err := solana.PublicKeyFromBase58(ac.Address)
	if err != nil {
		return nil, fmt.Errorf("error getting account from constant: %w", err)
	}
	return []*solana.AccountMeta{
		{
			PublicKey:  address,
			IsSigner:   ac.IsSigner,
			IsWritable: ac.IsWritable,
		},
	}, nil
}

func (al AccountLookup) Resolve(_ context.Context, args any, _ map[string]map[string][]*solana.AccountMeta, _ client.Reader, _ string) ([]*solana.AccountMeta, error) {
	derivedValues, err := GetValuesAtLocation(args, al.Location)
	if err != nil {
		return nil, fmt.Errorf("error getting account from lookup: %w", err)
	}

	var metas []*solana.AccountMeta
	signerIndexes, err := resolveBitMap(al.IsSigner, args, len(derivedValues))
	if err != nil {
		return nil, err
	}

	writerIndexes, err := resolveBitMap(al.IsWritable, args, len(derivedValues))
	if err != nil {
		return nil, err
	}

	for i, address := range derivedValues {
		// Resolve isSigner for this particular pubkey
		isSigner := signerIndexes[i]

		// Resolve isWritable
		isWritable := writerIndexes[i]

		metas = append(metas, &solana.AccountMeta{
			PublicKey:  solana.PublicKeyFromBytes(address),
			IsSigner:   isSigner,
			IsWritable: isWritable,
		})
	}
	return metas, nil
}

func resolveBitMap(mb MetaBool, args any, length int) ([]bool, error) {
	result := make([]bool, length)
	if mb.BitmapLocation == "" {
		for i := 0; i < length; i++ {
			result[i] = mb.Value
		}
		return result, nil
	}

	bitmapVals, err := GetValuesAtLocation(args, mb.BitmapLocation)
	if err != nil {
		return []bool{}, fmt.Errorf("error reading bitmap from location '%s': %w", mb.BitmapLocation, err)
	}

	if len(bitmapVals) != 1 {
		return []bool{}, fmt.Errorf("bitmap value is not a single value: %v, length: %d", bitmapVals, len(bitmapVals))
	}

	bitmapInt := binary.LittleEndian.Uint64(bitmapVals[0])
	for i := 0; i < length; i++ {
		result[i] = bitmapInt&(1<<i) > 0
	}

	return result, nil
}

func (alt AccountsFromLookupTable) Resolve(_ context.Context, _ any, derivedTableMap map[string]map[string][]*solana.AccountMeta, _ client.Reader, _ string) ([]*solana.AccountMeta, error) {
	// Fetch the inner map for the specified lookup table name
	innerMap, ok := derivedTableMap[alt.LookupTableName]
	if !ok {
		return nil, fmt.Errorf("lookup table not found: %s", alt.LookupTableName)
	}

	var result []*solana.AccountMeta

	// If no indices are specified, include all addresses
	if len(alt.IncludeIndexes) == 0 {
		for _, metas := range innerMap {
			result = append(result, metas...)
		}
		return result, nil
	}

	// Otherwise, include only addresses at the specified indices
	for publicKey, metas := range innerMap {
		for _, index := range alt.IncludeIndexes {
			if index < 0 || index >= len(metas) {
				return nil, fmt.Errorf("invalid index %d for account %s in lookup table %s", index, publicKey, alt.LookupTableName)
			}
			result = append(result, metas[index])
		}
	}

	return result, nil
}

func (pda PDALookups) Resolve(ctx context.Context, args any, derivedTableMap map[string]map[string][]*solana.AccountMeta, reader client.Reader, idl string) ([]*solana.AccountMeta, error) {
	publicKeys, err := GetAddresses(ctx, args, []Lookup{pda.PublicKey}, derivedTableMap, reader, idl)
	if err != nil {
		return nil, fmt.Errorf("error getting public key for PDALookups: %w", err)
	}

	seeds, err := getSeedBytesCombinations(ctx, pda, args, derivedTableMap, reader)
	if err != nil {
		return nil, fmt.Errorf("error getting seeds for PDALookups: %w", err)
	}

	pdas, err := generatePDAs(publicKeys, seeds, pda)
	if err != nil {
		return nil, fmt.Errorf("error generating PDAs: %w", err)
	}

	if pda.InternalField.Location == "" {
		return pdas, nil
	}

	// If a decoded location is specified, fetch the data at that location
	var result []*solana.AccountMeta
	for _, accountMeta := range pdas {
		accountInfo, err := reader.GetAccountInfoWithOpts(ctx, accountMeta.PublicKey, &rpc.GetAccountInfoOpts{
			Encoding:   "base64",
			Commitment: rpc.CommitmentFinalized,
		})

		if err != nil || accountInfo == nil || accountInfo.Value == nil {
			return nil, fmt.Errorf("error fetching account info for PDA account: %s, error: %w", accountMeta.PublicKey.String(), err)
		}

		var idlCodec codec.IDL
		if err = json.Unmarshal([]byte(idl), &idlCodec); err != nil {
			return nil, fmt.Errorf("failed to unmarshal IDL for PDA: %s, error: %w", pda.Name, err)
		}

		internalType := pda.InternalField.TypeName

		idlDef, err := codec.FindDefinitionFromIDL(codec.ChainConfigTypeAccountDef, internalType, idlCodec)
		if err != nil {
			return nil, fmt.Errorf("error finding definition for type %s: %w", internalType, err)
		}

		input, err := codec.CreateCodecEntry(idlDef, internalType, idlCodec, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create codec entry for method %s, error: %w", internalType, err)
		}

		decoded, _, err := input.Decode(accountInfo.Value.Data.GetBinary())
		if err != nil {
			return nil, fmt.Errorf("error decoding account data: %w", err)
		}

		value, err := GetValuesAtLocation(decoded, pda.InternalField.Location)
		if err != nil {
			return nil, fmt.Errorf("error getting value at location: %w", err)
		}
		if len(value) > 1 {
			return nil, fmt.Errorf("multiple values found at location: %s", pda.InternalField.Location)
		}

		result = append(result, &solana.AccountMeta{
			PublicKey:  solana.PublicKeyFromBytes(value[0]),
			IsSigner:   accountMeta.IsSigner,
			IsWritable: accountMeta.IsWritable,
		})
	}
	return result, nil
}

// getSeedBytesCombinations extracts the seeds for the PDALookups.
// The return type is [][][]byte, where each element of the outer slice is
// one combination of seeds. This handles the case where one seed can resolve
// to multiple addresses, multiplying the combinations accordingly.
func getSeedBytesCombinations(
	ctx context.Context,
	lookup PDALookups,
	args any,
	derivedTableMap map[string]map[string][]*solana.AccountMeta,
	reader client.Reader,
) ([][][]byte, error) {
	allCombinations := [][][]byte{
		{},
	}

	// For each seed in the definition, expand the current list of combinations
	// by all possible values for this seed.
	for _, seed := range lookup.Seeds {
		expansions := make([][]byte, 0)
		if seed.Static != nil {
			expansions = append(expansions, seed.Static)
			// Static and Dynamic seeds are mutually exclusive
		} else if seed.Dynamic != nil {
			dynamicSeed := seed.Dynamic
			if lookupSeed, ok := dynamicSeed.(AccountLookup); ok {
				// Get value from a location (This doens't have to be an address, it can be any value)
				bytes, err := GetValuesAtLocation(args, lookupSeed.Location)
				if err != nil {
					return nil, fmt.Errorf("error getting address seed for location %q: %w", lookupSeed.Location, err)
				}
				// append each byte array to the expansions
				for _, b := range bytes {
					// validate seed length
					if len(b) > solana.MaxSeedLength {
						return nil, fmt.Errorf("seed byte array exceeds maximum length of %d: got %d bytes", solana.MaxSeedLength, len(b))
					}
					expansions = append(expansions, b)
				}
			} else {
				// Get address seeds from the lookup
				seedAddresses, err := GetAddresses(ctx, args, []Lookup{dynamicSeed}, derivedTableMap, reader, "")
				if err != nil {
					return nil, fmt.Errorf("error getting address seed: %w", err)
				}
				// Add each address seed to the expansions
				for _, addrMeta := range seedAddresses {
					b := addrMeta.PublicKey.Bytes()
					if len(b) > solana.MaxSeedLength {
						return nil, fmt.Errorf("seed byte array exceeds maximum length of %d: got %d bytes", solana.MaxSeedLength, len(b))
					}
					expansions = append(expansions, b)
				}
			}
		}

		// expansions is the list of possible seed bytes for this single seed lookup.
		// Multiply the existing combinations in allCombinations by each item in expansions.
		newCombinations := make([][][]byte, 0, len(allCombinations)*len(expansions))
		for _, existingCombo := range allCombinations {
			for _, expandedSeed := range expansions {
				comboCopy := make([][]byte, len(existingCombo)+1)
				copy(comboCopy, existingCombo)
				comboCopy[len(existingCombo)] = expandedSeed
				newCombinations = append(newCombinations, comboCopy)
			}
		}

		allCombinations = newCombinations
	}

	return allCombinations, nil
}

// generatePDAs generates program-derived addresses (PDAs) from public keys and seeds.
// it will result in a list of PDAs whose length is the product of the number of public keys
// and the number of seed combinations.
func generatePDAs(
	publicKeys []*solana.AccountMeta,
	seedCombos [][][]byte,
	lookup PDALookups,
) ([]*solana.AccountMeta, error) {
	var results []*solana.AccountMeta
	for _, publicKeyMeta := range publicKeys {
		for _, combo := range seedCombos {
			if len(combo) > solana.MaxSeeds {
				return nil, fmt.Errorf("seed maximum exceeded: %d", len(combo))
			}
			address, _, err := solana.FindProgramAddress(combo, publicKeyMeta.PublicKey)
			if err != nil {
				return nil, fmt.Errorf("error finding program address: %w", err)
			}
			results = append(results, &solana.AccountMeta{
				PublicKey:  address,
				IsSigner:   lookup.IsSigner,
				IsWritable: lookup.IsWritable,
			})
		}
	}

	return results, nil
}
