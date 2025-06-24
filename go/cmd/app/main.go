package main

import (
	"fmt"
	"log"
	"time"

	db_config "github.com/Kovokar/codigo-no-minimo-suspeito/go/internal/db"
	queries "github.com/Kovokar/codigo-no-minimo-suspeito/go/internal/repository"
	"github.com/joho/godotenv"
)

func main() {
	// Carrega o .env
	Conn()

	intervalDelete := 10 * time.Second
	intervalCount := 3 * time.Second

	go func() {
		ticker := time.NewTicker(intervalDelete)
		defer ticker.Stop()

		for range ticker.C {
			Delete()
		}
	}()

	go func() {
		ticker := time.NewTicker(intervalCount)
		defer ticker.Stop()

		for range ticker.C {
			Count()
		}
	}()

	select {}
}

func Conn() {
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

func Delete() {
	acessos, err := queries.DeleteAll()
	if err != nil {
		log.Fatalf("Erro ao executar delete: %v", err)
	}

	fmt.Println(acessos)
}

func Count() {
	acessos, err := queries.Count()
	if err != nil {
		log.Fatalf("Erro ao executar count: %v", err)
	}

	fmt.Printf("Há %v registro(s) na tabela acessos\n", acessos)
}
