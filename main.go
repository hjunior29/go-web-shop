package main

import (
	"Web/routes"
	"log"
	"net/http"
	"Web/migrate"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	migrate.MigrateDB()

	routes.CarregaRotas()
	http.ListenAndServe(":8080", nil)
}
