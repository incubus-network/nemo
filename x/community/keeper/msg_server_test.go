package keeper_test

import (
	"testing"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"

	"github.com/incubus-network/nemo/app"
	"github.com/incubus-network/nemo/x/community/keeper"
	"github.com/incubus-network/nemo/x/community/testutil"
	"github.com/incubus-network/nemo/x/community/types"
)

type msgServerTestSuite struct {
	testutil.Suite

	communityPool sdk.AccAddress
	msgServer     types.MsgServer
}

func (suite *msgServerTestSuite) SetupTest() {
	suite.Suite.SetupTest()

	suite.communityPool = suite.App.GetAccountKeeper().GetModuleAddress(types.ModuleAccountName)
	suite.msgServer = keeper.NewMsgServerImpl(suite.Keeper)
}

func TestMsgServerTestSuite(t *testing.T) {
	suite.Run(t, new(msgServerTestSuite))
}

func (suite *msgServerTestSuite) TestMsgFundCommunityPool() {
	singleCoin := sdk.NewCoins(sdk.NewCoin("ufury", sdkmath.NewInt(2e6)))
	multipleCoins := sdk.NewCoins(
		sdk.NewCoin("ufury", sdkmath.NewInt(3e6)),
		sdk.NewCoin("musd", sdkmath.NewInt(1e7)),
	)
	testCases := []struct {
		name            string
		setup           func() *types.MsgFundCommunityPool
		expectedBalance sdk.Coins
		shouldPass      bool
	}{
		{
			name: "valid funding of single coin",
			setup: func() *types.MsgFundCommunityPool {
				sender := app.RandomAddress()
				suite.App.FundAccount(suite.Ctx, sender, singleCoin)
				return &types.MsgFundCommunityPool{
					Amount:    singleCoin,
					Depositor: sender.String(),
				}
			},
			expectedBalance: singleCoin,
			shouldPass:      true,
		},
		{
			name: "valid funding of multiple coins",
			setup: func() *types.MsgFundCommunityPool {
				sender := app.RandomAddress()
				suite.App.FundAccount(suite.Ctx, sender, multipleCoins)
				return &types.MsgFundCommunityPool{
					Amount:    multipleCoins,
					Depositor: sender.String(),
				}
			},
			expectedBalance: multipleCoins,
			shouldPass:      true,
		},
		{
			name: "invalid - failing message validation",
			setup: func() *types.MsgFundCommunityPool {
				return &types.MsgFundCommunityPool{
					Amount:    sdk.NewCoins(),
					Depositor: app.RandomAddress().String(),
				}
			},
			expectedBalance: sdk.NewCoins(),
			shouldPass:      false,
		},
		{
			name: "invalid - failing tx, insufficient funds",
			setup: func() *types.MsgFundCommunityPool {
				return &types.MsgFundCommunityPool{
					Amount:    singleCoin,
					Depositor: app.RandomAddress().String(),
				}
			},
			expectedBalance: sdk.NewCoins(),
			shouldPass:      false,
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			suite.SetupTest()

			msg := tc.setup()
			_, err := suite.msgServer.FundCommunityPool(sdk.WrapSDKContext(suite.Ctx), msg)
			if tc.shouldPass {
				suite.NoError(err)
			} else {
				suite.Error(err)
			}

			balance := suite.Keeper.GetModuleAccountBalance(suite.Ctx)
			suite.App.CheckBalance(suite.T(), suite.Ctx, suite.communityPool, balance)
		})
	}
}
