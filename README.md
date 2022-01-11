### What is it?

Simple client server application written to demonstrate how gRPC can be used to exchange data. 

## How to Run?

go run server/server.go

go run client/client.go

## Want to update weather.proto?

protoc -I. --go_-_out=. --go--grpc_out=. weather.proto 


