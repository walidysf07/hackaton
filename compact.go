package piscine

func Compact(ptr *[]string) int {
	count := 0

	for _, s := range *ptr {
		if s != "" {
			(*ptr)[count] = s

			count++
		}
	}

	*ptr = (*ptr)[:count]
	return count
}
