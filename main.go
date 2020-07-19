package main

import (
	"fmt"
	"golangRestApi/database"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dataBaseConnection := database.InitDb()
	defer dataBaseConnection.Close()

	fmt.Println(dataBaseConnection)

	// r := chi.NewRouter()
	// r.Use(middleware.Logger)
	// r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("welcome"))
	// })
	// http.ListenAndServe(":3000", r)
}
