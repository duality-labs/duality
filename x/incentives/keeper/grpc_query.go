package keeper

import (
	"github.com/duality-labs/duality/x/incentives/types"
)

var _ types.QueryServer = Keeper{}
