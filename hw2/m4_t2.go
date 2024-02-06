/*
Написать функцию.
Входные аргументы функции: количество бутылок от 0 до 200.
Функция должна вернуть количество и слово бутыл<?> с правильным окончанием.
Пример :
In: 5
Out: 5 бутылок

In: 1
Out: 1 бутылка

In: 22
Out: 22 бутылки
*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Введите количество бутылок от 0 до 200:")

	number, err := getInt()
	if err != nil || number < 0 || number > 200 {
		log.Fatal("Вы ввели не число от 0 до 200")
	}

	fmt.Println(number, "бутыл"+getEnding(number))
}

func getInt() (int, error) {

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	input = strings.TrimSpace(input)
	number, err := strconv.Atoi(input)
	if err != nil {
		return 0, err
	}

	return number, err
}

func getEnding(number int) string {

	switch number % 100 / 10 {

	case 1:
		return "ок"
	}

	switch number % 10 {

	case 0, 5, 6, 7, 8, 9:
		return "ок"

	case 1:
		return "ка"

	case 2, 3, 4:
		return "ки"
	}

	return ""
}
