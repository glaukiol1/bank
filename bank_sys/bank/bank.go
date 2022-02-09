package bank

import "github.com/glaukiol1/bank/bank_sys/card"

type Bank struct {
	Customers []*Customer
}

func (bank *Bank) NewCustomer(name, password string) *Customer {
	n := &Customer{name, createBin(name), createAddresses(), []*card.Card{}, 0, password}
	bank.Customers = append(bank.Customers, n)
	return n
}

var GagBank = &Bank{}
