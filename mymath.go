package mymath

func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}

func MinMax(xs []float64) (min, max float64) {
	min = float64(xs[0])
	max = float64(xs[0])
	for _, x := range xs {
		if x < min {
			min = x
		} else if x > max {
			max = x
		}
	}
	return min, max
}