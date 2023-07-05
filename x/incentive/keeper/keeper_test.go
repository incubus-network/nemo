package keeper_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"
	tmprototypes "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/incubus-network/nemo/app"
	"github.com/incubus-network/nemo/x/incentive/keeper"
	"github.com/incubus-network/nemo/x/incentive/types"
)

// Test suite used for all keeper tests
type KeeperTestSuite struct {
	suite.Suite

	keeper keeper.Keeper

	app app.TestApp
	ctx sdk.Context

	genesisTime time.Time
	addrs       []sdk.AccAddress
}

// SetupTest is run automatically before each suite test
func (suite *KeeperTestSuite) SetupTest() {
	config := sdk.GetConfig()
	app.SetBech32AddressPrefixes(config)

	_, suite.addrs = app.GeneratePrivKeyAddressPairs(5)

	suite.genesisTime = time.Date(2020, 12, 15, 14, 0, 0, 0, time.UTC)
}

func (suite *KeeperTestSuite) SetupApp() {
	suite.app = app.NewTestApp()

	suite.keeper = suite.app.GetIncentiveKeeper()

	suite.ctx = suite.app.NewContext(true, tmprototypes.Header{Time: suite.genesisTime})
}

func (suite *KeeperTestSuite) TestGetSetDeleteMUSDMintingClaim() {
	suite.SetupApp()
	c := types.NewMUSDMintingClaim(suite.addrs[0], c("ufury", 1000000), types.RewardIndexes{types.NewRewardIndex("bnb-a", sdk.ZeroDec())})
	_, found := suite.keeper.GetMUSDMintingClaim(suite.ctx, suite.addrs[0])
	suite.Require().False(found)
	suite.Require().NotPanics(func() {
		suite.keeper.SetMUSDMintingClaim(suite.ctx, c)
	})
	testC, found := suite.keeper.GetMUSDMintingClaim(suite.ctx, suite.addrs[0])
	suite.Require().True(found)
	suite.Require().Equal(c, testC)
	suite.Require().NotPanics(func() {
		suite.keeper.DeleteMUSDMintingClaim(suite.ctx, suite.addrs[0])
	})
	_, found = suite.keeper.GetMUSDMintingClaim(suite.ctx, suite.addrs[0])
	suite.Require().False(found)
}

func (suite *KeeperTestSuite) TestIterateMUSDMintingClaims() {
	suite.SetupApp()
	for i := 0; i < len(suite.addrs); i++ {
		c := types.NewMUSDMintingClaim(suite.addrs[i], c("ufury", 100000), types.RewardIndexes{types.NewRewardIndex("bnb-a", sdk.ZeroDec())})
		suite.Require().NotPanics(func() {
			suite.keeper.SetMUSDMintingClaim(suite.ctx, c)
		})
	}
	claims := types.MUSDMintingClaims{}
	suite.keeper.IterateMUSDMintingClaims(suite.ctx, func(c types.MUSDMintingClaim) bool {
		claims = append(claims, c)
		return false
	})
	suite.Require().Equal(len(suite.addrs), len(claims))

	claims = suite.keeper.GetAllMUSDMintingClaims(suite.ctx)
	suite.Require().Equal(len(suite.addrs), len(claims))
}

func (suite *KeeperTestSuite) TestGetSetDeleteSwapClaims() {
	suite.SetupApp()
	c := types.NewSwapClaim(suite.addrs[0], arbitraryCoins(), nonEmptyMultiRewardIndexes)

	_, found := suite.keeper.GetSwapClaim(suite.ctx, suite.addrs[0])
	suite.Require().False(found)

	suite.Require().NotPanics(func() {
		suite.keeper.SetSwapClaim(suite.ctx, c)
	})
	testC, found := suite.keeper.GetSwapClaim(suite.ctx, suite.addrs[0])
	suite.Require().True(found)
	suite.Require().Equal(c, testC)

	suite.Require().NotPanics(func() {
		suite.keeper.DeleteSwapClaim(suite.ctx, suite.addrs[0])
	})
	_, found = suite.keeper.GetSwapClaim(suite.ctx, suite.addrs[0])
	suite.Require().False(found)
}

