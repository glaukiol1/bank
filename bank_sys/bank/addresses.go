package bank

import "crypto/rand"

type Address struct {
	Ticker  string `json:"ticker"`
	Address string `json:"address"`
}

func createAddresses() []Address {
	ran1, _ := rand.Prime(rand.Reader, 220)
	ran2, _ := rand.Prime(rand.Reader, 220)
	var a []Address
	return append(a, Address{"USD", "USD" + ran1.String()[0:20]}, Address{"EUR", "EUR" + ran2.String()[0:20]})
}
