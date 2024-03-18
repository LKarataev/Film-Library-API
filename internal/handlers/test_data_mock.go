package handlers

import (
	"fmt"

	"github.com/LKarataev/Film-Library/internal/dao"
)

type AccountsRepositoryMock struct {
	MockAccounts []dao.Account
}

func NewAccountsRepositoryMock() *AccountsRepositoryMock {
	mockAccounts := []dao.Account{
		dao.Account{Username: "admin", Password: "admin_password"},
		dao.Account{Username: "user", Password: "user_password"},
	}
	arMock := AccountsRepositoryMock{MockAccounts: mockAccounts}
	return &arMock
}

func (accr AccountsRepositoryMock) GetAccount(username string) (*dao.Account, error) {
	switch username {
	case "admin":
		return &accr.MockAccounts[0], nil
	case "user":
		return &accr.MockAccounts[1], nil
	}
	return nil, fmt.Errorf("GetAccount error")
}

type ActorsRepositoryMock struct {
	MockActors []dao.Actor
}

func NewActorsRepositoryMock() *ActorsRepositoryMock {
	mockActors := []dao.Actor{
		dao.Actor{Id: 1, Name: "Thomas Hanks", Gender: "M", Birthday: "1956-07-09"},
		dao.Actor{Id: 2, Name: "Leonardo DiCaprio", Gender: "M", Birthday: "1974-11-11"},
		dao.Actor{Id: 3, Name: "Tobey Maguire", Gender: "M", Birthday: "1975-06-27"},
		dao.Actor{Id: 4, Name: "Keanu Reeves", Gender: "M", Birthday: "1964-09-02"},
		dao.Actor{Id: 5, Name: "Anne Hathaway", Gender: "F", Birthday: "1982-11-12"},
	}
	prMock := ActorsRepositoryMock{MockActors: mockActors}
	return &prMock
}

func (ar ActorsRepositoryMock) GetActorById(id int) (*dao.Actor, error) {
	switch id {
	case 1, 2, 3, 4, 5:
		return &ar.MockActors[id-1], nil
	}
	return nil, fmt.Errorf("GetActorById error")
}

func (ar ActorsRepositoryMock) InsertActor(values map[string]string) error {
	if len(values) == 0 {
		return fmt.Errorf("InsertActor error")
	}
	return nil
}

func (ar ActorsRepositoryMock) UpdateActor(id int, values map[string]string) error {
	if len(values) == 0 {
		return fmt.Errorf("UpdateActor error")
	}
	return nil
}

func (ar ActorsRepositoryMock) DeleteActor(id int) error {
	if id == 0 {
		return fmt.Errorf("DeleteActor error")
	}
	return nil
}

func (ar ActorsRepositoryMock) GetActors(limit, offset int) ([]dao.Actor, error) {
	if limit == 1 && offset == 0 {
		return ar.MockActors[0:1], nil
	}
	if limit == 2 && offset == 0 {
		return ar.MockActors[0:2], nil
	}
	return nil, fmt.Errorf("GetActors error")
}

func (ar ActorsRepositoryMock) GetActorsByFilmId(id int) ([]dao.Actor, error) {
	switch id {
	case 1, 2, 3, 4:
		return ar.MockActors[0:1], nil
	case 5, 6, 8:
		return ar.MockActors[1:2], nil
	case 7:
		return ar.MockActors[1:3], nil
	case 9, 10:
		return ar.MockActors[2:3], nil
	case 11, 12, 13:
		return ar.MockActors[3:4], nil
	case 14, 15, 16:
		return ar.MockActors[4:5], nil
	case 17:
		return nil, fmt.Errorf("GetActorsByFilmId error")
	}
	return nil, fmt.Errorf("GetActorsByFilmId error")
}

type FilmsRepositoryMock struct {
	MockFilms []dao.Film
}

