# Shopping Cart in Golang

## Installation

#### Clone Repository
```bash
git clone https://github.com/MatiasAdrian4/shopping_cart_in_golang.git
```

#### Compile Protobuf
```bash
protoc pb/shopping_cart.proto --go_out=plugins=grpc:.
```

#### Build
```bash
go build cmd/main_server.go
```

#### Run Server
```bash
go run cmd/main_server.go
```

#### Run Client
```bash
go run cmd/main_client.go <transport_protocol> <service_name> <arguments>
```
###### Example HTTP
```bash
go run cmd/main_client.go http add_cart 5
```
###### Example GRPC
```bash
go run cmd/main_client.go grpc add_item 2 ball 750
```

