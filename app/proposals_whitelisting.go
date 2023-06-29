//nolint:gochecknoglobals
package app

import (
	govv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	"github.com/cosmos/cosmos-sdk/x/params/types/proposal"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
)

func IsProposalWhitelisted(content govv1beta1.Content) bool {
	switch c := content.(type) {
	case *proposal.ParameterChangeProposal:
		return isParamChangeWhitelisted(c.Changes)
	case *upgradetypes.SoftwareUpgradeProposal,
		*upgradetypes.CancelSoftwareUpgradeProposal:
		return true

	default:
		return false
	}
}

func isParamChangeWhitelisted(paramChanges []proposal.ParamChange) bool {
	for _, paramChange := range paramChanges {
		_, found := WhitelistedParams[paramChangeKey{Subspace: paramChange.Subspace, Key: paramChange.Key}]
		if !found {
			return false
		}
	}

	return true
}

type paramChangeKey struct {
	Subspace, Key string
}

var WhitelistedParams = map[paramChangeKey]struct{}{
	// bank
	// TODO: confirm this can be removed? Send enabled for bank is deprecated
	// {Subspace: banktypes.ModuleName, Key: "SendEnabled"}: {},
	// ibc transfer
	{Subspace: ibctransfertypes.ModuleName, Key: "SendEnabled"}:    {},
	{Subspace: ibctransfertypes.ModuleName, Key: "ReceiveEnabled"}: {},
}
