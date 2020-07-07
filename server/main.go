package main

import (
	"fmt"
	"net/http"
	"stolencarsproject/server/app"
	carsController "stolencarsproject/server/app/cars"
	userController "stolencarsproject/server/app/users"
	"stolencarsproject/server/config"
	"stolencarsproject/server/internal/command"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	// Route => handler

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to Stolen Cars Project")
	})

	e.GET("/user/:email", userController.HandleGetUser)
	e.GET("/available/pos", userController.HandleGetAvailablePO)
	e.GET("/stolen/cars/:regID/:userType/:uID", carsController.HandleGetStolenCarDetails)
	e.GET("/stolen/cars/:uID/:userType", carsController.HandleGetAllStolenCarsForUser)

	e.POST("/user/new/:userType", userController.HandleCreateUser)
	e.POST("/stolen/car", carsController.HandleCreateStolenCarRecord)
	e.PATCH("/stolen/car/status", carsController.HandleChangeStolenCarStatus)
	e.POST("/stolen/car/assign/:uID", carsController.HandleAssignUnassignedStolenCar)

	// Server
	app.SetCommands()
	command.RunApp()
	appConfig := config.GetConfig()
	fmt.Println("Starting server at ", appConfig.Port)
	e.Run(standard.New(fmt.Sprintf(":%d", appConfig.Port)))

}
