package keeper

import (
	"github.com/duality-labs/duality/x/cwhooks/types"
)

var _ types.QueryServer = Keeper{}
