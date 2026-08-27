package piscine

func ActiveBits(n int) int {
	count := 0
	for n != 0 {
		bit := n % 2
		n = n / 2
		if bit == 1 {
			count++
		}
	}
	return count
}
