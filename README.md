# Wallet Generator with Identity Prefix Matching

This Go program generates random wallets and checks if the identity of the wallet starts with a specified prefix (`QFUNDS`). The program runs multiple worker goroutines concurrently to speed up the process. When a matching wallet is found, the program outputs the details of the generated wallet.

## Features

- Concurrent wallet generation using goroutines.
- Prefix matching for wallet identities.
- Configurable number of workers.
- Displays progress every 1,000,000 attempts.

## Prerequisites

- Go 1.16 or higher

## Installation

1. Clone the repository:

   ```sh
   git clone https://github.com/your-username/wallet-generator.git
   cd wallet-generator
   ```

2. Install the required dependencies:

   ```sh
   go get ./...
   ```

## Usage

1. Build the program:

   ```sh
   go build -o wallet-generator
   ```

2. Run the program:

   ```sh
   ./wallet-generator
   ```

## Configuration

- `targetPrefix`: The prefix to match in the wallet identity. Default is `QFUNDS`.
- `numWorkers`: The number of worker goroutines to run concurrently. Default is 1000.

These values can be adjusted by modifying the corresponding constants in the `main.go` file:

```go
const (
    targetPrefix = "QFUNDS"
    numWorkers   = 1000
)
```

## Output

When a wallet with the desired prefix is found, the program prints the following details:

- Seed
- Private Key
- Public Key
- Identity

Example output:

```
Seed: <seed_value>
Private Key: <private_key_hex>
Public Key: <public_key_hex>
Identity: QFUNDS...
```

## Notes

- The program will print the total number of attempts every 1,000,000 tries to track the progress.
- The program will run until it finds a wallet that matches the specified prefix.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.

## Acknowledgements

This project uses the `go-node-connector` library for wallet generation. Special thanks to the authors and maintainers of this library.
