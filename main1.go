package main

import (
	"fmt"
	"log"
	"strings"
	"sync/atomic"

	"github.com/qubic/go-node-connector/types"
)

var totalTries uint64

func main() {
	targetPrefix := "GIFT"
	var wallet types.Wallet
	var seed string
	var err error

	for {
		// Generate a random seed
		seed = types.GenerateRandomSeed()
		count := atomic.AddUint64(&totalTries, 1)

		// Print the total tries every 100,000 attempts
		if count%1000000 == 0 {
			fmt.Printf("Total attempts: %d\n", count)
		}
		// Create a new wallet using the generated seed
		wallet, err = types.NewWallet(seed)
		if err != nil {
			log.Fatalf("Failed to create wallet: %v", err)
		}

		// Check if the Identity starts with the desired prefix
		if strings.HasPrefix(string(wallet.Identity), targetPrefix) {
			break
		}
	}

	// Print the generated wallet details
	fmt.Printf("Seed: %s\n", seed)
	fmt.Printf("Private Key: %x\n", wallet.PrivKey)
	fmt.Printf("Public Key: %x\n", wallet.PubKey)
	fmt.Printf("Identity: %s\n", wallet.Identity)
}
