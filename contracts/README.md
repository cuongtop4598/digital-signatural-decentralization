# Step by step to build a private ethereum network
---
## 1.Go into the technical nuances
1. [Learn about Ethereum](https://ethereum.org/en/developers/tutorials/)
## 2.Preparation
1. [Install Geth](https://geth.ethereum.org/docs/install-and-build/installing-geth)
## 3.Create the Datadir folder
## 4.Create genesis block
As you probably know, the blockchain is a distributed digital register in which all transactions are recorded in chronological order in the form of blocks. Blocks are generated during the mining process and added to the blockchain chain. The number of blocks is unlimited, but there is always one separate block that gave rise to the whole chain - the genesis-block.
To create a private blockchain, you first need to generate a genesis block. To do this, you need to create a Genesis file, write the necessary commands (attributes) into it and use it with Geth.
- Cmd: 
```
geth init /path/to/genesisfile.json --datadir /path/to/data/directory
```
- genesis.json file like this:
```
{
    "config": {
        "chainId": 120999,
        "homesteadBlock": 0,
        "eip150Block": 0,
        "eip155Block": 0,
        "eip158Block": 0,
        "ethash": {}
    },
    "difficulty":"1",
    "gasLimit": "0x8000000",
    "alloc": {}
}
```
## 5.Connect to IPC
- Command:
```
geth attach /path/to/datadir/geth.ipc
```
- Create new account:
```
personal.newAccount()
```
- Get Balance: 
```
eth.getBalance("0x...")
```
## 6. Start miner
Before starting miner, you must set etherbase account for miner, like this:
```
miner.setEtherbase("0x...")
```
- Start miner:
```
miner.start()
```
- Stop miner:
```
miner.stop()
```
- Check ethers you mined:
```
eth.getBalance("0x...")
```
---
## Expose Network to use by HTTP based JSON RPC API
Command:
```
geth --datadir ./ --networkid 120999 --http --http.addr 0.0.0.0 --http.port 8545 --http.corsdomain "*" --http.api "web3,personal,eth,net,admin"
```
