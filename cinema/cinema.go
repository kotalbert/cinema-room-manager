package cinema

import "math"

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
	s := c.Rows * c.Seats
	if s < 60 {
		return s * priceSmallRoom
	} else {
		frontRowsProfit := math.Floor(float64(c.Rows)/2) * float64(c.Seats*priceSmallRoom)
		backRowsProfit := math.Ceil(float64(c.Rows)/2) * float64(c.Seats*priceBigRoom)
		return int(frontRowsProfit + backRowsProfit)
	}

}
