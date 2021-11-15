package tmdata

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jaekwon/testify/require"
)

func RegisterTestCodec(cdc *codec.LegacyAmino) {
	// register Tx, Msg
	sdk.RegisterLegacyAminoCodec(cdc)

	// register test types
	cdc.RegisterConcrete(&txTest{}, "cosmos-sdk/baseapp/txTest", nil)
	cdc.RegisterConcrete(&msgCounter{}, "cosmos-sdk/baseapp/msgCounter", nil)
	cdc.RegisterConcrete(&msgCounter2{}, "cosmos-sdk/baseapp/msgCounter2", nil)
	cdc.RegisterConcrete(&msgKeyValue{}, "cosmos-sdk/baseapp/msgKeyValue", nil)
	cdc.RegisterConcrete(&msgNoRoute{}, "cosmos-sdk/baseapp/msgNoRoute", nil)
}

func TestTxDecoder(t *testing.T) {
	codec := codec.NewLegacyAmino()
	RegisterTestCodec(codec)

	app := newBaseApp(t.Name())
	tx := NewTxCounter(1, 0)
	txBytes := codec.MustMarshalBinaryBare(tx)

	dTx, err := app.txDecoder(txBytes)
	require.NoError(t, err)

	cTx := dTx.(txTest)
	require.Equal(t, tx.Counter, cTx.Counter)
}
