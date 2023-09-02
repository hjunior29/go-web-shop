package main

import (
	"go-web-shop/migrate"
	"go-web-shop/routes"
	"log"
	"net/http"

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
