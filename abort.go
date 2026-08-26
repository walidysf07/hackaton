package piscine

func Abort(a, b, c, d, e int) int {
	slice := []int{a, b, c, d, e}
	SortIntegerTable(slice)
	return slice[len(slice)/2]
}
