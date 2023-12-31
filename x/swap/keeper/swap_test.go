package keeper_test

import (
	"errors"
	"fmt"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/incubus-network/nemo/x/swap/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmtime "github.com/tendermint/tendermint/types/time"
)

func (suite *keeperTestSuite) TestSwapExactForTokens() {
	suite.Keeper.SetParams(suite.Ctx, types.Params{
		SwapFee: sdk.MustNewDecFromStr("0.0025"),
	})
	owner := suite.CreateAccount(sdk.Coins{})
	reserves := sdk.NewCoins(
		sdk.NewCoin("ufury", sdkmath.NewInt(1000e6)),
		sdk.NewCoin("musd", sdkmath.NewInt(5000e6)),
	)
	totalShares := sdkmath.NewInt(30e6)
	poolID := suite.setupPool(reserves, totalShares, owner.GetAddress())

	balance := sdk.NewCoins(
		sdk.NewCoin("ufury", sdkmath.NewInt(10e6)),
	)
	requester := suite.NewAccountFromAddr(sdk.AccAddress("requester-----------"), balance)
	coinA := sdk.NewCoin("ufury", sdkmath.NewInt(1e6))
	coinB := sdk.NewCoin("musd", sdkmath.NewInt(5e6))

	err := suite.Keeper.SwapExactForTokens(suite.Ctx, requester.GetAddress(), coinA, coinB, sdk.MustNewDecFromStr("0.01"))
	suite.Require().NoError(err)

	expectedOutput := sdk.NewCoin("musd", sdkmath.NewInt(4982529))

	suite.AccountBalanceEqual(requester.GetAddress(), balance.Sub(coinA).Add(expectedOutput))
	suite.ModuleAccountBalanceEqual(reserves.Add(coinA).Sub(expectedOutput))
	suite.PoolLiquidityEqual(reserves.Add(coinA).Sub(expectedOutput))

	suite.EventsContains(suite.Ctx.EventManager().Events(), sdk.NewEvent(
		types.EventTypeSwapTrade,
		sdk.NewAttribute(types.AttributeKeyPoolID, poolID),
		sdk.NewAttribute(types.AttributeKeyRequester, requester.GetAddress().String()),
		sdk.NewAttribute(types.AttributeKeySwapInput, coinA.String()),
		sdk.NewAttribute(types.AttributeKeySwapOutput, expectedOutput.String()),
		sdk.NewAttribute(types.AttributeKeyFeePaid, "2500ufury"),
		sdk.NewAttribute(types.AttributeKeyExactDirection, "input"),
	))
}

func (suite *keeperTestSuite) TestSwapExactForTokens_OutputGreaterThanZero() {
	owner := suite.CreateAccount(sdk.Coins{})
	reserves := sdk.NewCoins(
		sdk.NewCoin("ufury", sdkmath.NewInt(10e6)),
		sdk.NewCoin("musd", sdkmath.NewInt(50e6)),
	)
	totalShares := sdkmath.NewInt(30e6)
	suite.setupPool(reserves, totalShares, owner.GetAddress())

	balance := sdk.NewCoins(
		sdk.NewCoin("musd", sdkmath.NewInt(10e6)),
	)
	requester := suite.NewAccountFromAddr(sdk.AccAddress("requester-----------"), balance)
	coinA := sdk.NewCoin("musd", sdkmath.NewInt(5))
	coinB := sdk.NewCoin("ufury", sdkmath.NewInt(1))

	err := suite.Keeper.SwapExactForTokens(suite.Ctx, requester.GetAddress(), coinA, coinB, sdk.MustNewDecFromStr("1"))
	suite.EqualError(err, "swap output rounds to zero, increase input amount: insufficient liquidity")
}

