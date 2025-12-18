package main

import (
	"example/db"
	"example/handlers"
	"log"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)





func main() {

     err:= godotenv.Load(); if err != nil{
	 log.Fatalf("Error loading .env file: %v", err) 
	 }
	 db.Conectar()
	 defer db.DB.Close()

	r:= gin.Default()
	r.Use(gin.Logger())
	handlers.MethodAsignment(r)
	log.Fatal(r.Run(":8080"))
	log.Println("Servidor iniciado en el puerto 8080")
}
