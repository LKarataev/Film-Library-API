package dao

import (
	"database/sql"
	"fmt"
	"time"
)

type ActorsRepository struct {
	DB *sql.DB
}

type Actor struct {
	Id       int
	Name     string
	Gender   string
	Birthday string
}

type ActorsRepositoryInterface interface {
	GetActors(limit, offset int) ([]Actor, error)
	GetActorsByFilmId(id int) ([]Actor, error)
	GetActorById(id int) (*Actor, error)
	InsertActor(values map[string]string) error
	UpdateActor(id int, values map[string]string) error
	DeleteActor(id int) error
}

func (ar ActorsRepository) GetActors(limit, offset int) ([]Actor, error) {
	var actors []Actor
	rows, err := ar.DB.Query("SELECT id, name, gender, birthday FROM actors LIMIT $1 OFFSET $2", limit, offset)
	if err != nil {
		return actors, err
	}
	var actor Actor
	for rows.Next() {
		if err := rows.Scan(&actor.Id, &actor.Name, &actor.Gender, &actor.Birthday); err != nil {
			return nil, err
		}
		parsedTime, _ := time.Parse(time.RFC3339, actor.Birthday)
		actor.Birthday = parsedTime.Format("2006-01-02")
		actors = append(actors, actor)
	}
	return actors, nil
}

func (ar ActorsRepository) GetActorsByFilmId(id int) ([]Actor, error) {
	var actors []Actor
	rows, err := ar.DB.Query("SELECT id, name, gender, birthday FROM actors p JOIN films_actors fp ON p.id = fp.actor_id WHERE fp.film_id = $1", id)
	if err != nil {
		return nil, err
	}
	var actor Actor
	for rows.Next() {
		if err := rows.Scan(&actor.Id, &actor.Name, &actor.Gender, &actor.Birthday); err != nil {
			return nil, err
		}
		parsedTime, _ := time.Parse(time.RFC3339, actor.Birthday)
		actor.Birthday = parsedTime.Format("2006-01-02")
		actors = append(actors, actor)
	}
	return actors, nil
}

func (ar ActorsRepository) GetActorById(id int) (*Actor, error) {
	row := ar.DB.QueryRow("SELECT id, name, gender, birthday FROM actors WHERE id = $1", id)

	var actor Actor
	err := row.Scan(&actor.Id, &actor.Name, &actor.Gender, &actor.Birthday)
	if err != nil {
		return nil, err
	}

	parsedTime, _ := time.Parse(time.RFC3339, actor.Birthday)
	actor.Birthday = parsedTime.Format("2006-01-02")

	return &actor, nil
}

func (ar ActorsRepository) InsertActor(values map[string]string) error {
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
	result, err := ar.DB.Exec("INSERT INTO actors"+queryColumns+queryValues, valuesOrdered...)

	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected != 1 {
		return fmt.Errorf("Can't insert this actor in database")
	}
	return nil
}

func (ar ActorsRepository) UpdateActor(id int, values map[string]string) error {
	var valuesOrdered []interface{}
	var actorid interface{}
	querys := ""
	columnPos := 1

	for key, val := range values {
		querys += fmt.Sprintf(`"%s" = $%d, `, key, columnPos)
		valuesOrdered = append(valuesOrdered, val)
		columnPos++
	}
	valuesOrdered = append(valuesOrdered, id)

	querys = querys[:len(querys)-2]
	query := "UPDATE actors SET " + querys + " WHERE id = $" + fmt.Sprintf("%d", len(values)+1)
	result, err := ar.DB.Exec(query, valuesOrdered...)

	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected != 1 {
		return fmt.Errorf("Actor id (%d) not found in database", actorid)
	}
	return nil
}

func (ar ActorsRepository) DeleteActor(id int) error {
	_, err := ar.DB.Exec("DELETE FROM films_actors WHERE actor_id = $1", id)
	if err != nil {
		return err
	}

	result, err := ar.DB.Exec("DELETE FROM actors WHERE id = $1", id)

	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected != 1 {
		return fmt.Errorf("Actor id (%d) not found in database", id)
	}
	return nil
}
