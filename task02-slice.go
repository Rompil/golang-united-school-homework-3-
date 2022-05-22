package homework

import "sort"

func reverse(input []int64) (result []int64) {
	//Place your code here
	sort.Slice(input, func(i, j int) bool {
		return (input[i] > input[j])
	})
	return input
}
