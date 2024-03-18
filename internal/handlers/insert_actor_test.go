package handlers

import (
	"context"
	"fmt"
	"testing"
)

type InsertActorCase struct {
	Ctx     context.Context
	Request InsertActorRequest
	Err     error
}

func TestInsertActor(t *testing.T) {
	cases := []InsertActorCase{
		InsertActorCase{
			Request: InsertActorRequest{
				Values: map[string]string{"name": "Matthew Paige Damon", "gender": "M", "birthday": "1960-10-08"},
			},
			Err: nil,
			Ctx: context.Background(),
		},
		InsertActorCase{
			Request: InsertActorRequest{
				Values: map[string]string{},
			},
			Err: fmt.Errorf("InsertActor error"),
			Ctx: context.Background(),
		},
	}

	runInsertActorCases(t, cases)
}

func runInsertActorCases(t *testing.T, cases []InsertActorCase) {
	actorsRepo := NewActorsRepositoryMock()
	filmsRepo := NewFilmsRepositoryMock()

	for idx, item := range cases {
		caseName := fmt.Sprintf("case %d", idx)

		expectedErr := item.Err
		resultErr := NewInsertActorHandler(actorsRepo, filmsRepo).Handle(item.Ctx, item.Request)

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
