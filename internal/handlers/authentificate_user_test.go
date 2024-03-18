package handlers

import (
	"context"
	"fmt"
	"testing"
)

type AuthenticateUserCase struct {
	Ctx     context.Context
	Request AuthenticateUserRequest
	Result  *AuthenticateUserResponse
	Err     error
}

func TestAuthenticateUser(t *testing.T) {
	cases := []AuthenticateUserCase{
		AuthenticateUserCase{
			Request: AuthenticateUserRequest{
				Username: "admin",
				Password: "admin_password",
			},
			Result: &AuthenticateUserResponse{
				signedToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9",
			},
			Err: nil,
			Ctx: context.Background(),
		},
		AuthenticateUserCase{
			Request: AuthenticateUserRequest{
				Username: "user",
				Password: "wrong_password",
			},
			Result: nil,
			Err:    fmt.Errorf("Error: password is wrong"),
			Ctx:    context.Background(),
		},
		AuthenticateUserCase{
			Request: AuthenticateUserRequest{
				Username: "user_not_exist",
			},
			Result: nil,
			Err:    fmt.Errorf("GetAccount error"),
			Ctx:    context.Background(),
		},
	}

	runAuthenticateUserCases(t, cases)
}

func runAuthenticateUserCases(t *testing.T, cases []AuthenticateUserCase) {
	accountsRepo := NewAccountsRepositoryMock()

	for idx, item := range cases {
		caseName := fmt.Sprintf("case %d", idx)

		result, err := NewAuthenticateUserHandler(accountsRepo).Handle(item.Ctx, item.Request)

		expectedErr := item.Err
		resultErr := err

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

		expected := item.Result.signedToken
		if result[:len(expected)] != expected {
			t.Fatalf("[%s] results not match\nGot : %#v\nWant: %#v", caseName, result, expected)
			continue
		}
	}
}
