package math

//godoc -http=":6060" - для поднятия веб сервера с документацией
// Найти среднее в массиве чисел.
func Average(xs []float64) float64 {
	if len(xs) == 0 {
		return 0
	}
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}

// Найти минимальное в массиве чисел.
func Min(xs []float64) float64 {
	var result float64 = xs[0]
	for i := 0; i < len(xs); i++ {
		if result > xs[i] {
			result = xs[i]
		}
	}
	return result
}

// Найти максимальное в массиве чисел.
func Max(xs []float64) float64 {
	var result = xs[0]
	for i := 0; i < len(xs); i++ {
		if result < xs[i] {
			result = xs[i]
		}
	}
	return result
}
