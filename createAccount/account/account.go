package account

import (
	"errors"
	"fmt"
	"math/rand"
	"net/url"
	"time"
)

type Account struct {
	login string
	pwd string
	url string
}

// Высосанный из пальца пример для понимания композиции в GO
type accountWithTimeStamp struct {
	createdAt time.Time
	updatedAt time.Time
	Account
}

var lettersRune = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!*-")

func (acc Account) OutputPassword() {
	fmt.Print(acc.login + " " + acc.pwd + " " + acc.url)
}

func (acc *Account) generatePassword(n int) {
	newPwd := make([]rune, n)

	for index := range newPwd {
		newPwd[index] = lettersRune[rand.Intn(len(lettersRune))]
	}

	acc.pwd = string(newPwd)
}

func NewAccountWithTimeStamp(login, pwd, urlString string) (*accountWithTimeStamp, error) {
	if 	login == "" {
		return nil, errors.New("INVALID_LOGIN")
	}

	_, err := url.ParseRequestURI(urlString)

	if err != nil {
		return nil, errors.New("INVALID_URL")
	}

	 newAccount := &accountWithTimeStamp {
		createdAt: time.Now(),
		updatedAt: time.Now(),
		Account: Account{
			login: login,
			pwd: pwd,
			url: urlString,
		},
	}

	if pwd == "" {
		newAccount.generatePassword(12)
	}

	return newAccount, nil
}

