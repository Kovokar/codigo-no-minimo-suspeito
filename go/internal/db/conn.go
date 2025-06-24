package db_config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

// OpenConn abre uma conexão com o banco de dados usando variáveis de ambiente
func OpenConn() (*sql.DB, error) {

	connStr := ConnectionString()

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	// Testa a conexão
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

// ConnectionString monta a string de conexão baseada nas variáveis de ambiente
func ConnectionString() string {

	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)
}
