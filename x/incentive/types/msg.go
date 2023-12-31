package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/migrations/legacytx"
)

const MaxDenomsToClaim = 1000

// ensure Msg interface compliance at compile time
var (
	_ sdk.Msg = &MsgClaimMUSDMintingReward{}
	_ sdk.Msg = &MsgClaimJinxReward{}
	_ sdk.Msg = &MsgClaimDelegatorReward{}
	_ sdk.Msg = &MsgClaimSwapReward{}
	_ sdk.Msg = &MsgClaimSavingsReward{}
	_ sdk.Msg = &MsgClaimEarnReward{}

	_ legacytx.LegacyMsg = &MsgClaimMUSDMintingReward{}
	_ legacytx.LegacyMsg = &MsgClaimJinxReward{}
	_ legacytx.LegacyMsg = &MsgClaimDelegatorReward{}
	_ legacytx.LegacyMsg = &MsgClaimSwapReward{}
	_ legacytx.LegacyMsg = &MsgClaimSavingsReward{}
	_ legacytx.LegacyMsg = &MsgClaimEarnReward{}
)

const (
	TypeMsgClaimMUSDMintingReward = "claim_musd_minting_reward"
	TypeMsgClaimJinxReward        = "claim_jinx_reward"
	TypeMsgClaimDelegatorReward   = "claim_delegator_reward"
	TypeMsgClaimSwapReward        = "claim_swap_reward"
	TypeMsgClaimSavingsReward     = "claim_savings_reward"
	TypeMsgClaimEarnReward        = "claim_earn_reward"
)

// NewMsgClaimMUSDMintingReward returns a new MsgClaimMUSDMintingReward.
func NewMsgClaimMUSDMintingReward(sender string, multiplierName string) MsgClaimMUSDMintingReward {
	return MsgClaimMUSDMintingReward{
		Sender:         sender,
		MultiplierName: multiplierName,
	}
}

// Route return the message type used for routing the message.
func (msg MsgClaimMUSDMintingReward) Route() string { return RouterKey }

// Type returns a human-readable string for the message, intended for utilization within tags.
func (msg MsgClaimMUSDMintingReward) Type() string { return TypeMsgClaimMUSDMintingReward }

// ValidateBasic does a simple validation check that doesn't require access to state.
func (msg MsgClaimMUSDMintingReward) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "sender address cannot be empty or invalid")
	}
	if msg.MultiplierName == "" {
		return errorsmod.Wrap(ErrInvalidMultiplier, "multiplier name cannot be empty")
	}
	return nil
}

// GetSignBytes gets the canonical byte representation of the Msg.
func (msg MsgClaimMUSDMintingReward) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgClaimMUSDMintingReward) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

// NewMsgClaimJinxReward returns a new MsgClaimJinxReward.
func NewMsgClaimJinxReward(sender string, denomsToClaim Selections) MsgClaimJinxReward {
	return MsgClaimJinxReward{
		Sender:        sender,
		DenomsToClaim: denomsToClaim,
	}
}

// Route return the message type used for routing the message.
func (msg MsgClaimJinxReward) Route() string { return RouterKey }

// Type returns a human-readable string for the message, intended for utilization within tags.
func (msg MsgClaimJinxReward) Type() string {
	return TypeMsgClaimJinxReward
}

// ValidateBasic does a simple validation check that doesn't require access to state.
func (msg MsgClaimJinxReward) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "sender address cannot be empty or invalid")
	}
	if err := msg.DenomsToClaim.Validate(); err != nil {
		return err
	}
	return nil
}

// GetSignBytes gets the canonical byte representation of the Msg.
func (msg MsgClaimJinxReward) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgClaimJinxReward) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

// NewMsgClaimDelegatorReward returns a new MsgClaimDelegatorReward.
func NewMsgClaimDelegatorReward(sender string, denomsToClaim Selections) MsgClaimDelegatorReward {
	return MsgClaimDelegatorReward{
		Sender:        sender,
		DenomsToClaim: denomsToClaim,
	}
}

