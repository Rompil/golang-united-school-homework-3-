package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	//Place your code here
	commonSize := len(input)
	keys := make([]int, 0, commonSize)
	for k := range input {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	output := make([]string, 0, commonSize)
	for _, val := range keys {
		output = append(output, input[val])
	}
	return output
}
