package servicesubject

import (
	"colegial_api/src/dtos"
	"database/sql"
)

type GradeService struct {
	DB *sql.DB
}

func NewGradeService(db *sql.DB) *GradeService {
	return &GradeService{
		DB: db,
	}
}

func GetGradeByStudent(service *GradeService, code int) ([]*dtos.GradeStudentDto, error) {
	rows, err := service.DB.Query(
		"SELECT d.nome, n.nota FROM disciplina d JOIN notas n ON d.id = n.id_disciplina WHERE n.id_aluno = $1", code)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var grades []*dtos.GradeStudentDto

	for rows.Next() {
		grade := &dtos.GradeStudentDto{}
		err := rows.Scan(&grade.Name, &grade.Grade)
		if err != nil {
			return nil, err
		}
		grades = append(grades, grade)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return grades, nil
}

func CreateNewGrade(service *GradeService, grade *dtos.GradeDto) error {
	stmt, err := service.DB.Prepare("INSERT INTO notas (id_aluno, id_disciplina, nota) VALUES ($1, $2, $3)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Executar a declaração SQL para inserir o aluno
	_, err = stmt.Exec(grade.IdStudent, grade.IdSubject, grade.Grade)
	if err != nil {
		return err
	}

	return nil
}