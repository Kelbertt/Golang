package main

import "log"

func main() {

	myNum := 100
	isFalse := false

	if myNum > 99 && isFalse == false{
		log.Println("myNum is greater than 99 and isFalse is set to true")
	}
}