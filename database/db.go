package database

import (
	"log"

	"github.com/murilorscampos/gin-api-rest/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {

	stringDeConexao := "host=localhost user=postgres password=root dbname=gin-api-rest port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexao))

	if err != nil {
		log.Panic("Erro ao conectar no banco de dados.")
	} else {
		log.Println("Banco de dados conectado...")
	}

	DB.AutoMigrate(&models.Aluno{})

}
