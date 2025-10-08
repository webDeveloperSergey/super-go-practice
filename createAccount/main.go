package main

import (
	"create-account/account"
	"create-account/files"
	"fmt"
)



func main() {
	userActions := map[int]func() {
		1: createAccount,
		2: findAccount,
		3: deleteAccount,
	}

	fmt.Println("__Менеджер паролей__")

	for {
		action, err := getMenu()

		if err != nil || action > 3 {
			return
		}

		userActions[action]()
	}
	
}

func getMenu() (int, error) {
	var action int

	fmt.Println("Выберите вариант:")
	fmt.Println("1. Создать аккаунт")
	fmt.Println("2. Найти аккаунт")
	fmt.Println("3. Найти аккаунт")
	fmt.Println("4. Выйти")

	_, err := fmt.Scan(&action)

	return action, err
}

func createAccount() {
	login := getPromptData("Введите логин: ")
	pwd := getPromptData("Введите пароль: ")
	url := getPromptData("Введите URL: ")

	myAccount, err := account.NewAccount(login, pwd, url)
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

func findAccount() {}
func deleteAccount() {}

func getPromptData(prompt string) string {
	var result string
	fmt.Print(prompt)
	fmt.Scanln(&result)

	return result
}