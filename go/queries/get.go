package queries

import (
	"fmt"

	"github.com/Kovokar/codigo-no-minimo-suspeito/go/db_config"
	"github.com/Kovokar/codigo-no-minimo-suspeito/go/models"
)

func GetAll() (sc []models.Acessos, err error) {

	conn, err := db_config.OpenConn()
	if err != nil {
		return
	}
	defer conn.Close()

	sql := `SELECT id, nome, momento FROM acessos`

	rows, err := conn.Query(sql)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var a models.Acessos
		err = rows.Scan(&a.Id, &a.Nome, &a.Momento)
		if err != nil {
			fmt.Println("Erro ao fazer scan:", err)
			continue
		}
		sc = append(sc, a)
	}

	err = rows.Err()
	return
}

func Count() (count int, err error) {
	conn, err := db_config.OpenConn()
	if err != nil {
		return
	}
	defer conn.Close()

	sql := `SELECT COUNT(*) FROM acessos`

	row := conn.QueryRow(sql)
	err = row.Scan(&count)

	if err != nil {
		return
	}

	return count, nil
}
