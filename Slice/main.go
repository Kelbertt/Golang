package main

import (
	"log"
	
)

func main() {
	var mySlice []string

	mySlice = append(mySlice, "Kaio")
	mySlice = append(mySlice, "Lucas")
	mySlice = append(mySlice, "Victor")

	log.Println(mySlice)

	var mySlice2 []int

	mySlice2 = append(mySlice2, 2)
	mySlice2 = append(mySlice2, 1)
	mySlice2 = append(mySlice2, 3)

	log.Println(mySlice2)
}