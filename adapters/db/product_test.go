package db_test

import (
	"appproduct/adapters/db"
	"appproduct/application"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	createProduct(Db)
}

func createTable(db *sql.DB) {
	table := `CREATE TABLE products (
		id string,
		name string,
		price string,
		status string
	);`

	stmt, err := db.Prepare(table)
	if err != nil {
		panic(err)
	}

	stmt.Exec()
}

func createProduct(db *sql.DB) {
	insert := `INSERT INTO products VALUES (
		"abc",
		"Product ABC",
        "100",
        "disabled"
	)`

	stmt, err := db.Prepare(insert)
	if err != nil {
		panic(err)
	}

	stmt.Exec()
}

func TestProductDb_Get(t *testing.T) {
	setUp()
	defer Db.Close()

	productDb := db.NewProductDb(Db)
	product, err := productDb.Get("abc")

	require.Nil(t, err)
	require.Equal(t, "Product ABC", product.GetName())
	require.Equal(t, 100.0, product.GetPrice())
	require.Equal(t, application.DISABLED, product.GetStatus())
}

func TestProductDb_Save(t *testing.T) {
	setUp()
	defer Db.Close()

	productDb := db.NewProductDb(Db)
	product := application.NewProduct()
	product.Name = "Product ABC"
	product.Price = 100.0

	productResult, err := productDb.Save(product)

	require.Nil(t, err)
	require.Equal(t, "Product ABC", productResult.GetName())
	require.Equal(t, 100.0, productResult.GetPrice())
	require.Equal(t, application.DISABLED, productResult.GetStatus())

	product.Status = application.ENABLED

	productResult, err = productDb.Save(product)
	
	require.Nil(t, err)
	require.Equal(t, "Product ABC", productResult.GetName())
	require.Equal(t, 100.0, productResult.GetPrice())
	require.Equal(t, application.ENABLED, productResult.GetStatus())
}
