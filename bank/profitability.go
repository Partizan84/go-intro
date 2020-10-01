package bank

import (
	"fmt"
	"math"
)

func Calculation() {

	var sumDeposit float64

	var percentYear float64

	const periodYears = 5

	fmt.Println("Введите сумму вашего вклада:")

	fmt.Scanln(&sumDeposit)

	fmt.Println("Введите годовой процент по вкладу:")

	fmt.Scanln(&percentYear)

	depositProfit := sumDeposit * math.Pow(1+percentYear/100, periodYears)

	fmt.Printf("Ваш вклад через %x лет составит: %.2f \n", periodYears, depositProfit)
}
