package main

import "fmt"

/*
Создать приложение, которое сначала выдаёт меню:

1. Посмотреть закладки
2. Добавить закладку
3. Удалить закладку
4. Выход

При 1 - Выводит закладки
При 2 - 2 поля ввода названия и адреса и после добавление
При 3 - Ввод названия и удаление по нему
При 4 - Завершение
*/

// * bk - bookmark

type bksMap = map[string]string

func main() {
	bkActions := map[int]func(bks bksMap){
		1: printBks,
		2: addBks,
		3: deleteBks,
	}

	bks := map[string]string{}

	fmt.Println("1. Посмотреть закладки")
	fmt.Println("2. Добавить закладку")
	fmt.Println("3. Удалить закладку")
	fmt.Println("4. Выход")

	for {
		userAction := getUserAction()

		if userAction == 4 {
			break
		}

		bkActions[userAction](bks)
	}
}

func getUserAction() int {
	var action int

	fmt.Print("Выберите действие: ")
	fmt.Scan(&action)

	return action
}

func printBks(bks bksMap) {
	if len(bks) == 0 {
		fmt.Println("Закладок нет")
		return
	}

	fmt.Println("\nВаши закладки:")
	for name, url := range bks {
		fmt.Printf("- %s: %s\n", name, url)
	}

	fmt.Println()
}

func addBks(bks bksMap) {
	var name string
	var url string

	fmt.Print("Введите название закладки: ")
	fmt.Scan(&name)

	fmt.Print("Введите адрес закладки: ")
	fmt.Scan(&url)

	bks[name] = url

	fmt.Println("Закладка добавлена!")
}

func deleteBks(bks bksMap) {
	var name string

	fmt.Print("Введите название закладки для удаления: ")
	fmt.Scan(&name)

	if _, exists := bks[name]; exists {
		delete(bks, name)
		fmt.Printf("\nЗакладка %s удалена!\n", name)
	} else {
		fmt.Printf("\nЗакладка %s не найдена\n", name)
	}
}