func (suite *keeperTestSuite) TestSwapExactForTokens_Slippage() {
	owner := suite.CreateAccount(sdk.Coins{})
	reserves := sdk.NewCoins(
		sdk.NewCoin("ufury", sdkmath.NewInt(100e6)),
		sdk.NewCoin("musd", sdkmath.NewInt(500e6)),
	)
	totalShares := sdkmath.NewInt(30e6)
	suite.setupPool(reserves, totalShares, owner.GetAddress())

	testCases := []struct {
		coinA      sdk.Coin
		coinB      sdk.Coin
		slippage   sdk.Dec
		fee        sdk.Dec
		shouldFail bool
	}{
		// positive slippage OK
		{sdk.NewCoin("ufury", sdkmath.NewInt(1e6)), sdk.NewCoin("musd", sdkmath.NewInt(2e6)), sdk.MustNewDecFromStr("0.01"), sdk.MustNewDecFromStr("0.0025"), false},
		{sdk.NewCoin("ufury", sdkmath.NewInt(1e6)), sdk.NewCoin("musd", sdkmath.NewInt(4e6)), sdk.MustNewDecFromStr("0.01"), sdk.MustNewDecFromStr("0.0025"), false},
		{sdk.NewCoin("musd", sdkmath.NewInt(50e6)), sdk.NewCoin("ufury", sdkmath.NewInt(5e6)), sdk.MustNewDecFromStr("0.01"), sdk.MustNewDecFromStr("0.0025"), false},
		{sdk.NewCoin("musd", sdkmath.NewInt(50e6)), sdk.NewCoin("ufury", sdkmath.NewInt(1e6)), sdk.MustNewDecFromStr("0.01"), sdk.MustNewDecFromStr("0.0025"), false},
		// positive slippage with zero slippage OK
		{sdk.NewCoin("ufury", sdkmath.NewInt(1e6)), sdk.NewCoin("musd", sdkmath.NewInt(2e6)), sdk.ZeroDec(), sdk.MustNewDecFromStr("0.0025"), false},
		{sdk.NewCoin("ufury", sdkmath.NewInt(1e6)), sdk.NewCoin("musd", sdkmath.NewInt(4e6)), sdk.ZeroDec(), sdk.MustNewDecFromStr("0.0025"), false},
		{sdk.NewCoin("musd", sdkmath.NewInt(50e6)), sdk.NewCoin("ufury", sdkmath.NewInt(5e6)), sdk.ZeroDec(), sdk.MustNewDecFromStr("0.0025"), false},
		{sdk.NewCoin("musd", sdkmath.NewInt(50e6)), sdk.NewCoin("ufury", sdkmath.NewInt(1e6)), sdk.ZeroDec(), sdk.MustNewDecFromStr("0.0025"), false},
		// exact zero slippage OK
		{sdk.NewCoin("ufury", sdkmath.NewInt(1e6)), sdk.NewCoin("musd", sdkmath.NewInt(4950495)), sdk.ZeroDec(), sdk.MustNewDecFromStr("0"), false},
		{sdk.NewCoin("ufury", sdkmath.NewInt(1e6)), sdk.NewCoin("musd", sdkmath.NewInt(4935790)), sdk.ZeroDec(), sdk.MustNewDecFromStr("0.003"), false},
		{sdk.NewCoin("ufury", sdkmath.NewInt(1e6)), sdk.NewCoin("musd", sdkmath.NewInt(4705299)), sdk.ZeroDec(), sdk.MustNewDecFromStr("0.05"), false},
		{sdk.NewCoin("musd", sdkmath.NewInt(5e6)), sdk.NewCoin("ufury", sdkmath.NewInt(990099)), sdk.ZeroDec(), sdk.MustNewDecFromStr("0"), false},
		{sdk.NewCoin("musd", sdkmath.NewInt(5e6)), sdk.NewCoin("ufury", sdkmath.NewInt(987158)), sdk.ZeroDec(), sdk.MustNewDecFromStr("0.003"), false},
		{sdk.NewCoin("musd", sdkmath.NewInt(5e6)), sdk.NewCoin("ufury", sdkmath.NewInt(941059)), sdk.ZeroDec(), sdk.MustNewDecFromStr("0.05"), false},
		// slippage failure, zero slippage tolerance
		{sdk.NewCoin("ufury", sdkmath.NewInt(1e6)), sdk.NewCoin("musd", sdkmath.NewInt(4950496)), sdk.ZeroDec(), sdk.MustNewDecFromStr("0"), true},
		{sdk.NewCoin("ufury", sdkmath.NewInt(1e6)), sdk.NewCoin("musd", sdkmath.NewInt(4935793)), sdk.ZeroDec(), sdk.MustNewDecFromStr("0.003"), true},
		{sdk.NewCoin("ufury", sdkmath.NewInt(1e6)), sdk.NewCoin("musd", sdkmath.NewInt(4705300)), sdk.ZeroDec(), sdk.MustNewDecFromStr("0.05"), true},
		{sdk.NewCoin("musd", sdkmath.NewInt(5e6)), sdk.NewCoin("ufury", sdkmath.NewInt(990100)), sdk.ZeroDec(), sdk.MustNewDecFromStr("0"), true},
		{sdk.NewCoin("musd", sdkmath.NewInt(5e6)), sdk.NewCoin("ufury", sdkmath.NewInt(987159)), sdk.ZeroDec(), sdk.MustNewDecFromStr("0.003"), true},
		{sdk.NewCoin("musd", sdkmath.NewInt(5e6)), sdk.NewCoin("ufury", sdkmath.NewInt(941060)), sdk.ZeroDec(), sdk.MustNewDecFromStr("0.05"), true},
		// slippage failure, 1 percent slippage
		{sdk.NewCoin("ufury", sdkmath.NewInt(1e6)), sdk.NewCoin("musd", sdkmath.NewInt(5000501)), sdk.MustNewDecFromStr("0.01"), sdk.MustNewDecFromStr("0"), true},
		{sdk.NewCoin("ufury", sdkmath.NewInt(1e6)), sdk.NewCoin("musd", sdkmath.NewInt(4985647)), sdk.MustNewDecFromStr("0.01"), sdk.MustNewDecFromStr("0.003"), true},
		{sdk.NewCoin("ufury", sdkmath.NewInt(1e6)), sdk.NewCoin("musd", sdkmath.NewInt(4752828)), sdk.MustNewDecFromStr("0.01"), sdk.MustNewDecFromStr("0.05"), true},
		{sdk.NewCoin("musd", sdkmath.NewInt(5e6)), sdk.NewCoin("ufury", sdkmath.NewInt(1000101)), sdk.MustNewDecFromStr("0.01"), sdk.MustNewDecFromStr("0"), true},
		{sdk.NewCoin("musd", sdkmath.NewInt(5e6)), sdk.NewCoin("ufury", sdkmath.NewInt(997130)), sdk.MustNewDecFromStr("0.01"), sdk.MustNewDecFromStr("0.003"), true},
		{sdk.NewCoin("musd", sdkmath.NewInt(5e6)), sdk.NewCoin("ufury", sdkmath.NewInt(950565)), sdk.MustNewDecFromStr("0.01"), sdk.MustNewDecFromStr("0.05"), true},
		// slippage OK, 1 percent slippage
		{sdk.NewCoin("ufury", sdkmath.NewInt(1e6)), sdk.NewCoin("musd", sdkmath.NewInt(5000500)), sdk.MustNewDecFromStr("0.01"), sdk.MustNewDecFromStr("0"), false},
		{sdk.NewCoin("ufury", sdkmath.NewInt(1e6)), sdk.NewCoin("musd", sdkmath.NewInt(4985646)), sdk.MustNewDecFromStr("0.01"), sdk.MustNewDecFromStr("0.003"), false},
		{sdk.NewCoin("ufury", sdkmath.NewInt(1e6)), sdk.NewCoin("musd", sdkmath.NewInt(4752827)), sdk.MustNewDecFromStr("0.01"), sdk.MustNewDecFromStr("0.05"), false},
		{sdk.NewCoin("musd", sdkmath.NewInt(5e6)), sdk.NewCoin("ufury", sdkmath.NewInt(1000100)), sdk.MustNewDecFromStr("0.01"), sdk.MustNewDecFromStr("0"), false},
		{sdk.NewCoin("musd", sdkmath.NewInt(5e6)), sdk.NewCoin("ufury", sdkmath.NewInt(997129)), sdk.MustNewDecFromStr("0.01"), sdk.MustNewDecFromStr("0.003"), false},
		{sdk.NewCoin("musd", sdkmath.NewInt(5e6)), sdk.NewCoin("ufury", sdkmath.NewInt(950564)), sdk.MustNewDecFromStr("0.01"), sdk.MustNewDecFromStr("0.05"), false},
	}

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("coinA=%s coinB=%s slippage=%s fee=%s", tc.coinA, tc.coinB, tc.slippage, tc.fee), func() {
			suite.SetupTest()
			suite.Keeper.SetParams(suite.Ctx, types.Params{
				SwapFee: tc.fee,
			})
			owner := suite.CreateAccount(sdk.Coins{})
			reserves := sdk.NewCoins(
				sdk.NewCoin("ufury", sdkmath.NewInt(100e6)),
				sdk.NewCoin("musd", sdkmath.NewInt(500e6)),
			)
			totalShares := sdkmath.NewInt(30e6)
			suite.setupPool(reserves, totalShares, owner.GetAddress())
			balance := sdk.NewCoins(
				sdk.NewCoin("ufury", sdkmath.NewInt(100e6)),
				sdk.NewCoin("musd", sdkmath.NewInt(100e6)),
			)
			requester := suite.NewAccountFromAddr(sdk.AccAddress("requester-----------"), balance)

			ctx := suite.App.NewContext(true, tmproto.Header{Height: 1, Time: tmtime.Now()})
			err := suite.Keeper.SwapExactForTokens(ctx, requester.GetAddress(), tc.coinA, tc.coinB, tc.slippage)

			if tc.shouldFail {
				suite.Require().Error(err)
				suite.Contains(err.Error(), "slippage exceeded")
			} else {
				suite.NoError(err)
			}
		})
	}
}

