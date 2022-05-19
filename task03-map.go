package homework

import "sort"

func sortMapValues(input map[int]string) ([]string) {
	keys := make([]int, len(input))
	vals := make([]string, len(input))
	for k, v := range input {
		keys = append(keys, k)
		vals = append(vals, v)
	}
	sort.Slice(vals, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	return vals
}
