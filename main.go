package main

import (
	"colegial_api/src/databases"
	routes "colegial_api/src/routes/students"
	route "colegial_api/src/routes/grades"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New();

	e.Use(middleware.CORS())

	databases.ConnectDb()
	routes.LoadRoutersStudent(e)
	route.LoadRoutersGrade(e)
	
	e.Start(":3000")
}