func (suite *KeeperTestSuite) TestIterateSwapClaims() {
	suite.SetupApp()
	claims := types.SwapClaims{
		types.NewSwapClaim(suite.addrs[0], arbitraryCoins(), nonEmptyMultiRewardIndexes),
		types.NewSwapClaim(suite.addrs[1], nil, nil), // different claim to the first
	}
	for _, claim := range claims {
		suite.keeper.SetSwapClaim(suite.ctx, claim)
	}

	var actualClaims types.SwapClaims
	suite.keeper.IterateSwapClaims(suite.ctx, func(c types.SwapClaim) bool {
		actualClaims = append(actualClaims, c)
		return false
	})

	suite.Require().Equal(claims, actualClaims)
}

func (suite *KeeperTestSuite) TestGetSetSwapRewardIndexes() {
	testCases := []struct {
		name      string
		poolName  string
		indexes   types.RewardIndexes
		wantIndex types.RewardIndexes
		panics    bool
	}{
		{
			name:     "two factors can be written and read",
			poolName: "btc/musd",
			indexes: types.RewardIndexes{
				{
					CollateralType: "jinx",
					RewardFactor:   d("0.02"),
				},
				{
					CollateralType: "ufury",
					RewardFactor:   d("0.04"),
				},
			},
			wantIndex: types.RewardIndexes{
				{
					CollateralType: "jinx",
					RewardFactor:   d("0.02"),
				},
				{
					CollateralType: "ufury",
					RewardFactor:   d("0.04"),
				},
			},
		},
		{
			name:     "indexes with empty pool name panics",
			poolName: "",
			indexes: types.RewardIndexes{
				{
					CollateralType: "jinx",
					RewardFactor:   d("0.02"),
				},
				{
					CollateralType: "ufury",
					RewardFactor:   d("0.04"),
				},
			},
			panics: true,
		},
		{
			// this test is to detect any changes in behavior
			name:     "setting empty indexes does not panic",
			poolName: "btc/musd",
			// Marshalling empty slice results in [] bytes, unmarshalling the []
			// empty bytes results in a nil slice instead of an empty slice
			indexes:   types.RewardIndexes{},
			wantIndex: nil,
			panics:    false,
		},
		{
			// this test is to detect any changes in behavior
			name:      "setting nil indexes does not panic",
			poolName:  "btc/musd",
			indexes:   nil,
			wantIndex: nil,
			panics:    false,
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			suite.SetupApp()

			_, found := suite.keeper.GetSwapRewardIndexes(suite.ctx, tc.poolName)
			suite.False(found)

			setFunc := func() { suite.keeper.SetSwapRewardIndexes(suite.ctx, tc.poolName, tc.indexes) }
			if tc.panics {
				suite.Panics(setFunc)
				return
			} else {
				suite.NotPanics(setFunc)
			}

			storedIndexes, found := suite.keeper.GetSwapRewardIndexes(suite.ctx, tc.poolName)
			suite.True(found)
			suite.Equal(tc.wantIndex, storedIndexes)
		})
	}
}

func (suite *KeeperTestSuite) TestIterateSwapRewardIndexes() {
	suite.SetupApp()
	multiIndexes := types.MultiRewardIndexes{
		{
			CollateralType: "bnb/musd",
			RewardIndexes: types.RewardIndexes{
				{
					CollateralType: "swap",
					RewardFactor:   d("0.0000002"),
				},
				{
					CollateralType: "ufury",
					RewardFactor:   d("0.04"),
				},
			},
		},
		{
			CollateralType: "btcb/musd",
			RewardIndexes: types.RewardIndexes{
				{
					CollateralType: "jinx",
					RewardFactor:   d("0.02"),
				},
			},
		},
	}
	for _, mi := range multiIndexes {
		suite.keeper.SetSwapRewardIndexes(suite.ctx, mi.CollateralType, mi.RewardIndexes)
	}

	var actualMultiIndexes types.MultiRewardIndexes
	suite.keeper.IterateSwapRewardIndexes(suite.ctx, func(poolID string, i types.RewardIndexes) bool {
		actualMultiIndexes = actualMultiIndexes.With(poolID, i)
		return false
	})

	suite.Require().Equal(multiIndexes, actualMultiIndexes)
}

