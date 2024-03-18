package handlers

import (
	"context"
	"log"

	"github.com/LKarataev/Film-Library/internal/dao"
)

type UpdateActorHandler struct {
	actorsRepo dao.ActorsRepositoryInterface
	filmsRepo  dao.FilmsRepositoryInterface
}

type UpdateActorResponse struct {
}

type UpdateActorRequest struct {
	Id     int
	Values map[string]string
}

func NewUpdateActorHandler(actorsRepo dao.ActorsRepositoryInterface, filmsRepo dao.FilmsRepositoryInterface) UpdateActorHandler {
	return UpdateActorHandler{actorsRepo: actorsRepo, filmsRepo: filmsRepo}
}

func (h UpdateActorHandler) Handle(ctx context.Context, req UpdateActorRequest) error {
	log.Println("UpdateActorRequest started")

	err := h.actorsRepo.UpdateActor(req.Id, req.Values)
	if err != nil {
		log.Println("UpdateActorRequest error: ", err)
		return err
	}

	log.Println("UpdateActorRequest successed")
	return nil
}
