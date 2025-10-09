package account

import (
	"create-account/files"
	"encoding/json"
	"time"

	"github.com/fatih/color"
)

type Vault struct {
	Accounts []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (vault *Vault) ToBytesJson() ([]byte, error) {
	jsonData, err := json.Marshal(vault)
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}

func NewVault() *Vault {
	file, err := files.ReadFile("data.json")

	if err != nil {
		return &Vault{
			Accounts: []Account{},
			UpdatedAt: time.Now(),
		}
	}

	var vault Vault
	err = json.Unmarshal(file, &vault)

	if err != nil {
		color.Red("Не смогли разобрать файл data.json")

		return &Vault{
			Accounts: []Account{},
			UpdatedAt: time.Now(),
		}
	}

	return &vault
}

func (vault *Vault) AddAccounts(acc Account) {
	vault.Accounts = append(vault.Accounts, acc)
	vault.UpdatedAt = time.Now()

	dataJson, err := vault.ToBytesJson()
	if err != nil {
		color.Red("Не удалось преобразовать")
	}

	files.WriteFile(dataJson, "data.json")
}

