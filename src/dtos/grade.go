package dtos

type GradeDto struct {
	IdStudent int `json:"id_aluno"`
	IdSubject int `json:"id_materia"`
	Grade float32 `json:"nota"`
}

type GradeStudentDto struct {
	Name string `json:"nome"`
	Grade float32 `json:"nota"`
}