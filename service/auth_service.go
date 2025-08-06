package service

import (
	"github.com/MdZunaed/bookshop/dto"
	"github.com/MdZunaed/bookshop/model"
	repo "github.com/MdZunaed/bookshop/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthServiceInterface interface {
	Login(loginDto dto.LoginDto, sessionContext mongo.SessionContext) (map[string]any, error)
}

type AuthService struct {
	repository  repo.Repository
	userService UserServiceInterface
}

func (as *AuthService) Login(loginDto dto.LoginDto, sessionContext mongo.SessionContext) (map[string]any, error) {
	res, err := as.userService.FindOneUserByEmail(loginDto.Email, sessionContext)
	if err != nil {
		return nil, &model.AppError{
			Source:     "AuthService_Login",
			StatusCode: 404,
			Message:    "No user found with this email",
			Err:        err,
		}
		//fmt.Errorf("no user found with this email")
	}
	user := res.(model.User)
	if user.Password != loginDto.Password {
		return nil, &model.AppError{
			Source:     "AuthService_Login",
			StatusCode: 400,
			Message:    "Invalid Credential",
			Err:        err,
		}
	}
	return map[string]any{
		"token": "",
		"user":  user,
	}, nil
}

func GetAuthService(repository repo.Repository, userService UserServiceInterface) AuthServiceInterface {
	return &AuthService{
		repository:  repository,
		userService: userService,
	}
}
