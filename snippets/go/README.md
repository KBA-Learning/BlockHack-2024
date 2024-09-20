# Snippets for Go

## Prerequisites

[![Go Badge](https://img.shields.io/badge/Go-00ADD8?logo=go&logoColor=fff&style=for-the-badge)](https://go.dev/)

## Run Locally

Install abigen:

```sh
go install github.com/ethereum/go-ethereum/cmd/abigen@latest
```

Generate Go binding:

```sh
abigen --bin /path/to/bytecode --abi /path/to/abi --pkg lib --type Storage --out lib/Storage.go
```

Deploy the Storage smart contract (if haven't already):

```sh
go run cmd/deploy.go
```

Interact with deployed smart contract:

```sh
go run .
```
