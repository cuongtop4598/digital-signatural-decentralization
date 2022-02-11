Create Go bindings

- abigen --sol TokenBalances.sol --pkg fetchtokenbalances --out fetch_token_balance.go

Create ABI

- solc --abi Document.sol -o build
