package main

import "fmt"

func main() {
	array := []int{11, 32, 323, -90, 89, 12, 43, -909, -43, 100}
	fmt.Println(array)
	insertionSort(array)
	fmt.Println(array)
}

func insertionSort(array []int) {
	for i:=0; i<len(array); i++ {
		j := i
		var tmp int
		for j>0 && (array[j] < array[j-1]) {
			tmp = array[j-1]
			array[j-1] = array[j]
			array[j] = tmp
			j--
		}
	}
}
