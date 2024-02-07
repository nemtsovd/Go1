/*
Задача №1
Написать функцию, которая расшифрует строку.
code = "220411112603141304"
Каждые две цифры - это либо буква латинского алфавита в нижнем регистре либо пробел.
Отчет с 00 -> 'a' и до 25 -> 'z', 26 -> ' '(пробел).
Вход: строка из цифр. Выход: Текст.
Проверка работы функции выполняется через вторую строку.

codeToString(code) -> "???????'
*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	src := "220411112603141304"
	codes := initMap()
	for i := 0; i < len(src); i = i + 2 {
		fmt.Print(codes[src[i:i+2]])
	}
	fmt.Println()
}

func initMap() map[string]string {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	result := make(map[string]string)
	code := ""
	for p, c := range alphabet {
		if p < 10 {
			code = "0" + strconv.Itoa(p)
		} else {
			code = strconv.Itoa(p)
		}
		result[code] = string(c)
	}
	return result
}
