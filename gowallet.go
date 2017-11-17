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
	mnemonic := "cannon ranch answer cherry chunk market muscle mesh rail rice wheel wrong"
	fKey, err := NewKeyFromMnemonic(mnemonic, TypeFactomFactoids, bip32.FirstHardenedChild, 0, 0)
	if err != nil {
		fmt.Println(err)
	}
	// if fKey.String() != "xprvA2vH8KdcBBKhMxhENJpJdbwiU5cUXSkaHR7QVTpBmusgYMR8NsZ4BFTNyRLUiaPHg7UYP8u92FJkSEAmmgu3PDQCoY7gBsdvpB7msWGCpXG" {
	// 	fmt.Printf("Invalid Factoid key - %v", fKey.String())
	// }
	fmt.Println(fKey.String())
	ecKey, err := NewKeyFromMnemonic(mnemonic, TypeFactomEntryCredits, bip32.FirstHardenedChild, 0, 0)
	if err != nil {
		fmt.Println(err)
	}
	// if ecKey.String() != "xprvA2ziNegvZRfAAUtDsjeS9LvCP1TFXfR3hUzMcJw7oYAhdPqZyiJTMf1ByyLRxvQmGvgbPcG6Q569m26ixWsmgTR3d3PwicrG7hGD7C7seJA" {
	// 	fmt.Printf("Invalid EC key - %v", ecKey.String())
	// }
	fmt.Println(ecKey.String())
}

func testNewKeyFromMasterKey() {
	mnemonic := "cannon ranch answer cherry chunk market muscle mesh rail rice wheel wrong"

	seed, err := bip39.NewSeedWithErrorChecking(mnemonic, "")
	if err != nil {
		fmt.Println(err)
	}

	masterKey, err := bip32.NewMasterKey(seed)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(masterKey)
	fKey, err := NewKeyFromMasterKey(masterKey, TypeFactomFactoids, bip32.FirstHardenedChild, 0, 0)
	if err != nil {
		fmt.Println(err)
	}
	// if fKey.String() != "xprvA2vH8KdcBBKhMxhENJpJdbwiU5cUXSkaHR7QVTpBmusgYMR8NsZ4BFTNyRLUiaPHg7UYP8u92FJkSEAmmgu3PDQCoY7gBsdvpB7msWGCpXG" {
	// 	fmt.Printf("Invalid Factoid key - %v", fKey.String())
	// }
	fmt.Println(fKey.String())
	ecKey, err := NewKeyFromMasterKey(masterKey, TypeFactomEntryCredits, bip32.FirstHardenedChild, 0, 0)
	if err != nil {
		fmt.Println(err)
	}
	// if ecKey.String() != "xprvA2ziNegvZRfAAUtDsjeS9LvCP1TFXfR3hUzMcJw7oYAhdPqZyiJTMf1ByyLRxvQmGvgbPcG6Q569m26ixWsmgTR3d3PwicrG7hGD7C7seJA" {
	// 	fmt.Printf("Invalid EC key - %v", ecKey.String())
	// }
	fmt.Println(ecKey.String())
}