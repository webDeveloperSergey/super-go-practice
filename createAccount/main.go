package main

import (
	"create-account/account"
	"create-account/files"
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

	jsonData, err := myAccount.ToBytesJson()
	if err != nil {
		fmt.Print("Не удалось преобразовать данные в JSON")
		return
	}

	files.WriteFile(jsonData, "data.json")

}

func getPromptData(prompt string) string {
	var result string
	fmt.Print(prompt)
	fmt.Scanln(&result)

	return result
}