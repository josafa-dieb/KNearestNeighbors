package modules

import "math"

func Dist(a []float64, b []float64) float64 {
	d := 0.0
	i := 0
	for i < len(a) {
		d += math.Pow((a[i] - b[i]), 2)
		i++
	}
	return float64((math.Pow(d, (1.0 / 2.0))))
}
