/*
Задача № 2. Получить реверсную запись трехзначного числа
Пример:
вход: 346, выход: 643
вход: 100, выход: 001
*/
package main

import (
	"fmt"
	"strconv"
)

func main() {

	var num int
	var rev string

	fmt.Print("Введите трехзначное число: ")
	fmt.Scanf("%d", &num)

	str := strconv.Itoa(num)
	for _, v := range str {
		rev = string(v) + rev
	}
	fmt.Println(rev)
}
