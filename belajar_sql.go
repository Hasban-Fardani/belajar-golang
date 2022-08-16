package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func connect_mysql() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:adminku123@tcp(localhost:3306)/latihan")
	if err != nil {
		return nil, err
	}
	fmt.Println("berhasil connect ke database...")
	return db, nil
}

func main() {
	db, err := connect_mysql()
	defer db.Close()

	if err != nil {
		fmt.Println(err)
		return
	}
}