func (suite *KeeperTestSuite) TestGetSetSwapRewardAccrualTimes() {
	testCases := []struct {
		name        string
		poolName    string
		accrualTime time.Time
		panics      bool
	}{
		{
			name:        "normal time can be written and read",
			poolName:    "btc/musd",
			accrualTime: time.Date(1998, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:        "zero time can be written and read",
			poolName:    "btc/musd",
			accrualTime: time.Time{},
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			suite.SetupApp()

			_, found := suite.keeper.GetSwapRewardAccrualTime(suite.ctx, tc.poolName)
			suite.False(found)

			setFunc := func() { suite.keeper.SetSwapRewardAccrualTime(suite.ctx, tc.poolName, tc.accrualTime) }
			if tc.panics {
				suite.Panics(setFunc)
				return
			} else {
				suite.NotPanics(setFunc)
			}

			storedTime, found := suite.keeper.GetSwapRewardAccrualTime(suite.ctx, tc.poolName)
			suite.True(found)
			suite.Equal(tc.accrualTime, storedTime)
		})
	}
}

func (suite *KeeperTestSuite) TestGetSetDeleteEarnClaims() {
	suite.SetupApp()
	c := types.NewEarnClaim(suite.addrs[0], arbitraryCoins(), nonEmptyMultiRewardIndexes)

	_, found := suite.keeper.GetEarnClaim(suite.ctx, suite.addrs[0])
	suite.Require().False(found)

	suite.Require().NotPanics(func() {
		suite.keeper.SetEarnClaim(suite.ctx, c)
	})
	testC, found := suite.keeper.GetEarnClaim(suite.ctx, suite.addrs[0])
	suite.Require().True(found)
	suite.Require().Equal(c, testC)

	suite.Require().NotPanics(func() {
		suite.keeper.DeleteEarnClaim(suite.ctx, suite.addrs[0])
	})
	_, found = suite.keeper.GetEarnClaim(suite.ctx, suite.addrs[0])
	suite.Require().False(found)
}

func (suite *KeeperTestSuite) TestIterateEarnClaims() {
	suite.SetupApp()
	claims := types.EarnClaims{
		types.NewEarnClaim(suite.addrs[0], arbitraryCoins(), nonEmptyMultiRewardIndexes),
		types.NewEarnClaim(suite.addrs[1], nil, nil), // different claim to the first
	}
	for _, claim := range claims {
		suite.keeper.SetEarnClaim(suite.ctx, claim)
	}

	var actualClaims types.EarnClaims
	suite.keeper.IterateEarnClaims(suite.ctx, func(c types.EarnClaim) bool {
		actualClaims = append(actualClaims, c)
		return false
	})

	suite.Require().Equal(claims, actualClaims)
}

func (suite *KeeperTestSuite) TestGetSetEarnRewardIndexes() {
	testCases := []struct {
		name       string
		vaultDenom string
		indexes    types.RewardIndexes
		wantIndex  types.RewardIndexes
		panics     bool
	}{
		{
			name:       "two factors can be written and read",
			vaultDenom: "musd",
			indexes: types.RewardIndexes{
				{
					CollateralType: "jinx",
					RewardFactor:   d("0.02"),
				},
				{
					CollateralType: "ufury",
					RewardFactor:   d("0.04"),
				},
			},
			wantIndex: types.RewardIndexes{
				{
					CollateralType: "jinx",
					RewardFactor:   d("0.02"),
				},
				{
					CollateralType: "ufury",
					RewardFactor:   d("0.04"),
				},
			},
		},
		{
			name:       "indexes with empty vault name panics",
			vaultDenom: "",
			indexes: types.RewardIndexes{
				{
					CollateralType: "jinx",
					RewardFactor:   d("0.02"),
				},
				{
					CollateralType: "ufury",
					RewardFactor:   d("0.04"),
				},
			},
			panics: true,
		},
		{
			// this test is to detect any changes in behavior
			name:       "setting empty indexes does not panic",
			vaultDenom: "musd",
			// Marshalling empty slice results in [] bytes, unmarshalling the []
			// empty bytes results in a nil slice instead of an empty slice
			indexes:   types.RewardIndexes{},
			wantIndex: nil,
			panics:    false,
		},
		{
			// this test is to detect any changes in behavior
			name:       "setting nil indexes does not panic",
			vaultDenom: "musd",
			indexes:    nil,
			wantIndex:  nil,
			panics:     false,
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			suite.SetupApp()

			_, found := suite.keeper.GetEarnRewardIndexes(suite.ctx, tc.vaultDenom)
			suite.False(found)

			setFunc := func() { suite.keeper.SetEarnRewardIndexes(suite.ctx, tc.vaultDenom, tc.indexes) }
			if tc.panics {
				suite.Panics(setFunc)
				return
			} else {
				suite.NotPanics(setFunc)
			}

			storedIndexes, found := suite.keeper.GetEarnRewardIndexes(suite.ctx, tc.vaultDenom)
			suite.True(found)
			suite.Equal(tc.wantIndex, storedIndexes)
		})
	}
}

