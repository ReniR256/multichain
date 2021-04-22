package bitblocks

import (
	"github.com/btcsuite/btcd/chaincfg"
)

func init() {
	if err := chaincfg.Register(&MainNetParams); err != nil {
		panic(err)
	}
	if err := chaincfg.Register(&TestNetParams); err != nil {
		panic(err)
	}
	if err := chaincfg.Register(&RegressionNetParams); err != nil {
		panic(err)
	}
}

// MainNetParams returns the chain configuration for mainnet.
var MainNetParams = chaincfg.Params{
	Name: "mainnet",
	Net:  0xc0c0c0c0,

	// Address encoding magics
	PubKeyHashAddrID: 25,
	ScriptHashAddrID: 5,
	PrivateKeyID:     153,

	// BIP32 hierarchical deterministic extended key magics
	HDPrivateKeyID: [4]byte{0x0, 0x0, 0x0, 0x0}, // not compatible
	HDPublicKeyID:  [4]byte{0x0, 0x0, 0x0, 0x0}, // not compatible

	// Human-readable part for Bech32 encoded segwit addresses, as defined in
	// BIP 173. Bitblocks does not actually support this, but we do not want to
	// collide with real addresses, so we specify it.
	Bech32HRPSegwit: "xbb",
}

// TestNetParams returns the chain configuration for testnet.
var TestNetParams = chaincfg.Params{
	Name: "testnet",
	Net:  0xfcc1b7dc,

	// Address encoding magics
	PubKeyHashAddrID: 113,
	ScriptHashAddrID: 196,
	PrivateKeyID:     241,

	// BIP32 hierarchical deterministic extended key magics
	HDPrivateKeyID: [4]byte{0x0, 0x0, 0x0, 0x0}, // not compatible
	HDPublicKeyID:  [4]byte{0x0, 0x0, 0x0, 0x0}, // not compatible

	// Human-readable part for Bech32 encoded segwit addresses, as defined in
	// BIP 173. Bitblocks does not actually support this, but we do not want to
	// collide with real addresses, so we specify it.
	Bech32HRPSegwit: "xbbt",
}

// RegressionNetParams returns the chain configuration for regression net.
var RegressionNetParams = chaincfg.Params{
	Name: "regtest",

	// Bitblocks has 0xdab5bffa as RegTest (same as Bitcoin's RegTest).
	// Setting it to an arbitrary value (leet_hex(bitblocks)), so that we can
	// register the regtest network.
	Net: 0xfabfb5da,

	// Address encoding magics
	PubKeyHashAddrID: 111,
	ScriptHashAddrID: 196,
	PrivateKeyID:     239,

	// BIP32 hierarchical deterministic extended key magics
	HDPrivateKeyID: [4]byte{0x0, 0x0, 0x0, 0x0}, // not compatible
	HDPublicKeyID:  [4]byte{0x0, 0x0, 0x0, 0x0}, // not compatible

	// Human-readable part for Bech32 encoded segwit addresses, as defined in
	// BIP 173. Bitblocks does not actually support this, but we do not want to
	// collide with real addresses, so we specify it.
	Bech32HRPSegwit: "xbbrt",
}
