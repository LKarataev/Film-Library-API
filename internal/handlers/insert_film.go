package handlers

import (
	"context"
	"log"

	"github.com/LKarataev/Film-Library-API/internal/dao"
)

type InsertFilmHandler struct {
	actorsRepo dao.ActorsRepositoryInterface
	filmsRepo  dao.FilmsRepositoryInterface
}

type InsertFilmRequest struct {
	Values   map[string]string
	ActorsId []int
}

func NewInsertFilmHandler(actorsRepo dao.ActorsRepositoryInterface, filmsRepo dao.FilmsRepositoryInterface) InsertFilmHandler {
	return InsertFilmHandler{actorsRepo: actorsRepo, filmsRepo: filmsRepo}
}

func (h InsertFilmHandler) Handle(ctx context.Context, req InsertFilmRequest) error {
	log.Println("InsertFilmRequest started")

	err := h.filmsRepo.InsertFilm(req.Values, req.ActorsId)
	if err != nil {
		log.Println("InsertFilmRequest error: ", err)
		return err
	}

	log.Println("InsertFilmRequest successed")
	return nil
}
