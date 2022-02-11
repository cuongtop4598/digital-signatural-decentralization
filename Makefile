build-sc:
	abigen --abi ./contracts/documents/build/Document.abi --bin ./contracts/documents/build/Document.bin --pkg document --out document.go
setup:
	docker-compose up -d 
run:
	go run main.go