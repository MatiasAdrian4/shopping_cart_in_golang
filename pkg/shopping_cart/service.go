package shopping_cart

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type CartService interface {
	AddCart(context.Context, int) (int, error)
	AddItem(context.Context, int, string, float64) (int, error)
}

func NewServer() CartService {
	var svc CartService
	svc = NewCartService()
	return svc
}

func NewCartService() CartService {
	return cartService{}
}

type cartService struct{}

func (s cartService) AddCart(_ context.Context, Id int) (int, error) {
	db, err := connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.QueryRow("INSERT INTO Cart (id) VALUES(" + strconv.Itoa(Id) + ")")

	return Id, nil
}

func (s cartService) AddItem(_ context.Context, Id int, Detail string, Price float64) (int, error) {
	db, err := connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.QueryRow("INSERT INTO Item (id,detail,price) VALUES(" + strconv.Itoa(Id) + ",'" + Detail + "'," + fmt.Sprintf("%f", Price) + ")")

	return Id, nil
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/shopping_cart")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return db, nil
}
