package evennumber

import

"fmt"

//Функция проверки четности числа.
func Calculation() {

var number int

fmt.Println("Введите число для проверки на четность:")

fmt.Scan(&number)

	if number % 2 == 0 {
		fmt.Println(number," - число четное.")
	} else {fmt.Println(number," - число нечетное.")
	}	
		
}