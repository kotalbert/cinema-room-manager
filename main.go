package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

const priceSmallRoom = 10
const priceBigRoom = 8

func main() {

	rows := getRows()
	seats := getSeats()
	c := NewCinema(rows, seats)
	for {
		row := getBookingRow()
		seat := getBookingSeat()
		if c.IsSeatBooked(row, seat) {
			fmt.Println("That ticket has already been purchased!")
		} else {
			c.BookSeat(row, seat)
			p := c.getTicketPrice(row)
			fmt.Printf("Ticket price: $%d\n", p)
			if c.Rows*c.Seats == len(c.Bookings) {
				fmt.Println("Cinema is full!")
				break
			}
		}
	}
	fmt.Println(c.ToString())

}

type Booking struct {
	Row  int
	Seat int
}

type Cinema struct {
	Rows     int
	Seats    int
	Bookings []Booking
}

func NewCinema(rows, seats int) *Cinema {
	return &Cinema{Rows: rows, Seats: seats}
}

func (c *Cinema) BookSeat(row, seat int) {
	c.Bookings = append(c.Bookings, Booking{Row: row, Seat: seat})
}

func (c *Cinema) IsSeatBooked(row, seat int) bool {
	for _, b := range c.Bookings {
		if b.Row == row && b.Seat == seat {
			return true
		}
	}
	return false
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
func getBookingRow() int {
	fmt.Println("Enter a row number:")
	row, err := getIntFromUser()
	checkError(err)
	return row
}

func getBookingSeat() int {
	fmt.Println("Enter a seat number in that row:")
	seat, err := getIntFromUser()
	checkError(err)
	return seat
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

func (c *Cinema) ToString() string {
	b := strings.Builder{}
	return b.String()
}

func (c *Cinema) getTicketPrice(row int) int {
	if c.Rows*c.Seats < 60 {
		return priceSmallRoom
	} else {
		if row <= c.Rows/2 {
			return priceSmallRoom
		} else {
			return priceBigRoom
		}
	}

}
