package piscine

func StringToIntSlice(str string) []int {
	if str == "" {
		return nil
	}

	result := []int{}
	for _, c := range str {
		result = append(result, int(c))
	}
	return result
}
