package palindrome_number

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func IsPalindromeNumber(x int) bool {
	if x < 0 {
		return false
	}

	number := x
	reverse := 0

	for number > 0 {
		temp := number % 10
		reverse = reverse*10 + temp
		number = number / 10
	}

	return reverse == x
}

func TestIsPalindromeNumber(t *testing.T) {
	t.Run("is palindrome number", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5, 7, 8, 9, 0, 11, 22, 33, 44, 55, 77, 121, 888, 999, 1001, 2002, 3003, 4004, 5005, 7007, 8008, 9009}

		for _, v := range input {
			got := IsPalindromeNumber(v)
			assert.Equal(t, true, got)
		}
	})

	t.Run("is palindrome not number", func(t *testing.T) {
		input := []int{-112, 112, 211, 301, 102, 778}

		for _, v := range input {
			got := IsPalindromeNumber(v)
			assert.Equal(t, false, got)
		}
	})
}
