package main

func Sum(arr []int) int {
	sum := 0

	for _, number := range arr {
		sum += number
	}

	return sum
}
