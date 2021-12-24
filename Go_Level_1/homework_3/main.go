package main

import (
	"fmt"
	"math"
)

func main() {
	var numberOne, numberTwo float64
	var operation string
	fmt.Println("Введите два числа и операцию (Например: 12 + 45):")
	_, err := fmt.Scan(&numberOne)
	if err != nil {
		fmt.Println("Ошибка ввода числа:", err.Error())
	}
	_, err = fmt.Scan(&operation)
	if err != nil {
		fmt.Println("Ошибка ввода операции:", err.Error())
	}
	_, err = fmt.Scan(&numberTwo)
	if err != nil {
		fmt.Println("Ошибка ввода числа:", err.Error())
	}
	fmt.Println("Результат = ", calculate(numberOne, numberTwo, operation))

	var num int
	fmt.Println("Введите число:")
	_, err = fmt.Scan(&num)
	if err != nil {
		fmt.Println("Ошибка:", err.Error())
		return
	}
	fmt.Println("Все простые числа меньшие", num)
	findAllSimpleNumbers(num)
}

func calculate(numberOne float64, numberTwo float64, operation string) (result float64) {
	switch operation {
		case "+":
			result = numberOne + numberTwo
		case "-":
			result = numberOne - numberTwo
		case "/":
			result = numberOne / numberTwo
		case "*":
			result = numberOne * numberTwo
		case "**":
			result = math.Pow(numberOne, numberTwo)
		//default:
		//	error("неизвестная операция")
	}
	return result
}

func findAllSimpleNumbers(number int) {
	for num := 2; num <= number; num++ {
		isSimple := true
		for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
			if num % i == 0 && num != i {
				isSimple = false
			}
		}
		if isSimple {
			fmt.Println(num)
		}
	}
}
