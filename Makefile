build-sc:
	solc --abi ./contracts/documents/Document.sol -o build && \
	abigen --abi ./contracts/documents/build/Document.abi --pkg document --out document.go 
setup:
	docker-compose up -d 
run:
	go run main.go