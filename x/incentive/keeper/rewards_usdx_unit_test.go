package keeper_test

import (
	"testing"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"

	cdptypes "github.com/incubus-network/nemo/x/cdp/types"
	"github.com/incubus-network/nemo/x/incentive/types"
)

// musdRewardsUnitTester contains common methods for running unit tests for keeper methods related to the MUSD minting rewards
type musdRewardsUnitTester struct {
	unitTester
}

func (suite *musdRewardsUnitTester) storeGlobalMUSDIndexes(indexes types.RewardIndexes) {
	for _, ri := range indexes {
		suite.keeper.SetMUSDMintingRewardFactor(suite.ctx, ri.CollateralType, ri.RewardFactor)
	}
}

func (suite *musdRewardsUnitTester) storeClaim(claim types.MUSDMintingClaim) {
	suite.keeper.SetMUSDMintingClaim(suite.ctx, claim)
}

type InitializeMUSDMintingClaimTests struct {
	musdRewardsUnitTester
}

func TestInitializeMUSDMintingClaims(t *testing.T) {
	suite.Run(t, new(InitializeMUSDMintingClaimTests))
}

func (suite *InitializeMUSDMintingClaimTests) TestClaimIndexIsSetWhenClaimDoesNotExist() {
	collateralType := "bnb-a"

	cdp := NewCDPBuilder(arbitraryAddress(), collateralType).Build()

	globalIndexes := types.RewardIndexes{{
		CollateralType: collateralType,
		RewardFactor:   d("0.2"),
	}}
	suite.storeGlobalMUSDIndexes(globalIndexes)

	suite.keeper.InitializeMUSDMintingClaim(suite.ctx, cdp)

	syncedClaim, f := suite.keeper.GetMUSDMintingClaim(suite.ctx, cdp.Owner)
	suite.True(f)
	suite.Equal(globalIndexes, syncedClaim.RewardIndexes)
}

func (suite *InitializeMUSDMintingClaimTests) TestClaimIndexIsSetWhenClaimExists() {
	collateralType := "bnb-a"

	claim := types.MUSDMintingClaim{
		BaseClaim: types.BaseClaim{
			Owner: arbitraryAddress(),
		},
		RewardIndexes: types.RewardIndexes{{
			CollateralType: collateralType,
			RewardFactor:   d("0.1"),
		}},
	}
	suite.storeClaim(claim)

	globalIndexes := types.RewardIndexes{{
		CollateralType: collateralType,
		RewardFactor:   d("0.2"),
	}}
	suite.storeGlobalMUSDIndexes(globalIndexes)

	cdp := NewCDPBuilder(claim.Owner, collateralType).Build()

	suite.keeper.InitializeMUSDMintingClaim(suite.ctx, cdp)

	syncedClaim, _ := suite.keeper.GetMUSDMintingClaim(suite.ctx, cdp.Owner)
	suite.Equal(globalIndexes, syncedClaim.RewardIndexes)
}

type SynchronizeMUSDMintingRewardTests struct {
	musdRewardsUnitTester
}

func TestSynchronizeMUSDMintingReward(t *testing.T) {
	suite.Run(t, new(SynchronizeMUSDMintingRewardTests))
}

func (suite *SynchronizeMUSDMintingRewardTests) TestRewardUnchangedWhenGlobalIndexesUnchanged() {
	unchangingRewardIndexes := nonEmptyRewardIndexes
	collateralType := extractFirstCollateralType(unchangingRewardIndexes)

	claim := types.MUSDMintingClaim{
		BaseClaim: types.BaseClaim{
			Owner:  arbitraryAddress(),
			Reward: c(types.MUSDMintingRewardDenom, 0),
		},
		RewardIndexes: unchangingRewardIndexes,
	}
	suite.storeClaim(claim)

	suite.storeGlobalMUSDIndexes(unchangingRewardIndexes)

	cdp := NewCDPBuilder(claim.Owner, collateralType).WithSourceShares(1e12).Build()

	suite.keeper.SynchronizeMUSDMintingReward(suite.ctx, cdp)

	syncedClaim, _ := suite.keeper.GetMUSDMintingClaim(suite.ctx, claim.Owner)
	suite.Equal(claim.Reward, syncedClaim.Reward)
}

