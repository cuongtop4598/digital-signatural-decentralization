Create Go bindings

- abigen --sol Document.sol --pkg document --out document.go

Create ABI

- solc --abi Document.sol -o build

Build smart contract and create bindings

    solc --abi ./contracts/documents/Document.sol -o build && abigen --abi ./contracts/documents/build/Document.abi --pkg document --out document.go

Run dev network
[https://programmerclick.com/article/73351747654/](https://programmerclick.com/article/73351747654/)
