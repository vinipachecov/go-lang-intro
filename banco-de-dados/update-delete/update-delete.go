package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
	return result
}

func main() {
	db, err := sql.Open("mysql", "root:example@/cursogo")

	if err != nil {
		panic(err)
	}

	defer db.Close()
	stmt, err := db.Prepare("update usuarios set nome = ? where id = ?")

	stmt.Exec("New name", 1)

}
