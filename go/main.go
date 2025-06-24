package main

import (
	"fmt"
	"log"

	"github.com/Kovokar/codigo-no-minimo-suspeito/go/db_config"
	"github.com/joho/godotenv"
)

func main() {
	// Carrega o .env
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Erro ao carregar .env: %v", err)
	}

	db, err := db_config.OpenConn()
	if err != nil {
		log.Fatalf("Erro ao conectar no banco: %v", err)
	}
	defer db.Close()

	fmt.Println("Conexão com o banco estabelecida com sucesso!")
}
