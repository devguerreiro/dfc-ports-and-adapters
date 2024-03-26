package db_test

import "database/sql"

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory")
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
