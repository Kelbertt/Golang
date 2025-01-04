package main

import "log"

func main() {

	myMap := make(map[string]string)

	myMap["dog"] = "Samson"

	myMap["other-dog"] = "Cassie"

	myMap["dog"] = "Bilbo"

	log.Println(myMap["other-dog"])
	log.Println(myMap["dog"])
}