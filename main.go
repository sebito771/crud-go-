package main

import (
	"example/db"
	"example/handlers"
	"log"
	"github.com/gin-gonic/gin"
)

// --- Constantes ---
const (
	ContentType       = "Content-Type"
	ApplicationJSON   = "application/json"
	ErrMetodoNoPermit = "Método no permitido"
	ErrDecodificar    = "Error al decodificar JSON"
	ErrCodificar      = "Error al codificar JSON"
	ErrNoEncontrado   = "Jugador no encontrado"
)


// // --- Router ---
// func MethodAsignment(w http.ResponseWriter, r *http.Request) {
// 	sttr := strings.TrimPrefix(r.URL.Path, "/jugadores")
// 	sttr = strings.TrimPrefix(sttr, "/")

// 	if sttr == "" {
// 		switch r.Method {
// 		case http.MethodPost:
// 			handlers.CrearJugador(w, r)
// 		case http.MethodGet:
// 			handlers.ConsultarJugadores(w, r)
// 		default:
// 			http.Error(w, ErrMetodoNoPermit, http.StatusMethodNotAllowed)
// 		}
// 	} else {
// 		contx := context.WithValue(r.Context(), "id", sttr)
// 		r = r.WithContext(contx)
// 		switch r.Method {
// 		case http.MethodGet:
// 			handlers.Consultarjugador(w, r)
// 		case http.MethodPut:
// 			handlers.ActualizarJugador(w, r)
// 		case http.MethodDelete:
// 			handlers.EliminarJugador(w, r)
// 		default:
// 			http.Error(w, ErrMetodoNoPermit, http.StatusMethodNotAllowed)
// 		}
// 	}
// }

func main() {

	 db.Conectar()
	 defer db.DB.Close()

	r:= gin.Default()
	r.Use(gin.Logger())
	handlers.MethodAsignment(r)
	log.Fatal(r.Run(":8080"))
	log.Println("Servidor iniciado en el puerto 8080")
}
