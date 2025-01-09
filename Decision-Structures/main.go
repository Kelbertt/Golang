package main

import "log"

func main()  {
	
	var isTrue bool
	isTrue = false

	if isTrue == true {
		log.Println("isTrue is", isTrue)
	} else {
		log.Println("isTrue is", isTrue)
	}

	cat := "cat2"

	if cat == "cat"{
		log.Println("cat is", cat)
	} else {
		log.Println("cat is not cat")
	}

	
	myNum := 100
	isFalse := false

	if myNum > 99 && !isFalse {
		log.Println("myNum is greater than 99 and isFalse is set to true")
	}

}