package hlp

import "fmt"

func main() {
	user := user{
		id:        "p0ZoB1FwH6",
		email:     "john.doe@example.com",
		firstName: "John",
		lastName:  "Doe",
		nick:      "johnnydoe",
		age:       42,
	}

	pu := &user

	fmt.Println(pu.email)
}

// Написати метод, що поверне повне імʼя Користувача. Як ви думаєте чому методом краще?
func (u user) fullName() string {
	return u.firstName + " " + u.lastName
}

// Написати метод, що збільшує вік Користувача на одиницю. Чому методом краще?
func (u *user) incrementAge() {
	u.age = u.age + 1
}

// Написати функцію, що збільшує вік Користувача на 1
func incrementAge(u *user) {
	u.age = u.age + 1
}

// Написати функцію, яка роздрукує короткі відомості про Користувача - нік та вік
func printUserData(u user) {
	fmt.Println(u.nick, u.age)
}

type user struct {
	id        string
	email     string
	firstName string
	lastName  string
	nick      string
	age       int

	pointerField *int
}

//func getUser(userID string) user {
//
//}
//
//func saveUser(user user) {
//
//}

//func getUser(userID string) (email, firstName, lastName, nick string, age int) {
//
//}
//
//func saveUser(email, firstName, lastName, nick string, age int) {
//
//}
