//nolint
package tmdata

import (
	"github.com/cosmos/cosmos-sdk/std"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sentinel-official/hub"
	"github.com/sentinel-official/hub/params"
)

func MakeEncodingConfig() params.EncodingConfig {
	config := params.MakeEncodingConfig()
	std.RegisterLegacyAminoCodec(config.Amino)
	std.RegisterInterfaces(config.InterfaceRegistry)
	hub.ModuleBasics.RegisterLegacyAminoCodec(config.Amino)
	hub.ModuleBasics.RegisterInterfaces(config.InterfaceRegistry)
	return config
}

func DecodeTx(txBytes []byte) (sdk.Tx, error) {
	encodingConfig := MakeEncodingConfig()
	tx, err := encodingConfig.TxConfig.TxDecoder()(txBytes)
	if err != nil {
		return nil, err
	}
	return tx, nil
}
