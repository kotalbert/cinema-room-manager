package cinema

type Cinema struct {
	Rows  int
	Seats int
}

func NewCinema(rows, seats int) *Cinema {
	return &Cinema{rows, seats}
}

const priceSmallRoom = 10
const priceBigRoom = 8

func (c *Cinema) CalculateProfit() int {
	seats := c.Rows * c.Seats
	if seats < 60 {
		return seats * priceSmallRoom
	} else {
		return seats / 2 * (priceBigRoom + priceSmallRoom)
	}

}
