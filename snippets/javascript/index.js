import { readFileSync } from 'fs'
import { JsonRpcProvider, Contract } from 'ethers'
import chalk from 'chalk'

const provider = new JsonRpcProvider('http://127.0.0.1:8545')

const signer = await provider.getSigner()
console.log(`Address: ${chalk.green(signer.address)}`)

const balance = await provider.getBalance(signer.address)
console.log(`Balance: ${chalk.green(balance)}`)

let latestBlock = await provider.getBlockNumber()
console.log(`Latest Block: ${chalk.green(latestBlock)}`)

const abi = readFileSync('../../common/contracts/Storage.abi').toString()

const contract = new Contract(
  '0x5FbDB2315678afecb367f032d93F642f64180aa3',
  abi,
  signer
)

const trx = await contract.store('101', 'Hello from JavaScript 🔰!')
console.log(`Transaction Hash: ${chalk.green(trx.hash)}`)

const value = await contract.collection('101')
console.log(`Value: ${chalk.green(value)}`)
