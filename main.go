package main

import (

	"go-intro/evennumber"

	"go-intro/division"

	"go-intro/pisano"

	"fmt"
	
)

func main() {

	fmt.Println("Задание № 1:")

	evennumber.Calculation()

	fmt.Println("Задание № 2:")

	division.Calculation()

	fmt.Println("Задание № 3:")

	pisano.Calculation(100)
	
	fmt.Println(pisano.Calculation(100))

}
