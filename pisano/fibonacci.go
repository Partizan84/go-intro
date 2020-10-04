package pisano

import ("fmt")

//Calculation Вычисление числа Фибоначчи.
func Calculation(N uint) uint{

	var 	FN []uint

	FN = make([]uint, N+1)

	FN[0] = 0

	FN[1] = 1

	for i := uint(2); i <= N; i++ {

		FN[i] = FN[i-1] + FN[i-2]
	}

	return FN[N]

}

func pisano() {

	fmt.Println(Calculation(100))
	
	}