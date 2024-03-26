package main

import (
	db2 "appproduct/adapters/db"
	"appproduct/application"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", "db.sqlite")

	productDbAdapter := db2.NewProductDb(db)
	productService := application.NewProductService(productDbAdapter)
	
	product, _ := productService.Create("Product", 100)

	productService.Enable(product)
}
