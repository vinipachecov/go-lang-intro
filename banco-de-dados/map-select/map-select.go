package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
	return result
}

type Usuario struct {
	ID   int
	Nome string
}

func main() {
	db, err := sql.Open("mysql", "root:example@/cursogo")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	rows, _ := db.Query("select * from usuarios where id > ?", 5)
	defer rows.Close()

	for rows.Next() {
		var user Usuario
		rows.Scan(&user.ID, &user.Nome)
		fmt.Println(user)
	}
}
