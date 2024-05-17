package utils

import (
	"math"
	"math/rand"
	"time"
)

func Max(data []float64) float64 {
	maxVal := data[0]
	for _, val := range data {
		if val > maxVal {
			maxVal = val
		}
	}
	return maxVal
}

func Min(data []float64) float64 {
	minVal := data[0]
	for _, val := range data {
		if val < minVal {
			minVal = val
		}
	}
	return minVal
}

func Mean(data []float64) float64 {
	sum := 0.0
	for _, val := range data {
		sum += val
	}
	return sum / float64(len(data))
}

func StdDev(data []float64) float64 {
	mean := Mean(data)
	sum := 0.0
	for _, val := range data {
		sum += math.Pow(val-mean, 2)
	}
	return math.Sqrt(sum / float64(len(data)))
}

func NormalizationZeroOne(data []float64) []float64 {
	minVal := Min(data)
	rangeVal := Max(data) - minVal
	result := make([]float64, len(data))
	for i, val := range data {
		result[i] = (val - minVal) / rangeVal
	}
	return result
}

func NormalizationOneOne(data []float64) []float64 {
	maxAbs := math.Abs(data[0])
	for _, val := range data {
		if math.Abs(val) > maxAbs {
			maxAbs = math.Abs(val)
		}
	}
	result := make([]float64, len(data))
	for i, val := range data {
		result[i] = val / maxAbs
	}
	return result
}

func Standardization(data []float64) []float64 {
	mean := Mean(data)
	stdDev := StdDev(data)
	result := make([]float64, len(data))
	for i, val := range data {
		result[i] = (val - mean) / stdDev
	}
	return result
}

func RandomBetween(lower, upper int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(upper-lower+1) + lower
}
