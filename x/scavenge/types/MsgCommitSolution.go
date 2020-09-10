package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// TODO: Describe your actions, these will implment the interface of `sdk.Msg`

// verify interface at compile time
var _ sdk.Msg = &MsgCommitSolution{}

// MsgCommitSolution - struct for unjailing jailed validator
type MsgCommitSolution struct {
	Scavenger             sdk.AccAddress `json:"scavenger" yaml:"scavenger"` // address of the validator operator
	SolutionHash          string         `json:"solutionHash" yaml:"solutionHash"`
	SolutionScavengerHash string         `json:"solutionScavengerHash" yaml:"solutionScavengerHash"`
}

// NewMsgMsgCreate creates a new MsgMsgCreate instance
func NewMsgCommitSolution(scavenger sdk.AccAddress, solutionHash, solutionScavengerHash string) MsgCommitSolution {
	return MsgCommitSolution{
		Scavenger:             scavenger,
		SolutionHash:          solutionHash,
		SolutionScavengerHash: solutionScavengerHash,
	}
}

const CommitSolutionConst = "CommitSolutionConst"

// nolint
func (msg MsgCommitSolution) Route() string { return RouterKey }
func (msg MsgCommitSolution) Type() string  { return CommitSolutionConst }
func (msg MsgCommitSolution) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Scavenger)}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgCommitSolution) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgCommitSolution) ValidateBasic() error {
	if msg.Scavenger.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "scavenger can't be empty")
	}
	if msg.SolutionHash == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "solutionHash can't be empty")
	}
	if msg.SolutionScavengerHash == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "solutionScavengerHash can't be empty")
	}
	return nil
}
