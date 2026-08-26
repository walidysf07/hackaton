package piscine

import "fmt"

func DealAPackOfCards(deck []int) {
	j := 1
	for i := 0; i < 12; i += 3 {

		fmt.Printf("Player %d: %d, %d, %d\n", j, deck[i], deck[i+1], deck[i+2])
		j++
	}
}
