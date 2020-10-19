package statistic

//Average Среднее.
func Average(xs []float64) float64 {
    total := float64(0)
    for _, x := range xs {
        total += x
    }
	return total / float64(len(xs))
}
// Summa Сумма всех элементов.
func Summa(elements []float64) (sum float64) {
	for _, element := range elements {
		sum += element
	}

	return sum
}