package migrate

import (
	"Web/db"
	"Web/models"
	"log"
)

func MigrateDB() {
	db := db.ConectaComBD()

	err := db.AutoMigrate(&models.Produto{})
	if err != nil {
		log.Fatalf("Erro na migração: %v", err)
	}
}
