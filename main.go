package main

import (
	"bufio"
	"encoding/base64"
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
	tx := "CqEBCpwBCjcvY29zbW9zLmRpc3RyaWJ1dGlvbi52MWJldGExLk1zZ1dpdGhkcmF3RGVsZWdhdG9yUmV3YXJkEmEKK2p1bm8xcGx2NHE2OGE5ZHVnNGxsczU2Znk3NHB6czZ5bTN5dmN1bjY2dG4SMmp1bm92YWxvcGVyMTN2NHNwc2FoODVwczR2dHJ3MDd2emVhMzdncTVsYTVna3Rsa2V1EgASZgpQCkYKHy9jb3Ntb3MuY3J5cHRvLnNlY3AyNTZrMS5QdWJLZXkSIwohA0YeP6kAW7305XbCN9ZA1hY25A46uBEs0gKP5BqCaT9ZEgQKAgh/GA0SEgoMCgV1anVubxIDMzUwEODFCBpASBNuIMKnD8BBHoVuDELkbTorc2tLb0yBImUqQadS021SXYeyefuRYVMyvFlddwGC1lA5zbOd+TXHCvOoBOjtDg=="

	tx_byte, err := base64.StdEncoding.DecodeString(tx)

	if err != nil {
		panic(err)
	}

	sdk_tx, err := tmdata.DecodeTx(tx_byte)
	if err != nil {
		panic(err)
	}
	for _, i := range sdk_tx.GetMsgs() {
		fmt.Printf("%+v", i)
	}

}
