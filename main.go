package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/notional-labs/test/tmdata"
)

func readHexStringFileOfParts(fileDir string) []string {
	file, err := os.Open(fileDir)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	hexStrings := []string{}
	for {
		hexString, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		hexString = hexString[:len(hexString)-1]
		// hexString = strings.ReplaceAll(hexString, "\n", "")
		// fmt.Println(strings.Contains(hexString, "\n"))
		hexStrings = append(hexStrings, hexString)
	}
	return hexStrings
}

func main() {
	hexStrings := readHexStringFileOfParts("evil_blocks/block0")
	ps, err := tmdata.GetPartSetFromHexStrings(hexStrings)
	if err != nil {
		fmt.Println("err")
	}
	block, err := tmdata.GetBlockFromPartSet(ps)
	if err != nil {
		panic(err)
	}
	// fmt.Printf("%+v\n", block)
	// for _, tx := range block.Data.Txs {
	// 	fmt.Println(tx)
	// }
	// tx, err := tmdata.DecodeTx(block.Data.Txs[0])
	// if err == nil {
	// 	fmt.Printf("%+v\n", tx)
	// } else {
	// 	fmt.Println(err, 12)
	// }

	// txBz, _ := hex.DecodeString(hexString)
	// fmt.Println(hex.EncodeToString(block.Data.Txs[0].Hash()))

	// tx, err := tmdata.DecodeTx(block.Data.Txs[0])
	length := 0
	for _, i := range block.Data.Txs {
		length += len(i)
	}
	fmt.Println(length)
	fmt.Println(int(999999999999999999))
	time.Sleep(1)
	fmt.Println(999)
}
