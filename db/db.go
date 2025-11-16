package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

// Conectar abre la conexión a la BD
func Conectar() {
	var err error
	dsn := "root:Sena2025*@tcp(127.0.0.1:3306)/crud_db"
	conex, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	err = conex.Ping()
	if err != nil {
		panic(err)
	}
	DB = conex
	fmt.Println("Conexión exitosa a la BD 🚀")
}
