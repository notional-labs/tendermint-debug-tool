package main

import (
	"bufio"
	"fmt"
	"os"

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
	homeDir, _ := os.UserHomeDir()
	hexStrings := readHexStringFileOfParts(homeDir + "/block0")
	ps, err := tmdata.GetPartSetFromHexStrings(hexStrings)
	if err != nil {
		fmt.Println("err")
	}
	block, err := tmdata.GetBlockFromPartSet(ps)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", block)
	// for _, tx := range block.Data.Txs {
	// 	fmt.Println(tx)
	// }
}