func (suite *keeperTestSuite) TestSwapExactForTokens_InsufficientFunds() {
	testCases := []struct {
		name     string
		balanceA sdk.Coin
		coinA    sdk.Coin
		coinB    sdk.Coin
	}{
		{"no ufury balance", sdk.NewCoin("ufury", sdk.ZeroInt()), sdk.NewCoin("ufury", sdkmath.NewInt(100)), sdk.NewCoin("musd", sdkmath.NewInt(500))},
		{"no musd balance", sdk.NewCoin("musd", sdk.ZeroInt()), sdk.NewCoin("musd", sdkmath.NewInt(500)), sdk.NewCoin("ufury", sdkmath.NewInt(100))},
		{"low ufury balance", sdk.NewCoin("ufury", sdkmath.NewInt(1000000)), sdk.NewCoin("ufury", sdkmath.NewInt(1000001)), sdk.NewCoin("musd", sdkmath.NewInt(5000000))},
		{"low ufury balance", sdk.NewCoin("musd", sdkmath.NewInt(5000000)), sdk.NewCoin("musd", sdkmath.NewInt(5000001)), sdk.NewCoin("ufury", sdkmath.NewInt(1000000))},
		{"large ufury balance difference", sdk.NewCoin("ufury", sdkmath.NewInt(100e6)), sdk.NewCoin("ufury", sdkmath.NewInt(1000e6)), sdk.NewCoin("musd", sdkmath.NewInt(5000e6))},
		{"large musd balance difference", sdk.NewCoin("musd", sdkmath.NewInt(500e6)), sdk.NewCoin("musd", sdkmath.NewInt(5000e6)), sdk.NewCoin("ufury", sdkmath.NewInt(1000e6))},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			suite.SetupTest()
			owner := suite.CreateAccount(sdk.Coins{})
			reserves := sdk.NewCoins(
				sdk.NewCoin("ufury", sdkmath.NewInt(100000e6)),
				sdk.NewCoin("musd", sdkmath.NewInt(500000e6)),
			)
			totalShares := sdkmath.NewInt(30000e6)
			suite.setupPool(reserves, totalShares, owner.GetAddress())
			balance := sdk.NewCoins(tc.balanceA)
			requester := suite.NewAccountFromAddr(sdk.AccAddress("requester-----------"), balance)

			ctx := suite.App.NewContext(true, tmproto.Header{Height: 1, Time: tmtime.Now()})
			err := suite.Keeper.SwapExactForTokens(ctx, requester.GetAddress(), tc.coinA, tc.coinB, sdk.MustNewDecFromStr("0.1"))
			suite.Require().True(errors.Is(err, sdkerrors.ErrInsufficientFunds), fmt.Sprintf("got err %s", err))
		})
	}
}

func (suite *keeperTestSuite) TestSwapExactForTokens_InsufficientFunds_Vesting() {
	testCases := []struct {
		name     string
		balanceA sdk.Coin
		vestingA sdk.Coin
		coinA    sdk.Coin
		coinB    sdk.Coin
	}{
		{"no ufury balance, vesting only", sdk.NewCoin("ufury", sdk.ZeroInt()), sdk.NewCoin("ufury", sdkmath.NewInt(100)), sdk.NewCoin("ufury", sdkmath.NewInt(100)), sdk.NewCoin("musd", sdkmath.NewInt(500))},
		{"no musd balance, vesting only", sdk.NewCoin("musd", sdk.ZeroInt()), sdk.NewCoin("musd", sdkmath.NewInt(500)), sdk.NewCoin("musd", sdkmath.NewInt(500)), sdk.NewCoin("ufury", sdkmath.NewInt(100))},
		{"low ufury balance, vesting matches exact", sdk.NewCoin("ufury", sdkmath.NewInt(1000000)), sdk.NewCoin("ufury", sdkmath.NewInt(1)), sdk.NewCoin("ufury", sdkmath.NewInt(1000001)), sdk.NewCoin("musd", sdkmath.NewInt(5000000))},
		{"low ufury balance, vesting matches exact", sdk.NewCoin("musd", sdkmath.NewInt(5000000)), sdk.NewCoin("musd", sdkmath.NewInt(1)), sdk.NewCoin("musd", sdkmath.NewInt(5000001)), sdk.NewCoin("ufury", sdkmath.NewInt(1000000))},
		{"large ufury balance difference, vesting covers difference", sdk.NewCoin("ufury", sdkmath.NewInt(100e6)), sdk.NewCoin("ufury", sdkmath.NewInt(1000e6)), sdk.NewCoin("ufury", sdkmath.NewInt(1000e6)), sdk.NewCoin("musd", sdkmath.NewInt(5000e6))},
		{"large musd balance difference, vesting covers difference", sdk.NewCoin("musd", sdkmath.NewInt(500e6)), sdk.NewCoin("musd", sdkmath.NewInt(5000e6)), sdk.NewCoin("musd", sdkmath.NewInt(5000e6)), sdk.NewCoin("ufury", sdkmath.NewInt(1000e6))},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			suite.SetupTest()
			owner := suite.CreateAccount(sdk.Coins{})
			reserves := sdk.NewCoins(
				sdk.NewCoin("ufury", sdkmath.NewInt(100000e6)),
				sdk.NewCoin("musd", sdkmath.NewInt(500000e6)),
			)
			totalShares := sdkmath.NewInt(30000e6)
			suite.setupPool(reserves, totalShares, owner.GetAddress())
			balance := sdk.NewCoins(tc.balanceA)
			vesting := sdk.NewCoins(tc.vestingA)
			requester := suite.CreateVestingAccount(balance, vesting)

			ctx := suite.App.NewContext(true, tmproto.Header{Height: 1, Time: tmtime.Now()})
			err := suite.Keeper.SwapExactForTokens(ctx, requester.GetAddress(), tc.coinA, tc.coinB, sdk.MustNewDecFromStr("0.1"))
			suite.Require().True(errors.Is(err, sdkerrors.ErrInsufficientFunds), fmt.Sprintf("got err %s", err))
		})
	}
}

