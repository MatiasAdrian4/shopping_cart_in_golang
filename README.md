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
go build cmd/main.go
```

#### Run Server
```bash
go run cmd/main.go
```
