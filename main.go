package main

import (

	"go-intro/implementation"

	"fmt"

	"go-intro/calculator"
)

func main() {
	
	implementation.Books()

	input := ""
    for {
        fmt.Print("> ")
        if _, err := fmt.Scanln(&input); err != nil {
            fmt.Println(err)
            continue
        }

        if input == "exit" {
            break
		}
		
		if input == "help" {
			fmt.Printf("Справка по работе калькулятора.\n Для вычисления бинарный выражений :\n 1) Введите первый аргумент.\n 2) Введите выражение:+-/*%&| \n 3) Введите второй аргумент.")
			break
		}

        if res, err := calculator.Calculate(input); err == nil {
            fmt.Printf("Результат: %v\n", res)
        } else {
            fmt.Println("Не удалось произвести вычисление")
        }
    }
}