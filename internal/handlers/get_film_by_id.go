package handlers

import (
	"context"
	"log"

	"github.com/LKarataev/Film-Library-API/internal/dao"
)

type GetFilmByIdHandler struct {
	actorsRepo dao.ActorsRepositoryInterface
	filmsRepo  dao.FilmsRepositoryInterface
}

type GetFilmByIdResponse struct {
	Film *Film
}

type GetFilmByIdRequest struct {
	Id int
}

func NewGetFilmByIdHandler(actorsRepo dao.ActorsRepositoryInterface, filmsRepo dao.FilmsRepositoryInterface) GetFilmByIdHandler {
	return GetFilmByIdHandler{actorsRepo: actorsRepo, filmsRepo: filmsRepo}
}

func (h GetFilmByIdHandler) Handle(ctx context.Context, req GetFilmByIdRequest) (*GetFilmByIdResponse, error) {
	log.Println("GetFilmByIdRequest started")

	film, err := h.filmsRepo.GetFilmById(req.Id)
	if err != nil {
		log.Println("GetFilmByIdRequest error: ", err)
		return nil, err
	}

	actors, err := h.actorsRepo.GetActorsByFilmId(film.Id)
	if err != nil {
		log.Println("GetFilmByIdRequest error: ", err)
		return nil, err
	}

	var names = make([]string, len(actors))
	for i, p := range actors {
		names[i] = p.Name
	}

	resp := GetFilmByIdResponse{}
	resp.Film = &Film{Id: film.Id, Name: film.Name, Year: film.Year, Description: film.Description, Rating: film.Rating, Actors: names}
	log.Println("GetFilmByIdRequest successed")
	return &resp, nil
}
