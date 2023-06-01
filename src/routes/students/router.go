package routes

import (

	"github.com/labstack/echo/v4"

	controllerstudent "colegial_api/src/controllers/student_controller"
)

func LoadRoutersStudent(api *echo.Echo) {
	api.GET("/students", func(c echo.Context) error {
		return controllerstudent.GetStudents(c)
	})

	api.GET("/students/:code", func(c echo.Context) error {
		return controllerstudent.GetStudentsByCode(c)
	})

	api.POST("/students/create", func(c echo.Context) error {
		return controllerstudent.CreateStudent(c)
	})
}

//TO DO => CADASTRAR AS NOTAS POR DISCIPLINA
// BUSCAR AS NOTAS POR ALUNO