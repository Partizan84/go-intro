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
		// Добавлена справка к калькулятору.
		if input == "help" {
			calculator.HelpCalc()
			continue
		}

        if res, err := calculator.Calculate(input); err == nil {
            fmt.Printf("Результат: %v\n", res)
        } else {
            fmt.Println("Не удалось произвести вычисление")
        }
    }
}