package main

import (
	"fmt"
	"math/rand/v2"
	"unicode"
)

var passwordRunes = []rune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz!@#$%^&-_<>?/")

type account struct {
	login    string
	password string
	url      string
}

func main() {
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")

	myAccount := account{
		login:    login,
		password: password,
		url:      url,
	}

	outputData(&myAccount)

	p := generatePassword(10)
	fmt.Print(p)
}

func promptData(prompt string) string {
	fmt.Printf("%s: ", prompt)
	var res string
	fmt.Scan(&res)
	return res
}

func outputData(acc *account) {
	fmt.Println(acc.login, acc.password, acc.url)
}

func generatePassword(length int) string {
	password := make([]rune, length)

	for {
		for i := range password {
			password[i] = passwordRunes[rand.IntN(len(passwordRunes))]
		}

		if validatePassword(string(password)) {
			return string(password)
		}
	}

}

func validatePassword(password string) bool {

	var lettersCount uint8 = 0
	var digitsCount uint8 = 0

	for i, char := range password {
		if i == 0 {
			if !unicode.IsLetter(char) && !unicode.IsDigit(char) {
				return false
			}
		}
		if unicode.IsLetter(char) {
			lettersCount++
		}
		if unicode.IsDigit(char) {
			digitsCount++
		}
	}

	if digitsCount == 0 || lettersCount == 0 {
		return false
	}

	if lettersCount+digitsCount < uint8(len(password)/2) {
		return false
	}

	return true

}
