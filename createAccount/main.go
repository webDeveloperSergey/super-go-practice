package main

import (
	"errors"
	"fmt"
	"math/rand"
	"net/url"
	"time"
)

type account struct {
	login string
	pwd string
	url string
}

// Высосанный из пальца пример для понимания композиции в GO
type accountWithTimeStamp struct {
	createdAt time.Time
	updatedAt time.Time
	account
}

func (acc account) outputPassword() {
	fmt.Print(acc.login + " " + acc.pwd + " " + acc.url)
}

func (acc *account) generatePassword(n int) {
	newPwd := make([]rune, n)

	for index := range newPwd {
		newPwd[index] = lettersRune[rand.Intn(len(lettersRune))]
	}

	acc.pwd = string(newPwd)
}

func newAccountWithTimeStamp(login, pwd, urlString string) (*accountWithTimeStamp, error) {
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
		account: account{
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

var lettersRune = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!*-")

func main() {

	login := getPromptData("Введите логин: ")
	pwd := getPromptData("Введите пароль: ")
	url := getPromptData("Введите URL: ")

	myAccount, err := newAccountWithTimeStamp(login, pwd, url)

	if err != nil {
		fmt.Print("Неверно введет URL или Логин")
		return
	}

	myAccount.generatePassword(12)
	myAccount.outputPassword()

	fmt.Print(myAccount)
}

func getPromptData(prompt string) string {
	var result string
	fmt.Print(prompt)
	fmt.Scanln(&result)

	return result
}