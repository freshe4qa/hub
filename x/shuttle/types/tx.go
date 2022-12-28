package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	sdktx "github.com/cosmos/cosmos-sdk/types/tx"
)

var (
	_ sdk.Msg = &MsgRegisterAccount{}
	_ sdk.Msg = &MsgSendFunds{}
	_ sdk.Msg = &MsgSendMessages{}
)

//------------------------------------------------------------------------------
// MsgRegisterAccount
//------------------------------------------------------------------------------

// ValidateBasic does a sanity check on the provided data
func (m *MsgRegisterAccount) ValidateBasic() error {
	// the authority address must be valid
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return sdkerrors.Wrap(ErrInvalidProposalAuthority, err.Error())
	}

	return nil
}

// GetSigners returns the expected signers for the message
func (m *MsgRegisterAccount) GetSigners() []sdk.AccAddress {
	// we have already asserted that the authority address is valid in
	// ValidateBasic, so can ignore the error here
	addr, _ := sdk.AccAddressFromBech32(m.Authority)
	return []sdk.AccAddress{addr}
}

//------------------------------------------------------------------------------
// MsgSendFunds
//------------------------------------------------------------------------------

// ValidateBasic does a sanity check on the provided data
func (m *MsgSendFunds) ValidateBasic() error {
	// the authority address must be valid
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return sdkerrors.Wrap(ErrInvalidProposalAuthority, err.Error())
	}

	// the coins amount must be valid
	if err := m.Amount.Validate(); err != nil {
		return sdkerrors.Wrap(ErrInvalidProposalAmount, err.Error())
	}

	return nil
}

// GetSigners returns the expected signers for the message
func (m *MsgSendFunds) GetSigners() []sdk.AccAddress {
	// we have already asserted that the authority address is valid in
	// ValidateBasic, so can ignore the error here
	addr, _ := sdk.AccAddressFromBech32(m.Authority)
	return []sdk.AccAddress{addr}
}

//------------------------------------------------------------------------------
// MsgSendMessages
//------------------------------------------------------------------------------

// ValidateBasic does a sanity check on the provided data
func (m *MsgSendMessages) ValidateBasic() error {
	// the authority address must be valid
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return sdkerrors.Wrap(ErrInvalidProposalAuthority, err.Error())
	}

	// the messages must each implement the sdk.Msg interface
	msgs, err := sdktx.GetMsgs(m.Messages, sdk.MsgTypeURL(m))
	if err != nil {
		return sdkerrors.Wrap(ErrInvalidProposalMsg, err.Error())
	}

	// all messages must be valid
	for _, msg := range msgs {
		if err = msg.ValidateBasic(); err != nil {
			return sdkerrors.Wrap(ErrInvalidProposalMsg, err.Error())
		}

		// TODO: should be check the messages' signers?
	}

	return nil
}

// GetSigners returns the expected signers for the message
func (m *MsgSendMessages) GetSigners() []sdk.AccAddress {
	// we have already asserted that the authority address is valid in
	// ValidateBasic, so can ignore the error here
	addr, _ := sdk.AccAddressFromBech32(m.Authority)
	return []sdk.AccAddress{addr}
}
