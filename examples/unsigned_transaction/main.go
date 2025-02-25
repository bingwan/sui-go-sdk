package main

import (
	"context"
	"fmt"

	"github.com/block-vision/sui-go-sdk/models"
	"github.com/block-vision/sui-go-sdk/sui"
)

//send unsigned transactions
func main() {
	cli := sui.NewSuiClient("https://gateway.devnet.sui.io:443")

	resp, err := cli.MoveCall(context.Background(), models.MoveCallRequest{
		Signer:          "0x4d6f1a54e805038f44ecd3112927af147e9b9ecb",
		PackageObjectId: "0x0000000000000000000000000000000000000002",
		Module:          "devnet_nft",
		Function:        "mint",
		TypeArguments:   []interface{}{},
		Arguments:       []interface{}{"blockvision", "blockvision", "testurl"},
		Gas:             "0x43c233ae92c79556b2642464ec5e16f506e9045a",
		GasBudget:       1000,
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(resp)

}
