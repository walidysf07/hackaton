package piscine

func Unmatch(a []int) int {
	count := make(map[int]int)
	for _, nbr := range a {
		count[nbr]++
	}
	for _, nbr := range a {
		if count[nbr]%2 != 0 {
			return nbr
		}
	}
	return -1
}
