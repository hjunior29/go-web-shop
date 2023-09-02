package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConectaComBD() *gorm.DB {

	conexaodb := fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_SSLMODE"),
	)

	db, err := gorm.Open(postgres.Open(conexaodb), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Erro ao conectar ao banco de dados: %v", err))
	}
	return db
}
