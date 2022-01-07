run-network:
	geth --datadir /home/cuongtop/.ethereum --miner.threads=1 --http --http.addr 0.0.0.0 --http.port 8545 --http.corsdomain "*" --http.api "web3,personal,eth,net,admin" --networkid=15
run:
	go run main.go
build-sc:
	abigen --abi ./contracts/documents/build/Document.abi --bin ./contracts/documents/build/Document.bin --pkg document --out document.go
deploy-sc:
	