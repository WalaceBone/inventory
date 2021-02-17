package model

type Component struct {
	ID       string `json:"id"`
	Type     string `json:"type"`
	Quantity int    `json:"quantity"`
}

type Recipe struct {
	ID                 string    `json:"id"`
	Name               string    `json:"name"`
	CurrentStock       int       `json:"currentStock"`
	MinimumStock       int       `json:"minimumStock"`
	MandatoryComponent Component `json:"mandatoryComponent"`
}