func (suite *SynchronizeMUSDMintingRewardTests) TestRewardIsIncrementedWhenGlobalIndexIncreased() {
	collateralType := "bnb-a"

	claim := types.MUSDMintingClaim{
		BaseClaim: types.BaseClaim{
			Owner:  arbitraryAddress(),
			Reward: c(types.MUSDMintingRewardDenom, 0),
		},
		RewardIndexes: types.RewardIndexes{
			{
				CollateralType: collateralType,
				RewardFactor:   d("0.1"),
			},
		},
	}
	suite.storeClaim(claim)

	globalIndexes := types.RewardIndexes{
		{
			CollateralType: collateralType,
			RewardFactor:   d("0.2"),
		},
	}
	suite.storeGlobalMUSDIndexes(globalIndexes)

	cdp := NewCDPBuilder(claim.Owner, collateralType).WithSourceShares(1e12).Build()

	suite.keeper.SynchronizeMUSDMintingReward(suite.ctx, cdp)

	syncedClaim, _ := suite.keeper.GetMUSDMintingClaim(suite.ctx, claim.Owner)
	// reward is ( new index - old index ) * cdp.TotalPrincipal
	suite.Equal(c(types.MUSDMintingRewardDenom, 1e11), syncedClaim.Reward)
}

func (suite *SynchronizeMUSDMintingRewardTests) TestClaimIndexIsUpdatedWhenGlobalIndexIncreased() {
	claimsRewardIndexes := nonEmptyRewardIndexes
	collateralType := extractFirstCollateralType(claimsRewardIndexes)

	claim := types.MUSDMintingClaim{
		BaseClaim: types.BaseClaim{
			Owner:  arbitraryAddress(),
			Reward: c(types.MUSDMintingRewardDenom, 0),
		},
		RewardIndexes: claimsRewardIndexes,
	}
	suite.storeClaim(claim)

	globalIndexes := increaseRewardFactors(claimsRewardIndexes)
	suite.storeGlobalMUSDIndexes(globalIndexes)

	cdp := NewCDPBuilder(claim.Owner, collateralType).Build()

	suite.keeper.SynchronizeMUSDMintingReward(suite.ctx, cdp)

	syncedClaim, _ := suite.keeper.GetMUSDMintingClaim(suite.ctx, claim.Owner)

	// Only the claim's index for `collateralType` should have been changed
	i, _ := globalIndexes.Get(collateralType)
	expectedIndexes := claimsRewardIndexes.With(collateralType, i)
	suite.Equal(expectedIndexes, syncedClaim.RewardIndexes)
}

func (suite *SynchronizeMUSDMintingRewardTests) TestClaimIndexIsUpdatedWhenNewRewardAddedAndClaimAlreadyExists() {
	claimsRewardIndexes := types.RewardIndexes{
		{
			CollateralType: "bnb-a",
			RewardFactor:   d("0.1"),
		},
		{
			CollateralType: "busd-b",
			RewardFactor:   d("0.4"),
		},
	}
	newRewardIndex := types.NewRewardIndex("xrp-a", d("0.0001"))

	claim := types.MUSDMintingClaim{
		BaseClaim: types.BaseClaim{
			Owner:  arbitraryAddress(),
			Reward: c(types.MUSDMintingRewardDenom, 0),
		},
		RewardIndexes: claimsRewardIndexes,
	}
	suite.storeClaim(claim)

	globalIndexes := increaseRewardFactors(claimsRewardIndexes)
	globalIndexes = append(globalIndexes, newRewardIndex)
	suite.storeGlobalMUSDIndexes(globalIndexes)

	cdp := NewCDPBuilder(claim.Owner, newRewardIndex.CollateralType).Build()

	suite.keeper.SynchronizeMUSDMintingReward(suite.ctx, cdp)

	syncedClaim, _ := suite.keeper.GetMUSDMintingClaim(suite.ctx, claim.Owner)

	// Only the claim's index for `collateralType` should have been changed
	expectedIndexes := claimsRewardIndexes.With(newRewardIndex.CollateralType, newRewardIndex.RewardFactor)
	suite.Equal(expectedIndexes, syncedClaim.RewardIndexes)
}

