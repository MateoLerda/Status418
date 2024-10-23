package main

import (
	"Status418/clients"
	"Status418/handlers"
	"Status418/middlewares"
	"Status418/repositories"
	"Status418/services"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	r               *gin.Engine
	foodHandler     *handlers.FoodHandler
	purchaseHandler *handlers.PurchaseHandler
	recipeHandler   *handlers.RecipeHandler
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	r = gin.Default()

	dependencies()
	routes()

	r.Run(":8080")
}

func routes() {
	authClient := clients.NewAuthClient()
	authMiddleware := middlewares.NewAuthMiddleware(authClient)

	foodRoutes := r.Group("/foods")
	foodRoutes.Use(authMiddleware.ValidateToken)
	foodRoutes.GET("/", foodHandler.GetAll)
	foodRoutes.GET("/:foodcode", foodHandler.GetByCode)
	foodRoutes.POST("/", foodHandler.Create)
	foodRoutes.DELETE("/:foodcode", foodHandler.Delete)
	foodRoutes.PUT("/:foodcode", foodHandler.Update)

	purchaseRoutes := r.Group("/purchases")
	purchaseRoutes.Use(authMiddleware.ValidateToken)
	purchaseRoutes.POST("/", purchaseHandler.Create)
	purchaseRoutes.GET("/", foodHandler.GetAll)

	recipesRoutes := r.Group("/recipes")
	recipesRoutes.Use(authMiddleware.ValidateToken)
	recipesRoutes.GET("/", recipeHandler.GetAll)
	recipesRoutes.DELETE("/:recipeid", recipeHandler.Delete)
	recipesRoutes.PUT("/:recipeid", recipeHandler.Update)
	recipesRoutes.POST("/", recipeHandler.Create)
	recipesRoutes.PUT("/cook/:recipeid",recipeHandler.Cook) //VER COMO HACER ESTO /STO ME PARECE QUE RECONTRA OUT AMIGO MATEITO PERDONAME
}

func dependencies() {
	var db repositories.DB

	var foodRepository repositories.FoodRepositoryInterface
	var foodService services.FoodServiceInterface

	db = repositories.NewMongoDB()

	foodRepository = repositories.NewFoodRepository(db)
	foodService = services.NewFoodService(foodRepository)
	foodHandler = handlers.NewFoodHandler(foodService)

	var purchaseRepository repositories.PurchaseRepositoryInterface
	var purchaseService services.PurchaseServiceInterface

	purchaseRepository = repositories.NewPurchaseRepository(db)
	purchaseService = services.NewPurchaseService(purchaseRepository)
	purchaseHandler = handlers.NewPurchaseHandler(purchaseService)

	var recipeRepository repositories.RecipeRepositoryInterface
	var recipeService services.RecipeServiceInterface

	recipeRepository = repositories.NewRecipeRepository(db)
	recipeService = services.NewRecipeService(recipeRepository)
	recipeHandler = handlers.NewRecipeHandler(recipeService)
}
