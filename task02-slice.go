package homework

func reverse(input []int64) ([]int64) {
	result := make([]int64, 0)
	for i := len(input)-1; i >= 0; i-- {
		result = append(result, input[i])
	}
	return result
}
