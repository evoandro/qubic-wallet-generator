package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/qubic/go-node-connector/types"
)

const (
	numWorkers = 100
)

type WalletResult struct {
	Seed   string
	Wallet types.Wallet
}

var totalTries uint64

func worker(ctx context.Context, id int, targetPrefix string, wg *sync.WaitGroup, resultChan chan<- WalletResult) {
	defer wg.Done()

	for {
		count := atomic.AddUint64(&totalTries, 1)
		// Print the total tries every 100,000 attempts
		if count%1000000 == 0 {
			fmt.Printf("Total attempts: %d\n", count)
		}
		select {
		case <-ctx.Done():
			return
		default:
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
				select {
				case resultChan <- WalletResult{Seed: seed, Wallet: wallet}:
				case <-ctx.Done():
				}
				return
			}
		}
	}
}

func main() {
	// Define and parse the command-line flag
	targetPrefix := flag.String("prefix", "", "The target prefix for the wallet identity")
	flag.Parse()

	// Check if the prefix is empty and return an error if it is
	if *targetPrefix == "" {
		fmt.Println("Error: No prefix provided. Please specify a prefix using the -prefix flag.")
		os.Exit(1)
	}

	var wg sync.WaitGroup
	resultChan := make(chan WalletResult)

	ctx, cancel := context.WithCancel(context.Background())

	// Start worker goroutines
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(ctx, i, *targetPrefix, &wg, resultChan)
	}

	// Wait for a result and then close the channel
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Get the result from the channel
	walletResult := <-resultChan
	cancel() // Cancel all workers

	// Print the generated wallet details
	fmt.Printf("Seed: %s\n", walletResult.Seed)
	fmt.Printf("Private Key: %x\n", walletResult.Wallet.PrivKey)
	fmt.Printf("Public Key: %x\n", walletResult.Wallet.PubKey)
	fmt.Printf("Identity: %s\n", walletResult.Wallet.Identity)
}
