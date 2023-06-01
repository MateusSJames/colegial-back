package controllergrade

import (
	"colegial_api/src/databases"
	servicesubject "colegial_api/src/services/service_subject"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"colegial_api/src/dtos"

	"github.com/labstack/echo/v4"
)

func GetGradesByStudent(c echo.Context) error {

	codeStudent:= c.Param("code")

	numCode, err:= strconv.Atoi(codeStudent)

	db, err := databases.ConnectDb()

	if err != nil {
		return err
	}

	serviceGrade:= servicesubject.NewGradeService(db);

	student, errStudent := servicesubject.GetGradeByStudent(serviceGrade, numCode)

	if errStudent != nil {
		fmt.Println(errStudent)
		return c.JSON(http.StatusInternalServerError, errStudent)
	}

	return c.JSON(http.StatusOK, student)
}

func CreateGrade(c echo.Context) error {

	body := c.Request().Body

	gradeDTO := new(dtos.GradeDto)
	if err := json.NewDecoder(body).Decode(gradeDTO); err != nil {
		return c.JSON(http.StatusBadRequest, "Erro ao decodificar o corpo da requisição")
	}

	db, err := databases.ConnectDb()

	if err != nil {
		return err
	}

	servicegrade:= servicesubject.NewGradeService(db);

	errgrade := servicesubject.CreateNewGrade(servicegrade, gradeDTO)

	if errgrade != nil {
		return c.JSON(http.StatusInternalServerError, errgrade)
	}

	return c.JSON(http.StatusCreated, "Nota cadastrada com sucesso!")
}
