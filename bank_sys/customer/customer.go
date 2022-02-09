package customer

import (
	"crypto/rand"
	"strings"

	"github.com/glaukiol1/bank/bank_sys/card"
)

type Customer struct {
	Name  string
	Bin   string
	Cards []*card.Card
}

func createBin(name string) string {
	RandomCrypto, _ := rand.Prime(rand.Reader, 220)
	return RandomCrypto.String()[0:20] + "-" + strings.Split(name, " ")[0] + "/" + strings.Split(name, " ")[1] + "-" + RandomCrypto.String()[21:60]
}

func NewCustomer(name string) *Customer {
	return &Customer{name, createBin(name), []*card.Card{}}
}

func (customer *Customer) NewCard() {
	customer.Cards = append(customer.Cards, card.NewCard(customer.Name, customer.Bin))
}
