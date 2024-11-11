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
	r.Use(middlewares.CORSMiddleware())
	dependencies()
	routes()

	r.Static("/pages", "./pages")
	r.Static("/assets", "./assets")

	r.GET("/", func(c *gin.Context) {
		c.File("./pages/home/home.html")
	})

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
	recipesRoutes.PUT("/:recipeid", recipeHandler.Update) // falta modificar en el repositori
	recipesRoutes.POST("/", recipeHandler.Create)
	// recipesRoutes.PUT("/cook/:recipeid", recipeHandler.Cook) //falta probar solo (el de cocinar y el de desacer)
}

func dependencies() {
	var db repositories.DB

	var foodRepository repositories.FoodRepositoryInterface
	var foodService services.FoodServiceInterface
	var purchaseRepository repositories.PurchaseRepositoryInterface
	var purchaseService services.PurchaseServiceInterface
	var recipeRepository repositories.RecipeRepositoryInterface
	var recipeService services.RecipeServiceInterface

	db = repositories.NewMongoDB()
	recipeRepository = repositories.NewRecipeRepository(db)
	foodRepository = repositories.NewFoodRepository(db)
	purchaseRepository = repositories.NewPurchaseRepository(db)

	foodService = services.NewFoodService(foodRepository, recipeRepository)
	purchaseService = services.NewPurchaseService(purchaseRepository, foodRepository)
	recipeService = services.NewRecipeService(recipeRepository, foodRepository)

	foodHandler = handlers.NewFoodHandler(foodService)
	purchaseHandler = handlers.NewPurchaseHandler(purchaseService)
	recipeHandler = handlers.NewRecipeHandler(recipeService)
}
