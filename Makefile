init-network:
	geth --http --http.addr "localhost" --http.port "8545"  --datadir "/home/cuongtop/Desktop/DigitalSignaturalNetwork" init "/home/cuongtop/Desktop/DigitalSignaturalNetwork/genesis.json" 
run-network:
	geth --http --http.addr "localhost" --http.port "8545"  --datadir "/home/cuongtop/Desktop/DigitalSignaturalNetwork" --mine --miner.threads=1 --networkid 45
build-sc:
	abigen --abi ./contracts/documents/build/Document.abi --bin ./contracts/documents/build/Document.bin --pkg document --out document.go
deploy-sc:
	