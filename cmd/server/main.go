package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"text/tabwriter"

	kitgrpc "github.com/go-kit/kit/transport/grpc"

	"github.com/oklog/oklog/pkg/group"
	"google.golang.org/grpc"

	sc "shopping_cart_in_golang_with_go_kit/pkg/shopping_cart"
	"shopping_cart_in_golang_with_go_kit/pb"
)

func main() {

	fs := flag.NewFlagSet("addsvc", flag.ExitOnError)
	var (
		httpAddr = fs.String("http-addr", ":8081", "HTTP listen address")
		grpcAddr = fs.String("grpc-addr", ":8082", "gRPC listen address")
	)
	fs.Usage = usageFor(fs, os.Args[0]+" [flags]")
	fs.Parse(os.Args[1:])

	var (
		service     = sc.NewShoppingCartServer()
		endpoints   = sc.MakeEndpoints(service)
		httpHandler = sc.NewHTTPHandler(endpoints)
		grpcServer  = sc.NewGRPCServer(endpoints)
	)

	var g group.Group

	httpListener, err := net.Listen("tcp", *httpAddr)
	if err != nil {
		os.Exit(1)
	}
	g.Add(func() error {
		return http.Serve(httpListener, httpHandler)
	}, func(error) {
		httpListener.Close()
	})

	grpcListener, err := net.Listen("tcp", *grpcAddr)
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

func usageFor(fs *flag.FlagSet, short string) func() {
	return func() {
		fmt.Fprintf(os.Stderr, "USAGE\n")
		fmt.Fprintf(os.Stderr, "  %s\n", short)
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "FLAGS\n")
		w := tabwriter.NewWriter(os.Stderr, 0, 2, 2, ' ', 0)
		fs.VisitAll(func(f *flag.Flag) {
			fmt.Fprintf(w, "\t-%s %s\t%s\n", f.Name, f.DefValue, f.Usage)
		})
		w.Flush()
		fmt.Fprintf(os.Stderr, "\n")
	}
}
