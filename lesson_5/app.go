package main

import (
	"fmt"
)

func main() {
	mapNumber := map[int]int{
		0: 0,
		1: 1,
	}
	var numbers int
	fmt.Println("Введите число:")
	fmt.Scan(&numbers)
	fmt.Printf("Число фибоначчи: %d", fibonacci(mapNumber, numbers))
}

func fibonacci(mapNumber map[int]int, number int) int {
	if mapNumber[number] == number {
		return mapNumber[number]
	}

	mapNumber[number] = fibonacci(mapNumber, number-1) + fibonacci(mapNumber, number-2)
	return mapNumber[number]
}
