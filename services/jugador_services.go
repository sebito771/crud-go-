package services

import (
	"example/models"
	"example/repository"
	"errors"
	"strings"
	"database/sql"
	"example/dto"
)

type JugadorServices struct{
	repo repository.JugadorRepo
}

func NewJugadorService(r *repository.JugadorRepo) *JugadorServices {
	return &JugadorServices{
		repo: *r,
	}
}

func (s *JugadorServices) CreateJugador(j models.Jugador) (int64, error) {
	if strings.TrimSpace(j.Nombre) == "" {
		return 0, errors.New("el nombre es obligatorio")
	}

	if j.Puntaje < 0 {
		return 0, errors.New("el puntaje no puede ser negativo")
	}

	id, err := s.repo.Create(j)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *JugadorServices) ConsultarJugadores() ([]models.Jugador, error) {
	jugadores, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	return jugadores, nil
}

func (s *JugadorServices) ConsultarJugador(id int64) (models.Jugador, error) {
	jugador, err := s.repo.GetById(id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Jugador{}, errors.New("jugador no existe")
		}
		return models.Jugador{}, err
	}

	return jugador, nil
}


func (s *JugadorServices)BorrarJugador(id int64)error{
	row,err:=s.repo.Delete(id)
	if err!=nil {
		return err
	}

	if row == 0{
		return  errors.New("Jugador no encontrado")
	}

	return nil
	
}

func (s *JugadorServices) ActualizarJugador(
	id int64,
	dto dto.JugadorPatchDTO,
) error {

	// 1️⃣ Validar que venga al menos un campo
	if dto.Nombre == nil && dto.Puntaje == nil {
		return errors.New("no se enviaron campos para actualizar")
	}

	// 2️⃣ Validaciones de negocio (opcionales pero recomendadas)
	if dto.Nombre != nil && strings.TrimSpace(*dto.Nombre) == "" {
		return errors.New("el nombre no puede estar vacío")
	}

	if dto.Puntaje != nil && *dto.Puntaje < 0 {
		return errors.New("el puntaje no puede ser negativo")
	}

	// 3️⃣ Ejecutar update
	rows, err := s.repo.UpdatePartial(id, dto)
	if err != nil {
		return err
	}

	// 4️⃣ Verificar existencia
	if rows == 0 {
		return errors.New("jugador no existe")
	}

	return nil
}
