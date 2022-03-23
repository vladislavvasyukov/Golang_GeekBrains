package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"time"
)

type CalculationError struct {
	message     string
	createdTime time.Time
	// посмотреть, как получить текущее время
}

func (e CalculationError) Error() string {
	return fmt.Sprintf("CalculationError<%s> %s", e.message, e.createdTime)
}

func main() {
	result, err := release()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
	fmt.Println("continue to work...")

	var filename string
	fmt.Println("Введите название файла:")
	fmt.Scan(&filename)

	createFile(filename)
}

func release() (result float64, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = CalculationError{message: " some error ", createdTime: time.Now()}
			fmt.Println(e)
		}
	}()
	result, _ = calculate(34, 0, "/")
	return result, nil
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
			panic(errors.New("divide by zero"))
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

func createFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		e := file.Close()
		if e != nil {
			log.Fatal(e)
		}
	}()
}
