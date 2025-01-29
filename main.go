package main

import (
	"account_wallet/account"
	"account_wallet/files"
	"fmt"
)

func main() {
	writeErr := files.WriteFile("test.txt", "New_file")
	if writeErr != nil {
		panic(writeErr)
	}

	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	urlString := promptData("Введите URL")

	myAccount, createErr := account.NewAccount(login, password, urlString)
	if createErr != nil {
		panic(createErr)
	}

	myAccount.DisplayAccountInfo()

}

func promptData(prompt string) string {
	fmt.Printf("%s: ", prompt)
	var res string
	_, scanErr := fmt.Scanln(&res)
	if scanErr != nil {
		return ""
	}
	return res
}
