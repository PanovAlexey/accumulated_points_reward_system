package luhn

import (
	"log"
)

func Example() {
	var correctOrderNumber, wrongOrderNumber int64

	correctOrderNumber = 79927398713
	wrongOrderNumber = 111222333

	luhnChecker := GetLuhnAlgorithmChecker()

	// process the correct order number
	err := luhnChecker.Validate(correctOrderNumber)
	log.Print("order number ")
	log.Print(correctOrderNumber)

	if err != nil {
		log.Print(" is wrong. ")
	} else {
		log.Print(" is correct. ")
	}

	// Output: ["order number 79927398713 is correct. "]

	log.Println("")

	// process the wrong order number
	err = luhnChecker.Validate(wrongOrderNumber)
	log.Print("order number ")
	log.Print(wrongOrderNumber)

	if err != nil {
		log.Println(" is wrong. ")
	} else {
		log.Println(" is correct. ")
	}

	// Output: ["order number 111222333 is wrong. "]
}
