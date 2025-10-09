package main

import (
	"create-account/account"
	"fmt"

	"github.com/fatih/color"
)



func main() {


	fmt.Println("__Менеджер паролей__")
	vault := account.NewVault()

	userActions := map[int]func(vault *account.Vault) {
		1: createAccount,
		2: findAccount,
		3: deleteAccount,
	}

	for {
		action, err := getMenu()

		if err != nil || action > 3 {
			return
		}

		userActions[action](vault)
	}
	
}

func getMenu() (int, error) {
	var action int

	fmt.Println("Выберите вариант:")
	fmt.Println("1. Создать аккаунт")
	fmt.Println("2. Найти аккаунт")
	fmt.Println("3. Удалить аккаунт")
	fmt.Println("4. Выйти")

	_, err := fmt.Scan(&action)

	return action, err
}

func createAccount(vault *account.Vault) {
	login := getPromptData("Введите логин: ")
	pwd := getPromptData("Введите пароль: ")
	url := getPromptData("Введите URL: ")

	myAccount, err := account.NewAccount(login, pwd, url)
	if err != nil {
		fmt.Print("Неверно введет URL или Логин")
		return
	}

	vault.AddAccount(*myAccount)
}

func findAccount(vault *account.Vault) {
	url := getPromptData("Введите URL для поиска: ")

	fmt.Println("")
	fmt.Println("Результаты поиска:")

	accounts := vault.FindAccountByUrl(url)
	if len(accounts) == 0 {
		color.Red("Акакаунтов не найдено")
		return
	}

	for index, account := range accounts {
		account.Output(index)
		fmt.Println("-----")
	}
}
func deleteAccount(vault *account.Vault) {
	url := getPromptData("Введите URL для удаления аккаунта: ")

	fmt.Println("")

	isDeleted := vault.DeleteAccountByUrl(url)

	if isDeleted {
		color.Green("Запись успешно удалена")
	} else {
		color.Red("Не удалось удалить или найти запись")
	}
}

func getPromptData(prompt string) string {
	var result string
	fmt.Print(prompt)
	fmt.Scanln(&result)

	return result
}