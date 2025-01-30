package account

import (
	"account_wallet/encrypter"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type DB interface {
	Read() ([]byte, error)
	Write([]byte) error
}

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"vault_updated_at"`
}

type VaultWithDB struct {
	Vault     Vault
	db        DB
	encrypter encrypter.Encrypter
}

func NewVault(db DB, encrypter encrypter.Encrypter) *VaultWithDB {
	data, readErr := db.Read()
	if readErr != nil {
		return &VaultWithDB{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db:        db,
			encrypter: encrypter,
		}
	}
	var vault Vault

	err := json.Unmarshal(data, &vault)
	if err != nil {
		fmt.Printf("parsing error: %v", err)
		return &VaultWithDB{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db:        db,
			encrypter: encrypter,
		}
	}
	return &VaultWithDB{
		Vault:     vault,
		db:        db,
		encrypter: encrypter,
	}

}

func (v *Vault) ConvertToBytes() ([]byte, error) {
	data, convertErr := json.Marshal(v)
	if convertErr != nil {
		return nil, fmt.Errorf("error converting to bytes: %w", convertErr)
	}

	return data, nil
}

func (v *VaultWithDB) AddAccount(account *Account) {
	v.Vault.Accounts = append(v.Vault.Accounts, *account)
	v.save()

}

func (v *VaultWithDB) FindAccounts(token string, checker func(Account, string) bool) []Account {
	var accounts []Account
	for _, account := range v.Vault.Accounts {
		if checker(account, token) {
			accounts = append(accounts, account)
		}
	}
	return accounts
}

func (v *VaultWithDB) DeleteAccountByUrl(url string) bool {
	var accounts []Account
	for _, account := range v.Vault.Accounts {
		if !strings.Contains(account.Url, url) {
			accounts = append(accounts, account)
		}
	}
	if len(accounts) == len(v.Vault.Accounts) {
		return false
	}
	v.Vault.Accounts = accounts

	v.save()

	return true

}
func (v *VaultWithDB) save() {
	v.Vault.UpdatedAt = time.Now()
	data, convertErr := v.Vault.ConvertToBytes()
	if convertErr != nil {
		panic(convertErr)
	}

	writeErr := v.db.Write(data)
	if writeErr != nil {
		panic(writeErr)
	}
}
