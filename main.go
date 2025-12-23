package main

import (
	"example/db"
	"example/handlers"
	"log"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"example/repository"
	"example/services"
	"github.com/gin-contrib/cors"
)





func main() {

     err:= godotenv.Load(); if err != nil{
	 log.Fatalf("Error loading .env file: %v", err) 
	 }

	 db.Conectar()
	 defer db.DB.Close()

    repo:= repository.NewJugadorRepo(db.DB)
	service:= services.NewJugadorService(*repo)

	r:= gin.Default()
	r.Use(gin.Logger())
	handlers.MethodAsignment(r,service)
	r.Use(cors.Default())
 
    
    log.Println("Servidor iniciado en el puerto 8080")
	log.Fatal(r.Run(":8080"))



}
