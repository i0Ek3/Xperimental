#!/bin/bash

# ref: https://segmentfault.com/a/1190000013339403


echo "This script will setup the basic RPC development env and build proto file."
sleep 5

echo "Now, we'll pull grpc and realated components, maybe you should use proxy."
go get -u google.golang.org/grpc
go get -u google.golang.org/protobuf/protoc-gen-go
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway

echo "Now, we'll checks if protoc and protoc-gen-go installed."
protoc --version
protoc-gen-go --version

echo "Build proto file with grpc support."
echo "If you want to build proto file without rpc support, run command:"
echo "      protoc --go_out=. *.proto"
echo "This two commands will generate different .pb.go files"
sleep 5
protoc --go-grpc_out=. ./proto/*.proto


echo "All done!"
