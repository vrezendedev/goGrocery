package db

import (
	"database/sql"
	"fmt"
	"goGrocery/configs"

	_ "github.com/lib/pq"
)

func OpenConnection() (*sql.DB, error) {
	cnf := configs.GetDB()

	sc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cnf.Host, cnf.Port, cnf.User, cnf.Password, cnf.Database,
	)

	cn, err := sql.Open("postgres", sc)

	if err != nil {
		panic(err)
	}

	err = cn.Ping()
	return cn, err
}
