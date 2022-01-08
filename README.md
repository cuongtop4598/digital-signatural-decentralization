## Digital Signature Decentralization

### Introduction

- A project helps people who are in the world can be mutual trust and make remote contracts. It is also used for stored documents and verifying signed documents.

### How to setup

- Install go 16.5
- go mod tidy
- go mod vendor

### How to install

- go install

### How to run

##### 1. Running blockchain network

```
1. copy genesis.json into your /datadir, change path of this file in init-network command on Makefile
2. replace keystore folder in datadir of your local machine with keystore folder in /wallets
3. remove old geth folder and geth.ipc in /datadir and keep genesis.json 
4. run make init-network
5. run make run-network
```

##### 2. Running application

```
make run
```

### Documents

- Run private network tutorial

  - https://geth.ethereum.org/docs/getting-started/private-net
  - https://medium.com/swlh/how-to-set-up-a-private-ethereum-blockchain-c0e74260492c#:~:text=To%20run%20a%20private%20network,with%20a%20custom%20genesis%20file.
- Contract complier tutorial

  - https://www.metachris.com/2021/05/creating-go-bindings-for-ethereum-smart-contracts/
