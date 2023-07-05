package types

import (
	"fmt"
	"time"
)

var (
	DefaultMUSDClaims         = MUSDMintingClaims{}
	DefaultJinxClaims         = JinxLiquidityProviderClaims{}
	DefaultDelegatorClaims    = DelegatorClaims{}
	DefaultSwapClaims         = SwapClaims{}
	DefaultSavingsClaims      = SavingsClaims{}
	DefaultGenesisRewardState = NewGenesisRewardState(
		AccumulationTimes{},
		MultiRewardIndexes{},
	)
	DefaultEarnClaims = EarnClaims{}
)

// NewGenesisState returns a new genesis state
func NewGenesisState(
	params Params,
	musdState, jinxSupplyState, jinxBorrowState, delegatorState, swapState, savingsState, earnState GenesisRewardState,
	c MUSDMintingClaims, hc JinxLiquidityProviderClaims, dc DelegatorClaims, sc SwapClaims, savingsc SavingsClaims,
	earnc EarnClaims,
) GenesisState {
	return GenesisState{
		Params: params,

		MUSDRewardState:       musdState,
		JinxSupplyRewardState: jinxSupplyState,
		JinxBorrowRewardState: jinxBorrowState,
		DelegatorRewardState:  delegatorState,
		SwapRewardState:       swapState,
		SavingsRewardState:    savingsState,
		EarnRewardState:       earnState,

		MUSDMintingClaims:           c,
		JinxLiquidityProviderClaims: hc,
		DelegatorClaims:             dc,
		SwapClaims:                  sc,
		SavingsClaims:               savingsc,
		EarnClaims:                  earnc,
	}
}

// DefaultGenesisState returns a default genesis state
func DefaultGenesisState() GenesisState {
	return GenesisState{
		Params:                      DefaultParams(),
		MUSDRewardState:             DefaultGenesisRewardState,
		JinxSupplyRewardState:       DefaultGenesisRewardState,
		JinxBorrowRewardState:       DefaultGenesisRewardState,
		DelegatorRewardState:        DefaultGenesisRewardState,
		SwapRewardState:             DefaultGenesisRewardState,
		SavingsRewardState:          DefaultGenesisRewardState,
		EarnRewardState:             DefaultGenesisRewardState,
		MUSDMintingClaims:           DefaultMUSDClaims,
		JinxLiquidityProviderClaims: DefaultJinxClaims,
		DelegatorClaims:             DefaultDelegatorClaims,
		SwapClaims:                  DefaultSwapClaims,
		SavingsClaims:               DefaultSavingsClaims,
		EarnClaims:                  DefaultEarnClaims,
	}
}

// Validate performs basic validation of genesis data returning an
// error for any failed validation criteria.
func (gs GenesisState) Validate() error {
	if err := gs.Params.Validate(); err != nil {
		return err
	}

	if err := gs.MUSDRewardState.Validate(); err != nil {
		return err
	}
	if err := gs.JinxSupplyRewardState.Validate(); err != nil {
		return err
	}
	if err := gs.JinxBorrowRewardState.Validate(); err != nil {
		return err
	}
	if err := gs.DelegatorRewardState.Validate(); err != nil {
		return err
	}
	if err := gs.SwapRewardState.Validate(); err != nil {
		return err
	}
	if err := gs.SavingsRewardState.Validate(); err != nil {
		return err
	}
	if err := gs.EarnRewardState.Validate(); err != nil {
		return err
	}

	if err := gs.MUSDMintingClaims.Validate(); err != nil {
		return err
	}
	if err := gs.JinxLiquidityProviderClaims.Validate(); err != nil {
		return err
	}
	if err := gs.DelegatorClaims.Validate(); err != nil {
		return err
	}
	if err := gs.SwapClaims.Validate(); err != nil {
		return err
	}

	if err := gs.SavingsClaims.Validate(); err != nil {
		return err
	}

	return gs.EarnClaims.Validate()
}

// NewGenesisRewardState returns a new GenesisRewardState
func NewGenesisRewardState(accumTimes AccumulationTimes, indexes MultiRewardIndexes) GenesisRewardState {
	return GenesisRewardState{
		AccumulationTimes:  accumTimes,
		MultiRewardIndexes: indexes,
	}
}

// Validate performs validation of a GenesisRewardState
func (grs GenesisRewardState) Validate() error {
	if err := grs.AccumulationTimes.Validate(); err != nil {
		return err
	}
	return grs.MultiRewardIndexes.Validate()
}

// NewAccumulationTime returns a new GenesisAccumulationTime
func NewAccumulationTime(ctype string, prevTime time.Time) AccumulationTime {
	return AccumulationTime{
		CollateralType:           ctype,
		PreviousAccumulationTime: prevTime,
	}
}

// Validate performs validation of GenesisAccumulationTime
func (gat AccumulationTime) Validate() error {
	if len(gat.CollateralType) == 0 {
		return fmt.Errorf("genesis accumulation time's collateral type must be defined")
	}
	return nil
}

// AccumulationTimes slice of GenesisAccumulationTime
type AccumulationTimes []AccumulationTime

// Validate performs validation of GenesisAccumulationTimes
func (gats AccumulationTimes) Validate() error {
	for _, gat := range gats {
		if err := gat.Validate(); err != nil {
			return err
		}
	}
	return nil
}
