package storeapp

import (
	"github.com/cosmos/cosmos-sdk/codec"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Keeper maintains the link to data storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	valueStoreKey sdk.StoreKey // Unexposed key to access name store from sdk.Context
	cdc           *codec.Codec // The wire codec for binary encoding/decoding.
}

// NewKeeper creates new instances of the nameservice Keeper
func NewKeeper(valueStoreKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{
		valueStoreKey: valueStoreKey,
		cdc:           cdc,
	}
}

// SetValue sets value for key
func (keeper Keeper) SetValue(ctx sdk.Context, key string, value string) {
	store := ctx.KVStore(keeper.valueStoreKey)
	store.Set([]byte(key), []byte(value))
}

// GetValue get value from key
func (keeper Keeper) GetValue(ctx sdk.Context, key string) string {
	store := ctx.KVStore(keeper.valueStoreKey)
	bz := store.Get([]byte(key))
	return string(bz)
}