// Route return the message type used for routing the message.
func (msg MsgClaimDelegatorReward) Route() string { return RouterKey }

// Type returns a human-readable string for the message, intended for utilization within tags.
func (msg MsgClaimDelegatorReward) Type() string {
	return TypeMsgClaimDelegatorReward
}

// ValidateBasic does a simple validation check that doesn't require access to state.
func (msg MsgClaimDelegatorReward) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "sender address cannot be empty or invalid")
	}
	if err := msg.DenomsToClaim.Validate(); err != nil {
		return err
	}
	return nil
}

// GetSignBytes gets the canonical byte representation of the Msg.
func (msg MsgClaimDelegatorReward) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgClaimDelegatorReward) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

// NewMsgClaimSwapReward returns a new MsgClaimSwapReward.
func NewMsgClaimSwapReward(sender string, denomsToClaim Selections) MsgClaimSwapReward {
	return MsgClaimSwapReward{
		Sender:        sender,
		DenomsToClaim: denomsToClaim,
	}
}

// Route return the message type used for routing the message.
func (msg MsgClaimSwapReward) Route() string { return RouterKey }

// Type returns a human-readable string for the message, intended for utilization within tags.
func (msg MsgClaimSwapReward) Type() string {
	return TypeMsgClaimSwapReward
}

// ValidateBasic does a simple validation check that doesn't require access to state.
func (msg MsgClaimSwapReward) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "sender address cannot be empty or invalid")
	}
	if err := msg.DenomsToClaim.Validate(); err != nil {
		return err
	}
	return nil
}

// GetSignBytes gets the canonical byte representation of the Msg.
func (msg MsgClaimSwapReward) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgClaimSwapReward) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

// NewMsgClaimSavingsReward returns a new MsgClaimSavingsReward.
func NewMsgClaimSavingsReward(sender string, denomsToClaim Selections) MsgClaimSavingsReward {
	return MsgClaimSavingsReward{
		Sender:        sender,
		DenomsToClaim: denomsToClaim,
	}
}

// Route return the message type used for routing the message.
func (msg MsgClaimSavingsReward) Route() string { return RouterKey }

// Type returns a human-readable string for the message, intended for utilization within tags.
func (msg MsgClaimSavingsReward) Type() string {
	return TypeMsgClaimSavingsReward
}

// ValidateBasic does a simple validation check that doesn't require access to state.
func (msg MsgClaimSavingsReward) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "sender address cannot be empty or invalid")
	}
	if err := msg.DenomsToClaim.Validate(); err != nil {
		return err
	}
	return nil
}

// GetSignBytes gets the canonical byte representation of the Msg.
func (msg MsgClaimSavingsReward) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgClaimSavingsReward) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

// NewMsgClaimEarnReward returns a new MsgClaimEarnReward.
func NewMsgClaimEarnReward(sender string, denomsToClaim Selections) MsgClaimEarnReward {
	return MsgClaimEarnReward{
		Sender:        sender,
		DenomsToClaim: denomsToClaim,
	}
}

// Route return the message type used for routing the message.
func (msg MsgClaimEarnReward) Route() string { return RouterKey }

// Type returns a human-readable string for the message, intended for utilization within tags.
func (msg MsgClaimEarnReward) Type() string {
	return TypeMsgClaimEarnReward
}

// ValidateBasic does a simple validation check that doesn't require access to state.
func (msg MsgClaimEarnReward) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "sender address cannot be empty or invalid")
	}
	if err := msg.DenomsToClaim.Validate(); err != nil {
		return err
	}
	return nil
}

// GetSignBytes gets the canonical byte representation of the Msg.
func (msg MsgClaimEarnReward) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners returns the addresses of signers that must sign.
func (msg MsgClaimEarnReward) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}
