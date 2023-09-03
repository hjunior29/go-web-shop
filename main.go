package main

import (
	"go-web-shop/migration"
	"go-web-shop/routes"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	migration.MigrateDB()

	routes.CarregaRotas()
	http.ListenAndServe(":8080", nil)
}
