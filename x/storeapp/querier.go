package storeapp

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// query endpoints supported by the governance Querier
const (
	QueryValue = "get-value"
)

// NewQuerier is the module level router for state queries
func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryValue:
			return queryValue(ctx, path[1:], req, keeper)
		default:
			return nil, sdk.ErrUnknownRequest("unknown nameservice query endpoint")
		}
	}
}

func queryValue(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) (res []byte, err sdk.Error) {
	key := path[0]
	value := keeper.GetValue(ctx, key)
	if value == "" {
		return []byte{}, sdk.ErrUnknownRequest("could not get value")
	}
	return []byte(value), nil
}
