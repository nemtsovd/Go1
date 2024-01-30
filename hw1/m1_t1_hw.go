/*
Задача №1
Вход:
    расстояние(50 - 10000 км),
    расход в литрах (5-25 литров) на 100 км и
    стоимость бензина(константа) = 48 руб

Выход: стоимость поездки в рублях
*/
package main

import "fmt"

func main() {
	var length float64
    var rate float64
    const cost float64 = 48

    for length < 50 || length > 10000 {
        fmt.Print("Введите расстояние (50 - 10000 км): ")
	    fmt.Scanf("%f", &length)
        if length < 50 || length > 10000 {
            fmt.Println("Расстояние должно быть от 50 до 10000 км")
        }
    }

    
    for rate < 5 || rate > 25 {
        fmt.Print("Введите расход (5 - 25 литров): ")
	    fmt.Scanf("%f", &rate)
        if length < 5 || length > 25 {
            fmt.Println("Расход должен быть от 5 до 25 литров")
        }
    }
    
    fmt.Println("")
    fmt.Printf("Стоимость поездки: %.2f рублей\n", length * rate / 100 * cost)
}