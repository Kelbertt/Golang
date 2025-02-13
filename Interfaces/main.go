package main

import "fmt"

type Animal interface {
	Says() 	string
	NumberofLegs() int
}

type Dog struct {
	Name 	string
	Breed 	string
}

type Gorilla struct {
	Name 	string
	Color 	string
}

func PrintInfo( a Animal) {
	fmt.Println("This animal says", a.Says(), "and has", a.NumberofLegs(), "legs")
}

func (d *Dog) Says() string {
	return "Woof"
}

func (d *Dog) NumberofLegs() int {
	return 4
}


func main() {
	dog := Dog{
		Name:  "Buddy",
		Breed: "German Shephered",
	}

	PrintInfo(&dog)
}