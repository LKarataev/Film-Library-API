package handlers

import (
	"context"
	"log"

	"github.com/LKarataev/Film-Library/internal/dao"
)

type GetFilmsHandler struct {
	actorsRepo dao.ActorsRepositoryInterface
	filmsRepo  dao.FilmsRepositoryInterface
}

type GetFilmsResponse struct {
	Films []Film
}

type GetFilmsRequest struct {
	Limit  int
	Offset int
	Search string
	Sort   string
	Order  string
}

func NewGetFilmsHandler(actorsRepo dao.ActorsRepositoryInterface, filmsRepo dao.FilmsRepositoryInterface) GetFilmsHandler {
	return GetFilmsHandler{actorsRepo: actorsRepo, filmsRepo: filmsRepo}
}

func (h GetFilmsHandler) Handle(ctx context.Context, req GetFilmsRequest) (*GetFilmsResponse, error) {
	log.Println("GetFilmsRequest started")
	resp := GetFilmsResponse{Films: []Film{}}

	films, err := h.filmsRepo.GetFilms(req.Search, req.Sort, req.Order, req.Limit, req.Offset)
	if err != nil {
		log.Println("GetFilmsRequest error: ", err)
		return nil, err
	}
	for _, film := range films {
		actors, err := h.actorsRepo.GetActorsByFilmId(film.Id)
		if err != nil {
			log.Println("GetFilmsRequest error: ", err)
			return nil, err
		}
		var actorsStr = make([]string, len(actors))
		for i, str := range actors {
			actorsStr[i] = str.Name
		}
		resp.Films = append(resp.Films, Film{Id: film.Id, Name: film.Name, Year: film.Year, Description: film.Description, Rating: film.Rating, Actors: actorsStr})
	}
	log.Println("GetFilmsRequest successed")
	return &resp, nil
}
