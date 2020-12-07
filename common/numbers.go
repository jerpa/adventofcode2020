package common

import "strconv"

// GetInts converts a slice of strings to a slice of ints
func GetInts(data []string) []int {
	result := []int{}
	for _, v := range data {
		val, err := strconv.Atoi(v)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, val)
	}
	return result
}

// GetInt converts a slice of strings to a slice of ints
func GetInt(data string) int {
	val, err := strconv.Atoi(data)
	if err != nil {
		panic(err.Error())
	}
	return val
}

// GetFloats converts a slice of strings to a slice of floats
func GetFloats(data []string) []float64 {
	result := []float64{}
	for _, v := range data {
		val, err := strconv.ParseFloat(v, 64)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, val)
	}
	return result
}
