package service

import (
	"fmt"

	"github.com/MdZunaed/bookshop/dto"
	"github.com/MdZunaed/bookshop/model"
	repo "github.com/MdZunaed/bookshop/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthServiceInterface interface {
	Login(loginDto dto.LoginDto, sessionContext mongo.SessionContext) (map[string]any, error)
}

type AuthService struct {
	repository repo.Repository
	userService UserServiceInterface
}

func (as *AuthService) Login(loginDto dto.LoginDto, sessionContext mongo.SessionContext) (map[string]any, error) {
	res, err:= as.userService.FindOneUserByEmail(loginDto.Email, sessionContext)
	if err != nil {
		return nil, err
	}
	user:= res.(model.User)
	if user.Password == loginDto.Password{
		return map[string]any{
			"token": "",
			"user": user,
		}, nil
	}
	return nil, fmt.Errorf("invalid credential")
}

func GetAuthService(repository repo.Repository, userService UserServiceInterface) AuthServiceInterface{
	return &AuthService{
		repository: repository,
		userService: userService,
	}
}