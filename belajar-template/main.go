package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
)

func main() {
	var ADDR = "localhost:8080"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		filePath := path.Join("html", "index.html")
		tmpl, err := template.ParseFiles(filePath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			// return
		}
		data := map[string]interface{}{
			"nama":            "Hasban Fardani",
			"jurusan":         "PPLG",
			"kelas":           "X-1",
			"ekstrakulikuler": []string{"KDA", "Voli"},
		}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			// return
		}
	})

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("assets"))))

	fmt.Printf("server started at http://%s\n", ADDR)
	http.ListenAndServe(ADDR, nil)
}
