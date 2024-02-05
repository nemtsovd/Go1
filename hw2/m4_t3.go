/*
Написать 3 функции.
Даны координаты трех точек(x1, y1, x2, y2, x3, y3), значения(целые) которых >= 0.
Первая функция проверяет, что можно построить треугольник по заданным точкам
Вторая функция вычисляет площадь треугольника.
Третья функция должна определить, является ли треугольник прямоугольным.
*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Введите x1:")
	x1 := getCoordinate()
	fmt.Println("Введите y1:")
	y1 := getCoordinate()
	fmt.Println("Введите x2:")
	x2 := getCoordinate()
	fmt.Println("Введите y2:")
	y2 := getCoordinate()
	fmt.Println("Введите x3:")
	x3 := getCoordinate()
	fmt.Println("Введите y3:")
	y3 := getCoordinate()

	fmt.Println("Можно построить треугольник: ", isTiangleAvailable(x1, y1, x2, y2, x3, y3))
	fmt.Println("Площадь треугольника: ", square(x1, y1, x2, y2, x3, y3))
	fmt.Println("Треугольник прямоугольный: ", isRightAngledTriangle(x1, y1, x2, y2, x3, y3))
}

func getCoordinate() int {

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Вы ввели не натуральное число")
	}

	input = strings.TrimSpace(input)
	number, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal("Вы ввели не натуральное число")
	}

	if number < 0 {
		log.Fatal("Вы ввели не натуральное число")
	}

	return number
}

func getLength(x1, y1, x2, y2 int) float64 {
	return math.Sqrt(math.Pow(float64(x2-x1), 2) + math.Pow(float64(y2-y1), 2))
}

func isTiangleAvailable(x1, y1, x2, y2, x3, y3 int) bool {
	a := getLength(x1, y1, x2, y2)
	b := getLength(x2, y2, x3, y3)
	c := getLength(x3, y3, x1, y1)

	return a+b > c && a+c > b && b+c > a
}

func square(x1, y1, x2, y2, x3, y3 int) float64 {
	a := getLength(x1, y1, x2, y2)
	b := getLength(x2, y2, x3, y3)
	c := getLength(x3, y3, x1, y1)
	p := float64(a+b+c) / 2

	return math.Sqrt(p * (p - a) * (p - b) * (p - c))
}

//как это должно работать непонятно int во float точно не конвертируются
func isRightAngledTriangle(x1, y1, x2, y2, x3, y3 int) bool {
	a := getLength(x1, y1, x2, y2)
	b := getLength(x2, y2, x3, y3)
	c := getLength(x3, y3, x1, y1)
	p := a + b + c
	min := math.Min(math.Min(a, b), c)
	max := math.Max(math.Max(a, b), c)
	mid := p - min - max

	return math.Pow(float64(min), 2)+math.Pow(float64(mid), 2) == math.Pow(float64(max), 2)
}
