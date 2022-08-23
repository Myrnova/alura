package models

import "gorm.io/gorm"

type Aluno struct {
	gorm.Model
	Nome string `json:"nome"`
	RG string `json:"rg"`
	CPF string `json:"cpf"`	
}