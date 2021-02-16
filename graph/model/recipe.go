package model

type Component struct {
	Type string `json:"type"`
	Recipe string `json:"recipe"`
	Ingredient string `json:"ingredient"`
}

type Recipe struct {
	ID                 string    `json:"id"`
	Name               string    `json:"name"`
	CurrentStock       int       `json:"currentStock"`
	MinimumStock       int       `json:"minimumStock"`
	MandatoryComponent Component `json:"mandatoryComponent"`
}
