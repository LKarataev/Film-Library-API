package handlers

import (
	"context"
	"log"

	"github.com/LKarataev/Film-Library-API/internal/dao"
)

type DeleteFilmHandler struct {
	actorsRepo dao.ActorsRepositoryInterface
	filmsRepo  dao.FilmsRepositoryInterface
}

type DeleteFilmRequest struct {
	Id int
}

func NewDeleteFilmHandler(actorsRepo dao.ActorsRepositoryInterface, filmsRepo dao.FilmsRepositoryInterface) DeleteFilmHandler {
	return DeleteFilmHandler{actorsRepo: actorsRepo, filmsRepo: filmsRepo}
}

func (h DeleteFilmHandler) Handle(ctx context.Context, req DeleteFilmRequest) error {
	log.Println("DeleteFilmRequest started")

	err := h.filmsRepo.DeleteFilm(req.Id)
	if err != nil {
		log.Println("DeleteFilmRequest error: ", err)
		return err
	}

	log.Println("DeleteFilmRequest successed")
	return nil
}
