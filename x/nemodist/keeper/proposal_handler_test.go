package keeper_test

import (
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/incubus-network/nemo/x/nemodist/keeper"
	"github.com/incubus-network/nemo/x/nemodist/types"
)

func (suite *keeperTestSuite) TestHandleCommunityPoolMultiSpendProposal() {
	addr, distrKeeper, ctx := suite.Addrs[0], suite.App.GetDistrKeeper(), suite.Ctx
	initBalances := suite.BankKeeper.GetAllBalances(ctx, addr)

	// add coins to the module account and fund fee pool
	macc := distrKeeper.GetDistributionAccount(ctx)
	fundAmount := sdk.NewCoins(sdk.NewInt64Coin("unemo", 1000000))
	suite.Require().NoError(suite.App.FundModuleAccount(ctx, macc.GetName(), fundAmount))
	feePool := distrKeeper.GetFeePool(ctx)
	feePool.CommunityPool = sdk.NewDecCoinsFromCoins(fundAmount...)
	distrKeeper.SetFeePool(ctx, feePool)

	proposalAmount1 := int64(1100)
	proposalAmount2 := int64(1200)
	proposal := types.NewCommunityPoolMultiSpendProposal("test title", "description", []types.MultiSpendRecipient{
		{
			Address: addr.String(),
			Amount:  sdk.NewCoins(sdk.NewInt64Coin("unemo", proposalAmount1)),
		},
		{
			Address: addr.String(),
			Amount:  sdk.NewCoins(sdk.NewInt64Coin("unemo", proposalAmount2)),
		},
	})
	err := keeper.HandleCommunityPoolMultiSpendProposal(ctx, suite.Keeper, proposal)
	suite.Require().Nil(err)

	balances := suite.BankKeeper.GetAllBalances(ctx, addr)
	expected := initBalances.AmountOf("unemo").Add(sdkmath.NewInt(proposalAmount1 + proposalAmount2))
	suite.Require().Equal(expected, balances.AmountOf("unemo"))
}
