package main

import "log"

type User struct {
	FirstName string
	LastName  string
}
func main() {

	myMap := make(map[string]string)

	myMap["dog"] = "Samson"

	myMap["other-dog"] = "Cassie"

	myMap["dog"] = "Bilbo"

	log.Println(myMap["other-dog"])
	log.Println(myMap["dog"])

	myMap2 := make(map[string]int)

	myMap2["First"] = 1
	myMap2["Second"] = 2

	log.Println(myMap2["First"], myMap2["Second"])

	myMapt := make(map[string]User)

	me := User {
		FirstName: "Kaio",
		LastName: "Rangel",
	}

	myMapt["me"] = me
	
	log.Println((myMapt)["me"].FirstName)
}