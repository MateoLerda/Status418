package handlers

import (
	"Status418/services"
	"github.com/gin-gonic/gin"
	"Status418/utils"
)


type ReportHandler struct {
	reportService services.ReportServiceInterface
}

func NewReportHandler(reportService services.ReportServiceInterface) *ReportHandler {
	return &ReportHandler{
		reportService: reportService,
	}
}

func (reportHandler *ReportHandler) GetRecipeMomentReport(c *gin.Context) {
	user := utils.GetUserInfoFromContext(c)
	reportHandler.reportService.GetRecipesReport(user.Code, true)
}

func (reportHandler *ReportHandler) GetRecipeFoodTypeReport(c *gin.Context) {
	user := utils.GetUserInfoFromContext(c)
	reportHandler.reportService.GetRecipesReport(user.Code, false)
}
