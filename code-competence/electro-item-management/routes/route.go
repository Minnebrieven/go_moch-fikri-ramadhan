package routes

import (
	"electro-item-management/configs"
	"electro-item-management/handlers"
	m "electro-item-management/middlewares"
	"electro-item-management/repositories"
	"electro-item-management/services"

	mid "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func NewRoute(db *gorm.DB, e *echo.Echo) {
	//REPOSITORIES
	userRepository := repositories.NewUserRepository(db)
	categoryRepository := repositories.NewCategoryRepository(db)
	itemRepository := repositories.NewItemRepository(db)

	//SERVICES
	userService := services.NewUserService(userRepository)
	categoryService := services.NewCategoryService(categoryRepository)
	itemService := services.NewItemService(itemRepository, categoryRepository)

	//Handlers
	userHandler := handlers.NewUserHandler(userService)
	categoryHandler := handlers.NewCategoryHandler(categoryService)
	itemHandler := handlers.NewItemHandler(itemService)

	//Log
	m.LogMiddleware(e)

	//=======================ROUTES=====================================

	//------------------NOT AUTHENTICATED--------------------------------
	// USERS ROUTES
	e.POST("/login", userHandler.Login)
	e.POST("/register", userHandler.CreateUser)

	//------------------AUTHENTICATED------------------------------------
	jwtKey := configs.AppConfig.JWTKey
	// USERS ROUTES
	usersJWT := e.Group("/users")
	usersJWT.Use(mid.JWT([]byte(jwtKey)))
	usersJWT.GET("", userHandler.GetUsers)
	usersJWT.GET("/:id", userHandler.GetUserByID)
	usersJWT.PUT("/:id", userHandler.EditUser)
	usersJWT.DELETE("/:id", userHandler.DeleteUser)

	// CATEGORY ROUTES
	categoryJWT := e.Group("/category")
	categoryJWT.Use(mid.JWT([]byte(jwtKey)))
	categoryJWT.POST("", categoryHandler.CreateCategory)
	categoryJWT.GET("", categoryHandler.GetCategories)
	categoryJWT.GET("/:id", categoryHandler.GetCategoryByID)
	categoryJWT.PUT("/:id", categoryHandler.EditCategory)
	categoryJWT.DELETE("/:id", categoryHandler.DeleteCategory)

	itemsJWT := e.Group("/items")
	itemsJWT.Use(mid.JWT([]byte(jwtKey)))
	itemsJWT.POST("", itemHandler.CreateItem)
	itemsJWT.GET("", itemHandler.GetItems)
	itemsJWT.GET("/:id", itemHandler.GetItemByID)
	itemsJWT.PUT("/:id", itemHandler.EditItem)
	itemsJWT.DELETE("/:id", itemHandler.DeleteItem)
	itemsJWT.GET("/category/:category_id", itemHandler.GetItemsByCategory)
}
