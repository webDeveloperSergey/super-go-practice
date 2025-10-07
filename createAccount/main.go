package main

import (
	"create-account/account"
	"fmt"
)



func main() {

	login := getPromptData("Введите логин: ")
	pwd := getPromptData("Введите пароль: ")
	url := getPromptData("Введите URL: ")

	myAccount, err := account.NewAccountWithTimeStamp(login, pwd, url)

	if err != nil {
		fmt.Print("Неверно введет URL или Логин")
		return
	}

	myAccount.OutputPassword()
	fmt.Print(myAccount)
}

func getPromptData(prompt string) string {
	var result string
	fmt.Print(prompt)
	fmt.Scanln(&result)

	return result
}