package handlers

import (
	"context"
	"log"

	"github.com/LKarataev/Film-Library/internal/dao"
)

type InsertActorHandler struct {
	actorsRepo dao.ActorsRepositoryInterface
	filmsRepo  dao.FilmsRepositoryInterface
}

type InsertActorRequest struct {
	Values map[string]string
}

func NewInsertActorHandler(actorsRepo dao.ActorsRepositoryInterface, filmsRepo dao.FilmsRepositoryInterface) InsertActorHandler {
	return InsertActorHandler{actorsRepo: actorsRepo, filmsRepo: filmsRepo}
}

func (h InsertActorHandler) Handle(ctx context.Context, req InsertActorRequest) error {
	log.Println("InsertActorRequest started")

	err := h.actorsRepo.InsertActor(req.Values)
	if err != nil {
		log.Println("InsertActorRequest error: ", err)
		return err
	}

	log.Println("InsertActorRequest successed")
	return nil
}