func (suite *keeperTestSuite) TestSwapExactForTokens_PoolNotFound() {
	owner := suite.CreateAccount(sdk.Coins{})
	reserves := sdk.NewCoins(
		sdk.NewCoin("ufury", sdkmath.NewInt(1000e6)),
		sdk.NewCoin("musd", sdkmath.NewInt(5000e6)),
	)
	totalShares := sdkmath.NewInt(3000e6)
	poolID := suite.setupPool(reserves, totalShares, owner.GetAddress())
	suite.Keeper.DeletePool(suite.Ctx, poolID)

	balance := sdk.NewCoins(
		sdk.NewCoin("ufury", sdkmath.NewInt(10e6)),
		sdk.NewCoin("musd", sdkmath.NewInt(10e6)),
	)
	requester := suite.NewAccountFromAddr(sdk.AccAddress("requester-----------"), balance)
	coinA := sdk.NewCoin("ufury", sdkmath.NewInt(1e6))
	coinB := sdk.NewCoin("musd", sdkmath.NewInt(5e6))

	err := suite.Keeper.SwapExactForTokens(suite.Ctx, requester.GetAddress(), coinA, coinB, sdk.MustNewDecFromStr("0.01"))
	suite.EqualError(err, "pool ufury:musd not found: invalid pool")

	err = suite.Keeper.SwapExactForTokens(suite.Ctx, requester.GetAddress(), coinB, coinA, sdk.MustNewDecFromStr("0.01"))
	suite.EqualError(err, "pool ufury:musd not found: invalid pool")
}

func (suite *keeperTestSuite) TestSwapExactForTokens_PanicOnInvalidPool() {
	owner := suite.CreateAccount(sdk.Coins{})
	reserves := sdk.NewCoins(
		sdk.NewCoin("ufury", sdkmath.NewInt(1000e6)),
		sdk.NewCoin("musd", sdkmath.NewInt(5000e6)),
	)
	totalShares := sdkmath.NewInt(3000e6)
	poolID := suite.setupPool(reserves, totalShares, owner.GetAddress())

	poolRecord, found := suite.Keeper.GetPool(suite.Ctx, poolID)
	suite.Require().True(found, "expected pool record to exist")

	poolRecord.TotalShares = sdk.ZeroInt()
	suite.Keeper.SetPool_Raw(suite.Ctx, poolRecord)

	balance := sdk.NewCoins(
		sdk.NewCoin("ufury", sdkmath.NewInt(10e6)),
		sdk.NewCoin("musd", sdkmath.NewInt(10e6)),
	)
	requester := suite.NewAccountFromAddr(sdk.AccAddress("requester-----------"), balance)
	coinA := sdk.NewCoin("ufury", sdkmath.NewInt(1e6))
	coinB := sdk.NewCoin("musd", sdkmath.NewInt(5e6))

	suite.PanicsWithValue("invalid pool ufury:musd: total shares must be greater than zero: invalid pool", func() {
		_ = suite.Keeper.SwapExactForTokens(suite.Ctx, requester.GetAddress(), coinA, coinB, sdk.MustNewDecFromStr("0.01"))
	}, "expected invalid pool record to panic")

	suite.PanicsWithValue("invalid pool ufury:musd: total shares must be greater than zero: invalid pool", func() {
		_ = suite.Keeper.SwapExactForTokens(suite.Ctx, requester.GetAddress(), coinB, coinA, sdk.MustNewDecFromStr("0.01"))
	}, "expected invalid pool record to panic")
}

func (suite *keeperTestSuite) TestSwapExactForTokens_PanicOnInsufficientModuleAccFunds() {
	owner := suite.CreateAccount(sdk.Coins{})
	reserves := sdk.NewCoins(
		sdk.NewCoin("ufury", sdkmath.NewInt(1000e6)),
		sdk.NewCoin("musd", sdkmath.NewInt(5000e6)),
	)
	totalShares := sdkmath.NewInt(3000e6)
	suite.setupPool(reserves, totalShares, owner.GetAddress())

	suite.RemoveCoinsFromModule(sdk.NewCoins(
		sdk.NewCoin("ufury", sdkmath.NewInt(1000e6)),
		sdk.NewCoin("musd", sdkmath.NewInt(5000e6)),
	))

	balance := sdk.NewCoins(
		sdk.NewCoin("ufury", sdkmath.NewInt(10e6)),
		sdk.NewCoin("musd", sdkmath.NewInt(10e6)),
	)
	requester := suite.NewAccountFromAddr(sdk.AccAddress("requester-----------"), balance)
	coinA := sdk.NewCoin("ufury", sdkmath.NewInt(1e6))
	coinB := sdk.NewCoin("musd", sdkmath.NewInt(5e6))

	suite.Panics(func() {
		_ = suite.Keeper.SwapExactForTokens(suite.Ctx, requester.GetAddress(), coinA, coinB, sdk.MustNewDecFromStr("0.01"))
	}, "expected panic when module account does not have enough funds")

	suite.Panics(func() {
		_ = suite.Keeper.SwapExactForTokens(suite.Ctx, requester.GetAddress(), coinA, coinB, sdk.MustNewDecFromStr("0.01"))
	}, "expected panic when module account does not have enough funds")
}

