package main

import "log"

func main() {

	for i := 0; i <= 5; i++ {
		log.Println(i)
	}

	animals := []string{"dog", "fish", "cat", "horse", "cow"}

	for i, animal := range animals {
		log.Println(i, animal)
	}

	for _, animal := range animals {
		log.Println(animal)
	}

	animalsmap := make(map[string]string)
	animalsmap["dog"] = "Fluffy"
	animalsmap["cat"] = "Doja"

	for animalType, animal := range animalsmap {
		log.Println(animalType, animal)
	}

	type User struct {
		FirstName 	string
		LastName 	string
		Email 		string
		Age 		int
	}

	var users	[]User
	users = append(users, User{"Kaio", "Rangel", "kaiomontecristo1@gmail.com", 22})
	users = append(users, User{"Leticia", "Marins", "lelemarinslopes@gmail.com", 22})
	users = append(users, User{"Karol", "Rangel", "karolmontecristo1@gmail.com", 17})

	for _, l := range users {
		log.Println(l.FirstName, l.LastName, l.Email, l.Age)
	}
}