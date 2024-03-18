package handlers

import (
	"context"
	"log"

	"github.com/LKarataev/Film-Library/internal/dao"
)

type DeleteActorHandler struct {
	actorsRepo dao.ActorsRepositoryInterface
	filmsRepo  dao.FilmsRepositoryInterface
}

type DeleteActorRequest struct {
	Id int
}

func NewDeleteActorHandler(actorsRepo dao.ActorsRepositoryInterface, filmsRepo dao.FilmsRepositoryInterface) DeleteActorHandler {
	return DeleteActorHandler{actorsRepo: actorsRepo, filmsRepo: filmsRepo}
}

func (h DeleteActorHandler) Handle(ctx context.Context, req DeleteActorRequest) error {
	log.Println("DeleteActorRequest started")

	err := h.actorsRepo.DeleteActor(req.Id)
	if err != nil {
		log.Println("DeleteActorRequest error: ", err)
		return err
	}

	log.Println("DeleteActorRequest successed")
	return nil
}
