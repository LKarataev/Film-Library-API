package handlers

import (
	"context"
	"log"

	"github.com/LKarataev/Film-Library/internal/dao"
)

type GetActorsHandler struct {
	actorsRepo dao.ActorsRepositoryInterface
	filmsRepo  dao.FilmsRepositoryInterface
}

type GetActorsResponse struct {
	Actors []Actor
}

type GetActorsRequest struct {
	Limit  int
	Offset int
}

func NewGetActorsHandler(actorsRepo dao.ActorsRepositoryInterface, filmsRepo dao.FilmsRepositoryInterface) GetActorsHandler {
	return GetActorsHandler{actorsRepo: actorsRepo, filmsRepo: filmsRepo}
}

func (h GetActorsHandler) Handle(ctx context.Context, req GetActorsRequest) (*GetActorsResponse, error) {
	log.Println("GetActorsRequest started")
	resp := GetActorsResponse{Actors: []Actor{}}

	actors, err := h.actorsRepo.GetActors(req.Limit, req.Offset)
	if err != nil {
		log.Println("GetActorsRequest error: ", err)
		return nil, err
	}

	for _, actor := range actors {
		films, err := h.filmsRepo.GetFilmsByActorId(actor.Id)
		if err != nil {
			log.Println("GetActorsRequest error: ", err)
			return nil, err
		}
		var filmsStr = make([]string, len(films))
		for i, str := range films {
			filmsStr[i] = str.Name
		}
		resp.Actors = append(resp.Actors, Actor{Id: actor.Id, Name: actor.Name, Gender: actor.Gender, Birthday: actor.Birthday, Films: filmsStr})
	}
	log.Println("GetActorsRequest successed")
	return &resp, nil
}
