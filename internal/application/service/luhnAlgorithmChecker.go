package service

import "errors"

type luhnAlgorithmChecker struct {
}

func GetLuhnAlgorithmChecker() luhnAlgorithmChecker {
	return luhnAlgorithmChecker{}
}

func (checker luhnAlgorithmChecker) Validate(number int) error {
	if (number%10+checksum(number/10))%10 != 0 {
		return errors.New("the number does not satisfy Luhn's algorithm")
	}

	return nil
}

func checksum(number int) int {
	var luhn int

	for i := 0; number > 0; i++ {
		cur := number % 10

		if i%2 == 0 {
			cur = cur * 2
			if cur > 9 {
				cur = cur%10 + cur/10
			}
		}

		luhn += cur
		number = number / 10
	}
	return luhn % 10
}
