package shopping_cart

import (
	"fmt"
	"context"
	"database/sql"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"

	"shopping_cart_in_golang_with_go_kit/pb"
	"shopping_cart_in_golang_with_go_kit/errs"
)

type ShoppingCartService interface {
	AddCart(context.Context, int) (int, error)
	GetCart(context.Context, int) (*pb.Cart, error)
	ListCarts(context.Context) ([]*pb.Cart, error)
	
	AddItem(context.Context, int, string, float64) (int, error)
	GetItem(context.Context, int) (*pb.Item, error)
	ListItems(context.Context) ([]*pb.Item, error)

	AddCartElement(context.Context, int, int, float64) (error)
	ListItemsByCart(context.Context, int) ([]*pb.Item, error)
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

	err = db.QueryRow("INSERT INTO shopping_cart.Cart (id) VALUES(" + strconv.Itoa(Id) + ")").Scan()
	if strings.Split(err.Error(), ":")[0] == "Error 1062" {
		return 0, errs.DuplicityCartError(err)
	} else {
		return Id, nil
	}
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
		return nil, errs.CartDoesNotExist(err)
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

	err = db.QueryRow("INSERT INTO shopping_cart.Item (id,detail,price) VALUES(" + strconv.Itoa(Id) + ",'" + Detail + "'," + fmt.Sprintf("%f", Price) + ")").Scan()
	if strings.Split(err.Error(), ":")[0] == "Error 1062" {
		return 0, errs.DuplicityItemError(err)
	} else {
		return Id, nil
	}
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
		return nil, errs.ItemDoesNotExist(err)
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

func (s shoppingCartService) AddCartElement(_ context.Context, Cart_id int, Item_id int, Quantity float64) (error) {
	db, err := connect()
	if err != nil {
		return err
	}
	defer db.Close()

	sqlStatement := `INSERT INTO shopping_cart.CartElement (cart_id, item_id, quantity) VALUES(?,?,?);`
	err = db.QueryRow(sqlStatement, Cart_id, Item_id, Quantity).Scan()
	if strings.Split(err.Error(), ":")[0] == "Error 1452" {
		return errs.CartOrItemDoesNotExist(err)
	} else {
		return nil
	}
}

func (s shoppingCartService) ListItemsByCart(_ context.Context, cart_id int) ([]*pb.Item, error) {
	db, err := connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStatement := `SELECT item_id FROM shopping_cart.CartElement WHERE cart_id=?`
	all_items, err := db.Query(sqlStatement, cart_id)
	if err != nil {
		return nil, err
	}

	var items []*pb.Item
	sqlStatement = `SELECT id, detail, price FROM shopping_cart.Item WHERE id=?;`
	
	for all_items.Next() {
		var item_id int
		all_items.Scan(&item_id)
		row := db.QueryRow(sqlStatement, item_id)
		var id int
		var detail string
		var price float64
		row.Scan(&id, &detail, &price)
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
