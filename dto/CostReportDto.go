package dto

import (
	"time"
)

type CostReportDto struct{
	Month string `json:"month"`
	Count int `json:"count"`
}

func NewCostReport() []CostReportDto{
	var month = []string{"January", "February", "March", "April", "May", "June","July", "August", "September", "October", "November", "December"}
	monthNow:= int(time.Now().Month())
	var costReport []CostReportDto
	for  i := 0; i <= monthNow-1;  i++ {
		costReport= append(costReport, CostReportDto{Month: month[i], Count: 0} )
	}
	return costReport
}