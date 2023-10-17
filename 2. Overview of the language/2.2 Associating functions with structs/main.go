package main

import "log"

type User struct {
	FirstName string
}

func (m *User) getFirstName() string {
	return m.FirstName
}
func main() {
	var user1 User
	user1.FirstName = "Abhishek"

	user2 := User{
		FirstName: "Aditya",
	}

	log.Println("user 1 ", user1.getFirstName());
	log.Println("user 2 ", user2.getFirstName());

}