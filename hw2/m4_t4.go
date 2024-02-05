/*
Задача №4. Шахматная доска
Вход: размер шахматной доски, от 0 до 20
Выход: вывести на экран эту доску, заполняя поля нулями и единицами

Пример:
Вход: 5
Выход:
    0 1 0 1 0
    1 0 1 0 1
    0 1 0 1 0
    1 0 1 0 1
    0 1 0 1 0
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

	fmt.Println("Введите размер шахматной доски от 0 до 20:")

	number, err := getInt()
	if err != nil || number < 0 || number > 20 {
		log.Fatal("Вы ввели не число от 0 до 20")
	}

	for i := 0; i < number; i++ {
		x := float64(i % 2)
		for j := 0; j < number; j++ {
			x = math.Round(math.Abs(float64(x - 1)))
			fmt.Printf("%.0f ", x)
		}
		fmt.Println("\n")
	}
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
