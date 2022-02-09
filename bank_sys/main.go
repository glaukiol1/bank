package main

import (
	"fmt"

	"github.com/glaukiol1/bank/bank_sys/customer"
)

func main() {
	glaukio := customer.NewCustomer("Glaukio L")
	fmt.Println(glaukio.Bin)
	glaukio.NewCard()
	println(glaukio.Addresses[0].Address, glaukio.Addresses[0].Ticker)
}
