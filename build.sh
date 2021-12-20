#!/bin/bash

#export GO111MODULE=on
go mod init
go build -o depocket_solana ./cmd/main.go
go build -o generate_secret ./cmd/generate_secret.go