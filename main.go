package main

import (
	"errors"
	"fmt"
	"math"
	"os"
)

const priceSmallRoom = 10
const priceBigRoom = 8

func main() {

	rows := getRows()
	seats := getSeats()

	c := NewCinema(rows, seats)
	profit := c.CalculateProfit()
	fmt.Printf("Total income:\n$%d", profit)

}

type Cinema struct {
	Rows  int
	Seats int
}

func NewCinema(rows, seats int) *Cinema {
	return &Cinema{rows, seats}
}

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
func getSeats() int {
	fmt.Println("Enter the number of seats in each row:")
	seats, err := getIntFromUser()
	checkError(err)
	return seats
}

func getRows() int {
	fmt.Println("Enter the number of rows:")
	rows, err := getIntFromUser()
	checkError(err)
	return rows
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func getIntFromUser() (int, error) {
	var input int
	_, err := fmt.Scan(&input)
	if err != nil {
		return -1, errors.New("error reading from user")
	}
	return input, nil
}
