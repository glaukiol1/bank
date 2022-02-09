package api

import "github.com/gin-gonic/gin"

func InitApi() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/getuserdata", userData)

	router.Run("localhost:8080")
}
