package main

import (
	"errors"
	"fmt"
	"math"
)

const IMTPower = 2


func main() {
	for {
		userHeight, userWeight := getUserInfo()
		IMT, err := calculateIMT(userHeight, userWeight)

		if err != nil {
			fmt.Println("Введены неверные данные")
			continue
		}

		outputIMT(IMT)

		if !isContinueScript() {
			break
		}
	}
	
}

func outputIMT(imt float64) {
	result := fmt.Sprintf("Your IMT is: %.0f", imt)
	fmt.Println(result)

	switch {
		case imt < 16:
			fmt.Println("Выраженный дефицит массы тела")
		case imt >= 16 && imt < 18.5:
			fmt.Println("Недостаточная масса тела")
		case imt >= 18.5 && imt < 25:
			fmt.Println("Нормальная масса тела")
		case imt >= 25 && imt < 30:
			fmt.Println("Избыточная масса тела (предожирение)")
		case imt >= 30 && imt < 35:
			fmt.Println("Ожирение 1 степени")
		case imt >= 35 && imt < 40:
			fmt.Println("Ожирение 2 степени")
		case imt >= 40:
			fmt.Println("Ожирение 3 степени")
	}
}


func calculateIMT(height, weight float64) (float64, error) {
	if height <= 0 || weight <= 0 {
		return 0, errors.New("NO_CORRECT_PARAMS")
	}

  IMT := weight / math.Pow(height / 100, IMTPower)

	return IMT, nil
}

func getUserInfo() (float64, float64) {
	var userHeight float64
	var userWeight float64

	fmt.Print("Введите свой рост в см: ")
	fmt.Scan(&userHeight)

	fmt.Print("Введите свой вес в кг: ")
	fmt.Scan(&userWeight)

	return userHeight, userWeight
}

func isContinueScript() bool {
	var userChoice string
	fmt.Print("Вы хотите продолжить (y/n): ")
	fmt.Scan(&userChoice)

	return userChoice == "y" || userChoice == "Y"
}