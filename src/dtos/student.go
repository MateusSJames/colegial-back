package dtos

type StudentDto struct {
	Codigo int `json:"codigo"`
	Nome string `json:"nome"`
	Idade int `json:"idade"`
	Email string `json:"email"`
}