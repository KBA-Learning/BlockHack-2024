import { readFileSync } from 'fs'
import { JsonRpcProvider, Contract } from 'ethers'

const provider = new JsonRpcProvider('http://127.0.0.1:8545')

const signer = await provider.getSigner()
console.log('Address:', signer.address)

const balance = await provider.getBalance(signer.address)
console.log('Balance:', balance)

let latestBlock = await provider.getBlockNumber()
console.log('Latest Block:', latestBlock)

const abi = readFileSync('../contracts/Storage.abi').toString()

const contract = new Contract(
  '0x5FbDB2315678afecb367f032d93F642f64180aa3',
  abi,
  signer
)

const trx = await contract.store('101', 'Hello, KBA!')
console.log('Transaction Hash:', trx.hash)

const value = await contract.collection('101')
console.log('Value:', value)
