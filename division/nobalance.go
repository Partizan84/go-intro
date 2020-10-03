package division

import

"fmt"

//Calculation Функция проверки деление числа на 3 без остатка
func Calculation() {

	var number int
	 
	fmt.Println("Введите число для верки деления на 3:")
	
	fmt.Scanln(&number)

	if number % 3 == 0 {
		fmt.Println(number," - делится нацело на 3.")
	} else {fmt.Println(number," - не делится нацело на 3.")
	}	


}