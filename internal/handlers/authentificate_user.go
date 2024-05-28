package handlers

import (
	"context"
	"fmt"
	"log"

	"github.com/LKarataev/Film-Library-API/internal/auth"
	"github.com/LKarataev/Film-Library-API/internal/dao"
)

type AuthenticateUserHandler struct {
	accountsRepo dao.AccountsRepositoryInterface
}

type AuthenticateUserResponse struct {
	signedToken string
}

type AuthenticateUserRequest struct {
	Username string
	Password string
}

func NewAuthenticateUserHandler(accountsRepo dao.AccountsRepositoryInterface) AuthenticateUserHandler {
	return AuthenticateUserHandler{accountsRepo: accountsRepo}
}

func (h AuthenticateUserHandler) Handle(ctx context.Context, req AuthenticateUserRequest) (string, error) {
	log.Println("AuthenticateUserRequest started")

	account, err := h.accountsRepo.GetAccount(req.Username)
	if err != nil {
		log.Println("AuthenticateUserRequest error: ", err)
		return "", err
	}

	if req.Password != account.Password {
		return "", fmt.Errorf("Error: password is wrong")
	}

	authAcc := auth.Account{Username: account.Username, Role: account.Role}
	signedToken, err := auth.GenerateToken(authAcc)
	if err != nil {
		log.Println("AuthenticateUserRequest error: ", err)
		return "", err
	}

	log.Println("AuthenticateUserRequest successed")
	return signedToken, nil
}
