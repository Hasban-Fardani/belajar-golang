package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	id        string
	name      string
	email     string
	password  string
	highScore string
}

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

	rows, err := db.Query("select * from users")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var results []User
	for rows.Next() {
		var each User
		var err = rows.Scan(&each.id, &each.name, &each.email, &each.password, &each.highScore)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		results = append(results, each)
	}
	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(results)
}
