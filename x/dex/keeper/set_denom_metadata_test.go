package keeper_test

import (
	"github.com/NicholasDotSol/duality/x/dex/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

var denomMetadata banktypes.Metadata = banktypes.Metadata{
	Description: "A Test Coin",
	Display:     "Test",
	Base:        "uTest",
	Name:        "Test Coin",
	Symbol:      "TES",
	DenomUnits: []*banktypes.DenomUnit{
		{
			Denom:    "uTest",
			Exponent: 0,
		},
		{
			Denom:    "Test",
			Exponent: 18,
		},
	},
}

func (s *MsgServerTestSuite) TestSetDenomMetadata() {

	_, err := s.msgServer.SetDenomMetadata(s.goCtx,
		&types.MsgSetDenomMetadata{
			Creator:  s.alice.String(),
			Metadata: denomMetadata,
		},
	)

	s.Assert().Nil(err)

	denom, found := s.app.BankKeeper.GetDenomMetaData(s.ctx, "uTest")

	s.Assert().True(found)
	s.Assert().Equal(denomMetadata, denom)
}

func (s *MsgServerTestSuite) TestSetDenomMetadataRepeatFails() {

	_, err := s.msgServer.SetDenomMetadata(s.goCtx,
		&types.MsgSetDenomMetadata{
			Creator:  s.alice.String(),
			Metadata: denomMetadata,
		},
	)
	s.Assert().Nil(err)

	denomMetadata2 := banktypes.Metadata{
		Description: "A Copy Cat Coin",
		Display:     "uTest",
		Base:        "uTest",
		Name:        "Fake Coin",
		Symbol:      "TES",
		DenomUnits: []*banktypes.DenomUnit{
			{
				Denom:    "uTest",
				Exponent: 0,
			},
		},
	}

	_, err = s.msgServer.SetDenomMetadata(s.goCtx,
		&types.MsgSetDenomMetadata{
			Creator:  s.alice.String(),
			Metadata: denomMetadata2,
		},
	)

	s.Assert().ErrorIs(err, types.ErrDenomAlreadyExists)
}

func (s *MsgServerTestSuite) TestSetDenomMetadataValidation() {

	invalidDenomMetadata := banktypes.Metadata{
		Description: "Baseless Denom",
		Display:     "uTest",
		Name:        "Fake Coin",
		Symbol:      "TES",
		DenomUnits: []*banktypes.DenomUnit{
			{
				Denom:    "uTest",
				Exponent: 0,
			},
		},
	}

	_, err := s.msgServer.SetDenomMetadata(s.goCtx,
		&types.MsgSetDenomMetadata{
			Creator:  s.alice.String(),
			Metadata: invalidDenomMetadata,
		},
	)

	s.Assert().Error(err)
}
