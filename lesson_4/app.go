package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var numbers string
	var listNumbers []int
	fmt.Println("Введите целые числа через запятую:")
	fmt.Scan(&numbers)
	arrayNumbers := strings.Split(numbers, ",")
	for _, v := range arrayNumbers {
		v, err := strconv.Atoi(v)
		if err != nil {
			fmt.Printf("Не является целым числом, ошибка: %s\n", err)
			os.Exit(1)
		}
		listNumbers = append(listNumbers, v)
	}

	for i := 1; i < len(listNumbers); i++ {
		k := i
		for k > 0 && listNumbers[k-1] > listNumbers[k] {
			tmp := listNumbers[k-1]
			listNumbers[k-1] = listNumbers[k]
			listNumbers[k] = tmp
			k--
		}
	}

	fmt.Printf("результат %v", listNumbers)
}
