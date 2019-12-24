package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"strings"

	kitgrpc "github.com/go-kit/kit/transport/grpc"

	"github.com/oklog/oklog/pkg/group"
	"google.golang.org/grpc"

	sc "shopping_cart_in_golang_with_go_kit/pkg/shopping_cart"
	"shopping_cart_in_golang_with_go_kit/pb"

	"github.com/gorilla/handlers"
)

func main() {

	httpAddr := ":8081"
	grpcAddr := ":8082"

	var (
		service     = sc.NewShoppingCartServer()
		endpoints   = sc.MakeEndpoints(service)
		httpHandler = sc.NewHTTPHandler(endpoints)
		grpcServer  = sc.NewGRPCServer(endpoints)
	)

	var g group.Group

	g.Add(func() error {
		fmt.Println("Running server on", strings.Split(httpAddr, ":")[1])
		return (http.ListenAndServe(httpAddr, handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(httpHandler)))
	}, func(error) {
		fmt.Println("Closing server on", strings.Split(httpAddr, ":")[1])
		os.Exit(1)
	})

	grpcListener, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		os.Exit(1)
	}

	g.Add(func() error {
		baseServer := grpc.NewServer(grpc.UnaryInterceptor(kitgrpc.Interceptor))
		pb.RegisterShoppingCartServer(baseServer, grpcServer)
		return baseServer.Serve(grpcListener)
	}, func(error) {
		grpcListener.Close()
	})

	cancelInterrupt := make(chan struct{})
	g.Add(func() error {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		select {
		case sig := <-c:
			return fmt.Errorf("received signal %s", sig)
		case <-cancelInterrupt:
			return nil
		}
	}, func(error) {
		close(cancelInterrupt)
	})

	g.Run()
}
