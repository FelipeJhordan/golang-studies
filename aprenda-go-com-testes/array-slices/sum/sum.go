package sum

func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(numbersForSum ...[]int) (sums []int) {

	for _, numbers := range numbersForSum {
		sums = append(sums, Sum(numbers))
	}

	return
}

func SumAllRest(numbersForSum ...[]int) (sums []int) {

	for _, numbers := range numbersForSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
			continue
		}
		final := numbers[1:]
		sums = append(sums, Sum(final))
	}

	return
}
