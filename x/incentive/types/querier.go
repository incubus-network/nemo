package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Querier routes for the incentive module
const (
	QueryGetHardRewards        = "hard-rewards"
	QueryGetMUSDMintingRewards = "musd-minting-rewards"
	QueryGetDelegatorRewards   = "delegator-rewards"
	QueryGetSwapRewards        = "swap-rewards"
	QueryGetSavingsRewards     = "savings-rewards"
	QueryGetEarnRewards        = "earn-rewards"
	QueryGetRewardFactors      = "reward-factors"
	QueryGetParams             = "parameters"
	QueryGetAPYs               = "apys"

	RestClaimCollateralType = "collateral_type"
	RestClaimOwner          = "owner"
	RestClaimType           = "type"
	RestUnsynced            = "unsynced"
)

// QueryRewardsParams params for query /incentive/rewards/<claim type>
type QueryRewardsParams struct {
	Page           int            `json:"page" yaml:"page"`
	Limit          int            `json:"limit" yaml:"limit"`
	Owner          sdk.AccAddress `json:"owner" yaml:"owner"`
	Unsynchronized bool           `json:"unsynchronized" yaml:"unsynchronized"`
}

// NewQueryRewardsParams returns QueryRewardsParams
func NewQueryRewardsParams(page, limit int, owner sdk.AccAddress, unsynchronized bool) QueryRewardsParams {
	return QueryRewardsParams{
		Page:           page,
		Limit:          limit,
		Owner:          owner,
		Unsynchronized: unsynchronized,
	}
}

// QueryGetRewardFactorsResponse holds the response to a reward factor query
type QueryGetRewardFactorsResponse struct {
	MUSDMintingRewardFactors RewardIndexes      `json:"musd_minting_reward_factors" yaml:"musd_minting_reward_factors"`
	HardSupplyRewardFactors  MultiRewardIndexes `json:"hard_supply_reward_factors" yaml:"hard_supply_reward_factors"`
	HardBorrowRewardFactors  MultiRewardIndexes `json:"hard_borrow_reward_factors" yaml:"hard_borrow_reward_factors"`
	DelegatorRewardFactors   MultiRewardIndexes `json:"delegator_reward_factors" yaml:"delegator_reward_factors"`
	SwapRewardFactors        MultiRewardIndexes `json:"swap_reward_factors" yaml:"swap_reward_factors"`
	SavingsRewardFactors     MultiRewardIndexes `json:"savings_reward_factors" yaml:"savings_reward_factors"`
	EarnRewardFactors        MultiRewardIndexes `json:"earn_reward_factors" yaml:"earn_reward_factors"`
}

// NewQueryGetRewardFactorsResponse returns a new instance of QueryAllRewardFactorsResponse
func NewQueryGetRewardFactorsResponse(musdMintingFactors RewardIndexes, supplyFactors,
	hardBorrowFactors, delegatorFactors, swapFactors, savingsFactors, earnFactors MultiRewardIndexes,
) QueryGetRewardFactorsResponse {
	return QueryGetRewardFactorsResponse{
		MUSDMintingRewardFactors: musdMintingFactors,
		HardSupplyRewardFactors:  supplyFactors,
		HardBorrowRewardFactors:  hardBorrowFactors,
		DelegatorRewardFactors:   delegatorFactors,
		SwapRewardFactors:        swapFactors,
		SavingsRewardFactors:     savingsFactors,
		EarnRewardFactors:        earnFactors,
	}
}

// QueryGetAPYsResponse holds the response to a APY query
type QueryGetAPYsResponse struct {
	Earn []Apy `json:"earn" yaml:"earn"`
}

// NewQueryGetAPYsResponse returns a new instance of QueryGetAPYsResponse
func NewQueryGetAPYsResponse(earn []Apy) QueryGetAPYsResponse {
	return QueryGetAPYsResponse{
		Earn: earn,
	}
}