func (suite *KeeperTestSuite) TestIterateEarnRewardIndexes() {
	suite.SetupApp()
	multiIndexes := types.MultiRewardIndexes{
		{
			CollateralType: "ufury",
			RewardIndexes: types.RewardIndexes{
				{
					CollateralType: "earn",
					RewardFactor:   d("0.0000002"),
				},
				{
					CollateralType: "ufury",
					RewardFactor:   d("0.04"),
				},
			},
		},
		{
			CollateralType: "musd",
			RewardIndexes: types.RewardIndexes{
				{
					CollateralType: "jinx",
					RewardFactor:   d("0.02"),
				},
			},
		},
	}
	for _, mi := range multiIndexes {
		suite.keeper.SetEarnRewardIndexes(suite.ctx, mi.CollateralType, mi.RewardIndexes)
	}

	var actualMultiIndexes types.MultiRewardIndexes
	suite.keeper.IterateEarnRewardIndexes(suite.ctx, func(vaultDenom string, i types.RewardIndexes) bool {
		actualMultiIndexes = actualMultiIndexes.With(vaultDenom, i)
		return false
	})

	suite.Require().Equal(multiIndexes, actualMultiIndexes)
}

func (suite *KeeperTestSuite) TestGetSetEarnRewardAccrualTimes() {
	testCases := []struct {
		name        string
		vaultDenom  string
		accrualTime time.Time
		panics      bool
	}{
		{
			name:        "normal time can be written and read",
			vaultDenom:  "musd",
			accrualTime: time.Date(1998, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:        "zero time can be written and read",
			vaultDenom:  "musd",
			accrualTime: time.Time{},
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			suite.SetupApp()

			_, found := suite.keeper.GetEarnRewardAccrualTime(suite.ctx, tc.vaultDenom)
			suite.False(found)

			setFunc := func() { suite.keeper.SetEarnRewardAccrualTime(suite.ctx, tc.vaultDenom, tc.accrualTime) }
			if tc.panics {
				suite.Panics(setFunc)
				return
			} else {
				suite.NotPanics(setFunc)
			}

			storedTime, found := suite.keeper.GetEarnRewardAccrualTime(suite.ctx, tc.vaultDenom)
			suite.True(found)
			suite.Equal(tc.accrualTime, storedTime)
		})
	}
}

type accrualtime struct {
	denom string
	time  time.Time
}

var nonEmptyAccrualTimes = []accrualtime{
	{
		denom: "btcb",
		time:  time.Date(1998, 1, 1, 0, 0, 0, 1, time.UTC),
	},
	{
		denom: "ufury",
		time:  time.Time{},
	},
}

