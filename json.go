package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string
	Age   int
	Kelas string `json:"Class"`
}

func main() {
	var jsonString = `{"Name": "Hasban Fardani", "Age": 16, "Class": "X PPLG-1"}`
	var jsonData = []byte(jsonString)

	var Me User

	var err = json.Unmarshal(jsonData, &Me)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(Me)
	fmt.Println(Me.Name)
	fmt.Println(Me.Age)
	fmt.Println(Me.Kelas)

	var data map[string]interface{}
	jsonString1 := `{"Name": "Budi", "Age": 16, "Class": "X PPLG-1"}`

	err = json.Unmarshal([]byte(jsonString1), &data)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return
	}
	fmt.Println("\n", data)
	fmt.Println(data["Name"])
	fmt.Println(data["Age"])
	fmt.Println(data["Class"])

	Object := User{Name: "Reyga", Age: 16, Kelas: "X PPLG-1"}
	jsonData1, err := json.Marshal(Object)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	jsonString2 := string(jsonData1)
	fmt.Println("json string: ", jsonString2)

}
