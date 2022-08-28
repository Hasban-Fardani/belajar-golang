package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func routeIndexGet(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	if r.Method == "GET" {
		var tmpl = template.Must(template.New("form").ParseFiles("form.html"))
		var err = tmpl.Execute(w, nil)
		if err != nil {
			fmt.Println(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		fmt.Println("error at routeIndex method:", r.Method)
	}
	// http.Error(w, "", http.StatusBadRequest)
}

func routeSubmitPost(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	if r.Method == "POST" {
		var tmpl = template.Must(template.New("result").ParseFiles("result.html"))
		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var nama = r.FormValue("nama")
		var kelas = r.FormValue("kelas")

		var data = map[string]string{"nama": nama, "kelas": kelas}
		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		fmt.Println("error at routeSubmit method:", r.Method)
	}
	// http.Error(w, "", http.StatusBadRequest)
}

func main() {
	var URL = "localhost:8080"
	http.HandleFunc("/", routeIndexGet)
	http.HandleFunc("/form", routeSubmitPost)

	fmt.Printf("server started at http://%s \n", URL)
	http.ListenAndServe(URL, nil)
}
