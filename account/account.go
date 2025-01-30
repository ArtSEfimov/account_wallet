package account

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"
	"unicode"
)

var passwordRunes = []rune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz!@#$%^&-_?")

type Account struct {
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"account_created_at"`
	UpdatedAt time.Time `json:"account_updated_at"`
}

func NewAccount(login string, password string, urlString string) (*Account, error) {
	if login == "" {
		return nil, errors.New("login required")
	}
	_, urlErr := url.ParseRequestURI(urlString)
	if urlErr != nil {
		return nil, fmt.Errorf("invalid url: %w", urlErr)
	}

	acc := Account{
		Login:     login,
		Url:       urlString,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if !validatePassword(password) {
		acc.generatePassword(10)
	} else {
		acc.Password = password
	}

	return &acc, nil
}

func (a *Account) DisplayAccountInfo() {
	fmt.Println(a.Login, a.Password, a.Url)
}

func (a *Account) generatePassword(length int) {
	password := make([]rune, length)

	for {
		for i := range password {
			password[i] = passwordRunes[rand.IntN(len(passwordRunes))]
		}

		if validatePassword(string(password)) {
			a.Password = string(password)
			return
		}
	}
}

func validatePassword(password string) bool {
	if len(password) == 0 {
		return false
	}

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
	if lettersCount+digitsCount == uint8(len(password)) {
		return false
	}

	return true

}
