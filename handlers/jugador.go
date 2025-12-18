package handlers

import (
	"database/sql"
	"strconv"
	"example/db"
	"example/models"
	"example/dto"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gin-gonic/gin"
)

const (
	ErrDecodificar    = "Error al decodificar JSON"
	ErrCodificar      = "Error al codificar JSON"
	ErrNoEncontrado   = "Jugador no encontrado"
)

// --- Router ---
func MethodAsignment(r *gin.Engine) {

	
    r.POST("/jugadores",func(c *gin.Context){
	var nuevoJugador models.Jugador
	if err := c.ShouldBindJSON(&nuevoJugador);err!=nil{
		c.JSON(400,gin.H{"error":ErrDecodificar})
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

	r.GET("/jugadores",func(c *gin.Context){
		rows,err:= db.DB.Query("SELECT id,nombre,puntaje from jugadores")
		if err != nil{
			c.JSON(500,gin.H{"error":"error al consultar la base de datos","err": err.Error()})
			return
		}
		defer rows.Close()

		var jugadores []models.Jugador

		for rows.Next(){

		var j models.Jugador
		if err := rows.Scan(&j.Id, &j.Nombre, &j.Puntaje); err != nil {
			c.JSON(500,gin.H{"error":"no se pudieron recorrer las filas"})
			return
		}
		  
          jugadores= append(jugadores, j)
		  c.JSON(200,jugadores)
		}
	})

    r.PATCH("/jugadores/:id", func(c *gin.Context) {
    // 1. Convertir ID
    strId := c.Param("id")
    id, err := strconv.Atoi(strId)
    if err != nil {
        c.JSON(400, gin.H{"error": "id no válido"})
        return
    }

    // 2. Verificar si el jugador existe
    var jugador models.Jugador
    err = db.DB.QueryRow("SELECT id, nombre, puntaje FROM jugadores WHERE id = ?", id).
        Scan(&jugador.Id, &jugador.Nombre, &jugador.Puntaje)

    if err == sql.ErrNoRows {
        c.JSON(404, gin.H{"error": "el jugador no existe"})
        return
    } else if err != nil {
        c.JSON(500, gin.H{"error": "error al consultar la base de datos"})
        return
    }

    // 3. Bind del JSON (DTO)
    var datos dto.JugadorDTO
    if err := c.ShouldBindJSON(&datos); err != nil {
        c.JSON(400, gin.H{"error": "error al decodificar JSON"})
        return
    }

    // 4. Merge de los datos
    if datos.Nombre != nil {
        jugador.Nombre = *datos.Nombre
    }
    if datos.Puntaje != nil {
        if *datos.Puntaje < 0 {
            c.JSON(400, gin.H{"error": "el puntaje no puede ser negativo"})
            return
        }
        jugador.Puntaje = *datos.Puntaje
    }

    // 5. Ejecutar UPDATE
    _, err = db.DB.Exec(
        "UPDATE jugadores SET nombre = ?, puntaje = ? WHERE id = ?",
        jugador.Nombre,
        jugador.Puntaje,
        id,
    )

    if err != nil {
        c.JSON(500, gin.H{"error": "error al actualizar el jugador"})
        return
    }

    // 6. Respuesta final
    c.JSON(200, gin.H{
        "mensaje": "jugador actualizado correctamente",
        "jugador": jugador,
    })
    })

    r.DELETE("/jugadores/:id",func(c *gin.Context){
		strId:= c.Param("id")
		id,err:= strconv.Atoi(strId); if err!=nil{
			c.JSON(400,gin.H{"error":"id no valido"})
			return
		}	
	result,err:=db.DB.Exec("DELETE FROM jugadores WHERE id= ?",id) 
	if err != nil {
		c.JSON(500,gin.H{"error":"error en la base de datos"})
		return
	}

	row,_:= result.RowsAffected()
	if row==0{
		c.JSON(404,gin.H{"error":"jugador no encontrado"})
		return
	}

	c.JSON(200,gin.H{"mensaje":"jugador borrado exitosamente"})
	})



}
     
