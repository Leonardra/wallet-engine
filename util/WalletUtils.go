package util

import (
	"math/rand"
	"time"
)

func GenerateAccountNumber() string{
	var letter = []rune("0123456789")

	rand.Seed(time.Now().UnixNano())
	accountNumber := make([]rune, 10)
	for i := range accountNumber {
		accountNumber[i] = letter[rand.Intn(len(letter))]
	}
	return string(accountNumber)
}