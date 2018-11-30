package storeapp

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// MsgSetValue defines a SetValue message
type MsgSetValue struct {
	Key    string
	Value  string
	Sender sdk.AccAddress
}

// NewMsgSetName is a constructor function for MsgSetName
func NewMsgSetName(key string, value string, sender sdk.AccAddress) MsgSetValue {
	return MsgSetValue{
		Key:    key,
		Value:  value,
		Sender: sender,
	}
}

// Route Implements Msg.
func (msg MsgSetValue) Route() string { return "storeapp" }

// Type Implements Msg.
func (msg MsgSetValue) Type() string { return "set_value" }

// ValidateBasic Implements Msg.
func (msg MsgSetValue) ValidateBasic() sdk.Error {
	if msg.Sender.Empty() {
		return sdk.ErrInvalidAddress(msg.Sender.String())
	}
	if len(msg.Key) == 0 || len(msg.Value) == 0 {
		return sdk.ErrUnknownRequest("Key and/or Value cannot be empty")
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgSetValue) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners Implements Msg.
func (msg MsgSetValue) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}
