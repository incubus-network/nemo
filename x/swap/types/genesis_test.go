package types_test

import (
	"encoding/json"
	"testing"

	"github.com/incubus-network/nemo/x/swap/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"sigs.k8s.io/yaml"
)

func TestGenesis_Default(t *testing.T) {
	defaultGenesis := types.DefaultGenesisState()

	require.NoError(t, defaultGenesis.Validate())

	defaultParams := types.DefaultParams()
	assert.Equal(t, defaultParams, defaultGenesis.Params)
}

func TestGenesis_Validate_SwapFee(t *testing.T) {
	type args struct {
		name      string
		swapFee   sdk.Dec
		expectErr bool
	}
	// More comprehensive swap fee tests are in prams_test.go
	testCases := []args{
		{
			"normal",
			sdk.MustNewDecFromStr("0.25"),
			false,
		},
		{
			"negative",
			sdk.MustNewDecFromStr("-0.5"),
			true,
		},
		{
			"greater than 1.0",
			sdk.MustNewDecFromStr("1.001"),
			true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			genesisState := types.GenesisState{
				Params: types.Params{
					AllowedPools: types.DefaultAllowedPools,
					SwapFee:      tc.swapFee,
				},
			}

			err := genesisState.Validate()
			if tc.expectErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestGenesis_Validate_AllowedPools(t *testing.T) {
	type args struct {
		name      string
		pairs     types.AllowedPools
		expectErr bool
	}
	// More comprehensive pair validation tests are in pair_test.go, params_test.go
	testCases := []args{
		{
			"normal",
			types.DefaultAllowedPools,
			false,
		},
		{
			"invalid",
			types.AllowedPools{
				{
					TokenA: "same",
					TokenB: "same",
				},
			},
			true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			genesisState := types.GenesisState{
				Params: types.Params{
					AllowedPools: tc.pairs,
					SwapFee:      types.DefaultSwapFee,
				},
			}

			err := genesisState.Validate()
			if tc.expectErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestGenesis_JSONEncoding(t *testing.T) {
	raw := `{
    "params": {
			"allowed_pools": [
			  {
			    "token_a": "ufury",
					"token_b": "musd"
				},
			  {
			    "token_a": "jinx",
					"token_b": "busd"
				}
			],
			"swap_fee": "0.003000000000000000"
		},
		"pool_records": [
		  {
				"pool_id": "ufury:musd",
			  "reserves_a": { "denom": "ufury", "amount": "1000000" },
			  "reserves_b": { "denom": "musd", "amount": "5000000" },
			  "total_shares": "3000000"
			},
		  {
			  "pool_id": "jinx:musd",
			  "reserves_a": { "denom": "ufury", "amount": "1000000" },
			  "reserves_b": { "denom": "musd", "amount": "2000000" },
			  "total_shares": "2000000"
			}
		],
		"share_records": [
		  {
		    "depositor": "fury1mq9qxlhze029lm0frzw2xr6hem8c3k9tu2gu2x",
		    "pool_id": "ufury:musd",
		    "shares_owned": "100000"
			},
		  {
		    "depositor": "fury1esagqd83rhqdtpy5sxhklaxgn58k2m3sa9wnu4",
		    "pool_id": "jinx:musd",
		    "shares_owned": "200000"
			}
		]
	}`

	var state types.GenesisState
	err := json.Unmarshal([]byte(raw), &state)
	require.NoError(t, err)

	assert.Equal(t, 2, len(state.Params.AllowedPools))
	assert.Equal(t, sdk.MustNewDecFromStr("0.003"), state.Params.SwapFee)
	assert.Equal(t, 2, len(state.PoolRecords))
	assert.Equal(t, 2, len(state.ShareRecords))
}

func TestGenesis_YAMLEncoding(t *testing.T) {
	expected := `params:
  allowed_pools:
  - token_a: ufury
    token_b: musd
  - token_a: jinx
    token_b: busd
  swap_fee: "0.003000000000000000"
pool_records:
- pool_id: ufury:musd
  reserves_a:
    amount: "1000000"
    denom: ufury
  reserves_b:
    amount: "5000000"
    denom: musd
  total_shares: "3000000"
- pool_id: jinx:musd
  reserves_a:
    amount: "1000000"
    denom: jinx
  reserves_b:
    amount: "2000000"
    denom: musd
  total_shares: "1500000"
share_records:
- depositor: fury1mq9qxlhze029lm0frzw2xr6hem8c3k9tu2gu2x
  pool_id: ufury:musd
  shares_owned: "100000"
- depositor: fury1esagqd83rhqdtpy5sxhklaxgn58k2m3sa9wnu4
  pool_id: jinx:musd
  shares_owned: "200000"
`

	depositor_1, err := sdk.AccAddressFromBech32("fury1mq9qxlhze029lm0frzw2xr6hem8c3k9tu2gu2x")
	require.NoError(t, err)
	depositor_2, err := sdk.AccAddressFromBech32("fury1esagqd83rhqdtpy5sxhklaxgn58k2m3sa9wnu4")
	require.NoError(t, err)

	state := types.NewGenesisState(
		types.NewParams(
			types.NewAllowedPools(
				types.NewAllowedPool("ufury", "musd"),
				types.NewAllowedPool("jinx", "busd"),
			),
			sdk.MustNewDecFromStr("0.003"),
		),
		types.PoolRecords{
			types.NewPoolRecord(sdk.NewCoins(ufury(1e6), musd(5e6)), i(3e6)),
			types.NewPoolRecord(sdk.NewCoins(jinx(1e6), musd(2e6)), i(15e5)),
		},
		types.ShareRecords{
			types.NewShareRecord(depositor_1, types.PoolID("ufury", "musd"), i(1e5)),
			types.NewShareRecord(depositor_2, types.PoolID("jinx", "musd"), i(2e5)),
		},
	)

	data, err := yaml.Marshal(state)
	require.NoError(t, err)

	assert.Equal(t, expected, string(data))
}

func TestGenesis_ValidatePoolRecords(t *testing.T) {
	invalidPoolRecord := types.NewPoolRecord(sdk.NewCoins(ufury(1e6), musd(5e6)), i(-1))

	state := types.NewGenesisState(
		types.DefaultParams(),
		types.PoolRecords{invalidPoolRecord},
		types.ShareRecords{},
	)

	assert.Error(t, state.Validate())
}

func TestGenesis_ValidateShareRecords(t *testing.T) {
	depositor, err := sdk.AccAddressFromBech32("fury1mq9qxlhze029lm0frzw2xr6hem8c3k9tu2gu2x")
	require.NoError(t, err)

	invalidShareRecord := types.NewShareRecord(depositor, "", i(-1))

	state := types.NewGenesisState(
		types.DefaultParams(),
		types.PoolRecords{},
		types.ShareRecords{invalidShareRecord},
	)

	assert.Error(t, state.Validate())
}

func TestGenesis_Validate_PoolShareIntegration(t *testing.T) {
	depositor_1, err := sdk.AccAddressFromBech32("fury1mq9qxlhze029lm0frzw2xr6hem8c3k9tu2gu2x")
	require.NoError(t, err)
	depositor_2, err := sdk.AccAddressFromBech32("fury1esagqd83rhqdtpy5sxhklaxgn58k2m3sa9wnu4")
	require.NoError(t, err)

	testCases := []struct {
		name         string
		poolRecords  types.PoolRecords
		shareRecords types.ShareRecords
		expectedErr  string
	}{
		{
			name: "single pool record, zero share records",
			poolRecords: types.PoolRecords{
				types.NewPoolRecord(sdk.NewCoins(ufury(1e6), musd(5e6)), i(3e6)),
			},
			shareRecords: types.ShareRecords{},
			expectedErr:  "total depositor shares 0 not equal to pool 'ufury:musd' total shares 3000000",
		},
		{
			name:        "zero pool records, one share record",
			poolRecords: types.PoolRecords{},
			shareRecords: types.ShareRecords{
				types.NewShareRecord(depositor_1, types.PoolID("ufury", "musd"), i(5e6)),
			},
			expectedErr: "total depositor shares 5000000 not equal to pool 'ufury:musd' total shares 0",
		},
		{
			name: "one pool record, one share record",
			poolRecords: types.PoolRecords{
				types.NewPoolRecord(sdk.NewCoins(ufury(1e6), musd(5e6)), i(3e6)),
			},
			shareRecords: types.ShareRecords{
				types.NewShareRecord(depositor_1, "ufury:musd", i(15e5)),
			},
			expectedErr: "total depositor shares 1500000 not equal to pool 'ufury:musd' total shares 3000000",
		},
		{
			name: "more than one pool records, more than one share record",
			poolRecords: types.PoolRecords{
				types.NewPoolRecord(sdk.NewCoins(ufury(1e6), musd(5e6)), i(3e6)),
				types.NewPoolRecord(sdk.NewCoins(jinx(1e6), musd(2e6)), i(2e6)),
			},
			shareRecords: types.ShareRecords{
				types.NewShareRecord(depositor_1, types.PoolID("ufury", "musd"), i(15e5)),
				types.NewShareRecord(depositor_2, types.PoolID("ufury", "musd"), i(15e5)),
				types.NewShareRecord(depositor_1, types.PoolID("jinx", "musd"), i(1e6)),
			},
			expectedErr: "total depositor shares 1000000 not equal to pool 'jinx:musd' total shares 2000000",
		},
		{
			name: "valid case with many pool records and share records",
			poolRecords: types.PoolRecords{
				types.NewPoolRecord(sdk.NewCoins(ufury(1e6), musd(5e6)), i(3e6)),
				types.NewPoolRecord(sdk.NewCoins(jinx(1e6), musd(2e6)), i(2e6)),
				types.NewPoolRecord(sdk.NewCoins(jinx(7e6), ufury(10e6)), i(8e6)),
			},
			shareRecords: types.ShareRecords{
				types.NewShareRecord(depositor_1, types.PoolID("ufury", "musd"), i(15e5)),
				types.NewShareRecord(depositor_2, types.PoolID("ufury", "musd"), i(15e5)),
				types.NewShareRecord(depositor_1, types.PoolID("jinx", "musd"), i(2e6)),
				types.NewShareRecord(depositor_1, types.PoolID("jinx", "ufury"), i(3e6)),
				types.NewShareRecord(depositor_2, types.PoolID("jinx", "ufury"), i(5e6)),
			},
			expectedErr: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			state := types.NewGenesisState(types.DefaultParams(), tc.poolRecords, tc.shareRecords)
			err := state.Validate()

			if tc.expectedErr == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedErr)
			}
		})
	}
}
