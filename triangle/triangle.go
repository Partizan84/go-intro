package triangle

import (
	"fmt"
	"math"
)

func Calculation() {
	a := 4.0
	b := 2.0

	areat := a * b / 2

	perimeter := a + b + c

	c := math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))

	fmt.Printf("Катеты прямоугольного треугольника a: %.2f, b: %.2f \n", a, b)

	fmt.Printf("Площадь треугольника: %.2f\n", areat)

	fmt.Printf("Периметр треугольника: %.2f\n", perimeter)

	fmt.Printf("Гипотенуза треугольника: %.2f\n", с)
}
