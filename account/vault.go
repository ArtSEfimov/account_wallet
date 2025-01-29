package account

import (
	"account_wallet/files"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (v *Vault) ConvertToBytes() ([]byte, error) {
	data, convertErr := json.Marshal(v)
	if convertErr != nil {
		return nil, fmt.Errorf("error converting to bytes: %w", convertErr)
	}

	return data, nil
}

func NewVault() *Vault {

	data, readErr := files.ReadFile("data.json")
	if readErr != nil {
		return &Vault{
			Accounts:  []Account{},
			UpdatedAt: time.Now(),
		}
	}
	var vault Vault

	err := json.Unmarshal(data, &vault)
	if err != nil {
		fmt.Printf("parsing error: %v", err)
		return &Vault{
			Accounts:  []Account{},
			UpdatedAt: time.Now(),
		}
	}
	return &vault

}

func (v *Vault) AddAccount(account *Account) {
	v.Accounts = append(v.Accounts, *account)
	v.save()

}

func (v *Vault) FindAccountsByUrl(url string) []Account {
	var accounts []Account
	for _, account := range v.Accounts {
		if strings.Contains(account.Url, url) {
			accounts = append(accounts, account)
		}
	}
	return accounts
}
func (v *Vault) DeleteAccountByUrl(url string) bool {
	var accounts []Account
	for _, account := range v.Accounts {
		if !strings.Contains(account.Url, url) {
			accounts = append(accounts, account)
		}
	}
	if len(accounts) == len(v.Accounts) {
		return false
	}
	v.Accounts = accounts

	v.save()

	return true

}
func (v *Vault) save() {
	v.UpdatedAt = time.Now()
	data, convertErr := v.ConvertToBytes()
	if convertErr != nil {
		panic(convertErr)
	}
	writeErr := files.WriteFile("data.json", data)
	if writeErr != nil {
		panic(writeErr)
	}
}
