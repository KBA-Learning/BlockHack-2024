import json
from web3 import Web3
from eth_account import Account

def main():
    w3 = Web3(Web3.HTTPProvider("http://127.0.0.1:8545"))

    with open("../../common/contracts/Storage.abi", "r", -1, "utf-8") as f:
        abi = json.load(f)

    account = Account.from_key(
        "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
    )

    print(f"Address: \033[32m{account.address}\033[0m")
    print(f"Balance: \033[32m{w3.eth.get_balance(account.address)}\033[0m")
    print(f"Latest Block: \033[32m{w3.eth.get_block_number()}\033[0m")

    contract = w3.eth.contract(
        address="0x5FbDB2315678afecb367f032d93F642f64180aa3", abi=abi
    )

    trx = contract.functions.store("1991", "Hello from Python 🐍!").transact(
        {"from": account.address}
    )
    print(f"Transaction Hash: \033[32m{trx.hex()}\033[0m")

    value = contract.functions.collection("1991").call()
    print(f"Value: \033[32m{value}\033[0m")


if __name__ == "__main__":
    main()
