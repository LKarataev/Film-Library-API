package handlers

import (
	"context"
	"fmt"
	"reflect"
	"testing"
)

type GetActorsCase struct {
	Ctx     context.Context
	Request GetActorsRequest
	Result  *GetActorsResponse
	Err     error
}

func TestGetActors(t *testing.T) {
	cases := []GetActorsCase{
		GetActorsCase{
			Request: GetActorsRequest{
				Limit:  1,
				Offset: 0,
			},
			Result: &GetActorsResponse{
				[]Actor{
					Actor{Id: 1, Name: "Thomas Hanks", Gender: "M", Birthday: "1956-07-09", Films: []string{"The Green Mile", "Forrest Gump", "Catch Me If You Can", "Saving Private Ryan"}},
				},
			},
			Err: nil,
			Ctx: context.Background(),
		},
		GetActorsCase{
			Request: GetActorsRequest{
				Limit:  2,
				Offset: 0,
			},
			Result: nil,
			Err:    fmt.Errorf("GetFilmsByActorId error"),
			Ctx:    context.Background(),
		},
		GetActorsCase{
			Request: GetActorsRequest{
				Limit:  30,
				Offset: 30,
			},
			Result: nil,
			Err:    fmt.Errorf("GetActors error"),
			Ctx:    context.Background(),
		},
	}

	runGetActorsCases(t, cases)
}

func runGetActorsCases(t *testing.T, cases []GetActorsCase) {
	actorsRepo := NewActorsRepositoryMock()
	filmsRepo := NewFilmsRepositoryMock()

	for idx, item := range cases {
		caseName := fmt.Sprintf("case %d", idx)

		expected := item.Result
		result, err := NewGetActorsHandler(actorsRepo, filmsRepo).Handle(item.Ctx, item.Request)

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
