package main

import (
	"fmt"
	"log"
	"flag"
	"os"
	"context"
	"strconv"

	"google.golang.org/grpc"
	
	sc "shopping_cart/pkg/shopping_cart"
)

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		fmt.Fprintln(os.Stderr, "missing service")
		os.Exit(1)
	}

	conn, err := grpc.Dial(":8082", grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect to backend: %v", err)
	}
	client := sc.NewGRPCClient(conn)

	switch cmd := flag.Arg(0); cmd {
		case "add_cart":
			id, err := strconv.Atoi(flag.Arg(1))
			if err != nil {
				log.Fatal("incorrect id")
			}
			fmt.Println(client.AddCart(context.Background(), id))
		case "add_item":
			id, err := strconv.Atoi(flag.Arg(1))
			if err != nil {
				log.Fatal("incorrect id")
			}
			detail := flag.Arg(2)

			price, err := strconv.ParseFloat(flag.Arg(3), 64)
			if err != nil {
				log.Fatal("incorrect price")
			}
			fmt.Println(client.AddItem(context.Background(), id, detail, price))
		default:
			err = fmt.Errorf("unknown service %s", cmd)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}