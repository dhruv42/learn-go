package arrays

func Sum(numbers []int) int {
	sum := 0
	for _, element := range numbers {
		sum += element
	}
	return sum
}
