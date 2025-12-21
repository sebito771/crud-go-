package repository

import(
	"database/sql"
	"example/models"
	"example/dto"
	"strings"
	
)

type JugadorRepo struct {
  DB *sql.DB
}


func (db *JugadorRepo) Create(j models.Jugador )  (int64,error) {
	query := "INSERT INTO jugadores (nombre, puntaje) VALUES (?, ?)"
	result, err := db.DB.Exec(query, j.Nombre, j.Puntaje)
		if err != nil{
			return  0,err
		}
	 return result.LastInsertId()
}

func (db *JugadorRepo) GetAll() ([]models.Jugador, error) {
	query := "SELECT id, nombre, puntaje FROM jugadores"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var jugadores []models.Jugador
	for rows.Next() {
		var j models.Jugador
		err := rows.Scan(&j.Id, &j.Nombre, &j.Puntaje)
		if err != nil {
			return nil, err
		}
		jugadores = append(jugadores, j)
	}
	return jugadores, nil
}

func (db *JugadorRepo) GetById(id int64) (models.Jugador, error) {
	query := "SELECT id, nombre, puntaje FROM jugadores WHERE id = ?"
	var j models.Jugador
	err := db.DB.QueryRow(query, id).Scan(&j.Id, &j.Nombre, &j.Puntaje)
	return j, err
}

func (db *JugadorRepo) Delete(id int64) ( int64 ,  error){
	query:= "DELETE FROM jugadores WHERE id = ?"
	result, err := db.DB.Exec(query,id)
	if err!= nil{
       return 0,err
	}
	row,err:= result.RowsAffected(); if err !=nil {
		return 0 , err
	}
	return row,nil
}

func (r *JugadorRepo) UpdatePartial(id int64,dto dto.JugadorPatchDTO) (int64, error) {

	query := "UPDATE jugadores SET "
	//estos slices tienen como proposito guardar los argumentos del SQL(args) y los campos que se actualizaran(sets)
	args := []any{}
	sets := []string{}

	if dto.Nombre != nil {
		sets = append(sets, "nombre = ?")
		args = append(args, *dto.Nombre)
	}

	if dto.Puntaje != nil {
		sets = append(sets, "puntaje = ?")
		args = append(args, *dto.Puntaje)
	}

	query += strings.Join(sets, ", ")
	query += " WHERE id = ?"
	args = append(args, id)

	result, err := r.DB.Exec(query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}
