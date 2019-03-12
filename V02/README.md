#  gRPC / GoLang / BlockChain example.

see NOTES.md on instructions.

This is a work in progress...

## todo
1. Explain what this system is suppose to do.
...* explain REST
...* Explain gRPC/PROTOBUF
2. Explain the lay out of the files.  there are client and servers.
3. Explain how to compile the PROTO file
4. DISCUSS how simple the BLOCKCHAIN example is suppose to be.
5. Discuss the PROTO file
6. Future... do FLUTTER, SWIFT, KOTLIN, C++, JAVA examples of the PROTO compiler.



Goal
Create gRPC Server and Client in Go. We will create a fictional blockchain service.

Steps
Install protoc compiler.
Install protoc-gen-go plugin: go get -u github.com/golang/protobuf/protoc-gen-go
Define service definition in .proto file.
Build Go bindings from .proto file. protoc --go_out=plugins=grpc:. proto/blockchain.proto
Install grpc Go package - go get -u google.golang.org/grpc.
Install context package - go get -u golang.org/x/net/context.
Install protobuf package - go get -u github.com/golang/protobuf/proto
Implement Server, interface BlockchainServer.
Create a client using BlockchainClient.
Run server first.
Run client.
Usage
Start server:

go run server/main.go


Add block as client:

go run client/main.go --add
get blockchain as client:

go run client/main.go --list