func (suite *KeeperTestSuite) TestIterateMUSDMintingAccrualTimes() {
	suite.SetupApp()

	expectedAccrualTimes := nonEmptyAccrualTimes

	for _, at := range expectedAccrualTimes {
		suite.keeper.SetPreviousMUSDMintingAccrualTime(suite.ctx, at.denom, at.time)
	}

	var actualAccrualTimes []accrualtime
	suite.keeper.IterateMUSDMintingAccrualTimes(suite.ctx, func(denom string, accrualTime time.Time) bool {
		actualAccrualTimes = append(actualAccrualTimes, accrualtime{denom: denom, time: accrualTime})
		return false
	})

	suite.Equal(expectedAccrualTimes, actualAccrualTimes)
}

func (suite *KeeperTestSuite) TestIterateJinxSupplyRewardAccrualTimes() {
	suite.SetupApp()

	expectedAccrualTimes := nonEmptyAccrualTimes

	for _, at := range expectedAccrualTimes {
		suite.keeper.SetPreviousJinxSupplyRewardAccrualTime(suite.ctx, at.denom, at.time)
	}

	var actualAccrualTimes []accrualtime
	suite.keeper.IterateJinxSupplyRewardAccrualTimes(suite.ctx, func(denom string, accrualTime time.Time) bool {
		actualAccrualTimes = append(actualAccrualTimes, accrualtime{denom: denom, time: accrualTime})
		return false
	})

	suite.Equal(expectedAccrualTimes, actualAccrualTimes)
}

func (suite *KeeperTestSuite) TestIterateJinxBorrowrRewardAccrualTimes() {
	suite.SetupApp()

	expectedAccrualTimes := nonEmptyAccrualTimes

	for _, at := range expectedAccrualTimes {
		suite.keeper.SetPreviousJinxBorrowRewardAccrualTime(suite.ctx, at.denom, at.time)
	}

	var actualAccrualTimes []accrualtime
	suite.keeper.IterateJinxBorrowRewardAccrualTimes(suite.ctx, func(denom string, accrualTime time.Time) bool {
		actualAccrualTimes = append(actualAccrualTimes, accrualtime{denom: denom, time: accrualTime})
		return false
	})

	suite.Equal(expectedAccrualTimes, actualAccrualTimes)
}

func (suite *KeeperTestSuite) TestIterateDelegatorRewardAccrualTimes() {
	suite.SetupApp()

	expectedAccrualTimes := nonEmptyAccrualTimes

	for _, at := range expectedAccrualTimes {
		suite.keeper.SetPreviousDelegatorRewardAccrualTime(suite.ctx, at.denom, at.time)
	}

	var actualAccrualTimes []accrualtime
	suite.keeper.IterateDelegatorRewardAccrualTimes(suite.ctx, func(denom string, accrualTime time.Time) bool {
		actualAccrualTimes = append(actualAccrualTimes, accrualtime{denom: denom, time: accrualTime})
		return false
	})

	suite.Equal(expectedAccrualTimes, actualAccrualTimes)
}

func (suite *KeeperTestSuite) TestIterateSwapRewardAccrualTimes() {
	suite.SetupApp()

	expectedAccrualTimes := nonEmptyAccrualTimes

	for _, at := range expectedAccrualTimes {
		suite.keeper.SetSwapRewardAccrualTime(suite.ctx, at.denom, at.time)
	}

	var actualAccrualTimes []accrualtime
	suite.keeper.IterateSwapRewardAccrualTimes(suite.ctx, func(denom string, accrualTime time.Time) bool {
		actualAccrualTimes = append(actualAccrualTimes, accrualtime{denom: denom, time: accrualTime})
		return false
	})

	suite.Equal(expectedAccrualTimes, actualAccrualTimes)
}

func (suite *KeeperTestSuite) TestIterateEarnRewardAccrualTimes() {
	suite.SetupApp()

	expectedAccrualTimes := nonEmptyAccrualTimes

	for _, at := range expectedAccrualTimes {
		suite.keeper.SetEarnRewardAccrualTime(suite.ctx, at.denom, at.time)
	}

	var actualAccrualTimes []accrualtime
	suite.keeper.IterateEarnRewardAccrualTimes(suite.ctx, func(denom string, accrualTime time.Time) bool {
		actualAccrualTimes = append(actualAccrualTimes, accrualtime{denom: denom, time: accrualTime})
		return false
	})

	suite.Equal(expectedAccrualTimes, actualAccrualTimes)
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
