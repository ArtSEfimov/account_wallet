package main

import (
	"account_wallet/account"
	"account_wallet/encrypter"
	"account_wallet/files"
	"fmt"
	"github.com/joho/godotenv"
	"strings"
)

var menu = map[string]func(db *account.VaultWithDB){
	"1": createAccount,
	"2": findAccountByURL,
	"3": findAccountByLogin,
	"4": deleteAccount,
}

func main() {
	loadErr := godotenv.Load()
	if loadErr != nil {
		panic(".env file loading error")
	}

	fmt.Println("Starting password manager...")

	vault := account.NewVault(files.NewJSONdb("data.json"), *encrypter.NewEncrypter())

Menu:
	for {
		variant := promptData([]string{
			"1. Create account",
			"2. Find account by URL",
			"3. Find account by login",
			"4. Delete account",
			"5. Exit",
			"Enter variant",
		}...)
		menuFunc, ok := menu[variant]
		if !ok {
			break Menu
		}
		menuFunc(vault)

	}

}

func createAccount(vault *account.VaultWithDB) {
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	urlString := promptData("Введите URL")

	newAccount, createErr := account.NewAccount(login, password, urlString)
	if createErr != nil {
		panic(createErr)
	}

	vault.AddAccount(newAccount)

}

func findAccountByURL(vault *account.VaultWithDB) {
	url := promptData("Введите URL для поиска")
	accounts := vault.FindAccounts(url, func(a account.Account, s string) bool {
		return strings.Contains(a.Url, url)
	})
	if len(accounts) == 0 {
		fmt.Println("No accounts found")
		return
	}
	for _, acc := range accounts {
		acc.DisplayAccountInfo()
	}
}

func findAccountByLogin(vault *account.VaultWithDB) {
	login := promptData("Введите login для поиска")
	accounts := vault.FindAccounts(login, func(a account.Account, s string) bool {
		return strings.Contains(a.Login, login)
	})
	if len(accounts) == 0 {
		fmt.Println("No accounts found")
		return
	}
	for _, acc := range accounts {
		acc.DisplayAccountInfo()
	}
}

func deleteAccount(vault *account.VaultWithDB) {
	url := promptData("Введите URL для удаления")
	delResult := vault.DeleteAccountByUrl(url)
	if delResult {
		fmt.Println("Account deleted")
		return
	}
	fmt.Println("No found any accounts")
}

func promptData(prompt ...string) string {
	for i, line := range prompt {
		if i == len(prompt)-1 {
			fmt.Printf("%v: ", line)
		} else {
			fmt.Println(line)
		}

	}
	var res string
	_, scanErr := fmt.Scanln(&res)
	if scanErr != nil {
		return ""
	}
	return res
}
