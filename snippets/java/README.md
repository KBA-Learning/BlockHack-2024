# Snippets for Java

## Prerequisites

[![OpenJDK Badge](https://img.shields.io/badge/OpenJDK-000?logo=openjdk&logoColor=fff&style=for-the-badge)](https://openjdk.org/)[![Web3j Badge](https://img.shields.io/badge/Web3j-3C3C3D?logo=ethereum&logoColor=fff&style=for-the-badge)](https://www.web3labs.com/web3j-sdk)

## Run Locally

Change directory:

```sh
cd Web3App
```

Generate the smart contract function wrappers:

```sh
web3j generate solidity -a /path/to/abi -b /path/to/bytecode -o /path/to/src -p <package-name>
```

Interact with deployed Smart Contract:

```sh
web3j run network wallet-path wallet-password
```
