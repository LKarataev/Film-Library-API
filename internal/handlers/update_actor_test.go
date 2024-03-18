package handlers

import (
	"context"
	"fmt"
	"testing"
)

type UpdateActorCase struct {
	Ctx     context.Context
	Request UpdateActorRequest
	Err     error
}

func TestUpdateActor(t *testing.T) {
	cases := []UpdateActorCase{
		UpdateActorCase{
			Request: UpdateActorRequest{
				Values: map[string]string{"birthday": "1970-10-08"},
				Id:     1,
			},
			Err: nil,
			Ctx: context.Background(),
		},
		UpdateActorCase{
			Request: UpdateActorRequest{
				Id: 0,
			},
			Err: fmt.Errorf("UpdateActor error"),
			Ctx: context.Background(),
		},
	}

	runUpdateActorCases(t, cases)
}

func runUpdateActorCases(t *testing.T, cases []UpdateActorCase) {
	actorsRepo := NewActorsRepositoryMock()
	filmsRepo := NewFilmsRepositoryMock()

	for idx, item := range cases {
		caseName := fmt.Sprintf("case %d", idx)

		expectedErr := item.Err
		resultErr := NewUpdateActorHandler(actorsRepo, filmsRepo).Handle(item.Ctx, item.Request)

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
