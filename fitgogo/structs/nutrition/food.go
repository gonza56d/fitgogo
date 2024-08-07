package nutrition

import (
	"github.com/google/uuid"
)

type Food struct {
	ID            uuid.UUID
	Name          string
	Calories      int
	Protein       float64
	Carbohydrates float64
	Fat           float64
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
