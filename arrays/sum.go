package main

func Sum(arr []int) int {
	sum := 0

	for _, number := range arr {
		sum += number
	}

	return sum
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}