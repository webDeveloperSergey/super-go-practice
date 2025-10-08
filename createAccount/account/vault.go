package account

import (
	"encoding/json"
	"time"
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

func (vault *Vault) AddAccounts(acc Account) {
	vault.Accounts = append(vault.Accounts, acc)
	vault.UpdatedAt = time.Now()
}

func NewVault() *Vault {
	return &Vault{
		Accounts: []Account{},
		UpdatedAt: time.Now(),
	}
}