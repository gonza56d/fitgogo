package structs

import (
	"github.com/google/uuid"
)

type Food struct {
	ID            uuid.UUID `json:"id"`
	Name          string    `json:"name"`
	Calories      int       `json:"calories"`
	Protein       float64   `json:"protein"`
	Carbohydrates float64   `json:"carbohydrates"`
	Fat           float64   `json:"fat"`
}

func NewFood(name string, calories int, protein, carbs, fat float64) *Food {
	return &Food{
		ID:            uuid.New(),
		Name:          name,
		Calories:      calories,
		Protein:       protein,
		Carbohydrates: carbs,
		Fat:           fat,
	}
}

type FoodRequest struct {
	Name          string  `json:"name" binding:"required"`
	Calories      int     `json:"calories" binding:"required"`
	Protein       float64 `json:"protein" binding:"required"`
	Carbohydrates float64 `json:"carbohydrates" binding:"required"`
	Fat           float64 `json:"fat" binding:"required"`
}
