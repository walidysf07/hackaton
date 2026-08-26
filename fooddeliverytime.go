package piscine

type food struct {
	preptime int
}

func FoodDeliveryTime(order string) int {
	m := food{}

	switch order {
	case "burger":
		m.preptime = 15
	case "chips":
		m.preptime = 10
	case "nuggets":
		m.preptime = 12
	default:
		return 404
	}

	return m.preptime
}
