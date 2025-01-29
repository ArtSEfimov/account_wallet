package main

import (
	"account_wallet/account"
	"fmt"
)

func main() {

	fmt.Println("Starting password manager...")

	vault := account.NewVault()
Menu:
	for {
		variant := getMenu()
		switch variant {
		case 1:
			createAccount(vault)
		case 2:
			findAccount(vault)
		case 3:
			deleteAccount(vault)
		default:
			break Menu
		}
	}

}
func getMenu() uint8 {
	var variant uint8
	fmt.Println("Enter variant: ")
	fmt.Println("1. Create account")
	fmt.Println("2. Find account")
	fmt.Println("3. Delete account")
	fmt.Println("4. Exit")
	_, scanErr := fmt.Scan(&variant)
	if scanErr != nil {
		panic(scanErr)
	}

	return variant
}

func createAccount(vault *account.Vault) {
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	urlString := promptData("Введите URL")

	newAccount, createErr := account.NewAccount(login, password, urlString)
	if createErr != nil {
		panic(createErr)
	}

	vault.AddAccount(newAccount)

}
func findAccount(vault *account.Vault) {
	url := promptData("Введите URL для поиска")
	accounts := vault.FindAccountsByUrl(url)
	if len(accounts) == 0 {
		fmt.Println("No accounts found")
		return
	}
	for _, acc := range accounts {
		acc.DisplayAccountInfo()
	}
}

func deleteAccount(vault *account.Vault) {
	url := promptData("Введите URL для удаления")
	delResult := vault.DeleteAccountByUrl(url)
	if delResult {
		fmt.Println("Account deleted")
		return
	}
	fmt.Println("No found any accounts")
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
