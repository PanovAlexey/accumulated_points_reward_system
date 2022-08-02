package luhn

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"testing"
)

func Example_validate() {
	var correctOrderNumber int64 = 79927398713

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

	// Output:Order number 79927398713 is correct.
}

func Example_validateWrong() {
	var wrongOrderNumber int64 = 111222333

	luhnChecker := GetLuhnAlgorithmChecker()

	// process the wrong order number
	err := luhnChecker.Validate(wrongOrderNumber)
	fmt.Print("Order number ")
	fmt.Print(wrongOrderNumber)

	if err != nil {
		fmt.Print(" is wrong. ")
	} else {
		fmt.Print(" is correct. ")
	}

	// Output:Order number 111222333 is wrong.
}

func Test_Validate(t *testing.T) {
	tests := []struct {
		name  string
		value int64
		want  bool
	}{
		{
			name:  "Test by zero order number",
			value: 0,
			want:  true,
		},
		{
			name:  "Test by wrong order number №1",
			value: 1,
			want:  false,
		},
		{
			name:  "Test by wrong order number №2",
			value: 2147483647,
			want:  false,
		},
		{
			name:  "Test by correct order number №1",
			value: 5110000134567579,
			want:  false,
		},
		{
			name:  "Test by correct order number №2",
			value: 63900343901164149,
			want:  false,
		},
		{
			name:  "Test by correct order number №3",
			value: 4000001234567899,
			want:  true,
		},
		{
			name:  "Test by correct order number №4",
			value: 79927398713,
			want:  true,
		},
	}

	luhnChecker := GetLuhnAlgorithmChecker()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, luhnChecker.Validate(tt.value) == nil, tt.want)
		})
	}
}

func BenchmarkValidate(b *testing.B) {
	luhnChecker := GetLuhnAlgorithmChecker()

	startCheckingNumber := 10000
	stopCheckingNumber := 100000

	b.Run("optimized", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for j := startCheckingNumber; j < stopCheckingNumber; j++ {
				err := luhnChecker.ValidateOptimized(int64(j))

				if err != nil {
					log.SetOutput(ioutil.Discard)
					log.Println(err.Error())
				}
			}
		}
	})

	b.Run("original", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for j := startCheckingNumber; j < stopCheckingNumber; j++ {
				err := luhnChecker.Validate(int64(j))

				if err != nil {
					log.SetOutput(ioutil.Discard)
					log.Println(err.Error())
				}
			}
		}
	})
}
