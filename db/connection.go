package db

import (
	"api-golang/configs"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func OpenConnection() (*sql.DB, error) {
	conf := configs.GetDB()

	sc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disabled",
		conf.Host, conf.Port, conf.User, conf.Pass, conf.Database)

	conn, err := sql.Open("postgress", sc)
	if err != nil {
		panic(err)
	}

	err = conn.Ping()

	return conn, err
}
