package main

import (
	"errors"
	"fmt"
	"math"
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
		fmt.Println("1. Show the seats")
		fmt.Println("2. Buy a ticket")
		fmt.Println("3. Statistics")
		fmt.Println("0. Exit")
		menuInput, err := getIntFromUser()
		checkError(err)
		switch menuInput {
		case 0:
			return
		case 1:
			fmt.Print(c.ToString())
		case 2:
			HandleBooking(c)
		case 3:
			HandleStatistics(c)
		}
	}

}

func HandleStatistics(c *Cinema) {
	fmt.Printf("Number of purchased tickets: %d\n", len(c.Bookings))
	percentage := float64(len(c.Bookings)) / float64(c.Rows*c.Seats) * 100
	fmt.Printf("Percentage: %.2f%%\n", percentage)
	fmt.Printf("Current income: $%d\n", c.getCurrentIncome())
	fmt.Printf("Total income: $%d\n", c.getTotalIncome())
}

func HandleBooking(c *Cinema) {
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
			return
		}
	}
	fmt.Print(c.ToString())
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

func (c *Cinema) getCinemaString() string {

	b := strings.Builder{}
	b.WriteByte(' ')
	for i := 0; i <= c.Rows; i++ {
		for j := 0; j <= c.Seats; j++ {
			if i == 0 {
				if j == 0 {
					b.Write([]byte("  "))
				} else {
					b.WriteString(fmt.Sprintf("%d ", j))
				}
			} else {
				if j == 0 {
					b.WriteString(fmt.Sprintf("%d ", i))
				} else {
					if c.IsSeatBooked(i, j) {
						b.Write([]byte("B "))
					} else {
						b.Write([]byte("S "))
					}
				}
			}
		}
	}

	return b.String()
}

// ToString takes a getCinemaString and inserts a new line after every row
//
//	The result should be a grid with row and seat numbers as well as the bookings
func (c *Cinema) ToString() string {
	cs := c.getCinemaString()
	b := strings.Builder{}
	b.WriteString("Cinema:")
	for i := 0; i < len(cs); i++ {
		b.WriteByte(cs[i])
		if i%(2*c.Seats+2) == 0 {
			b.WriteByte('\n')
		}
	}
	b.WriteByte('\n')
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

func (c *Cinema) getTotalIncome() int {
	s := c.Rows * c.Seats
	if s < 60 {
		return s * priceSmallRoom
	} else {
		frontRowsProfit := math.Floor(float64(c.Rows)/2) * float64(c.Seats*priceSmallRoom)
		backRowsProfit := math.Ceil(float64(c.Rows)/2) * float64(c.Seats*priceBigRoom)
		return int(frontRowsProfit + backRowsProfit)
	}
}

func (c *Cinema) getCurrentIncome() int {
	var sum int
	for _, b := range c.Bookings {
		sum += c.getTicketPrice(b.Row)
	}
	return sum
}
