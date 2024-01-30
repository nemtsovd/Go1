/*
Задача № 4. Проверить, является ли четырехзначное число палиндромом
Пример:
Вход: 1221  Выход: 1221 - палиндром
Вход: 1234  Выход: 1234 - не палиндром
*/
package main

import (
    "fmt"
)

func main() {

	var src int
	var rev int

	fmt.Print("Введите число: ")
	fmt.Scanf("%d", &src)

	num := src
    for num > 0 {
		tmp := num % 10
		rev = rev * 10 + tmp
		num /= 10
	}

	if src == rev {
		fmt.Println("Палиндром")
	} else {
		fmt.Println("Не палиндром")
	}
}
