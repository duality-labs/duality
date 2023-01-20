package keeper

import (
	"github.com/duality-labs/duality/x/dex/types"
)

var _ types.QueryServer = Keeper{}
