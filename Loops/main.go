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
}