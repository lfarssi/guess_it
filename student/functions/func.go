package functions

import "math"

func Average(n []float64) float64 {
	var res float64
	for _, item := range n {
		nb := item
		res += float64(nb)
	}
	res /= float64(len(n))
	return res
}
func Variance(arr []float64) float64 {
	var variance float64
	var sum float64
	avg := Average(arr)
	for i := 0; i < len(arr); i++ {
		nb := arr[i]

		sum += (float64(nb) - avg) * (float64(nb) - avg)
	}
	variance = sum / float64(len(arr))
	return variance
}
func StandardDeviation(arr []float64) float64 {
	var res float64
	variance := Variance(arr)
	res = math.Sqrt(variance)
	return res
}

func Interval(data []float64) (int, int) {
	n := len(data) - 4
	if n < 0 {
		n = 0
	}
	lastFour := data[n:]
	avg := Average(lastFour)
	StandardDeviation := StandardDeviation(lastFour)
	nb1 := avg - (StandardDeviation * 2)
	nb2 := avg + (StandardDeviation * 2)
	return int(math.Round(nb1)), int(math.Round(nb2))
}
