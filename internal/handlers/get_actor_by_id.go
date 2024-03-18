package handlers

import (
	"context"
	"log"

	"github.com/LKarataev/Film-Library/internal/dao"
)

type GetActorByIdHandler struct {
	actorsRepo dao.ActorsRepositoryInterface
	filmsRepo  dao.FilmsRepositoryInterface
}

type GetActorByIdResponse struct {
	Actor *Actor
}

type GetActorByIdRequest struct {
	Id int
}

func NewGetActorByIdHandler(actorsRepo dao.ActorsRepositoryInterface, filmsRepo dao.FilmsRepositoryInterface) GetActorByIdHandler {
	return GetActorByIdHandler{actorsRepo: actorsRepo, filmsRepo: filmsRepo}
}

func (h GetActorByIdHandler) Handle(ctx context.Context, req GetActorByIdRequest) (*GetActorByIdResponse, error) {
	log.Println("GetActorByIdRequest started")

	actor, err := h.actorsRepo.GetActorById(req.Id)
	if err != nil {
		log.Println("GetActorByIdRequest error: ", err)
		return nil, err
	}

	films, err := h.filmsRepo.GetFilmsByActorId(actor.Id)
	if err != nil {
		log.Println("GetActorByIdRequest error: ", err)
		return nil, err
	}

	var names = make([]string, len(films))
	for i, f := range films {
		names[i] = f.Name
	}

	resp := GetActorByIdResponse{}
	resp.Actor = &Actor{Id: actor.Id, Name: actor.Name, Gender: actor.Gender, Birthday: actor.Birthday, Films: names}
	log.Println("GetActorByIdRequest successed")
	return &resp, nil
}
