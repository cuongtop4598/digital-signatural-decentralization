run-network:
	geth --rpcapi eth,web3,personal --rpc --networkid=15
run:
	go run main.go