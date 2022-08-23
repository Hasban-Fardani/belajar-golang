package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type user struct {
	id        string
	name      string
	email     string
	password  string
	highScore uint
}

func connectDB() (*sql.DB, error) {
	url := "postgres://postgres:admin@localhost:5432/latihan"
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	fmt.Println("berhasil connect ke database")
	return db, nil
}

func get_user(db *sql.DB, command string) ([]user, error) {
	var results []user
	rows, err := db.Query(command)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		each := user{}
		// fmt.Println(rows.Columns())
		err = rows.Scan(&each.id, &each.name, &each.email, &each.password, &each.highScore)
		if err != nil {
			return nil, err
		}
		results = append(results, each)
	}
	return results, nil
}

func InsetToTable(tableName string, KV map[string]interface{}) {

}

func main() {
	db, err := connectDB()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()
	all_user, err := get_user(db, "select * from \"Users\";")
	fmt.Println(all_user, err)
}
