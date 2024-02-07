/*
Задача №2
Вход:
Пользователь должен ввести правильный пароль, состоящий из:
цифр,
букв латинского алфавита(строчные и прописные) и
специальных символов  special = "_!@#$%^&"

Всего 4 набора различных символов.
В пароле обязательно должен быть хотя бы один символ из каждого набора.
Длина пароля от 8(мин) до 15(макс) символов.
Максимальное количество попыток ввода неправильного пароля - 5.
Каждый раз выводим номер попытки.
*Желательно выводить пояснение, почему пароль не принят и что нужно исправить.

digits = "0123456789"
lowercase = "abcdefghiklmnopqrstvxyz"
uppercase = "ABCDEFGHIKLMNOPQRSTVXYZ"
special = "_!@#$%^&"

Выход:
Написать, что ввели правильный пароль.

Пример:
хороший пароль -> o58anuahaunH!
хороший пароль -> aaaAAA111!!!
плохой пароль -> saucacAusacu8
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < 6; i++ {
		fmt.Printf("Введите пароль (попытка %d): ", i)
		input, err := getString()

		if err != nil {
			fmt.Println("Неверный пароль")
			continue
		}

		if len(input) < 8 || len(input) > 15 {
			fmt.Println("Неверный пароль. Длина пароля должна быть от 8 до 15 символов.")
			continue
		}

		res := checkInput(input)
		if res != "" {
			fmt.Println("Неверный пароль." + res)
			continue
		}

		fmt.Println("Верный пароль!")
		break
	}

}

func getString() (string, error) {

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(input), err
}

func checkInput(input string) string {

	digits := "0123456789"
	lowercase := "abcdefghiklmnopqrstvxyz"
	uppercase := "ABCDEFGHIKLMNOPQRSTVXYZ"
	special := "_!@#$%^&"

	var result string
	var consistsDigit, consistsLower, consistsUpper, consistsSpecial bool

	for _, c := range input {
		if strings.Contains(digits, string(c)) {
			consistsDigit = true
		}
		if strings.Contains(lowercase, string(c)) {
			consistsLower = true
		}
		if strings.Contains(uppercase, string(c)) {
			consistsUpper = true
		}
		if strings.Contains(special, string(c)) {
			consistsSpecial = true
		}
	}

	if !consistsDigit {
		result += " Пароль должен содержать хотя бы одну цифру."
	}
	if !consistsLower {
		result += " Пароль должен содержать хотя бы одну строчную букву."
	}
	if !consistsUpper {
		result += " Пароль должен содержать хотя бы одну прописную букву."
	}
	if !consistsSpecial {
		result += " Пароль должен содержать хотя бы один специальный символ."
	}

	return result
}
