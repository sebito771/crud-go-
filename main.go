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

		//CORS
	r.Use(cors.New(cors.Config{
    AllowOrigins: []string{
        "http://localhost:5500",
        "http://127.0.0.1:5500",
    },
    AllowMethods: []string{"GET", "POST", "PATCH", "DELETE"},
    AllowHeaders: []string{"Content-Type"},
}))
	handlers.MethodAsignment(r,service)



 
    
    log.Println("Servidor iniciado en el puerto 8080")
	log.Fatal(r.Run(":8080"))



}
