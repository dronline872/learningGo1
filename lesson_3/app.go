package main

import (
	"fmt"
	"math"
	"math/big"
	"os"
)

func main() {
	calculator()
	probablyPrime()
}

func calculator() {
	var a, b, res float64
	var op string
	fmt.Print("Введите первое число: ")
	_, err := fmt.Scanln(&a)
	if err != nil {
		fmt.Printf("Введены неверные данные первого числа, ошибка: %s\n", err)
		os.Exit(1)
	}

	fmt.Print("Введите второе число: ")
	_, err = fmt.Scanln(&b)
	if err != nil {
		fmt.Printf("Введены неверные данные второго числа, ошибка: %s\n", err)
		os.Exit(1)
	}

	fmt.Print("Введите арифметическую операцию (+, -, *, /, ^, Min, Max): ")
	fmt.Scanln(&op)

	switch op {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		res = a / b
	case "^":
		res = math.Pow(a, b)
	case "Min":
		res = math.Min(a, b)
	case "Max":
		res = math.Max(a, b)
	default:
		fmt.Printf("Операция выбрана неверно")
		os.Exit(1)
	}

	fmt.Printf("Результат выполнения операции: %f\n", res)
}

func probablyPrime() {
	var a int
	fmt.Print("Введите число: ")
	_, err := fmt.Scanln(&a)
	if err != nil {
		fmt.Printf("Введены неверные данные, ошибка: %s\n", err)
		os.Exit(1)
	}

	for i := 0; i <= a; i++ {
		if big.NewInt(int64(i)).ProbablyPrime(0) {
			fmt.Println(i, "- простое число")
		}
	}

}
