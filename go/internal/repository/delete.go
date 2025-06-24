package repository

import (
	db_config "github.com/Kovokar/codigo-no-minimo-suspeito/go/internal/db"
)

func DeleteAll() (msg string, err error) {
	conn, err := db_config.OpenConn()
	if err != nil {
		return
	}
	defer conn.Close()

	sql := `DELETE FROM acessos`

	_, err = conn.Exec(sql)
	if err != nil {
		return
	}

	msg = "Dados deletados com sucesso"
	return msg, nil
}
