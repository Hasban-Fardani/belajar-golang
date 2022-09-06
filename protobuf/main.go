package main

import (
	model "belajar/protobuf/model"
	"encoding/json"
	"fmt"
	"reflect"
)

var user1 = &model.User{
	Id:       "u001",
	Name:     "Hasban",
	Password: "test 0n3 2 3",
	Gender:   model.UserGender_FEMALE,
}

var userList = &model.UserList{
	List: []*model.User{
		user1,
	},
}

var garage1 = &model.Garage{
	Id:   "g001",
	Name: "bengkel amanah",
	Coordinate: &model.GarageCoordinate{
		Latitude:   23.123123123,
		Longtitude: 51.12098731,
	},
}

var garageList = &model.GarageList{
	List: []*model.Garage{
		garage1,
	},
}

func main() {
	// var buf bytes.Buffer
	// err := (&jsonpb.Marshaler{}).Marshal(&buf, garageList)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	// fmt.Println(buf.String())
	// fmt.Println(r)

	fmt.Println(user1, " type:", reflect.TypeOf(user1))
	fmt.Println(user1.String(), " type:", reflect.TypeOf(user1.String()))
	fmt.Println(userList)
	result, err := json.Marshal(user1)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(result)
}
