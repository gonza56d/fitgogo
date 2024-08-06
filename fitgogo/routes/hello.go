package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func addHelloRoutes(rg *gin.RouterGroup) {
	ping := rg.Group("/hello")

	ping.GET("/world", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Hello world greet")
	})
}
