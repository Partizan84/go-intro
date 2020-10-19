package main

import (
	"fmt"
	"go-intro/image"
	"go-intro/statistic"
	"go-intro/hello"
)

func main() {

	fmt.Println("Задание № 1./n")
	statistic.Average()
	statistic.Summa()

	fmt.Println("Задание № 2./n")
	image.Image()

	fmt.Println("Задание № 3./n")
	hello.Createserver()
}