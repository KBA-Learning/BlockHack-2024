package main

import (
	"context"
	"fmt"
	"snippets/go/helpers"
	"snippets/go/lib"

	"github.com/ethereum/go-ethereum/common"
)

func main() {
	client, err := helpers.DialClient()
	if err != nil {
		panic(err)
	}

	instance, err := lib.NewStorage(common.HexToAddress("0x5FbDB2315678afecb367f032d93F642f64180aa3"), client)
	if err != nil {
		panic(err)
	}

	auth, err := helpers.AuthGenerator(client, "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Address: \033[32m%s\033[0m\n", auth.From)

	balance, err := client.BalanceAt(context.Background(), auth.From, nil)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Balance: \033[32m%s\033[0m\n", balance)

	latest, err := client.BlockNumber(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Printf("Latest Block: \033[32m%d\033[0m\n", latest)

	trx, err := instance.Store(auth, "101", "Hello from Go 🐹!")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Transaction Hash: \033[32m%s\033[0m\n", trx.Hash())

	value, err := instance.Collection(nil, "101")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Value: \033[32m%s\033[0m\n", value)
}
