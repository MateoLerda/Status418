package enums

type FoodType int

const (
	Vegetable FoodType= iota
	Fruit
	Cheese
	Dairy
	Meat
)

func (f FoodType) String() string{
	return[]string{"Vegetable", "Fruit", "Cheese", "Dairy", "Meat"}[f]
}


