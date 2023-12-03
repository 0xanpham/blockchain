package main

import (
	"fmt"
	"log"

	"github.com/0xanpham/blockchain/wallet"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	w := wallet.NewWallet()
	fmt.Println(w.PublicKeyStr())
	fmt.Println(w.PrivateKeyStr())
	fmt.Println(w.BlockchainAddress())
}
