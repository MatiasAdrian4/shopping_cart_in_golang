CREATE TABLE shopping_cart.Cart (
	id	INT,
	PRIMARY KEY (id)
)

CREATE TABLE shopping_cart.Item (
	id	INT,
	detail VARCHAR(255),
	price FLOAT,
	PRIMARY KEY (id)
)

CREATE TABLE shopping_cart.CartElement (
	cart_id INT,
	item_id INT,
	quantity INT,
	FOREIGN KEY (cart_id) REFERENCES shopping_cart.Cart(id),
	FOREIGN KEY (item_id) REFERENCES shopping_cart.Item(id)
)