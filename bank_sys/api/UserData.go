package api

import (
	"github.com/gin-gonic/gin"
	"github.com/glaukiol1/bank/bank_sys/bank"
)

func userData(c *gin.Context) {
	var documentId = c.Query("documentid")
	var password = c.Query("password")
	for _, customer := range bank.GagBank.Customers {
		if customer.DocumentId == documentId {
			if customer.Password == password {
				c.IndentedJSON(200, customer)
				return
			}
		}
	}

	c.IndentedJSON(400, APIError{"Invalid Credentials"})
}
