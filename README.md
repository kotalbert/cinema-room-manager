# Cinema Room Manager

A Hyperskill project to create a simple cinema room manager. The program will help manage the seats in the cinema room.
It will display the seats, mark the seats that are already booked, and help a user buy tickets. The program will also
display statistics about the number of seats and the income from the tickets.

## Test suite

```go
package main

import "testing"

func TestCalculateProfitForSmallRoom(t *testing.T) {
	c := NewCinema(4, 5)
	if c.CalculateProfit() != 200 {
		t.Error("Expected 200 but got", c.CalculateProfit())
	}
}

func TestCalculateProfitForBigRoom(t *testing.T) {
	c := main.NewCinema(8, 9)
	if c.CalculateProfit() != 648 {
		t.Error("Expected 648 but got", c.CalculateProfit())
	}
}

func TestCalculateProfitForOddNumberOfRows(t *testing.T) {
	c := main.NewCinema(9, 7)
	if c.CalculateProfit() != 560 {
		t.Error("Expected 560 but got", c.CalculateProfit())
	}
}
```