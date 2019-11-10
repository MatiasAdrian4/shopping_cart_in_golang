package main

import (
	"fmt"
	"log"
	"flag"
	"os"
	"context"
	"strconv"

	"google.golang.org/grpc"
	
	sc "shopping_cart_in_golang_with_go_kit/pkg/shopping_cart"
)

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		fmt.Fprintln(os.Stderr, "missing protocol transfer")
		os.Exit(1)
	}

	var client sc.ShoppingCartService
	var err error

	switch cmd := flag.Arg(0); cmd {
	case "http":
		client, err = sc.NewHTTPClient("localhost:8081")
		if err != nil {
			log.Fatal("could not connect to backend: %v", err)
		}
	case "grpc":
		conn, err := grpc.Dial(":8082", grpc.WithInsecure())
		if err != nil {
			log.Fatal("could not connect to backend: %v", err)
		}
		client = sc.NewGRPCClient(conn)
	default:
		err = fmt.Errorf("unknown protocol transport %s", cmd)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

	switch cmd := flag.Arg(1); cmd {
		case "add_cart":
			if flag.NArg() < 3 {
				log.Fatal("missing arguments")
			}
			id, err := strconv.Atoi(flag.Arg(2))
			if err != nil {
				log.Fatal("incorrect id")
			}
			fmt.Println(client.AddCart(context.Background(), id))
		case "add_item":
			if flag.NArg() < 5 {
				log.Fatal("missing arguments")
			}
			id, err := strconv.Atoi(flag.Arg(2))
			if err != nil {
				log.Fatal("incorrect id")
			}
			
			detail := flag.Arg(3)

			price, err := strconv.ParseFloat(flag.Arg(4), 64)
			if err != nil {
				log.Fatal("incorrect price")
			}
			fmt.Println(client.AddItem(context.Background(), id, detail, price))
		default:
			err = fmt.Errorf("unknown service %s", cmd)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
	}
}