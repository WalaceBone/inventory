package main

import "fmt"

type Ingredient struct {
	Name        string
	Id          string
	Quantity    uint32
	MinQuantity uint32
}

func (i *Ingredient) Increment(incr uint32) {
	i.Quantity += incr
}

func (i *Ingredient) Decrement(incr uint32) {
	i.Quantity -= incr
}

func (i *Ingredient) SetMinStock(stock uint32) {
	i.MinQuantity = stock
}

func CreateIngredient(name string, quantity uint32, min uint32) Ingredient {

	ingredient := Ingredient{
		Name:        name,
		Id:          "",
		Quantity:    quantity,
		MinQuantity: min,
	}
	return ingredient
}

func (i *Ingredient) dump() {
	fmt.Print(i.MinQuantity, i.Name, i.Quantity)
}

func main() {

	test := CreateIngredient("test", 32, 23)
	test.Increment(3)
	test.dump()
}
