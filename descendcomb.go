package piscine

import "github.com/01-edu/z01"

func DescendComb() {
	for a := 9; a >= 0; a-- {
		for b := 9; b >= 0; b-- {
			for c := 9; c >= 0; c-- {
				for d := 9; d >= 0; d-- {
					if (a*10 + b) > (c*10 + d) {
						z01.PrintRune(rune(a) + '0')
						z01.PrintRune(rune(b) + '0')
						z01.PrintRune(' ')
						z01.PrintRune(rune(c) + '0')
						z01.PrintRune(rune(d) + '0')

						if !(a == 0 && b == 1 && c == 0 && d == 0) {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
					}
				}
			}
		}
	}
}