func NewFilmsRepositoryMock() *FilmsRepositoryMock {
	mockFilms := []dao.Film{
		{Id: 1, Name: "The Green Mile", Year: 1999, Description: "drama", Rating: 9.1},
		{Id: 2, Name: "Forrest Gump", Year: 1994, Description: "drama, war, comedy, history", Rating: 8.9},
		{Id: 3, Name: "Catch Me If You Can", Year: 2002, Description: "crime, biography, comedy", Rating: 8.5},
		{Id: 4, Name: "Saving Private Ryan", Year: 2002, Description: "war, action, history", Rating: 8.2},
		{Id: 5, Name: "Inception", Year: 2010, Description: "sci-fi, action, thriller", Rating: 8.7},
		{Id: 6, Name: "Django Unchained", Year: 2012, Description: "western, action, drama, comedy", Rating: 8.2},
		{Id: 7, Name: "The Great Gatsby", Year: 2013, Description: "drama, romance", Rating: 7.9},
		{Id: 8, Name: "Titanic", Year: 1997, Description: "romance, history, drama", Rating: 8.4},
		{Id: 9, Name: "SpIder-Man", Year: 2002, Description: "sci-fi, action, adventure", Rating: 7.7},
		{Id: 10, Name: "Fear and Loathing in Las Vegas", Year: 1998, Description: "drama, comedy", Rating: 7.6},
		{Id: 11, Name: "The Matrix", Year: 1999, Description: "sci-fi, action", Rating: 8.5},
		{Id: 12, Name: "John Wick", Year: 2014, Description: "action, thriller, crime", Rating: 7.0},
		{Id: 13, Name: "Point Break", Year: 1991, Description: "action, thriller, crime", Rating: 7.8},
		{Id: 14, Name: "Interstellar", Year: 2014, Description: "sci-fi, drama, adventure", Rating: 8.6},
		{Id: 15, Name: "Les Miserables", Year: 2012, Description: "musical, drama, romance", Rating: 7.9},
		{Id: 16, Name: "The Devil Wears Prada", Year: 2006, Description: "drama, comedy", Rating: 7.7},
		{Id: 17, Name: "Edward Scissorhands", Year: 1990, Description: "fantasy, drama", Rating: 8.0},
	}
	prMock := FilmsRepositoryMock{MockFilms: mockFilms}
	return &prMock
}

func (fr FilmsRepositoryMock) GetFilms(search, sort, order string, limit, offset int) ([]dao.Film, error) {
	if search == "Gatsby" && limit > 0 && offset == 0 {
		return fr.MockFilms[6:7], nil
	}
	if search == "Scissorhands" && limit > 0 && offset == 0 {
		return fr.MockFilms[16:17], nil
	}
	if limit == 1 && offset == 0 {
		return fr.MockFilms[0:1], nil
	}
	return nil, fmt.Errorf("GetFilms error")
}

func (fr FilmsRepositoryMock) InsertFilm(values map[string]string, actorsId []int) error {
	if len(values) == 0 {
		return fmt.Errorf("InsertFilm error")
	}
	return nil
}

func (fr FilmsRepositoryMock) UpdateFilm(id int, values map[string]string) error {
	if len(values) == 0 {
		return fmt.Errorf("UpdateFilm error")
	}
	return nil
}

func (fr FilmsRepositoryMock) DeleteFilm(id int) error {
	if id == 0 {
		return fmt.Errorf("DeleteFilm error")
	}
	return nil
}

func (fr FilmsRepositoryMock) GetFilmsByActorId(id int) ([]dao.Film, error) {
	if id == 1 {
		return fr.MockFilms[0:4], nil
	}
	if id == 2 {
		return nil, fmt.Errorf("GetFilmsByActorId error")
	}
	return nil, fmt.Errorf("GetFilmsByActorId error")
}

func (fr FilmsRepositoryMock) GetFilmById(id int) (*dao.Film, error) {
	if id >= 1 && id <= 17 {
		return &fr.MockFilms[id-1], nil
	}
	return nil, fmt.Errorf("GetFilmById error")
}
