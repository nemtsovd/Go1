/*
Задача № 3. Вывести на экран в порядке возрастания три введенных числа
Пример:
Вход: 1, 9, 2
Выход: 1, 2, 9
*/
package main

import (
	"fmt"
)

func main() {

	var num1 int
	var num2 int
	var num3 int
	var tmp int

	fmt.Println("Введите три числа: ")
	fmt.Scanf("%d", &num1)
	fmt.Scanf("%d", &num2)
	fmt.Scanf("%d", &num3)

	if num1 > num2 {
		tmp = num1
		num1 = num2
		num2 = tmp
	}
	if num2 > num3 {
		tmp = num2
		num2 = num3
		num3 = tmp
	}
	if num1 > num2 {
		tmp = num1
		num1 = num2
		num2 = tmp
	}

	fmt.Printf("%d, %d, %d\n", num1, num2, num3)
}
