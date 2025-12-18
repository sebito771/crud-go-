package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

// Conectar abre la conexión a la BD
func Conectar() {
	var err error
	//tomamos las variables de entorno
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	// Verificamos que las variables no estén vacías
	if user == "" || pass == "" || host == "" || port == "" || name == "" {
		panic("Las variables de entorno de la BD no están configuradas")
	}
	//formateamos el dsn
dsn := fmt.Sprintf(
	"%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local",
	user, pass, host, port, name,
)

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
