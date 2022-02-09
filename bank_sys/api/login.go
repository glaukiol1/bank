package api

import (
	"github.com/gin-gonic/gin"
	"github.com/glaukiol1/bank/bank_sys/bank"
)

func userData(c *gin.Context) {
	var name = c.Query("name")
	var password = c.Query("password")
	for _, customer := range bank.GagBank.Customers {
		if customer.Name == name {
			if customer.Password == password {
				c.IndentedJSON(200, customer)
				return
			}
		}
	}
	c.IndentedJSON(400, "{\"message\":\"Couldn't find this user "+name+"\"}")
}
