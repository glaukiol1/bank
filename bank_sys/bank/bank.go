package bank

import "github.com/glaukiol1/bank/bank_sys/card"

type Bank struct {
	Customers []*Customer
}

func (bank *Bank) NewCustomer(name, password, documentId string) *Customer {
	n := &Customer{name, documentId, createBin(name), createAddresses(), []*card.Card{}, password}
	bank.Customers = append(bank.Customers, n)
	return n
}

var GagBank = &Bank{}
