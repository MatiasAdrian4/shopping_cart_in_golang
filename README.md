# Shopping Cart in Golang With Go-Kit

## Introduction

A simple shopping cart developed in Golang, based on [addsvc](https://github.com/go-kit/kit/tree/master/examples/addsvc) example.

## Installation

#### Clone Repository
```bash
git clone https://github.com/MatiasAdrian4/shopping_cart_in_golang_with_go_kit.git
```

#### Compile Protobuf
```bash
protoc pb/shopping_cart.proto --go_out=plugins=grpc:.
```

#### Build
```bash
go build cmd/server/main.go
```

#### Run Server
```bash
go run cmd/server/main.go
```

#### Run Client (Only add_cart and add_item services)
```bash
go run cmd/client/main.go <transport_protocol> <service_name> <arguments>
```
###### Example HTTP
```bash
go run cmd/client/main.go http add_cart 5
```
###### Example GRPC
```bash
go run cmd/client/main.go grpc add_item 2 ball 750
```

