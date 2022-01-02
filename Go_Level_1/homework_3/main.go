package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	var numberOne, numberTwo float64
	var operation string
	fmt.Println("Введите два числа и операцию (Например: 12 + 45):")
	_, err := fmt.Scan(&numberOne, &operation, &numberTwo)
	if err != nil {
		fmt.Println("Ошибка ввода:", err.Error())
	}

	result, calculateError := calculate(numberOne, numberTwo, operation)
	if calculateError != nil {
		fmt.Println("Ошибка:", calculateError.Error())
		return
	}

	fmt.Println("Результат = ", result)

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

func calculate(numberOne float64, numberTwo float64, operation string) (float64, error) {
	var result float64
	var err error

	switch operation {
		case "+":
			result = numberOne + numberTwo
		case "-":
			result = numberOne - numberTwo
		case "/":
			if numberTwo == 0 {
				err = errors.New("делить на ноль нельзя")
				break
			}
			result = numberOne / numberTwo
		case "*":
			result = numberOne * numberTwo
		case "**":
			result = math.Pow(numberOne, numberTwo)
		default:
			err = errors.New("неизвестная операция")
	}
	return result, err
}

func findAllSimpleNumbers(number int) {
	for num := 2; num <= number; num++ {
		isSimple := true
		for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
			if num % i == 0 && num != i {
				isSimple = false
				break
			}
		}
		if isSimple {
			fmt.Println(num)
		}
	}
}