func (suite *keeperTestSuite) TestSwapForExactTokens() {
	suite.Keeper.SetParams(suite.Ctx, types.Params{
		SwapFee: sdk.MustNewDecFromStr("0.0025"),
	})
	owner := suite.CreateAccount(sdk.Coins{})
	reserves := sdk.NewCoins(
		sdk.NewCoin("ufury", sdkmath.NewInt(1000e6)),
		sdk.NewCoin("musd", sdkmath.NewInt(5000e6)),
	)
	totalShares := sdkmath.NewInt(30e6)
	poolID := suite.setupPool(reserves, totalShares, owner.GetAddress())

	balance := sdk.NewCoins(
		sdk.NewCoin("ufury", sdkmath.NewInt(10e6)),
	)
	requester := suite.NewAccountFromAddr(sdk.AccAddress("requester-----------"), balance)
	coinA := sdk.NewCoin("ufury", sdkmath.NewInt(1e6))
	coinB := sdk.NewCoin("musd", sdkmath.NewInt(5e6))

	err := suite.Keeper.SwapForExactTokens(suite.Ctx, requester.GetAddress(), coinA, coinB, sdk.MustNewDecFromStr("0.01"))
	suite.Require().NoError(err)

	expectedInput := sdk.NewCoin("ufury", sdkmath.NewInt(1003511))

	suite.AccountBalanceEqual(requester.GetAddress(), balance.Sub(expectedInput).Add(coinB))
	suite.ModuleAccountBalanceEqual(reserves.Add(expectedInput).Sub(coinB))
	suite.PoolLiquidityEqual(reserves.Add(expectedInput).Sub(coinB))

	suite.EventsContains(suite.Ctx.EventManager().Events(), sdk.NewEvent(
		types.EventTypeSwapTrade,
		sdk.NewAttribute(types.AttributeKeyPoolID, poolID),
		sdk.NewAttribute(types.AttributeKeyRequester, requester.GetAddress().String()),
		sdk.NewAttribute(types.AttributeKeySwapInput, expectedInput.String()),
		sdk.NewAttribute(types.AttributeKeySwapOutput, coinB.String()),
		sdk.NewAttribute(types.AttributeKeyFeePaid, "2509ufury"),
		sdk.NewAttribute(types.AttributeKeyExactDirection, "output"),
	))
}

func (suite *keeperTestSuite) TestSwapForExactTokens_OutputLessThanPoolReserves() {
	owner := suite.CreateAccount(sdk.Coins{})
	reserves := sdk.NewCoins(
		sdk.NewCoin("ufury", sdkmath.NewInt(100e6)),
		sdk.NewCoin("musd", sdkmath.NewInt(500e6)),
	)
	totalShares := sdkmath.NewInt(300e6)
	suite.setupPool(reserves, totalShares, owner.GetAddress())

	balance := sdk.NewCoins(
		sdk.NewCoin("ufury", sdkmath.NewInt(1000e6)),
	)
	requester := suite.NewAccountFromAddr(sdk.AccAddress("requester-----------"), balance)
	coinA := sdk.NewCoin("ufury", sdkmath.NewInt(1e6))

	coinB := sdk.NewCoin("musd", sdkmath.NewInt(500e6).Add(sdk.OneInt()))
	err := suite.Keeper.SwapForExactTokens(suite.Ctx, requester.GetAddress(), coinA, coinB, sdk.MustNewDecFromStr("0.01"))
	suite.EqualError(err, "output 500000001 >= pool reserves 500000000: insufficient liquidity")

	coinB = sdk.NewCoin("musd", sdkmath.NewInt(500e6))
	err = suite.Keeper.SwapForExactTokens(suite.Ctx, requester.GetAddress(), coinA, coinB, sdk.MustNewDecFromStr("0.01"))
	suite.EqualError(err, "output 500000000 >= pool reserves 500000000: insufficient liquidity")
}

