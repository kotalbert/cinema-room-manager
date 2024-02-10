package main

import (
	"errors"
	"fmt"
	"kotalbert/cinema-room-manager/cinema"
	"os"
)

func main() {

	rows := getRows()
	seats := getSeats()

	c := cinema.NewCinema(rows, seats)
	profit := c.CalculateProfit()
	fmt.Printf("Total income:\n$%d", profit)

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
