package routes

import (
	mysql "day-2/config"
	"day-2/controllers"

	"github.com/labstack/echo/v4"
)

func Api() {

	mysql.ConnectDB()
	route := echo.New()

	// Restricted group
	userRoute := route.Group("/user")
	userRoute.GET("", controllers.UserList)
	userRoute.POST("/store", controllers.UserStore)
	userRoute.GET("/show/:id", controllers.UserShow)
	userRoute.PUT("/update/:id", controllers.UserUpdate)
	userRoute.DELETE("/delete/:id", controllers.UserDelete)

	bookRoute := route.Group("/book")
	bookRoute.GET("", controllers.BookList)
	bookRoute.POST("/store", controllers.BookStore)
	bookRoute.GET("/show/:id", controllers.BookShow)
	bookRoute.PUT("/update/:id", controllers.BookUpdate)
	bookRoute.DELETE("/delete/:id", controllers.BookDelete)

	route.Start(":9000")
}
