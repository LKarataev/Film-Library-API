package handlers

import (
	"context"
	"fmt"
	"testing"
)

type InsertFilmCase struct {
	Ctx     context.Context
	Request InsertFilmRequest
	Err     error
}

func TestInsertFilm(t *testing.T) {
	cases := []InsertFilmCase{
		InsertFilmCase{
			Request: InsertFilmRequest{
				Values:   map[string]string{"name": "The Green Mile", "year": "1999", "description": "drama", "rating": "9.1"},
				ActorsId: []int{2, 3},
			},
			Err: nil,
			Ctx: context.Background(),
		},
		InsertFilmCase{
			Request: InsertFilmRequest{
				Values: map[string]string{},
			},
			Err: fmt.Errorf("InsertFilm error"),
			Ctx: context.Background(),
		},
	}

	runInsertFilmCases(t, cases)
}

func runInsertFilmCases(t *testing.T, cases []InsertFilmCase) {
	actorsRepo := NewActorsRepositoryMock()
	filmsRepo := NewFilmsRepositoryMock()

	for idx, item := range cases {
		caseName := fmt.Sprintf("case %d", idx)

		expectedErr := item.Err
		resultErr := NewInsertFilmHandler(actorsRepo, filmsRepo).Handle(item.Ctx, item.Request)

		if expectedErr != nil && resultErr != nil {
			if resultErr.Error() != expectedErr.Error() {
				t.Fatalf("[%s] results not match\nGot : %#v\nWant: %#v", caseName, resultErr, expectedErr)
			}
			continue
		}

		if !(expectedErr == nil && resultErr == nil) {
			t.Fatalf("[%s] results not match\nGot : %#v\nWant: %#v", caseName, resultErr, expectedErr)
			continue
		}
	}
}