func (suite *SynchronizeMUSDMintingRewardTests) TestClaimIsUnchangedWhenGlobalFactorMissing() {
	claimsRewardIndexes := nonEmptyRewardIndexes
	claim := types.MUSDMintingClaim{
		BaseClaim: types.BaseClaim{
			Owner:  arbitraryAddress(),
			Reward: c(types.MUSDMintingRewardDenom, 0),
		},
		RewardIndexes: claimsRewardIndexes,
	}
	suite.storeClaim(claim)
	// don't store any reward indexes

	// create a cdp with collateral type that doesn't exist in the claim's indexes, and does not have a corresponding global factor
	cdp := NewCDPBuilder(claim.Owner, "unrewardedcollateral").WithSourceShares(1e12).Build()

	suite.keeper.SynchronizeMUSDMintingReward(suite.ctx, cdp)

	syncedClaim, _ := suite.keeper.GetMUSDMintingClaim(suite.ctx, claim.Owner)
	suite.Equal(claim.RewardIndexes, syncedClaim.RewardIndexes)
	suite.Equal(claim.Reward, syncedClaim.Reward)
}

// CDPBuilder is a tool for creating a CDP in tests.
// The builder inherits from cdp.CDP, so fields can be accessed directly if a helper method doesn't exist.
type CDPBuilder struct {
	cdptypes.CDP
}

// NewCDPBuilder creates a CdpBuilder containing a CDP with owner and collateral type set.
func NewCDPBuilder(owner sdk.AccAddress, collateralType string) CDPBuilder {
	return CDPBuilder{
		CDP: cdptypes.CDP{
			Owner: owner,
			Type:  collateralType,
			// The zero value of Principal and AccumulatedFees (type sdk.Coin) is invalid as the denom is ""
			// Set them to the default denom, but with 0 amount.
			Principal:       c(cdptypes.DefaultStableDenom, 0),
			AccumulatedFees: c(cdptypes.DefaultStableDenom, 0),
			// zero value of sdk.Dec causes nil pointer panics
			InterestFactor: sdk.OneDec(),
		},
	}
}

// Build assembles and returns the final deposit.
func (builder CDPBuilder) Build() cdptypes.CDP { return builder.CDP }

// WithSourceShares adds a principal amount and interest factor such that the source shares for this CDP is equal to specified.
// With a factor of 1, the total principal is the source shares. This picks an arbitrary factor to ensure factors are accounted for in production code.
func (builder CDPBuilder) WithSourceShares(shares int64) CDPBuilder {
	if !builder.GetTotalPrincipal().Amount.Equal(sdk.ZeroInt()) {
		panic("setting source shares on cdp with existing principal or fees not implemented")
	}
	if !(builder.InterestFactor.IsNil() || builder.InterestFactor.Equal(sdk.OneDec())) {
		panic("setting source shares on cdp with existing interest factor not implemented")
	}
	// pick arbitrary interest factor
	factor := sdkmath.NewInt(2)

	// Calculate deposit amount that would equal the requested source shares given the above factor.
	principal := sdkmath.NewInt(shares).Mul(factor)

	builder.Principal = sdk.NewCoin(cdptypes.DefaultStableDenom, principal)
	builder.InterestFactor = sdk.NewDecFromInt(factor)

	return builder
}

func (builder CDPBuilder) WithPrincipal(principal sdkmath.Int) CDPBuilder {
	builder.Principal = sdk.NewCoin(cdptypes.DefaultStableDenom, principal)
	return builder
}

var nonEmptyRewardIndexes = types.RewardIndexes{
	{
		CollateralType: "bnb-a",
		RewardFactor:   d("0.1"),
	},
	{
		CollateralType: "busd-b",
		RewardFactor:   d("0.4"),
	},
}

func extractFirstCollateralType(indexes types.RewardIndexes) string {
	if len(indexes) == 0 {
		panic("cannot extract a collateral type from 0 length RewardIndexes")
	}
	return indexes[0].CollateralType
}
