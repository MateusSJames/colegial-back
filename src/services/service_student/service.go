package servicestudent

import (
	"colegial_api/src/dtos"
	"database/sql"
)

type StudentService struct {
	DB *sql.DB
}

func NewStudentService(db *sql.DB) *StudentService {
	return &StudentService{
		DB: db,
	}
}

func GetStudents(service *StudentService) ([]*dtos.StudentDto, error) {
	rows, err := service.DB.Query("Select codigo, nome, idade, email FROM alunos")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var students []*dtos.StudentDto

	for rows.Next() {
		student := &dtos.StudentDto{}

		err := rows.Scan(&student.Codigo, &student.Nome, &student.Idade, &student.Email)

		if err != nil {
			return nil, err
		}

		students = append(students, student)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return students, nil
}

func GetStudentsByCode(service *StudentService, code int) (*dtos.StudentDto, error) {
	student:= &dtos.StudentDto{} 
	err := service.DB.QueryRow("Select codigo, nome, idade, email from alunos Where codigo = $1", code).Scan(&student.Codigo, &student.Nome, &student.Idade, &student.Email)

	if(err != nil) {
		return nil, err
	}

	return student, nil
}

func CreateNewStudent(service *StudentService, student *dtos.StudentDto) error {
	stmt, err := service.DB.Prepare("INSERT INTO alunos (codigo, nome, idade, email) VALUES ($1, $2, $3, $4)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Executar a declaração SQL para inserir o aluno
	_, err = stmt.Exec(student.Codigo, student.Nome, student.Idade,student.Email)
	if err != nil {
		return err
	}

	return nil
}