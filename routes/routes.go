package routes

import (
	"vuejs/collabs/constants"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New(e *echo.Echo) {

	e.POST("/users", controllers.CreateUserController)

	e.POST("/login", controllers.LoginUsersController)

	//JWT Group
	r := e.Group("/jwt")
	r.Use(middleware.JWT([]byte(constants.SECRET_JWT)))

	r.GET("/users/:id", controllers.GetUserDetailController)
	r.GET("/users", controllers.GetUserControllers)
	r.PUT("/users/:id", controllers.UpdateUserController)
	r.DELETE("/users/:id", controllers.DeleteUserController)

}
