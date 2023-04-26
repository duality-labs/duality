package keeper

import (
	"github.com/duality-labs/duality/x/mev/types"
)

var _ types.QueryServer = Keeper{}
