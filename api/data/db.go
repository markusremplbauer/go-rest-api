package data

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS products
(
    id SERIAL,
    name TEXT NOT NULL,
    price NUMERIC(10,2) NOT NULL DEFAULT 0.00,
    CONSTRAINT products_pkey PRIMARY KEY (id)
)`

func Initialize(user, password, host, dbname string) {
	connectionString := fmt.Sprintf("user=%s password=%s host=%s dbname=%s sslmode=disable", user, password, host, dbname)

	var err error
	DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	ensureTableExists()
}

func ensureTableExists() {
	if _, err := DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func (p *Product) GetProduct(db *sql.DB) error {
	return db.QueryRow("SELECT name, price FROM products WHERE id=$1", p.ID).Scan(&p.Name, &p.Price)
}

func (p *Product) UpdateProduct(db *sql.DB) error {
	_, err := db.Exec("UPDATE products SET name=$1, price=$2 WHERE id=$3", p.Name, p.Price, p.ID)
	return err
}

func (p *Product) DeleteProduct(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM products WHERE id=$1", p.ID)

	return err
}

func (p *Product) CreateProduct(db *sql.DB) error {
	err := db.QueryRow("INSERT INTO products(name, price) VALUES($1, $2) RETURNING id", p.Name, p.Price).Scan(&p.ID)

	if err != nil {
		return err
	}

	return nil
}

func GetProducts(db *sql.DB, offset, limit int) ([]Product, error) {
	rows, err := db.Query("SELECT id, name,  price FROM products LIMIT $1 OFFSET $2", limit, offset)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	products := []Product{}

	for rows.Next() {
		var p Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}
