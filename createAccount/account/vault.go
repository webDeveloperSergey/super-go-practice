package account

import (
	"create-account/files"
	"encoding/json"
	"strings"
	"time"

	"github.com/fatih/color"
)

type Vault struct {
	Accounts []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updated_at"`
}

type VaultWithDb struct {
	Vault
	db files.DbJson
}

func (vault *Vault) ToBytesJson() ([]byte, error) {
	jsonData, err := json.Marshal(vault)
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}

func NewVault(db *files.DbJson) *VaultWithDb {
	file, err := db.Read()

	if err != nil {
		return &VaultWithDb{
			Vault: Vault{
				Accounts: []Account{},
				UpdatedAt: time.Now(),
			},
			db: *db,
		}
	}

	var vault Vault
	err = json.Unmarshal(file, &vault)

	if err != nil {
		color.Red("Не смогли разобрать файл data.json")

		return &VaultWithDb{
			Vault: Vault{
				Accounts: []Account{},
				UpdatedAt: time.Now(),
			},
			db: *db,
		}
	}

	return &VaultWithDb{
		Vault: vault,
		db: *db,
	}
}

func (vault *VaultWithDb) AddAccount(acc Account) {
	vault.Accounts = append(vault.Accounts, acc)
	vault.save()
}

func (vault *Vault) FindAccountByUrl(urlString string) []Account {
	var finedAccounts []Account

	for _, account := range vault.Accounts {
		isMatched := strings.Contains(account.Url, urlString)

		if isMatched {
			finedAccounts = append(finedAccounts, account)
		}
	}

	return finedAccounts
}

func (vault *VaultWithDb) DeleteAccountByUrl(urlString string) bool {
	isDeleted := false
	var notDeletedAcc []Account

		for _, account := range vault.Accounts {
		isMatched := strings.Contains(account.Url, urlString)

		if !isMatched {
			notDeletedAcc = append(notDeletedAcc, account)
			continue
		}

		isDeleted = true
	}

	vault.Accounts = notDeletedAcc
	vault.save()

	return isDeleted
}

func (vault *VaultWithDb) save() {
	vault.UpdatedAt = time.Now()

	dataJson, err := vault.Vault.ToBytesJson()
	if err != nil {
		color.Red("Не удалось преобразовать")
	}

	vault.db.Write(dataJson)
}
