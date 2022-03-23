package main

import (
	"fmt"
	"math"
)

type CalculationError struct {
	message string
}

func (e CalculationError) Error() string {
	return fmt.Sprintf("Something bad happened: %s", e.message)
}

func main() {
	value, err := tmp()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("value-", value)
}

func tmp() (string, error) {
	var err error = nil
	var value string = "initial"
	defer func() {
		if e := recover(); e != nil {
			err = CalculationError{message: "I do not know how to add error"}
			fmt.Println("something bad happs")
			return err
		}
	}()
	value = showCalculationExample()
	return value, err
}

func showCalculationExample() string {
	panic("all bad")
}

func calculate(numberOne float64, numberTwo float64, operation string) float64 {
	var result float64

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
	default:
		panic("Unknown operation")
	}
	return result
}
