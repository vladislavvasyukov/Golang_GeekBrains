package main

import "fmt"

type Fibonacci struct {
	cache map[int64]int64
}

func main()  {
	var number int64
	fmt.Println("Введите число >= 0")
	_, err := fmt.Scan(&number)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if number < 0 {
		fmt.Println("Число должно быть >= 0")
		return
	}

	fib := Fibonacci{
		cache: make(map[int64]int64),
	}
	fmt.Println("F(", number, ") =", fib.calculate(number))
}

func (f Fibonacci) calculate(number int64) int64 {
	if number == 0 || number == 1 {
		return number
	}

	n1, ok1 := f.cache[number-1]
	if !ok1 {
		n1 = f.calculate(number-1)
		f.cache[number-1] = n1
	}

	n2, ok2 := f.cache[number-2]
	if !ok2 {
		n2 = f.calculate(number-2)
		f.cache[number-2] = n2
	}

	return n1 + n2
}