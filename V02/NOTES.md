protoc --go_out=plugins=grpc:. proto/blockchain.proto


go get google.golang.org/grpc

go get -u github.com/golang/protobuf/proto






### server
go run server/main.go

### client
go run client/main.go --list
go run client/main.go --add
go run client/main.go --list
go run client/main.go --add
go run client/main.go --list
go run client/main.go --add

###todo...  

clean up the client output

get the DART code working for a client

dockerize it

