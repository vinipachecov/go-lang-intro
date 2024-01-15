package main

import (
	"database/sql"
	"fmt"
	"log"

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
	tx, err := db.Begin()
	if err != nil {
		fmt.Println("Could not open db")
		log.Fatal()
	}

	stmt, err := tx.Prepare("insert into usuarios(id, nome) values(?, ?)")

	if err != nil {
		fmt.Println("Could not prepare")
		log.Fatal()
	}

	stmt.Exec(2000, "Rodrigo")
	stmt.Exec(2001, "Leandro")

	_, err = stmt.Exec(1, "Tiago")

	if err != nil {
		tx.Rollback()
		log.Fatal()
	}

	fmt.Println("Success")

	tx.Commit()
}
