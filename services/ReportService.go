package services

import (
	"Status418/dto"
	"Status418/models"
	"Status418/repositories"
)

type ReportServiceInterface interface {
	GetRecipesReport(userCode string, groupFilter bool) ([]dto.RecipeReportDto, error)
}

type ReportService struct {
	recipeRepository   repositories.RecipeRepositoryInterface
	foodRepository     repositories.FoodRepositoryInterface
	purchaseRepository repositories.PurchaseRepositoryInterface
}

func NewReportService(recipeRepository repositories.RecipeRepositoryInterface, foodRepository repositories.FoodRepositoryInterface, purchaseRepository repositories.PurchaseRepositoryInterface) *ReportService {
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

func (ReportService *ReportService) groupByRecipeMoment(recipes []models.Recipe) []dto.RecipeReportDto {
	var reports = dto.NewMomentReport()
	for _, recipe := range recipes {
		for i := range reports {
			if reports[i].Moment == recipe.Moment.String() {
				reports[i].Count++
				break
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
			for i := range reports {
				if reports[i].Type == food.Type.String() {
					reports[i].Count++
					break 
				}
			}
		}
	}
	return reports, nil
}
