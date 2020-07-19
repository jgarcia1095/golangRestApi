package database

import "database/sql"

func InitDb() *sql.DB {
	connectionString := "root:Qwerty.1@tcp(localhost:3306)/northwind"
	dataBaseConnection, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err.Error())
	}
	return dataBaseConnection
}
