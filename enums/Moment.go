package enums

type Moment int

const (
	InvalidMoment Moment= iota
	Breakfast
	Lunch 
	Snack
	Dinner
	
)

func (m Moment) String() string{
	return[]string {"Breakfast","Lunch", "Snack", "Dinner"}[m]
}


func GetMomentEnum(c string) Moment{
	switch c {
	case "Lunch":
		return Lunch
	case "Snack":
		return Snack
	case "Dinner":
		return Dinner
	default:
		return InvalidMoment
	}
}