package mysql

import (
	"database/sql"
	"errors"
	"fmt"
)

type Message struct {
}

const (
	mysqlMessageCreateDatabase = iota
	mysqlMessageCreateTable
	mysqlMessageInsert
)

var (
	messageSQLString = []string{
		`CREATE DATABASE IF NOT EXISTS detail`,
		`CREATE TABLE IF NOT EXISTS detail.message (
			id INT NOT NULL AUTO_INCREMENT, 			
		    mobile VARCHAR(32) NOT NULL,
			code VARCHAR(32) NOT NULL,
			PRIMARY KEY(id)
		)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;`,
		`INSERT INTO detail.message (mobile, code) VALUES (?, ?)`,
	}
)

func CreateDatabase(db *sql.DB) error {
	_, err := db.Exec(messageSQLString[mysqlMessageCreateDatabase])
	return err
}

func CreateTable(db *sql.DB) {
	_, err := db.Exec(messageSQLString[mysqlMessageCreateTable])
	if err != nil {
		fmt.Println(err)
	}
}

func Insert(db *sql.DB, mobile, code string) {
	result, err := db.Exec(messageSQLString[mysqlMessageInsert], mobile, code)
	if err != nil {
		fmt.Println("errInsert", err)
	}

	if rows, _ := result.RowsAffected(); rows == 0 {
		fmt.Println(errors.New("errInvalidInsertTwo"))
	}
}
