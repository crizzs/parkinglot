#!/usr/bin/env bash
go get github.com/stretchr/testify
go test -v
cd ./src
go test -v
cd ..
go run main.go sample_file/file_inputs.txt
go run main.go