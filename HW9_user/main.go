package main

import (
	"fmt"
	"github.com/bxcodec/faker/v3"
)

type user struct {
	Id        string `faker:"uuid_hyphenated"`
	Email     string `faker:"email"`
	FirstName string `faker:"oneof:  Sasha,Ben,John"`
	LastName  string `faker:"last_name"`
	Nick      string `faker:"first_name"`
	Age       int    `faker:"oneof: 15, 27, 61"`
}

func (u user) getNameData() string {
	str := fmt.Sprintf("User`s full name is %v and nick is %v",
		u.FirstName+" "+u.LastName, u.Nick)
	return str
}

func main() {
	Users := make([]user, 100, 100)

	for i, v := range Users {
		v = user{}
		err := faker.FakeData(&v)
		if err != nil {
			fmt.Println(err)
		}
		Users[i] = v
	}
	for i, v := range Users {
		if v.FirstName == "John" && v.Age > 20 {
			fmt.Println(Users[i].getNameData(), "\n")
		}
	}
}
