package model

type Ingredient struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	CurrentStock int    `json:"currentStock"`
	MinimumStock int    `json:"minimumStock"`
}