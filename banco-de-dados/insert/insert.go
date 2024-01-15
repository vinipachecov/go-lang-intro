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
	db, err := sql.Open("mysql", "root:example@/")

	if err != nil {
		panic(err)
	}

	defer db.Close()
	stmt, _ := db.Prepare("insert into usuarios(nome) values(?)")
	stmt.Exec("Maria")
	// stmt.Exec("João")

	// res, _ := stmt.Exec("Pedro")

	// id, _ := res.LastInsertId()

	// fmt.Println(id)

	// linhas, _ := res.RowsAffected()

	// fmt.Println(linhas)

}
