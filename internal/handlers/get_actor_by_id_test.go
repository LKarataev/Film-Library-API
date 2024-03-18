package handlers

import (
	"context"
	"fmt"
	"reflect"
	"testing"
)

type GetActorByIdCase struct {
	Ctx     context.Context
	Request GetActorByIdRequest
	Result  *GetActorByIdResponse
	Err     error
}

func TestGetActorById(t *testing.T) {
	cases := []GetActorByIdCase{
		GetActorByIdCase{
			Request: GetActorByIdRequest{
				Id: 1,
			},
			Result: &GetActorByIdResponse{
				&Actor{Id: 1, Name: "Thomas Hanks", Gender: "M", Birthday: "1956-07-09", Films: []string{"The Green Mile", "Forrest Gump", "Catch Me If You Can", "Saving Private Ryan"}},
			},
			Err: nil,
			Ctx: context.Background(),
		},
		GetActorByIdCase{
			Request: GetActorByIdRequest{
				Id: 0,
			},
			Result: nil,
			Err:    fmt.Errorf("GetActorById error"),
			Ctx:    context.Background(),
		},
		GetActorByIdCase{
			Request: GetActorByIdRequest{
				Id: 5,
			},
			Result: nil,
			Err:    fmt.Errorf("GetFilmsByActorId error"),
			Ctx:    context.Background(),
		},
	}

	runGetActorByIdCases(t, cases)
}

func runGetActorByIdCases(t *testing.T, cases []GetActorByIdCase) {
	actorsRepo := NewActorsRepositoryMock()
	filmsRepo := NewFilmsRepositoryMock()

	for idx, item := range cases {
		caseName := fmt.Sprintf("case %d", idx)

		expected := item.Result
		result, err := NewGetActorByIdHandler(actorsRepo, filmsRepo).Handle(item.Ctx, item.Request)

		if result == nil && expected != nil {
			t.Fatalf("[%s] results not match\nGot : %#v\nWant: %#v", caseName, result, expected)
			continue
		}

		if expected == nil {
			if result != expected {
				t.Fatalf("[%s] results not match\nGot : %#v\nWant: %#v", caseName, result, expected)
				continue
			}
		} else if !reflect.DeepEqual(*result.Actor, *expected.Actor) {
			t.Fatalf("[%s] results not match\nGot : %#v\nWant: %#v", caseName, *result.Actor, *expected.Actor)
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