func (suite *keeperTestSuite) TestSwapForExactTokens_Slippage() {
	owner := suite.CreateAccount(sdk.Coins{})
	reserves := sdk.NewCoins(
		sdk.NewCoin("ufury", sdkmath.NewInt(100e6)),
		sdk.NewCoin("musd", sdkmath.NewInt(500e6)),
	)
	totalShares := sdkmath.NewInt(30e6)
	suite.setupPool(reserves, totalShares, owner.GetAddress())

	testCases := []struct {
		coinA      sdk.Coin
		coinB      sdk.Coin
		slippage   sdk.Dec
		fee        sdk.Dec
		shouldFail bool
	}{
		// positive slippage OK
		{sdk.NewCoin("ufury", sdkmath.NewInt(5e6)), sdk.NewCoin("musd", sdkmath.NewInt(5e6)), sdk.MustNewDecFromStr("0.01"), sdk.MustNewDecFromStr("0.0025"), false},
		{sdk.NewCoin("ufury", sdkmath.NewInt(5e6)), sdk.NewCoin("musd", sdkmath.NewInt(5e6)), sdk.MustNewDecFromStr("0.01"), sdk.MustNewDecFromStr("0.0025"), false},
		{sdk.NewCoin("musd", sdkmath.NewInt(100e6)), sdk.NewCoin("ufury", sdkmath.NewInt(10e6)), sdk.MustNewDecFromStr("0.01"), sdk.MustNewDecFromStr("0.0025"), false},
		{sdk.NewCoin("musd", sdkmath.NewInt(100e6)), sdk.NewCoin("ufury", sdkmath.NewInt(10e6)), sdk.MustNewDecFromStr("0.01"), sdk.MustNewDecFromStr("0.0025"), false},
		// positive slippage with zero slippage OK
		{sdk.NewCoin("ufury", sdkmath.NewInt(5e6)), sdk.NewCoin("musd", sdkmath.NewInt(5e6)), sdk.ZeroDec(), sdk.MustNewDecFromStr("0.0025"), false},
		{sdk.NewCoin("ufury", sdkmath.NewInt(5e6)), sdk.NewCoin("musd", sdkmath.NewInt(5e6)), sdk.ZeroDec(), sdk.MustNewDecFromStr("0.0025"), false},
		{sdk.NewCoin("musd", sdkmath.NewInt(100e6)), sdk.NewCoin("ufury", sdkmath.NewInt(10e6)), sdk.ZeroDec(), sdk.MustNewDecFromStr("0.0025"), false},
		{sdk.NewCoin("musd", sdkmath.NewInt(100e6)), sdk.NewCoin("ufury", sdkmath.NewInt(10e6)), sdk.ZeroDec(), sdk.MustNewDecFromStr("0.0025"), false},
		// exact zero slippage OK
		{sdk.NewCoin("ufury", sdkmath.NewInt(1010102)), sdk.NewCoin("musd", sdkmath.NewInt(5e6)), sdk.ZeroDec(), sdk.MustNewDecFromStr("0"), false},
		{sdk.NewCoin("ufury", sdkmath.NewInt(1010102)), sdk.NewCoin("musd", sdkmath.NewInt(5e6)), sdk.ZeroDec(), sdk.MustNewDecFromStr("0.003"), false},
		{sdk.NewCoin("ufury", sdkmath.NewInt(1010102)), sdk.NewCoin("musd", sdkmath.NewInt(5e6)), sdk.ZeroDec(), sdk.MustNewDecFromStr("0.05"), false},
		{sdk.NewCoin("musd", sdkmath.NewInt(5050506)), sdk.NewCoin("ufury", sdkmath.NewInt(1e6)), sdk.ZeroDec(), sdk.MustNewDecFromStr("0"), false},
		{sdk.NewCoin("musd", sdkmath.NewInt(5050506)), sdk.NewCoin("ufury", sdkmath.NewInt(1e6)), sdk.ZeroDec(), sdk.MustNewDecFromStr("0.003"), false},
		{sdk.NewCoin("musd", sdkmath.NewInt(5050506)), sdk.NewCoin("ufury", sdkmath.NewInt(1e6)), sdk.ZeroDec(), sdk.MustNewDecFromStr("0.05"), false},
		// slippage failure, zero slippage tolerance
		{sdk.NewCoin("ufury", sdkmath.NewInt(1010101)), sdk.NewCoin("musd", sdkmath.NewInt(5e6)), sdk.ZeroDec(), sdk.MustNewDecFromStr("0"), true},
		{sdk.NewCoin("ufury", sdkmath.NewInt(1010101)), sdk.NewCoin("musd", sdkmath.NewInt(5e6)), sdk.ZeroDec(), sdk.MustNewDecFromStr("0.003"), true},
		{sdk.NewCoin("ufury", sdkmath.NewInt(1010101)), sdk.NewCoin("musd", sdkmath.NewInt(5e6)), sdk.ZeroDec(), sdk.MustNewDecFromStr("0.05"), true},
		{sdk.NewCoin("musd", sdkmath.NewInt(5050505)), sdk.NewCoin("ufury", sdkmath.NewInt(1e6)), sdk.ZeroDec(), sdk.MustNewDecFromStr("0"), true},
		{sdk.NewCoin("musd", sdkmath.NewInt(5050505)), sdk.NewCoin("ufury", sdkmath.NewInt(1e6)), sdk.ZeroDec(), sdk.MustNewDecFromStr("0.003"), true},
		{sdk.NewCoin("musd", sdkmath.NewInt(5050505)), sdk.NewCoin("ufury", sdkmath.NewInt(1e6)), sdk.ZeroDec(), sdk.MustNewDecFromStr("0.05"), true},
		// slippage failure, 1 percent slippage
		{sdk.NewCoin("ufury", sdkmath.NewInt(1000000)), sdk.NewCoin("musd", sdkmath.NewInt(5e6)), sdk.MustNewDecFromStr("0.01"), sdk.MustNewDecFromStr("0"), true},
		{sdk.NewCoin("ufury", sdkmath.NewInt(1000000)), sdk.NewCoin("musd", sdkmath.NewInt(5e6)), sdk.MustNewDecFromStr("0.01"), sdk.MustNewDecFromStr("0.003"), true},
		{sdk.NewCoin("ufury", sdkmath.NewInt(1000000)), sdk.NewCoin("musd", sdkmath.NewInt(5e6)), sdk.MustNewDecFromStr("0.01"), sdk.MustNewDecFromStr("0.05"), true},
		{sdk.NewCoin("musd", sdkmath.NewInt(5000000)), sdk.NewCoin("ufury", sdkmath.NewInt(1e6)), sdk.MustNewDecFromStr("0.01"), sdk.MustNewDecFromStr("0"), true},
		{sdk.NewCoin("musd", sdkmath.NewInt(5000000)), sdk.NewCoin("ufury", sdkmath.NewInt(1e6)), sdk.MustNewDecFromStr("0.01"), sdk.MustNewDecFromStr("0.003"), true},
		{sdk.NewCoin("musd", sdkmath.NewInt(5000000)), sdk.NewCoin("ufury", sdkmath.NewInt(1e6)), sdk.MustNewDecFromStr("0.01"), sdk.MustNewDecFromStr("0.05"), true},
		// slippage OK, 1 percent slippage
		{sdk.NewCoin("ufury", sdkmath.NewInt(1000001)), sdk.NewCoin("musd", sdkmath.NewInt(5e6)), sdk.MustNewDecFromStr("0.01"), sdk.MustNewDecFromStr("0"), false},
		{sdk.NewCoin("ufury", sdkmath.NewInt(1000001)), sdk.NewCoin("musd", sdkmath.NewInt(5e6)), sdk.MustNewDecFromStr("0.01"), sdk.MustNewDecFromStr("0.003"), false},
		{sdk.NewCoin("ufury", sdkmath.NewInt(1000001)), sdk.NewCoin("musd", sdkmath.NewInt(5e6)), sdk.MustNewDecFromStr("0.01"), sdk.MustNewDecFromStr("0.05"), false},
		{sdk.NewCoin("musd", sdkmath.NewInt(5000001)), sdk.NewCoin("ufury", sdkmath.NewInt(1e6)), sdk.MustNewDecFromStr("0.01"), sdk.MustNewDecFromStr("0"), false},
		{sdk.NewCoin("musd", sdkmath.NewInt(5000001)), sdk.NewCoin("ufury", sdkmath.NewInt(1e6)), sdk.MustNewDecFromStr("0.01"), sdk.MustNewDecFromStr("0.003"), false},
		{sdk.NewCoin("musd", sdkmath.NewInt(5000001)), sdk.NewCoin("ufury", sdkmath.NewInt(1e6)), sdk.MustNewDecFromStr("0.01"), sdk.MustNewDecFromStr("0.05"), false},
	}

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("coinA=%s coinB=%s slippage=%s fee=%s", tc.coinA, tc.coinB, tc.slippage, tc.fee), func() {
			suite.SetupTest()
			suite.Keeper.SetParams(suite.Ctx, types.Params{
				SwapFee: tc.fee,
			})
			owner := suite.CreateAccount(sdk.Coins{})
			reserves := sdk.NewCoins(
				sdk.NewCoin("ufury", sdkmath.NewInt(100e6)),
				sdk.NewCoin("musd", sdkmath.NewInt(500e6)),
			)
			totalShares := sdkmath.NewInt(30e6)
			suite.setupPool(reserves, totalShares, owner.GetAddress())
			balance := sdk.NewCoins(
				sdk.NewCoin("ufury", sdkmath.NewInt(100e6)),
				sdk.NewCoin("musd", sdkmath.NewInt(100e6)),
			)
			requester := suite.NewAccountFromAddr(sdk.AccAddress("requester-----------"), balance)

			ctx := suite.App.NewContext(true, tmproto.Header{Height: 1, Time: tmtime.Now()})
			err := suite.Keeper.SwapForExactTokens(ctx, requester.GetAddress(), tc.coinA, tc.coinB, tc.slippage)

			if tc.shouldFail {
				suite.Require().Error(err)
				suite.Contains(err.Error(), "slippage exceeded")
			} else {
				suite.NoError(err)
			}
		})
	}
}

