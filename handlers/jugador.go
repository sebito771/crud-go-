package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"example/db"
	"example/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gin-gonic/gin"
)

const (
	ContentType       = "Content-Type"
	ApplicationJSON   = "application/json"
	ErrMetodoNoPermit = "Método no permitido"
	ErrDecodificar    = "Error al decodificar JSON"
	ErrCodificar      = "Error al codificar JSON"
	ErrNoEncontrado   = "Jugador no encontrado"

)

// --- Router ---
func MethodAsignment(r *gin.Engine) {

	
    r.POST("/jugadores",func(c *gin.Context){
	var nuevoJugador models.Jugador
	if err := c.ShouldBindJSON(&nuevoJugador);err!=nil{
		c.JSON(400,gin.H{"error":err.Error()})
		return
	}
	
	if nuevoJugador.Nombre==""{
		c.JSON(400,gin.H{"error":"El nombre del jugador no puede ir vacio"})
		return
	}

	if nuevoJugador.Puntaje < 0 {
	  c.JSON(400,gin.H{"error":"el puntaje del jugador no puede ser negativo"})
	  return
	}

	result, err := db.DB.Exec("INSERT INTO jugadores (nombre, puntaje) VALUES (?, ?)", nuevoJugador.Nombre, nuevoJugador.Puntaje)
	if err != nil {
		c.JSON(500,gin.H{"error":err.Error()})
		return
	}

    id, err := result.LastInsertId()
    if err != nil {
        c.JSON(500, gin.H{"error": "no se pudo obtener el ID"})
        return
    }

	nuevoJugador.Id = int(id)
	c.JSON(201,gin.H{
		"mensaje":"jugador creado exitosamente",
        "jugador": nuevoJugador,
		})
	})

	r.GET("/jugadores/:id",func(c *gin.Context){
		strId:= c.Param("id")
		jugadorID, err := strconv.Atoi(strId)
        if err != nil {
    c.JSON(400, gin.H{"error": "el id debe ser un número"})
    return
}  

	  var player models.Jugador
	   err = db.DB.QueryRow("SELECT id, nombre, puntaje FROM jugadores WHERE id = ?", jugadorID).Scan(&player.Id, &player.Nombre, &player.Puntaje)
	if err == sql.ErrNoRows {
		c.JSON(404,gin.H{"error":"el id no existe en la base de datos"})
		return
	} else if err != nil {
	   c.JSON(500,gin.H{"error":err.Error()})
		return
	}
	
	c.JSON(200,gin.H{
		"jugador":player,
	})
	})
}

// --- Función auxiliar ---
func responderJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set(ContentType, ApplicationJSON)
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// --- Handlers ---
func CrearJugador(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, ErrMetodoNoPermit, http.StatusMethodNotAllowed)
		return
	}

	var nuevoJugador models.Jugador
	if err := json.NewDecoder(r.Body).Decode(&nuevoJugador); err != nil {
		http.Error(w, ErrDecodificar, http.StatusBadRequest)
		return
	}

	if nuevoJugador.Nombre == "" {
		http.Error(w, "El nombre no puede estar vacío", http.StatusBadRequest)
		return
	}
	if nuevoJugador.Puntaje < 0 {
		http.Error(w, "El puntaje no puede ser negativo", http.StatusBadRequest)
		return
	}

	result, err := db.DB.Exec("INSERT INTO jugadores (nombre, puntaje) VALUES (?, ?)", nuevoJugador.Nombre, nuevoJugador.Puntaje)
	if err != nil {
		http.Error(w, "Error al insertar en DB", http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	nuevoJugador.Id = int(id)

	responderJSON(w, http.StatusCreated, nuevoJugador)
}

func Consultarjugador(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, ErrMetodoNoPermit, http.StatusMethodNotAllowed)
		return
	}

	jugadorIDStr := r.Context().Value("id").(string)
	jugadorID, err := strconv.Atoi(jugadorIDStr)
	if err != nil {
		http.Error(w, ErrDecodificar, http.StatusBadRequest)
		return
	}

	var player models.Jugador
	err = db.DB.QueryRow("SELECT id, nombre, puntaje FROM jugadores WHERE id = ?", jugadorID).Scan(&player.Id, &player.Nombre, &player.Puntaje)
	if err == sql.ErrNoRows {
		http.Error(w, ErrNoEncontrado, http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Error al consultar DB", http.StatusInternalServerError)
		return
	}

	responderJSON(w, http.StatusOK, player)
}

func ConsultarJugadores(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, ErrMetodoNoPermit, http.StatusMethodNotAllowed)
		return
	}

	rows, err := db.DB.Query("SELECT id, nombre, puntaje FROM jugadores")
	if err != nil {
		http.Error(w, "Error al consultar DB", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var lista []models.Jugador
	for rows.Next() {
		var j models.Jugador
		if err := rows.Scan(&j.Id, &j.Nombre, &j.Puntaje); err != nil {
			http.Error(w, "Error al leer datos", http.StatusInternalServerError)
			return
		}
		lista = append(lista, j)
	}

	if len(lista) == 0 {
		http.Error(w, "No hay jugadores registrados", http.StatusNotFound)
		return
	}

	responderJSON(w, http.StatusOK, lista)
}

func ActualizarJugador(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, ErrMetodoNoPermit, http.StatusMethodNotAllowed)
		return
	}

	jugadorIDStr := r.Context().Value("id").(string)
	jugadorID, _ := strconv.Atoi(jugadorIDStr)

	var jugadorAct models.Jugador
	if err := json.NewDecoder(r.Body).Decode(&jugadorAct); err != nil {
		http.Error(w, ErrDecodificar, http.StatusBadRequest)
		return
	}

	if jugadorAct.Nombre == "" || jugadorAct.Puntaje < 0 {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	_, err := db.DB.Exec("UPDATE jugadores SET nombre = ?, puntaje = ? WHERE id = ?", jugadorAct.Nombre, jugadorAct.Puntaje, jugadorID)
	if err != nil {
		http.Error(w, "Error al actualizar DB", http.StatusInternalServerError)
		return
	}

	jugadorAct.Id = jugadorID
	responderJSON(w, http.StatusOK, jugadorAct)
}

func EliminarJugador(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, ErrMetodoNoPermit, http.StatusMethodNotAllowed)
		return
	}

	jugadorIDStr := r.Context().Value("id").(string)
	jugadorID, _ := strconv.Atoi(jugadorIDStr)

	result, err := db.DB.Exec("DELETE FROM jugadores WHERE id = ?", jugadorID)
	if err != nil {
		http.Error(w, "Error al eliminar DB", http.StatusInternalServerError)
		return
	}

	rowsAff, _ := result.RowsAffected()
	if rowsAff == 0 {
		http.Error(w, ErrNoEncontrado, http.StatusNotFound)
		return
	}

	responderJSON(w, http.StatusOK, map[string]string{"message": "Jugador eliminado"})
}

