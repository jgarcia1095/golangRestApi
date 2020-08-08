package database

import (
	"database/sql"

	"github.com/golangRestApi/helper"
)

func InitDb() *sql.DB {
	connectionString := "root:Qwerty.1@tcp(localhost:3306)/northwind"
	dataBaseConnection, err := sql.Open("mysql", connectionString)
	helper.Catch(err)
	return dataBaseConnection
}
