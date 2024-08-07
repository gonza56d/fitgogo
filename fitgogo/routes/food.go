package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"fitgogo/services"
	"fitgogo/structs"
)

func addFoodRoutes(rg *gin.RouterGroup) {
	foodRoute := rg.Group("/food")

	foodRoute.POST("/", func(c *gin.Context) {
		var foodRequest structs.FoodRequest

		if err := c.ShouldBindJSON(&foodRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		food := structs.NewFood(
			foodRequest.Name,
			foodRequest.Calories,
			foodRequest.Protein,
			foodRequest.Carbohydrates,
			foodRequest.Fat,
		)
		services.CreateFood(food)
	})
}
