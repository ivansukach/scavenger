package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// TODO: Describe your actions, these will implment the interface of `sdk.Msg`

// verify interface at compile time
var _ sdk.Msg = &MsgCreateScavenge{}

// MsgCommitSolution - struct for unjailing jailed validator
type MsgCreateScavenge struct {
	Creator      sdk.AccAddress `json:"creator" yaml:"creator"` // address of the validator operator
	Description  string         `json:"description" yaml:"description"`
	SolutionHash string         `json:"solutionHash" yaml:"solutionHash"`
	Reward       sdk.Coins      `json:"reward" yaml:"reward"`
}

// NewMsgMsgCreate creates a new MsgMsgCreate instance
func NewMsgCreateScavenge(creator sdk.AccAddress, description, solutionHash string, reward sdk.Coins) MsgCreateScavenge {
	return MsgCreateScavenge{
		Creator:      creator,
		Description:  description,
		SolutionHash: solutionHash,
		Reward:       reward,
	}
}

const CreateScavengeConst = "CreateScavenge"

// nolint
func (msg MsgCreateScavenge) Route() string { return RouterKey }
func (msg MsgCreateScavenge) Type() string  { return CreateScavengeConst }
func (msg MsgCreateScavenge) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgCreateScavenge) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgCreateScavenge) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	if msg.SolutionHash == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "solutionHash can't be empty")
	}
	return nil
}
