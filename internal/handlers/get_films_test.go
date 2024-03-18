package handlers

import (
	"context"
	"fmt"
	"reflect"
	"testing"
)

type GetFilmsCase struct {
	Ctx     context.Context
	Request GetFilmsRequest
	Result  *GetFilmsResponse
	Err     error
}

func TestGetFilms(t *testing.T) {
	cases := []GetFilmsCase{
		GetFilmsCase{
			Request: GetFilmsRequest{
				Limit:  1,
				Offset: 0,
				Search: "",
				Sort:   "",
				Order:  "",
			},
			Result: &GetFilmsResponse{
				[]Film{
					Film{Id: 1, Name: "The Green Mile", Year: 1999, Description: "drama", Rating: 9.1, Actors: []string{"Thomas Hanks"}},
				},
			},
			Err: nil,
			Ctx: context.Background(),
		},
		GetFilmsCase{
			Request: GetFilmsRequest{
				Limit:  4,
				Offset: 0,
				Search: "Gatsby",
				Sort:   "",
				Order:  "",
			},
			Result: &GetFilmsResponse{
				[]Film{
					Film{Id: 7, Name: "The Great Gatsby", Year: 2013, Description: "drama, romance", Rating: 7.9, Actors: []string{"Leonardo DiCaprio", "Tobey Maguire"}},
				},
			},
			Err: nil,
			Ctx: context.Background(),
		},
		GetFilmsCase{
			Request: GetFilmsRequest{
				Limit:  1,
				Offset: -1,
				Search: "",
				Sort:   "",
				Order:  "",
			},
			Result: nil,
			Err:    fmt.Errorf("GetFilms error"),
			Ctx:    context.Background(),
		},
		GetFilmsCase{
			Request: GetFilmsRequest{
				Limit:  1,
				Offset: 0,
				Search: "Scissorhands",
				Sort:   "",
				Order:  "",
			},
			Result: nil,
			Err:    fmt.Errorf("GetActorsByFilmId error"),
			Ctx:    context.Background(),
		},
	}

	runGetFilmsCases(t, cases)
}

func runGetFilmsCases(t *testing.T, cases []GetFilmsCase) {
	actorsRepo := NewActorsRepositoryMock()
	filmsRepo := NewFilmsRepositoryMock()

	for idx, item := range cases {
		caseName := fmt.Sprintf("case %d", idx)

		expected := item.Result
		result, err := NewGetFilmsHandler(actorsRepo, filmsRepo).Handle(item.Ctx, item.Request)

		if !reflect.DeepEqual(result, expected) {
			t.Fatalf("[%s] results not match\nGot : %#v\nWant: %#v", caseName, result, expected)
			continue
		}

		expectedErr := item.Err
		resultErr := err

		if !reflect.DeepEqual(resultErr, expectedErr) {
			t.Fatalf("[%s] results not match\nGot : %#v\nWant: %#v", caseName, resultErr, expectedErr)
			continue
		}
	}
}
