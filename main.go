package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (user User) ToMapString() (userMap map[string]interface{}, err error) {
	data, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &userMap)
	return
}

func main() {

	var myUser User
	myUser.Name = "Dewok Satria"
	myUser.Email = "dewok.satria@gmail.com"

	userMap, err := myUser.ToMapString()
	if err != nil {
		fmt.Printf("Err: %v\n", err)
	}

	fmt.Printf("User Struct: %v\n", myUser)
	fmt.Printf("User Map String: %v\n", userMap)
}
