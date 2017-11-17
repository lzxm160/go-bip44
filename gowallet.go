package main

import (
	bip32"./go-bip32"
	bip39"./go-bip39"
	"fmt"
	
)

func main() {
	testNewKeyFromMnemonic()
	testNewKeyFromMasterKey()
}
func testNewKeyFromMnemonic() {
	mnemonic := "yellow yellow yellow yellow yellow yellow yellow yellow yellow yellow yellow yellow"
	fKey, err := NewKeyFromMnemonic(mnemonic, TypeFactomFactoids, bip32.FirstHardenedChild, 0, 0)
	if err != nil {
		panic(err)
	}
	if fKey.String() != "xprvA2vH8KdcBBKhMxhENJpJdbwiU5cUXSkaHR7QVTpBmusgYMR8NsZ4BFTNyRLUiaPHg7UYP8u92FJkSEAmmgu3PDQCoY7gBsdvpB7msWGCpXG" {
		fmt.Printf("Invalid Factoid key - %v", fKey.String())
	}

	ecKey, err := NewKeyFromMnemonic(mnemonic, TypeFactomEntryCredits, bip32.FirstHardenedChild, 0, 0)
	if err != nil {
		panic(err)
	}
	if ecKey.String() != "xprvA2ziNegvZRfAAUtDsjeS9LvCP1TFXfR3hUzMcJw7oYAhdPqZyiJTMf1ByyLRxvQmGvgbPcG6Q569m26ixWsmgTR3d3PwicrG7hGD7C7seJA" {
		fmt.Printf("Invalid EC key - %v", ecKey.String())
	}
}

func testNewKeyFromMasterKey() {
	mnemonic := "yellow yellow yellow yellow yellow yellow yellow yellow yellow yellow yellow yellow"

	seed, err := bip39.NewSeedWithErrorChecking(mnemonic, "")
	if err != nil {
		panic(err)
	}

	masterKey, err := bip32.NewMasterKey(seed)
	if err != nil {
		panic(err)
	}

	fKey, err := NewKeyFromMasterKey(masterKey, TypeFactomFactoids, bip32.FirstHardenedChild, 0, 0)
	if err != nil {
		panic(err)
	}
	if fKey.String() != "xprvA2vH8KdcBBKhMxhENJpJdbwiU5cUXSkaHR7QVTpBmusgYMR8NsZ4BFTNyRLUiaPHg7UYP8u92FJkSEAmmgu3PDQCoY7gBsdvpB7msWGCpXG" {
		fmt.Printf("Invalid Factoid key - %v", fKey.String())
	}

	ecKey, err := NewKeyFromMasterKey(masterKey, TypeFactomEntryCredits, bip32.FirstHardenedChild, 0, 0)
	if err != nil {
		panic(err)
	}
	if ecKey.String() != "xprvA2ziNegvZRfAAUtDsjeS9LvCP1TFXfR3hUzMcJw7oYAhdPqZyiJTMf1ByyLRxvQmGvgbPcG6Q569m26ixWsmgTR3d3PwicrG7hGD7C7seJA" {
		fmt.Printf("Invalid EC key - %v", ecKey.String())
	}
}