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
	highScore uint
}

func connect_mysql() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:adminku123@tcp(localhost:3306)/latihan")
	if err != nil {
		fmt.Println("gagal connect ke database, Error:\n", err.Error())
		return nil, err
	}
	fmt.Println("berhasil connect ke database...")
	return db, nil
}

func getUsers(db *sql.DB, command string) ([]User, error) {
	if command == "" {
		command = "select * from users"
	}

	var results []User
	rows, err := db.Query(command)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		each := User{}
		err = rows.Scan(&each.id, &each.name, &each.email, &each.password, &each.highScore)
		if err != nil {
			return nil, err
		}
		results = append(results, each)
		if err = rows.Err(); err != nil {
			return nil, err
		}
	}
	return results, nil
}

func printAllUser(db *sql.DB) ([]User, error) {
	return getUsers(db, "select * from users")
}

func getUsersById(db *sql.DB, id string) ([]User, error) {
	command := fmt.Sprintf("select * from users where id = %s", id)
	return getUsers(db, command)
}

func insertUser(db *sql.DB) {

}

func main() {
	db, err := connect_mysql()
	defer db.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	allUser, err := printAllUser(db)
	fmt.Println("Users:", allUser)
}
