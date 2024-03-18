package handlers

import (
	"context"
	"fmt"
	"testing"
)

type UpdateFilmCase struct {
	Ctx     context.Context
	Request UpdateFilmRequest
	Err     error
}

func TestUpdateFilm(t *testing.T) {
	cases := []UpdateFilmCase{
		UpdateFilmCase{
			Request: UpdateFilmRequest{
				Values: map[string]string{"name": "The Green Kilometer"},
				Id:     1,
			},
			Err: nil,
			Ctx: context.Background(),
		},
		UpdateFilmCase{
			Request: UpdateFilmRequest{
				Id: 0,
			},
			Err: fmt.Errorf("UpdateFilm error"),
			Ctx: context.Background(),
		},
	}

	runUpdateFilmCases(t, cases)
}

func runUpdateFilmCases(t *testing.T, cases []UpdateFilmCase) {
	actorsRepo := NewActorsRepositoryMock()
	filmsRepo := NewFilmsRepositoryMock()

	for idx, item := range cases {
		caseName := fmt.Sprintf("case %d", idx)

		expectedErr := item.Err
		resultErr := NewUpdateFilmHandler(actorsRepo, filmsRepo).Handle(item.Ctx, item.Request)

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
