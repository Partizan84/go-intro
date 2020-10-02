package evennumber

import

"fmt"

//Функция проверки четности числа.
func Calculation() {

var number int

fmt.Println("Введите число для проверки на четность:")

fmt.Scanln(number)

	if number % 4 == 0 {
		fmt.Println("%n.2f делится на 4", number)
	}
		
}