func (suite *keeperTestSuite) TestSwapForExactTokens_InsufficientFunds() {
	testCases := []struct {
		name     string
		balanceA sdk.Coin
		coinA    sdk.Coin
		coinB    sdk.Coin
	}{
		{"no ufury balance", sdk.NewCoin("ufury", sdk.ZeroInt()), sdk.NewCoin("ufury", sdkmath.NewInt(100)), sdk.NewCoin("musd", sdkmath.NewInt(500))},
		{"no musd balance", sdk.NewCoin("musd", sdk.ZeroInt()), sdk.NewCoin("musd", sdkmath.NewInt(500)), sdk.NewCoin("ufury", sdkmath.NewInt(100))},
		{"low ufury balance", sdk.NewCoin("ufury", sdkmath.NewInt(1000000)), sdk.NewCoin("ufury", sdkmath.NewInt(1000000)), sdk.NewCoin("musd", sdkmath.NewInt(5000000))},
		{"low ufury balance", sdk.NewCoin("musd", sdkmath.NewInt(5000000)), sdk.NewCoin("musd", sdkmath.NewInt(5000000)), sdk.NewCoin("ufury", sdkmath.NewInt(1000000))},
		{"large ufury balance difference", sdk.NewCoin("ufury", sdkmath.NewInt(100e6)), sdk.NewCoin("ufury", sdkmath.NewInt(1000e6)), sdk.NewCoin("musd", sdkmath.NewInt(5000e6))},
		{"large musd balance difference", sdk.NewCoin("musd", sdkmath.NewInt(500e6)), sdk.NewCoin("musd", sdkmath.NewInt(5000e6)), sdk.NewCoin("ufury", sdkmath.NewInt(1000e6))},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			suite.SetupTest()
			owner := suite.CreateAccount(sdk.Coins{})
			reserves := sdk.NewCoins(
				sdk.NewCoin("ufury", sdkmath.NewInt(100000e6)),
				sdk.NewCoin("musd", sdkmath.NewInt(500000e6)),
			)
			totalShares := sdkmath.NewInt(30000e6)
			suite.setupPool(reserves, totalShares, owner.GetAddress())
			balance := sdk.NewCoins(tc.balanceA)
			requester := suite.NewAccountFromAddr(sdk.AccAddress("requester-----------"), balance)

			ctx := suite.App.NewContext(true, tmproto.Header{Height: 1, Time: tmtime.Now()})
			err := suite.Keeper.SwapForExactTokens(ctx, requester.GetAddress(), tc.coinA, tc.coinB, sdk.MustNewDecFromStr("0.1"))
			suite.Require().True(errors.Is(err, sdkerrors.ErrInsufficientFunds), fmt.Sprintf("got err %s", err))
		})
	}
}

func (suite *keeperTestSuite) TestSwapForExactTokens_InsufficientFunds_Vesting() {
	testCases := []struct {
		name     string
		balanceA sdk.Coin
		vestingA sdk.Coin
		coinA    sdk.Coin
		coinB    sdk.Coin
	}{
		{"no ufury balance, vesting only", sdk.NewCoin("ufury", sdk.ZeroInt()), sdk.NewCoin("ufury", sdkmath.NewInt(100)), sdk.NewCoin("ufury", sdkmath.NewInt(1000)), sdk.NewCoin("musd", sdkmath.NewInt(500))},
		{"no musd balance, vesting only", sdk.NewCoin("musd", sdk.ZeroInt()), sdk.NewCoin("musd", sdkmath.NewInt(500)), sdk.NewCoin("musd", sdkmath.NewInt(5000)), sdk.NewCoin("ufury", sdkmath.NewInt(100))},
		{"low ufury balance, vesting matches exact", sdk.NewCoin("ufury", sdkmath.NewInt(1000000)), sdk.NewCoin("ufury", sdkmath.NewInt(100000)), sdk.NewCoin("ufury", sdkmath.NewInt(1000000)), sdk.NewCoin("musd", sdkmath.NewInt(5000000))},
		{"low ufury balance, vesting matches exact", sdk.NewCoin("musd", sdkmath.NewInt(5000000)), sdk.NewCoin("musd", sdkmath.NewInt(500000)), sdk.NewCoin("musd", sdkmath.NewInt(5000000)), sdk.NewCoin("ufury", sdkmath.NewInt(1000000))},
		{"large ufury balance difference, vesting covers difference", sdk.NewCoin("ufury", sdkmath.NewInt(100e6)), sdk.NewCoin("ufury", sdkmath.NewInt(10000e6)), sdk.NewCoin("ufury", sdkmath.NewInt(1000e6)), sdk.NewCoin("musd", sdkmath.NewInt(5000e6))},
		{"large musd balance difference, vesting covers difference", sdk.NewCoin("musd", sdkmath.NewInt(500e6)), sdk.NewCoin("musd", sdkmath.NewInt(500000e6)), sdk.NewCoin("musd", sdkmath.NewInt(5000e6)), sdk.NewCoin("ufury", sdkmath.NewInt(1000e6))},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			suite.SetupTest()
			owner := suite.CreateAccount(sdk.Coins{})
			reserves := sdk.NewCoins(
				sdk.NewCoin("ufury", sdkmath.NewInt(100000e6)),
				sdk.NewCoin("musd", sdkmath.NewInt(500000e6)),
			)
			totalShares := sdkmath.NewInt(30000e6)
			suite.setupPool(reserves, totalShares, owner.GetAddress())
			balance := sdk.NewCoins(tc.balanceA)
			vesting := sdk.NewCoins(tc.vestingA)
			requester := suite.CreateVestingAccount(balance, vesting)

			ctx := suite.App.NewContext(true, tmproto.Header{Height: 1, Time: tmtime.Now()})
			err := suite.Keeper.SwapForExactTokens(ctx, requester.GetAddress(), tc.coinA, tc.coinB, sdk.MustNewDecFromStr("0.1"))
			suite.Require().True(errors.Is(err, sdkerrors.ErrInsufficientFunds), fmt.Sprintf("got err %s", err))
		})
	}
}

