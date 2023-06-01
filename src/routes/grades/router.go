package routes

import (

	"github.com/labstack/echo/v4"

	controllergrade "colegial_api/src/controllers/grade_controller"
)

func LoadRoutersGrade(api *echo.Echo) {
	api.GET("/grades/:code", func(c echo.Context) error {
		return controllergrade.GetGradesByStudent(c)
	})

	api.POST("/grade/create", func(c echo.Context) error {
		return controllergrade.CreateGrade(c)
	})
}

//TO DO -> PEGAR DISCIPLINAS PELO ID