package enums

type Moment int

const (
	Breakfast Moment= iota
	Lunch 
	Snack
	Dinner
)

func (m Moment) String() string{
	return[]string {"Breakfast","Lunch", "Snack", "Dinner"}[m]
}