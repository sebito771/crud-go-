package handlers

import (
	
	
	"example/dto"
	"example/models"
	"example/services"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

const (
	ErrDecodificar  = "Error al decodificar JSON"
	ErrCodificar    = "Error al codificar JSON"
	ErrNoEncontrado = "Jugador no encontrado"
)

// --- Router ---
func MethodAsignment(r *gin.Engine, service *services.JugadorServices) {

	r.POST("/jugadores", func(c *gin.Context) {
	var nuevoJugador models.Jugador

	if err := c.ShouldBindJSON(&nuevoJugador); err != nil {
		c.JSON(400, gin.H{"error": ErrDecodificar})
		return
	}

	id, err := service.CreateJugador(nuevoJugador)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	nuevoJugador.Id = int(id)

	c.JSON(201, gin.H{
		"mensaje": "jugador creado exitosamente",
		"jugador": nuevoJugador,
	})
      })

    r.GET("/jugadores/:id", func(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "id no válido"})
		return
	}

	jugador, err := service.ConsultarJugador(id)
	if err != nil {
		if err.Error() == "jugador no existe" {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, jugador)
      })

    r.GET("/jugadores", func(c *gin.Context) {
	jugadores, err := service.ConsultarJugadores()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, jugadores)
      })


    r.PATCH("/jugadores/:id", func(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "id no válido"})
		return
	}

	var dto dto.JugadorPatchDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(400, gin.H{"error": ErrDecodificar})
		return
	}

	err = service.ActualizarJugador(id, dto)
	if err != nil {
		switch err.Error() {
		case "jugador no existe":
			c.JSON(404, gin.H{"error": err.Error()})
		default:
			c.JSON(400, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(200, gin.H{"mensaje": "jugador actualizado correctamente"})
      })


	r.DELETE("/jugadores/:id", func(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "id no válido"})
		return
	}

	err = service.BorrarJugador(id)
	if err != nil {
		if err.Error() == "Jugador no encontrado" || err.Error() == "jugador no existe" {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"mensaje": "jugador borrado exitosamente"})
    })


}
