package division

import

"fmt"

//Calculation Функция проверки деление числа на 3 без остатка
func Calculation() {

	var num int
	 
	fmt.Println("Введите число для верки деления на 3:")
	
	fmt.Scanln(&num)

	if num % 3 == 0 {
		fmt.Println(num," - делится нацело на 3.")
	} else {fmt.Println(num," - не делится нацело на 3.")
	}	


}