package db_test

import (
	"appproduct/adapters/db"
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
        "DISABLED"
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
	require.Equal(t, "DISABLED", product.GetStatus())
}
