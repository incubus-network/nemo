package keeper_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"

	"github.com/incubus-network/nemo/x/incentive/types"
)

func (suite *HandlerTestSuite) TestPayoutDelegatorClaimMultiDenom() {
	userAddr := suite.addrs[0]
	receiverAddr := suite.addrs[1]

	authBulder := suite.authBuilder().
		WithSimpleAccount(userAddr, cs(c("ufury", 1e12))).
		WithSimpleAccount(receiverAddr, nil)

	incentBuilder := suite.incentiveBuilder().
		WithSimpleDelegatorRewardPeriod(types.BondDenom, cs(c("jinx", 1e6), c("swap", 1e6)))

	suite.SetupWithGenState(authBulder, incentBuilder)

	// create a delegation (need to create a validator first, which will have a self delegation)
	suite.NoError(
		suite.DeliverMsgCreateValidator(sdk.ValAddress(userAddr), c("ufury", 1e9)),
	)

	// Delete genesis validator to not influence rewards
	suite.App.DeleteGenesisValidator(suite.T(), suite.Ctx)

	// new block required to bond validator
	suite.NextBlockAfter(7 * time.Second)
	// Now the delegation is bonded, accumulate some delegator rewards
	suite.NextBlockAfter(7 * time.Second)

	preClaimBal := suite.GetBalance(userAddr)

	msg := types.NewMsgClaimDelegatorReward(
		userAddr.String(),
		types.Selections{
			types.NewSelection("jinx", "small"),
			types.NewSelection("swap", "medium"),
		},
	)

	// Claim denoms
	err := suite.DeliverIncentiveMsg(&msg)
	suite.NoError(err)

	// Check rewards were paid out
	expectedRewardsJinx := c("jinx", int64(0.2*float64(2*7*1e6)))
	expectedRewardsSwap := c("swap", int64(0.5*float64(2*7*1e6)))
	suite.BalanceEquals(userAddr, preClaimBal.Add(expectedRewardsJinx, expectedRewardsSwap))

	suite.VestingPeriodsEqual(userAddr, []vestingtypes.Period{
		{Length: (17+31)*secondsPerDay - 2*7, Amount: cs(expectedRewardsJinx)},
		{Length: (28 + 31 + 30 + 31 + 30) * secondsPerDay, Amount: cs(expectedRewardsSwap)}, // second length is stacked on top of the first
	})
	// Check that claimed coins have been removed from a claim's reward
	suite.DelegatorRewardEquals(userAddr, nil)
}

func (suite *HandlerTestSuite) TestPayoutDelegatorClaimSingleDenom() {
	userAddr := suite.addrs[0]

	authBulder := suite.authBuilder().
		WithSimpleAccount(userAddr, cs(c("ufury", 1e12)))

	incentBuilder := suite.incentiveBuilder().
		WithSimpleDelegatorRewardPeriod(types.BondDenom, cs(c("jinx", 1e6), c("swap", 1e6)))

	suite.SetupWithGenState(authBulder, incentBuilder)

	// create a delegation (need to create a validator first, which will have a self delegation)
	suite.NoError(
		suite.DeliverMsgCreateValidator(sdk.ValAddress(userAddr), c("ufury", 1e9)),
	)

	// Delete genesis validator to not influence rewards
	suite.App.DeleteGenesisValidator(suite.T(), suite.Ctx)

	// new block required to bond validator
	suite.NextBlockAfter(7 * time.Second)
	// Now the delegation is bonded, accumulate some delegator rewards
	suite.NextBlockAfter(7 * time.Second)

	preClaimBal := suite.GetBalance(userAddr)

	msg := types.NewMsgClaimDelegatorReward(
		userAddr.String(),
		types.Selections{
			types.NewSelection("swap", "large"),
		},
	)

	// Claim rewards
	err := suite.DeliverIncentiveMsg(&msg)
	suite.NoError(err)

	// Check rewards were paid out
	expectedRewards := c("swap", 2*7*1e6)
	suite.BalanceEquals(userAddr, preClaimBal.Add(expectedRewards))

	suite.VestingPeriodsEqual(userAddr, []vestingtypes.Period{
		{Length: (17+31+28+31+30+31+30+31+31+30+31+30+31)*secondsPerDay - 2*7, Amount: cs(expectedRewards)},
	})

	// Check that claimed coins have been removed from a claim's reward
	suite.DelegatorRewardEquals(userAddr, cs(c("jinx", 2*7*1e6)))
}
