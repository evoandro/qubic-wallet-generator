package main

import (
	"fmt"
	"log"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/qubic/go-node-connector/types"
)

const (
	targetPrefix = "D"
	numWorkers   = 1048575
)

var totalTries uint64

type WalletResult struct {
	Seed   string
	Wallet types.Wallet
}

func worker(id int, wg *sync.WaitGroup, resultChan chan<- WalletResult) {
	defer wg.Done()

	for {
		// Generate a random seed
		seed := types.GenerateRandomSeed()

		// Create a new wallet using the generated seed
		wallet, err := types.NewWallet(seed)
		if err != nil {
			log.Printf("Worker %d: Failed to create wallet: %v", id, err)
			continue
		}

		// Check if the Identity starts with the desired prefix
		if strings.HasPrefix(string(wallet.Identity), targetPrefix) {
			fmt.Printf("Identity: %s\n", wallet.Identity)
			resultChan <- WalletResult{Seed: seed, Wallet: wallet}
			return
		}

		// Increment the total number of tries
		count := atomic.AddUint64(&totalTries, 1)

		// Print the total tries every 100,000 attempts
		if count%1000000 == 0 {
			fmt.Printf("Total attempts: %d\n", count)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	resultChan := make(chan WalletResult)

	// Start worker goroutines
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, &wg, resultChan)
	}

	// Wait for a result and then close the channel
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Get the result from the channel
	walletResult := <-resultChan

	// Print the generated wallet details
	fmt.Printf("Seed: %s\n", walletResult.Seed)
	fmt.Printf("Private Key: %x\n", walletResult.Wallet.PrivKey)
	fmt.Printf("Public Key: %x\n", walletResult.Wallet.PubKey)
	fmt.Printf("Identity: %s\n", walletResult.Wallet.Identity)
}
