package router

import (
	"assignment-2/controllers"
	"assignment-2/database"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func StartApp() *gin.Engine {

	db := database.ConnectDB()
	router := gin.Default()
	orderController := controllers.NewOrderController(db)
	itemController := controllers.NewItemController(db)

	itemGroup := router.Group("/items")
	{
		itemGroup.GET("/", itemController.FindItems)
		itemGroup.GET("/:id", itemController.FindItemById)
		//itemGroup.POST("/", orderController.CreateOrder)
		itemGroup.PUT("/:id", itemController.UpdateItem)
		itemGroup.DELETE("/:id", itemController.DeleteItem)
	}

	orderGroup := router.Group("/orders")
	{
		orderGroup.POST("/", orderController.CreateOrder)
		orderGroup.GET("/", orderController.FindOrders)
		orderGroup.GET("/:id", orderController.FindOrderById)
		orderGroup.PUT("/:id", orderController.UpdateOrder)
		orderGroup.DELETE("/:id", orderController.DeleteOrder)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
