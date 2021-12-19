package main

import (
	"fmt"
	"math"
)


func main() {
	var a, b int
	fmt.Println("Введите длины сторон прямоугольника, чтобы найти его площадь (a, b):")
	fmt.Scan(&a, &b)
	fmt.Println("Площадь прямоугольника равна ", getRectangleArea(a, b))

	var circleArea float64
	fmt.Println("Введите площадь круга, чтобы рассчитать диаметр и длину окружности")
	fmt.Scan(&circleArea)
	fmt.Println("Диаметр равен ", getCircleDiameter(circleArea))
	fmt.Println("Длина окружности равна ", getCircumference(circleArea))

}

func getRectangleArea(a int, b int) (area int) {
	return a * b
}

func getCircleDiameter(circleArea float64) (diameter float64) {
	return math.Sqrt(circleArea / math.Pi) * 2
}

func getCircumference(circleArea float64) (circumference float64) {
	return math.Sqrt(circleArea * math.Pi * 4)
}
