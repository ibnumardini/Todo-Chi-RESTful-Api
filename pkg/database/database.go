package database

import (
	"todo-chi-rest-api/pkg/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func OpenDbConnection() (*sqlx.DB, error) {
	conn, err := sqlx.Open(config.C.Db.Driver, config.C.Db.ConnStr())
	if err != nil {
		return nil, err
	}

	if err := conn.Ping(); err != nil {
		return nil, err
	}

	return conn, nil
}
