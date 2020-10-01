package converter

import "fmt"

func Calculation() {

	const rate float64 = 75.55 //курс рубля к доллару

	var rub float64

	fmt.Println("Введите сумму в рублях: ")

	fmt.Scanln(&rub)

	usd := rub / rate

	fmt.Println("Ваша сумма в долларах: ")

	fmt.Printf("$%.2f\n", usd)

}
