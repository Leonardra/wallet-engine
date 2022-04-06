package configs

import "math/rand"

func GenerateAccountNumber() string{
	var letter = []rune("0123456789")

	accountNumber := make([]rune, 10)
	for i := range accountNumber {
		accountNumber[i] = letter[rand.Intn(len(letter))]
	}
	return string(accountNumber)
}