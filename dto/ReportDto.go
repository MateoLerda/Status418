package dto

type ReportDto struct {
	Moment string `json:"moment"`
	Type   string `json:"type"`
	Count  int    `json:"count"`
}

func NewMomentReport() []ReportDto {
	var moments = []string{"Breakfast", "Lunch", "Snack", "Dinner"}
	var momentReport []ReportDto
	for _, moment := range moments {
		momentReport = append(momentReport, ReportDto{
			Type:   "",
			Moment: moment,
			Count:  0,
		})
	}
	return momentReport
}

func NewFoodReport() []ReportDto {
	var types = []string{"Vegetable", "Fruit", "Cheese", "Dairy", "Meat"}
	var foodReport []ReportDto

	for _, ftype := range types {
		foodReport = append(foodReport, ReportDto{
			Type:   ftype,
			Moment: "",
			Count:  0,
		})
	}
	return foodReport
}
