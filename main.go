package main

import (
	"Status418/handlers"
	"Status418/repositories"
	"Status418/services"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	r *gin.Engine
	foodHandler *handlers.FoodHandler
	purchaseHandler *handlers.PurchaseHandler
	userHandler *handlers.UserHandler
	recipeHandler *handlers.RecipeHandler
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	
	r = gin.Default()
		
	routes()
	dependencies()

	r.Run(":8080")
}

func routes() {
	foodRoutes := r.Group("/foods") 
	foodRoutes.GET("/:userId", foodHandler.GetAll)
	foodRoutes.GET("/:userId/:code", foodHandler.GetByCode)
	foodRoutes.POST("/", foodHandler.Create)
	foodRoutes.DELETE("/:code", foodHandler.Delete)
	foodRoutes.PUT("/:code", foodHandler.Update)
	
	purchaseRoutes := r.Group("/purchases")
	purchaseRoutes.POST("/:userId")
	purchaseRoutes.GET("/:userId")

	userRoutes:= r.Group("/users")
	userRoutes.GET("/")
	userRoutes.GET("/:userId")
	userRoutes.POST("/")
	userRoutes.PUT("/:userId")
	userRoutes.DELETE("/:userId")	

	recipesRoutes := r.Group("/recipes")
	recipesRoutes.GET("/:userId")
	recipesRoutes.DELETE("/:id")
	recipesRoutes.PUT("/:id")
	recipesRoutes.POST("/")
	
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
	
	var userRepository repositories.UserRepositoryInterface
	var userService services.UserServiceInterface

	userRepository = repositories.NewUserRepository(db)
	userService = services.NewUserService(userRepository)
	userHandler = handlers.NewUserHandler(userService)
	
	var recipeRepository repositories.RecipeRepositoryInterface
	var recipeService services.RecipeServiceInterface
	

	recipeRepository = repositories.NewRecipeRepository(db)
	recipeService = services.NewRecipeService(recipeRepository)
	recipeHandler = handlers.NewRecipeHandler(recipeService)
}
