package bank

import (
	"crypto/rand"
	"strings"

	"github.com/glaukiol1/bank/bank_sys/card"
)

type Customer struct {
	Name       string       `json:"name"`
	DocumentId string       `json:"documentid"`
	Bin        string       `json:"bin"`
	Addresses  []Address    `json:"addresses"`
	Cards      []*card.Card `json:"cards"`
	Balance    int          `json:"balance"`
	Password   string       `json:"password"`
}

func createBin(name string) string {
	RandomCrypto, _ := rand.Prime(rand.Reader, 220)
	return RandomCrypto.String()[0:20] + "-" + strings.ReplaceAll(name, " ", "/") + "-" + RandomCrypto.String()[21:25]
}

func (customer *Customer) NewCard() (int, string) {
	if len(customer.Cards) < 3 {
		customer.Cards = append(customer.Cards, card.NewCard(customer.Name, customer.Bin))
		return 0, "Success!"
	} else {
		return 1, "Maximum amount of cards reached."
	}
}
