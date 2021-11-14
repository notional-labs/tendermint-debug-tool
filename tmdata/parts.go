package tmdata

import (
	"encoding/hex"
	"io/ioutil"

	"github.com/gogo/protobuf/proto"
	tmbytes "github.com/tendermint/tendermint/libs/bytes"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/tendermint/tendermint/types"
)

func GetPartSetFromHexStrings(hexStrings []string) (*types.PartSet, error) {
	partSet := &types.PartSet{}
	lenHexStrings := len(hexStrings)
	partSet.SetPartsLength(lenHexStrings)
	partSet.SetTotal(uint32(lenHexStrings))
	for id, hexString := range hexStrings {
		bz, _ := hex.DecodeString(hexString)
		tmbytes := tmbytes.HexBytes(bz)

		part := &types.Part{
			Index: uint32(id),
			Bytes: tmbytes,
		}
		_, err := partSet.AddPart(part)
		if err != nil {
			return nil, err
		}
	}
	return partSet, nil
}

func GetBlockFromPartSet(partSet *types.PartSet) (*types.Block, error) {
	bz, err := ioutil.ReadAll(partSet.GetReader())
	if err != nil {
		return nil, err
	}

	var pbb = new(tmproto.Block)
	err = proto.Unmarshal(bz, pbb)
	if err != nil {
		return nil, err
	}

	block, err := types.BlockFromProto(pbb)
	if err != nil {
		return nil, err
	}
	return block, nil
}
