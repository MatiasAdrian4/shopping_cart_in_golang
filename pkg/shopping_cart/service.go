package shopping_cart

import (
	"fmt"
	"context"
	"database/sql"
	"strconv"

	_ "github.com/go-sql-driver/mysql"

	"shopping_cart/pb"
)

type ShoppingCartService interface {
	AddCart(context.Context, int) (int, error)
	GetCart(context.Context, int) (*pb.Cart, error)
	ListCarts(context.Context) ([]*pb.Cart, error)
	
	AddItem(context.Context, int, string, float64) (int, error)
	GetItem(context.Context, int) (*pb.Item, error)
	ListItems(context.Context) ([]*pb.Item, error)
}

func NewShoppingCartServer() ShoppingCartService {
	return shoppingCartService{}
}

type shoppingCartService struct{}

func (s shoppingCartService) AddCart(_ context.Context, Id int) (int, error) {
	db, err := connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.QueryRow("INSERT INTO shopping_cart.Cart (id) VALUES(" + strconv.Itoa(Id) + ")")

	return Id, nil
}

func (s shoppingCartService) GetCart(_ context.Context, Id int) (*pb.Cart, error) {
	db, err := connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStatement := `SELECT id FROM shopping_cart.Cart WHERE id=?;`
	row := db.QueryRow(sqlStatement, Id)
	var id int
	err = row.Scan(&id)
	if err != nil {
		return nil, err
	}
	return &pb.Cart{
		Id: int64(id),
	}, nil
}

func (s shoppingCartService) ListCarts(_ context.Context) ([]*pb.Cart, error) {
	db, err := connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id FROM shopping_cart.Cart")
	if err != nil {
		return nil, err
	}

	var carts []*pb.Cart
	for rows.Next() {
		var id int
		rows.Scan(&id)
		carts = append(carts, &pb.Cart{
			Id: int64(id),
		})
	}

	return carts, nil
}

func (s shoppingCartService) AddItem(_ context.Context, Id int, Detail string, Price float64) (int, error) {
	db, err := connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.QueryRow("INSERT INTO shopping_cart.Item (id,detail,price) VALUES(" + strconv.Itoa(Id) + ",'" + Detail + "'," + fmt.Sprintf("%f", Price) + ")")

	return Id, nil
}

func (s shoppingCartService) GetItem(_ context.Context, Id int) (*pb.Item, error) {
	db, err := connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStatement := `SELECT id, detail, price FROM shopping_cart.Item WHERE id=?;`
	row := db.QueryRow(sqlStatement, Id)
	var id int
	var detail string
	var price float64
	err = row.Scan(&id, &detail, &price)
	if err != nil {
		return nil, err
	}
	return &pb.Item{
		Id: int64(id),
		Detail: detail,
		Price: price,
	}, nil
}

func (s shoppingCartService) ListItems(_ context.Context) ([]*pb.Item, error) {
	db, err := connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, detail, price FROM shopping_cart.Item")
	if err != nil {
		return nil, err
	}

	var items []*pb.Item
	for rows.Next() {
		var id int
		var detail string
		var price float64
		rows.Scan(&id, &detail, &price)
		items = append(items, &pb.Item{
			Id: int64(id),
			Detail: detail,
			Price: price,
		})
	}

	return items, nil
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/shopping_cart")
	if err != nil {
		return nil, err
	}
	return db, nil
}
