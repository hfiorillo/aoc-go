package helpers

import "strconv"

func ToInt(s string) (int, error) {
	result, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func FindMax(arr []int) (max int) {
	max = arr[0]
	for _, value := range arr {
		if value > max {
			max = value
		}
	}
	return max
}

func ToIntList(string_list []string) []int {
	var int_list = make([]int, len(string_list))
	for idx, v := range string_list {
		j, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		int_list[idx] = j
	}
	return int_list
}
