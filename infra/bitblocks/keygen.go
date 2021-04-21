package main

import (
	"fmt"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcutil"

	"github.com/ReniR256/id"
	"github.com/ReniR256/multichain/chain/bitblocks"
)

func main() {
	privKey := id.NewPrivKey()
	wif, err := btcutil.NewWIF((*btcec.PrivateKey)(privKey), &bitblocks.RegressionNetParams, true)
	if err != nil {
		panic(err)
	}
	addrPubKeyHash, err := btcutil.NewAddressPubKeyHash(btcutil.Hash160(wif.SerializePubKey()), &bitblocks.RegressionNetParams)
	if err != nil {
		panic(err)
	}
	fmt.Printf("BITBLOCKS_PK=%v\n", wif)
	fmt.Printf("BITBLOCKS_ADDRESS=%v\n", addrPubKeyHash)
}
