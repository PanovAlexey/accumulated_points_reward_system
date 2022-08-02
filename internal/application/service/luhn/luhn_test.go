package luhn

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
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
