package database

import (
	"gin-api-rest/src/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	stringDeConexao := "host=localhost user=root password=root dbname=root port=5433 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexao), &gorm.Config{})
	if err != nil {
		log.Fatal("Erro ao conectar com o banco de dados: ", err)
	}
	DB.AutoMigrate(&models.Aluno{})
}
