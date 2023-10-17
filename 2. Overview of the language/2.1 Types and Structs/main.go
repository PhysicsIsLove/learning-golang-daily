package main

import (
	"log"
	"time"
)

type User struct {
	FirstName   string
	LastName    string
	Age         int
	PhoneNumber string
	dob			time.Time
}

func main() {

	log.Println("yoyoy");

	user := User {
		FirstName: "Abhishek",
		LastName: "Anand",
		dob: time.Now(),
	}

	log.Println("user information ", user.FirstName, user.LastName, user.Age, user.dob);
}