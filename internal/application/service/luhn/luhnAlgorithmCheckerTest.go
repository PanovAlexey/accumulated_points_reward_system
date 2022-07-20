package luhn

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Validate(t *testing.T) {
	tests := []struct {
		name  string
		value int64
		want  bool
	}{
		{
			name:  "Test by zero order number",
			value: 0,
			want:  false,
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
			want:  true,
		},
		{
			name:  "Test by correct order number №2",
			value: 63900343901164149,
			want:  true,
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
