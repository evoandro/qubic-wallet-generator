# QUBIC Wallet Generator

This Go program generates wallets with a specified prefix for the wallet identity. It leverages concurrency to quickly find a wallet that matches the desired prefix using multiple worker goroutines.

## Features

- Generates wallets concurrently using worker goroutines.
- Supports specifying a target prefix for the wallet identity.
- Prints the total number of attempts every 1,000,000 tries.
- Exits with an error if no prefix is provided.

## Prerequisites

- Go 1.20 or later installed on your machine.
- `github.com/qubic/go-node-connector/types` package installed.

## Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/evoandro/qubic-wallet-generator.git
    cd qubic-wallet-generator
    ```

2. Install dependencies:

    ```sh
    go get -u github.com/qubic/go-node-connector/types
    ```

## Usage

To run the program, use the following command, specifying your desired prefix:

```sh
go run main.go -prefix YOUR_PREFIX
```

Replace `YOUR_PREFIX` with the desired prefix for the wallet identity.

### Example

```sh
go run main.go -prefix QUBIC
```

## Build

To build the program, use the following command:

```sh
go build -o wallet-generator
```

After building, you can run the program with the desired prefix as follows:

```sh
./wallet-generator -prefix YOUR_PREFIX
```

Replace `YOUR_PREFIX` with the desired prefix for the wallet identity.

### Example

```sh
./wallet-generator -prefix QUBIC
```

## Code Overview

### `main.go`

- `main()`: Parses the command-line flag for the prefix and initializes worker goroutines. Waits for a wallet result and prints the wallet details.
- `worker()`: A worker function that generates random seeds, creates wallets, and checks if the wallet identity starts with the specified prefix.

## Error Handling

The program will exit with an error if no prefix is provided. Ensure to specify the `-prefix` flag when running the program.

## Contributing

Contributions are welcome! Please fork the repository and open a pull request with your changes.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contact

For any questions or inquiries, please open an issue or contact the repository owner.
