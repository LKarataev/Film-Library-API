package handlers

import (
	"context"
	"log"

	"github.com/LKarataev/Film-Library-API/internal/dao"
)

type UpdateFilmHandler struct {
	actorsRepo dao.ActorsRepositoryInterface
	filmsRepo  dao.FilmsRepositoryInterface
}

type UpdateFilmRequest struct {
	Id     int
	Values map[string]string
}

func NewUpdateFilmHandler(actorsRepo dao.ActorsRepositoryInterface, filmsRepo dao.FilmsRepositoryInterface) UpdateFilmHandler {
	return UpdateFilmHandler{actorsRepo: actorsRepo, filmsRepo: filmsRepo}
}

func (h UpdateFilmHandler) Handle(ctx context.Context, req UpdateFilmRequest) error {
	log.Println("UpdateFilmRequest started")

	err := h.filmsRepo.UpdateFilm(req.Id, req.Values)
	if err != nil {
		log.Println("UpdateFilmRequest error: ", err)
		return err
	}

	log.Println("UpdateFilmRequest successed")
	return nil
}
