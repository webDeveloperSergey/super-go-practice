package main

import (
	"create-account/account"
	"create-account/encrypter"
	"create-account/files"
	"create-account/output"
	"fmt"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

var menuActions = []string{
	"1. Создать аккаунт",
	"2. Найти аккаунт",
	"3. Удалить аккаунт",
	"4. Выйти",
	"Выберете вариант",
}

func main() {
	err := godotenv.Load()
  if err != nil {
    output.PrintError("Не удалось загрузить файл .env")
  }

	
	fmt.Println("__Менеджер паролей__")
	vault := account.NewVault(files.NewDbJson("data.vault"), *encrypter.NewEncrypter())
	// vault := account.NewVault(cloud.NewDbCloud("https://test.ru"))

	userActions := map[string]func(vault *account.VaultWithDb) {
		"1": createAccount,
		"2": findAccount,
		"3": deleteAccount,
	}

Menu:	
	for {
		action := getPromptData(menuActions...)

		currentAction := userActions[action]

		if currentAction == nil {
			break Menu
		}

		currentAction(vault)
	}
	
}

func createAccount(vault *account.VaultWithDb) {
	login := getPromptData("Введите логин ")
	pwd := getPromptData("Введите пароль ")
	url := getPromptData("Введите URL ")

	myAccount, err := account.NewAccount(login, pwd, url)
	if err != nil {
		output.PrintError("Неверно введет URL или Логин")
		return
	}

	vault.AddAccount(*myAccount)
}

func findAccount(vault *account.VaultWithDb) {
	url := getPromptData("Введите URL для поиска: ")

	fmt.Println("")
	fmt.Println("Результаты поиска:")

	accounts := vault.FindAccountByUrl(url)
	if len(accounts) == 0 {
		output.PrintError("Аккаунтов не найдено")
		return
	}

	for index, account := range accounts {
		account.Output(index)
		fmt.Println("-----")
	}
}
func deleteAccount(vault *account.VaultWithDb) {
	url := getPromptData("Введите URL для удаления аккаунта: ")

	fmt.Println("")

	isDeleted := vault.DeleteAccountByUrl(url)

	if isDeleted {
		color.Green("Запись успешно удалена")
	} else {
		output.PrintError("Не удалось удалить или найти запись")
	}
}

func getPromptData(prompt ...string) string {
	for i, line := range prompt {
		if i == len(prompt)-1 {
			fmt.Printf("%v: ", line)
		} else {
			fmt.Println(line)
		}
	}
	
	var result string
	fmt.Scanln(&result)

	return result
}