package luhn

import (
	"fmt"
)

func Example() {
	var correctOrderNumber, wrongOrderNumber int64

	correctOrderNumber = 79927398713
	wrongOrderNumber = 111222333

	luhnChecker := GetLuhnAlgorithmChecker()

	// process the correct order number
	err := luhnChecker.Validate(correctOrderNumber)
	fmt.Print("Order number ")
	fmt.Print(correctOrderNumber)

	if err != nil {
		fmt.Print(" is wrong. ")
	} else {
		fmt.Print(" is correct. ")
	}

	// process the wrong order number
	err = luhnChecker.Validate(wrongOrderNumber)
	fmt.Print("Order number ")
	fmt.Print(wrongOrderNumber)

	if err != nil {
		fmt.Print(" is wrong. ")
	} else {
		fmt.Print(" is correct. ")
	}

	// Output:Order number 79927398713 is correct. Order number 111222333 is wrong.
}
