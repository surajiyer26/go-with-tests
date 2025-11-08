package main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(numbersToTailSum ...[]int) []int {
	var tailSum []int
	for _, numbers := range numbersToTailSum {
		if len(numbers) == 0 {
			tailSum = append(tailSum, 0)
		} else {
			tailSum = append(tailSum, Sum(numbers[1:]))
		}
	}

	return tailSum
}
