package main

import (
	"fmt"

	"github.com/glaukiol1/bank/bank_sys/api"
	"github.com/glaukiol1/bank/bank_sys/bank"
)

func main() {
	bank := bank.GagBank
	glaukio := bank.NewCustomer("Glaukio L", "1234")
	fmt.Println(glaukio.Bin)
	glaukio.NewCard()
	println(glaukio.Addresses[0].Address, glaukio.Addresses[0].Ticker)
	api.InitApi()
}
