package keeper

import (
	"github.com/NicholasDotSol/duality/x/mevdummy/types"
)

var _ types.QueryServer = Keeper{}
