package card

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Card struct {
	CardNumber     int
	CVV            int
	ExpMonth       int
	ExpYear        int
	CardHolderName string
	CardHolderBin  string
}

func getExp() (int, int) {
	i, _ := strconv.ParseInt(fmt.Sprint(time.Now().Year())[2:4], 0, 0)
	return int(i + 7), 5
}

func NewCard(CardHolderName, CardHolderBin string) *Card {
	p, _ := rand.Prime(rand.Reader, 1000)
	cn := "8" + p.String()[4:19]
	cardNumber, _ := strconv.ParseInt(cn, 0, 0)
	cvv, _ := strconv.ParseInt(p.String()[0:3], 0, 0)
	yearExp, monthExp := getExp()
	return &Card{int(cardNumber), int(cvv), monthExp, yearExp, CardHolderName, CardHolderBin}
}

func (card *Card) parseData() string {
	return fmt.Sprint(card.CardNumber) + strings.ReplaceAll(card.CardHolderName, " ", "") + card.CardHolderBin + fmt.Sprint(card.CVV) + fmt.Sprint(card.ExpYear) + fmt.Sprint(card.ExpMonth)
}

func (card *Card) CardHash() string {
	hasher := sha256.New()
	hasher.Write([]byte(card.parseData()))
	println(card.parseData())
	sha := hasher.Sum(nil)
	dst := make([]byte, hex.EncodedLen(len(sha)))
	hex.Encode(dst, sha)
	return string(dst)
}
