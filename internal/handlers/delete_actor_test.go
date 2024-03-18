package handlers

import (
	"context"
	"fmt"
	"testing"
)

type DeleteActorCase struct {
	Ctx     context.Context
	Request DeleteActorRequest
	Err     error
}

func TestDeleteActor(t *testing.T) {
	cases := []DeleteActorCase{
		DeleteActorCase{
			Request: DeleteActorRequest{
				Id: 1,
			},
			Err: nil,
			Ctx: context.Background(),
		},
		DeleteActorCase{
			Request: DeleteActorRequest{
				Id: 0,
			},
			Err: fmt.Errorf("DeleteActor error"),
			Ctx: context.Background(),
		},
	}

	runDeleteActorCases(t, cases)
}

func runDeleteActorCases(t *testing.T, cases []DeleteActorCase) {
	actorsRepo := NewActorsRepositoryMock()
	filmsRepo := NewFilmsRepositoryMock()

	for idx, item := range cases {
		caseName := fmt.Sprintf("case %d", idx)

		expectedErr := item.Err
		resultErr := NewDeleteActorHandler(actorsRepo, filmsRepo).Handle(item.Ctx, item.Request)

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
