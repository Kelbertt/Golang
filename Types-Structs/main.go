package main

import (
	"log"
	"time"
)

var s = "seven"

type User struct {
	FirstName 	string
	LastName 	string
	PhoneNumber string
	Age 		int
	BirthDate 	time.Time
}
func main() {

	user := User {
		FirstName: "Kaio",
		LastName: "Rangel",
		PhoneNumber: "22997644900",
	}

	log.Println(user.FirstName, user.LastName, "Birthdate", user.BirthDate, user.PhoneNumber )
	

	var s2 = "five"

	log.Println("s is", s)
	log.Println("s2 is", s2)

	saySomething("xxx")

}

func saySomething(s string) (string, string) {
	log.Println("s from the func is", s)
	return s, "World"
}