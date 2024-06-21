package db

import (
	"database/sql"
	"errors"
	constant "jwtAuth/constants"
	"log"
)

func NewDBconnection() (*sql.DB, error) {
	dbFileName := constant.GetEnvConstFromViper("SQLITE_DB_NAME")
	if dbFileName == "" {
		return nil, errors.New("viper fetch env variable failed")
	}
	db, err := sql.Open("sqlite3", dbFileName)
	if err != nil {
		log.Fatalf("Unable to open a new database connection", err)
		return nil, err
	}
	return db, nil
}
