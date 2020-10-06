# About 
This is a simple "Hello World" equivalent for creating [gRPC](https://grpc.io/) APIs. 

## Challenge
There's a proto file defined which is machine and human readable. To know more about what protobuf is, [visit here](https://developers.google.com/protocol-buffers)

Proto file can be used to generate server stubs and client library based on the definition. 
In this example, I've done it using Golang but there are more than ten other [languages](https://grpc.io/docs/languages/) that you can do it in. 

If you're looking for something like cURL for gRPC, check this [out.](https://github.com/fullstorydev/grpcurl) 

If you wanna learn more about gRPC, would highly recommend [this.](https://www.youtube.com/watch?v=pMgty_RYIOc&list=PLmD8u-IFdreyyTx93jJ5GkijwDXFqyr3T)


## Solution
Server stub has been implemented in [s_server.go](./s_server.go) and example of client is in [s_client.go](./s_client.go)

 