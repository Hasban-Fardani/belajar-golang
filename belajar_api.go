package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

var (
	db  *sqlx.DB
	err error
	URL string
)

type users struct {
	Id        string `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	Email     string `json:"email" db:"email"`
	Password  string `json:"password" db:"password"`
	HighScore int    `json:"high_score" db:"high_score"`
}

func (u users) JsonString() string {
	jsn, err := json.Marshal(u)
	if err != nil {
		return string(jsn)
	} else {
		return ""
	}
}

type dict map[string]interface{}

func checkError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
		return true
	} else {
		return false
	}
}

func StopError(code int, msg string) {
	fmt.Println(msg)
	os.Exit(code)
}

func connect_db() (*sqlx.DB, error) {
	url := "postgres://postgres:admin@localhost:5432/latihan"
	db, err := sqlx.Connect("postgres", url)
	if err != nil {
		return nil, err
	}
	fmt.Println("berhasil connect ke database")
	return db, nil
}

func ReadBody(r *http.Request) (string, error) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return "", err
	}
	return string(bodyBytes), nil
}

func StringToJson(str string, v dict) (dict, error) {
	byteData := []byte(str)
	err := json.Unmarshal(byteData, &v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func get_user_db(command string) ([]users, error) {
	rows, err := db.Query(command)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var results []users
	for rows.Next() {
		each := users{}
		err = rows.Scan(&each.Id, &each.Name, &each.Email, &each.Password, &each.HighScore)
		if err != nil {
			return nil, err
		}
		results = append(results, each)
	}
	return results, nil
}

func main() {
	URL := "localhost:1313"
	db, err = connect_db()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello user"))
	})

	http.HandleFunc("/all", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			res, err := get_user_db("select * from \"User\"")
			if err != nil {
				fmt.Println(err.Error())
			}
			result, err := json.Marshal(res)
			w.Write(result)
		} else {
			fmt.Fprintln(w, "method not allowed")
		}
	})

	http.HandleFunc("/user/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method == "GET" {
			path := strings.Split(r.URL.Path, "/")
			l_path := len(path)
			if l_path > 3 {
				fmt.Fprintln(w, "incorect url")
				fmt.Fprintf(w, "the correct is %s/user/<name>\n", r.Host)
				return
			}
			fmt.Println(path, r.URL.Path)
			// name := path[l_path-1]
			// cmd := fmt.Sprintf("select * from \"User\" where name = '%s'", name)
			// user_, err := get_user_db(cmd)
			// if err != nil {
			// 	fmt.Fprintln(w, "error", err.Error())
			// 	return
			// }
			// jsn, err := json.Marshal(user_[0])
			// fmt.Fprintln(w, string(jsn))
		}
	})

	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		log.SetOutput(w)
		if r.Method == "POST" {
			var data users
			dataStr, err := ReadBody(r)
			if checkError(err) {
				log.Panicln(err.Error())
				return
			}
			err = json.Unmarshal([]byte(dataStr), &data)
			if checkError(err) {
				log.Println(err.Error())
				return
			}

			query := `INSERT INTO "User" VALUES (:id, :name, :email, :password, :high_score)`
			if _, err = db.NamedExec(query, data); checkError(err) {
				log.Println(err.Error())
				return
			}
			log.Println("succes add user", data.Name)
		} else {
			log.Println("method not allowed")
		}
	})
	http.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request) {
		var user users
		// var result sql.Result
		if r.Method == "DELETE" {
			dataStr, err := ReadBody(r)
			if checkError(err) {
				log.Println(err.Error())
			}
			err = json.Unmarshal([]byte(dataStr), &user)
			if checkError(err) {
				log.Println(err.Error())
				return
			}
			query := `DELETE FROM "User" WHERE id=:id`
			if _, err = db.NamedExec(query, user); checkError(err) {
				log.Println(err.Error())
				return
			}
			fmt.Println("lewat sini")
			log.Println("succes deleted", user.Id)
		} else {
			log.Println("method not allowed")
			return
		}
		// log.Println(result)
	})

	fmt.Printf("serve on http://%s \n", URL)
	http.ListenAndServe(URL, nil)
}
