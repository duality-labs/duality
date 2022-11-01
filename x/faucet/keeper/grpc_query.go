package keeper

import (
	"github.com/NicholasDotSol/duality/x/faucet/types"
)

var _ types.QueryServer = Keeper{}
