package cinema

import "testing"

func TestCalculateProfitForSmallRoom(t *testing.T) {
	c := NewCinema(4, 5)
	if c.CalculateProfit() != 200 {
		t.Error("Expected 200 but got", c.CalculateProfit())
	}
}

func TestCalculateProfitForBigRoom(t *testing.T) {
	c := NewCinema(8, 9)
	if c.CalculateProfit() != 648 {
		t.Error("Expected 648 but got", c.CalculateProfit())
	}
}

func TestCalculateProfitForOddNumberOfRows(t *testing.T) {
	c := NewCinema(9, 7)
	if c.CalculateProfit() != 560 {
		t.Error("Expected 560 but got", c.CalculateProfit())
	}
}
