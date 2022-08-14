package main

import (
	"fmt"
	"net/url"
)

func main() {
	var urlString = "http://kalipare.com:80/hello?name=john wick&age=27"
	var u, e = url.Parse(urlString)
	if e != nil {
		fmt.Println(e.Error())
		return
	}
	fmt.Println(u.Host)
	fmt.Println(u.Hostname())
	fmt.Println(u.Scheme)

	fmt.Println(u.Query())
	fmt.Println(u.Query()["name"])
}
