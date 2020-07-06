package main

import (
	"go-rest-test/controllers"
	"go-rest-test/internal/db"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	db.InitInMemoryDB()
	e := echo.New()

	// Middelware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	// Routes
	// e.POST("/users", controllers.CreateUser)
	e.GET("/users/:id", controllers.GetUser)
	e.GET("/users", controllers.GetAllUser)
	e.PUT("/users/:id", controllers.UpdateUser)
	e.DELETE("/users/:id", controllers.DeleteUser)

	// Start Server
	e.Logger.Fatal(e.Start(":8080"))
}
