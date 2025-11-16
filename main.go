package main

import (
	"context"
	"example/db"
	"example/handlers"
	"log"
	"net/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"

	
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


// --- Router ---
func MethodAsignment(w http.ResponseWriter, r *http.Request) {
	sttr := strings.TrimPrefix(r.URL.Path, "/jugadores")
	sttr = strings.TrimPrefix(sttr, "/")

	if sttr == "" {
		switch r.Method {
		case http.MethodPost:
			handlers.CrearJugador(w, r)
		case http.MethodGet:
			handlers.ConsultarJugadores(w, r)
		default:
			http.Error(w, ErrMetodoNoPermit, http.StatusMethodNotAllowed)
		}
	} else {
		contx := context.WithValue(r.Context(), "id", sttr)
		r = r.WithContext(contx)
		switch r.Method {
		case http.MethodGet:
			handlers.Consultarjugador(w, r)
		case http.MethodPut:
			handlers.ActualizarJugador(w, r)
		case http.MethodDelete:
			handlers.EliminarJugador(w, r)
		default:
			http.Error(w, ErrMetodoNoPermit, http.StatusMethodNotAllowed)
		}
	}
}

func main() {
	

	db.Conectar()
	defer db.DB.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("/jugadores", MethodAsignment)
	mux.HandleFunc("/jugadores/", MethodAsignment)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
