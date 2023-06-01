package controllerstudent

import (
	"colegial_api/src/databases"
	servicestudent "colegial_api/src/services/service_student"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"colegial_api/src/dtos"

	"github.com/labstack/echo/v4"
)

func GetStudents(c echo.Context) error {
	db, err := databases.ConnectDb()
	
	if err != nil {		
		return err
	}

	serviceStudent := servicestudent.NewStudentService(db);

	listStudents, errStudent := servicestudent.GetStudents(serviceStudent);

	if errStudent != nil {
		fmt.Println(errStudent)
		return c.JSON(http.StatusInternalServerError, errStudent)
	}
	
	return c.JSON(http.StatusOK, listStudents);
}

func GetStudentsByCode(c echo.Context) error {

	codeStudent:= c.Param("code")

	numCode, err:= strconv.Atoi(codeStudent)

	db, err := databases.ConnectDb()

	if err != nil {
		return err
	}

	serviceStudent:= servicestudent.NewStudentService(db);

	student, errStudent := servicestudent.GetStudentsByCode(serviceStudent, numCode)

	if errStudent != nil {
		fmt.Println(errStudent)
		return c.JSON(http.StatusInternalServerError, errStudent)
	}

	return c.JSON(http.StatusOK, student)
}

func CreateStudent(c echo.Context) error {

	body := c.Request().Body

	studentDTO := new(dtos.StudentDto)
	if err := json.NewDecoder(body).Decode(studentDTO); err != nil {
		return c.JSON(http.StatusBadRequest, "Erro ao decodificar o corpo da requisição")
	}

	db, err := databases.ConnectDb()

	if err != nil {
		return err
	}

	serviceStudent:= servicestudent.NewStudentService(db);

	errStudent := servicestudent.CreateNewStudent(serviceStudent, studentDTO)

	if errStudent != nil {
		return c.JSON(http.StatusInternalServerError, errStudent)
	}

	return c.JSON(http.StatusCreated, "Estudante cadastrado com sucesso!")
}
