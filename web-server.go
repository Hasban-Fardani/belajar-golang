package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type test struct {
	name string
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World from go server")
	fmt.Println("method:", r.Method, "from:", r.RemoteAddr)

}
func main() {
	http.HandleFunc("/", index)

	// penggunaan func secara asynchronous + template
	http.HandleFunc("/hasban", func(w http.ResponseWriter, r *http.Request) {
		data := map[string]string{
			"Nama":  "hasban",
			"Kelas": "X PPLG",
		}
		tem, err := template.ParseFiles("html_template/person.html")
		if err != nil {
			fmt.Println("error: \n", err.Error())
		}

		tem.Execute(w, data)
	})

	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
