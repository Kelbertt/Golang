package main

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
func main() {

}