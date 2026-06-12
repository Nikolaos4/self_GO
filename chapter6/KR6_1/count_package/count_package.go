package count_package

func factorial(number int) int{
	if (number <= 1) {
		return number
	}
	return number*factorial(number-1)
}

func combinations(current, all_count int) int {
	return factorial(all_count)/(factorial(current)*factorial(all_count-current))
}

func Probability(count_of_pair, all_people int) int {
	N := combinations(2, all_people)
	M := combinations(count_of_pair, all_people)
	return M/N
}