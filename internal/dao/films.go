package dao

import (
	"database/sql"
	"fmt"
)

type FilmsRepository struct {
	DB *sql.DB
}

type Film struct {
	Id          int
	Name        string
	Year        int
	Description string
	Rating      float64
}

type FilmsRepositoryInterface interface {
	GetFilmsByActorId(id int) ([]Film, error)
	GetFilmById(id int) (*Film, error)
	GetFilms(search, sort, order string, limit, offset int) ([]Film, error)
	InsertFilm(values map[string]string, actorsId []int) error
	UpdateFilm(id int, values map[string]string) error
	DeleteFilm(id int) error
}

func (fr FilmsRepository) GetFilmsByActorId(id int) ([]Film, error) {
	var films []Film
	rows, err := fr.DB.Query(`SELECT f.id, f.name, f.year, f.description, f.rating FROM films f JOIN films_actors fp ON f.id = fp.film_id WHERE fp.actor_id = $1`, id)
	if err != nil {
		return films, err
	}
	var film Film
	for rows.Next() {
		if err := rows.Scan(&film.Id, &film.Name, &film.Year, &film.Description, &film.Rating); err != nil {
			return nil, err
		}
		films = append(films, film)
	}
	return films, nil
}

func (fr FilmsRepository) GetFilmById(id int) (*Film, error) {
	row := fr.DB.QueryRow("SELECT id, name, year, description, rating FROM films WHERE id = $1", id)

	var film Film
	err := row.Scan(&film.Id, &film.Name, &film.Year, &film.Description, &film.Rating)
	if err != nil {
		return nil, err
	}

	return &film, nil
}

func (fr FilmsRepository) GetFilms(search, sort, order string, limit, offset int) ([]Film, error) {
	var films []Film

	query := "SELECT DISTINCT f.id, f.name, f.year, f.description, f.rating " +
		"FROM films f JOIN films_actors fp ON f.id = fp.film_id "
	if search != "" {
		query += "WHERE f.name LIKE '%" + search + "%' OR fp.actor_id IN " +
			"( SELECT id FROM actors WHERE name LIKE '%" + search + "%' ) "
	}

	query += "ORDER BY "

	switch sort {
	case "rating":
		query += "f.rating "
	case "name":
		query += "f.name "
	case "year":
		query += "f.year "
	default:
		query += "f.rating "
	}

	switch order {
	case "asc":
		query += "ASC "
	case "desc":
		query += "DESC "
	default:
		query += "ASC "
	}

	query += "LIMIT $1 OFFSET $2"

	rows, err := fr.DB.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var f Film
	for rows.Next() {
		if err := rows.Scan(&f.Id, &f.Name, &f.Year, &f.Description, &f.Rating); err != nil {
			return nil, err
		}
		films = append(films, f)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return films, nil
}

func (fr FilmsRepository) InsertFilm(values map[string]string, actorsId []int) error {
	var valuesOrdered []interface{}

	queryColumns := " ("
	queryValues := " VALUES ("
	columnPos := 1

	for key, val := range values {
		queryColumns += fmt.Sprintf(`"%s", `, key)
		queryValues += fmt.Sprintf("$%d, ", columnPos)
		valuesOrdered = append(valuesOrdered, val)
		columnPos++
	}

	queryColumns = queryColumns[:len(queryColumns)-2] + ")"
	queryValues = queryValues[:len(queryValues)-2] + ")"
	query := "INSERT INTO films" + queryColumns + queryValues + " RETURNING id"

	row := fr.DB.QueryRow(query, valuesOrdered...)

	var filmId int = -1
	err := row.Scan(&filmId)
	if err != nil {
		return err
	}

	for _, actorId := range actorsId {
		fr.DB.QueryRow("INSERT INTO films_actors (film_id, actor_id) VALUES ($1, $2)", filmId, actorId)
	}
	return nil
}

func (fr FilmsRepository) UpdateFilm(id int, values map[string]string) error {
	querys := ""
	columnPos := 1
	var valuesOrdered []interface{}
	for key, val := range values {
		querys += fmt.Sprintf(`"%s" = $%d, `, key, columnPos)
		valuesOrdered = append(valuesOrdered, val)
		columnPos++
	}
	valuesOrdered = append(valuesOrdered, id)
	query := "UPDATE films SET " + querys[:len(querys)-2] + " WHERE id = $" + fmt.Sprintf("%d", len(values)+1)
	result, err := fr.DB.Exec(query, valuesOrdered...)

	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected != 1 {
		return fmt.Errorf("Film id (%v) not found in database", id)
	}
	return nil
}

func (fr FilmsRepository) DeleteFilm(id int) error {

	_, err := fr.DB.Exec("DELETE FROM films_actors WHERE film_id = $1", id)
	if err != nil {
		return err
	}

	result, err := fr.DB.Exec("DELETE FROM films WHERE id = $1", id)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected != 1 {
		return fmt.Errorf("Film id (%v) not found in database", id)
	}

	return nil
}
