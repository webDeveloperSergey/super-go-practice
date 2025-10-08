package account

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/url"
	"time"
)

type Account struct {
	Login string `json:"login"`
	Pwd string `json:"pwd"`
	Url string `json:"url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

var lettersRune = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!*-")

func (acc Account) OutputPassword() {
	fmt.Print(acc.Login + " " + acc.Pwd + " " + acc.Url)
}

func (acc *Account) generatePassword(n int) {
	newPwd := make([]rune, n)

	for index := range newPwd {
		newPwd[index] = lettersRune[rand.Intn(len(lettersRune))]
	}

	acc.Pwd = string(newPwd)
}

func (acc *Account) ToBytesJson() ([]byte, error) {
	jsonData, err := json.Marshal(acc)
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}

func NewAccount(login, pwd, urlString string) (*Account, error) {
	if 	login == "" {
		return nil, errors.New("INVALID_LOGIN")
	}

	_, err := url.ParseRequestURI(urlString)

	if err != nil {
		return nil, errors.New("INVALID_URL")
	}

	 newAccount := &Account{
		Login: login,
		Pwd: pwd,
		Url: urlString,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if pwd == "" {
		newAccount.generatePassword(12)
	}

	return newAccount, nil
}