func (suite *keeperTestSuite) TestSwapForExactTokens_PoolNotFound() {
	owner := suite.CreateAccount(sdk.Coins{})
	reserves := sdk.NewCoins(
		sdk.NewCoin("ufury", sdkmath.NewInt(1000e6)),
		sdk.NewCoin("musd", sdkmath.NewInt(5000e6)),
	)
	totalShares := sdkmath.NewInt(3000e6)
	poolID := suite.setupPool(reserves, totalShares, owner.GetAddress())
	suite.Keeper.DeletePool(suite.Ctx, poolID)

	balance := sdk.NewCoins(
		sdk.NewCoin("ufury", sdkmath.NewInt(10e6)),
		sdk.NewCoin("musd", sdkmath.NewInt(10e6)),
	)
	requester := suite.NewAccountFromAddr(sdk.AccAddress("requester-----------"), balance)
	coinA := sdk.NewCoin("ufury", sdkmath.NewInt(1e6))
	coinB := sdk.NewCoin("musd", sdkmath.NewInt(5e6))

	err := suite.Keeper.SwapForExactTokens(suite.Ctx, requester.GetAddress(), coinA, coinB, sdk.MustNewDecFromStr("0.01"))
	suite.EqualError(err, "pool ufury:musd not found: invalid pool")

	err = suite.Keeper.SwapForExactTokens(suite.Ctx, requester.GetAddress(), coinB, coinA, sdk.MustNewDecFromStr("0.01"))
	suite.EqualError(err, "pool ufury:musd not found: invalid pool")
}

func (suite *keeperTestSuite) TestSwapForExactTokens_PanicOnInvalidPool() {
	owner := suite.CreateAccount(sdk.Coins{})
	reserves := sdk.NewCoins(
		sdk.NewCoin("ufury", sdkmath.NewInt(1000e6)),
		sdk.NewCoin("musd", sdkmath.NewInt(5000e6)),
	)
	totalShares := sdkmath.NewInt(3000e6)
	poolID := suite.setupPool(reserves, totalShares, owner.GetAddress())

	poolRecord, found := suite.Keeper.GetPool(suite.Ctx, poolID)
	suite.Require().True(found, "expected pool record to exist")

	poolRecord.TotalShares = sdk.ZeroInt()
	suite.Keeper.SetPool_Raw(suite.Ctx, poolRecord)

	balance := sdk.NewCoins(
		sdk.NewCoin("ufury", sdkmath.NewInt(10e6)),
		sdk.NewCoin("musd", sdkmath.NewInt(10e6)),
	)
	requester := suite.NewAccountFromAddr(sdk.AccAddress("requester-----------"), balance)
	coinA := sdk.NewCoin("ufury", sdkmath.NewInt(1e6))
	coinB := sdk.NewCoin("musd", sdkmath.NewInt(5e6))

	suite.PanicsWithValue("invalid pool ufury:musd: total shares must be greater than zero: invalid pool", func() {
		_ = suite.Keeper.SwapForExactTokens(suite.Ctx, requester.GetAddress(), coinA, coinB, sdk.MustNewDecFromStr("0.01"))
	}, "expected invalid pool record to panic")

	suite.PanicsWithValue("invalid pool ufury:musd: total shares must be greater than zero: invalid pool", func() {
		_ = suite.Keeper.SwapForExactTokens(suite.Ctx, requester.GetAddress(), coinB, coinA, sdk.MustNewDecFromStr("0.01"))
	}, "expected invalid pool record to panic")
}

func (suite *keeperTestSuite) TestSwapForExactTokens_PanicOnInsufficientModuleAccFunds() {
	owner := suite.CreateAccount(sdk.Coins{})
	reserves := sdk.NewCoins(
		sdk.NewCoin("ufury", sdkmath.NewInt(1000e6)),
		sdk.NewCoin("musd", sdkmath.NewInt(5000e6)),
	)
	totalShares := sdkmath.NewInt(3000e6)
	suite.setupPool(reserves, totalShares, owner.GetAddress())

	suite.RemoveCoinsFromModule(sdk.NewCoins(
		sdk.NewCoin("ufury", sdkmath.NewInt(1000e6)),
		sdk.NewCoin("musd", sdkmath.NewInt(5000e6)),
	))

	balance := sdk.NewCoins(
		sdk.NewCoin("ufury", sdkmath.NewInt(10e6)),
		sdk.NewCoin("musd", sdkmath.NewInt(10e6)),
	)
	requester := suite.NewAccountFromAddr(sdk.AccAddress("requester-----------"), balance)
	coinA := sdk.NewCoin("ufury", sdkmath.NewInt(1e6))
	coinB := sdk.NewCoin("musd", sdkmath.NewInt(5e6))

	suite.Panics(func() {
		_ = suite.Keeper.SwapForExactTokens(suite.Ctx, requester.GetAddress(), coinA, coinB, sdk.MustNewDecFromStr("0.01"))
	}, "expected panic when module account does not have enough funds")

	suite.Panics(func() {
		_ = suite.Keeper.SwapForExactTokens(suite.Ctx, requester.GetAddress(), coinA, coinB, sdk.MustNewDecFromStr("0.01"))
	}, "expected panic when module account does not have enough funds")
}
