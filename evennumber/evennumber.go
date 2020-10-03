package evennumber

import

"fmt"

//Функция проверки четности числа.
func Calculation() {

var number int

fmt.Println("Введите число для проверки на четность:")

fmt.Scanln(number)

	if number % 2 == 0 {
		fmt.Println("%n.2f делится на 2 без остатка, число четное.", number)
	}
		
}