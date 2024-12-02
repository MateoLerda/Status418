package services

import (
	"Status418/dto"
	"Status418/models"
	"Status418/repositories"
)

type ReportServiceInterface interface {
	GetRecipesReport(userCode string, groupFilter bool) ([]dto.RecipeReportDto, error)
	GetCostRepositor(userCode string) ([]dto.CostReportDto, error)
}

type ReportService struct {
	recipeRepository   repositories.RecipeRepositoryInterface
	foodRepository     repositories.FoodRepositoryInterface
	purchaseRepository repositories.PurchaseRepository
}

func NewReportService(recipeRepository repositories.RecipeRepositoryInterface, foodRepository repositories.FoodRepositoryInterface, purchaseRepository repositories.PurchaseRepository) *ReportService {
	return &ReportService{
		recipeRepository:   recipeRepository,
		foodRepository:     foodRepository,
		purchaseRepository: purchaseRepository,
	}
}

func (reportService *ReportService) GetRecipesReport(userCode string, groupFilter bool) ([]dto.RecipeReportDto, error) {
	recipes, err := reportService.recipeRepository.GetAll(userCode, models.Filter{})
	var reports []dto.RecipeReportDto
	if err != nil {
		return nil, err
	}

	if groupFilter {
		reports = reportService.groupByRecipeMoment(recipes)
	} else {
		reports, err = reportService.groupRecipesByFoodType(recipes)
		if err != nil {
			return nil, err
		}
	}
	return reports, nil
}

func (reportService *ReportService) GetCostRepositor(userCode string) ([]dto.CostReportDto, error) {
	purchase, err := reportService.purchaseRepository.GetAll(userCode)
}

func (ReportService *ReportService) groupByRecipeMoment(recipes []models.Recipe) []dto.RecipeReportDto {
	var reports = dto.NewMomentReport()
	for _, recipe := range recipes {
		for _, report := range reports {
			if report.Moment == recipe.Moment.String() {
				report.Count++
			}
		}
	}
	return reports
}

func (ReportService *ReportService) groupRecipesByFoodType(recipes []models.Recipe) ([]dto.RecipeReportDto, error) {
	var reports = dto.NewFoodReport()

	for _, recipe := range recipes {
		for _, ingredient := range recipe.Ingredients {
			food, err := ReportService.foodRepository.GetByCode(ingredient.FoodCode, recipe.UserCode)
			if err != nil {
				return nil, err
			}
			for _, report := range reports {
				if report.Type == food.Type.String() {
					report.Count++
				}
			}
		}
	}
	return reports, nil
}
