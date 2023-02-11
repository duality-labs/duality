//go:build !test_amino
// +build !test_amino

package params

// "github.com/tendermint/spm/cosmoscmd"

// MakeTestEncodingConfig creates an EncodingConfig for a non-amino based test configuration.
// This function should be used only internally (in the SDK).
// App user should'nt create new codecs - use the app.AppCodec instead.
// [DEPRECATED]
// func MakeTestEncodingConfig() cosmoscmd.EncodingConfig {
// 	cdc := codec.NewLegacyAmino()
// 	interfaceRegistry := types.NewInterfaceRegistry()
// 	marshaler := codec.NewProtoCodec(interfaceRegistry)

// 	return cosmoscmd.EncodingConfig{
// 		InterfaceRegistry: interfaceRegistry,
// 		Marshaler:         marshaler,
// 		TxConfig:          tx.NewTxConfig(marshaler, tx.DefaultSignModes),
// 		Amino:             cdc,
// 	}
// }
