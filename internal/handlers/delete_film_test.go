package handlers

import (
	"context"
	"fmt"
	"testing"
)

type DeleteFilmCase struct {
	Ctx     context.Context
	Request DeleteFilmRequest
	Err     error
}

func TestDeleteFilm(t *testing.T) {
	cases := []DeleteFilmCase{
		DeleteFilmCase{
			Request: DeleteFilmRequest{
				Id: 1,
			},
			Err: nil,
			Ctx: context.Background(),
		},
		DeleteFilmCase{
			Request: DeleteFilmRequest{
				Id: 0,
			},
			Err: fmt.Errorf("DeleteFilm error"),
			Ctx: context.Background(),
		},
	}

	runDeleteFilmCases(t, cases)
}

func runDeleteFilmCases(t *testing.T, cases []DeleteFilmCase) {
	actorsRepo := NewActorsRepositoryMock()
	filmsRepo := NewFilmsRepositoryMock()

	for idx, item := range cases {
		caseName := fmt.Sprintf("case %d", idx)

		expectedErr := item.Err
		resultErr := NewDeleteFilmHandler(actorsRepo, filmsRepo).Handle(item.Ctx, item.Request)

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
