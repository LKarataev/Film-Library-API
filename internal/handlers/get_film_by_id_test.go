package handlers

import (
	"context"
	"fmt"
	"reflect"
	"testing"
)

type GetFilmByIdCase struct {
	Ctx     context.Context
	Request GetFilmByIdRequest
	Result  *GetFilmByIdResponse
	Err     error
}

func TestGetFilmById(t *testing.T) {
	cases := []GetFilmByIdCase{
		GetFilmByIdCase{
			Request: GetFilmByIdRequest{
				Id: 1,
			},
			Result: &GetFilmByIdResponse{
				&Film{Id: 1, Name: "The Green Mile", Year: 1999, Description: "drama", Rating: 9.1, Actors: []string{"Thomas Hanks"}},
			},
			Err: nil,
			Ctx: context.Background(),
		},
		GetFilmByIdCase{
			Request: GetFilmByIdRequest{
				Id: 0,
			},
			Result: nil,
			Err:    fmt.Errorf("GetFilmById error"),
			Ctx:    context.Background(),
		},
		GetFilmByIdCase{
			Request: GetFilmByIdRequest{
				Id: 17,
			},
			Result: nil,
			Err:    fmt.Errorf("GetActorsByFilmId error"),
			Ctx:    context.Background(),
		},
	}

	runGetFilmByIdCases(t, cases)
}

func runGetFilmByIdCases(t *testing.T, cases []GetFilmByIdCase) {
	actorsRepo := NewActorsRepositoryMock()
	filmsRepo := NewFilmsRepositoryMock()

	for idx, item := range cases {
		caseName := fmt.Sprintf("case %d", idx)

		expected := item.Result
		result, err := NewGetFilmByIdHandler(actorsRepo, filmsRepo).Handle(item.Ctx, item.Request)

		if result == nil && expected != nil {
			t.Fatalf("[%s] results not match\nGot : %#v\nWant: %#v", caseName, result, expected)
			continue
		}

		if expected == nil {
			if result != expected {
				t.Fatalf("[%s] results not match\nGot : %#v\nWant: %#v", caseName, result, expected)
				continue
			}
		} else if !reflect.DeepEqual(*result.Film, *expected.Film) {
			t.Fatalf("[%s] results not match\nGot : %#v\nWant: %#v", caseName, *result.Film, *expected.Film)
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
