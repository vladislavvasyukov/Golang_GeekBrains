package main

import (
	"errors"
	"fmt"
)

func main() {
	// что такое ошибка? (это интерфейс)
	// как реализовать свой тип ошибок
	// оборачивание ошибки
	err := errors.New("some test")
	wrappedErr := fmt.Errorf("wrapped err: %w", err)
	fmt.Println(wrappedErr)
	err = errors.Unwrap(wrappedErr)
	fmt.Println(err)

	// error.Is
	// error.As
	// отложенный вызов ф-ции - что это и как работает (defer - вызывается при выходе из ф-ции???)

	for i := 1; i < 10; i++ { // сколько раз выведется defer??? defer вызываются в обратном порядке
		defer close(i)
	}

	// как вызвать panic?
	// когда использовать paniс, а когда обычную ошибку?
	// как написать восстановление от panic?
	dir := 1
	if dir == 1 {
		panic("some panic")
	}

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic occured", err)
		}
	}()

	// ioutil.ReadDir(*dirPath)
	// ioutil.ReadAll(file)
	// strings.Contains
}

func close(i int) {
	fmt.Println("Goodbye, world ", i)
}
