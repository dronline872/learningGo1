package main

import (
	"fmt"
	"math"
)

const minNumber = 100
const maxNumber = 999

func main() {
	fmt.Println("Хотите узнать площадь треугольника?")
	areaTringle()

	fmt.Println("Введите площадь окружности чтобы узнать его длину и диаметр:")
	diameterAndLenghtCircle()

	fmt.Println("Введите трехзначное число:")
	parsingNumber()
}

//площадь треугольника
func areaTringle() {
	var a, b, c float64
	inputSide(&a, "a")
	inputSide(&b, "b")
	inputSide(&c, "c")
	if a > 0 && b > 0 && c > 0 && (a+b) > c && (a+c) > b && (b+c) > a {
		p := (a + b + c) / 2
		s := math.Sqrt(p * (p - a) * (p - b) * (p - c))
		fmt.Printf("Площадь треугольника: %f\n", s)
	} else {
		err := fmt.Errorf("неправильные значения сторон треугольника: %f, %f, %f", a, b, c)
		fmt.Println(err)
	}
}

func inputSide(side *float64, name string) {
	fmt.Printf("Введите значение стороны %s:", name)
	fmt.Scan(side)
}

//Диаметр и длина окружности по его площади
func diameterAndLenghtCircle() {
	var s float64
	fmt.Scan(&s)
	lenght := lengthCirce(s)
	diameter := diameter(lenght)
	fmt.Printf("Диаметр:%f, Длина:%f\n", diameter, lenght)
}

func diameter(lenght float64) float64 {
	d := lenght / math.Pi
	return d
}

func lengthCirce(s float64) float64 {
	p := math.Sqrt(s * 4 * math.Pi)
	return p
}

//Количество сотен, десятков и единиц
func parsingNumber() {
	number := 0
	fmt.Scan(&number)
	for number < minNumber || maxNumber < number {
		fmt.Println("Число не является трехзначным, введите корректное число:")
		fmt.Scan(&number)
	}

	hundreds := number / 100
	number = number - (hundreds * 100)
	dozens := number / 10
	units := number - (dozens * 10)
	fmt.Printf("Сотен:%d, десятков:%d, единиц:%d", hundreds, dozens, units)
}
