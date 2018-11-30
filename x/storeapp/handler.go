package storeapp

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewHandler returns a handler for "nameservice" type messages.
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case MsgSetValue:
			return handleMsgSetValue(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized nameservice Msg type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

// Handle MsgSetValue
func handleMsgSetValue(ctx sdk.Context, keeper Keeper, msg MsgSetValue) sdk.Result {
	keeper.SetValue(ctx, msg.Key, msg.Value) // If so, set the name to the value specified in the msg.
	return sdk.Result{}                      // return
}
