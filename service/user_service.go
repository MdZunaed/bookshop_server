package service

import (
	"github.com/MdZunaed/bookshop/model"
	repo "github.com/MdZunaed/bookshop/repository"
	"github.com/MdZunaed/bookshop/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserServiceInterface interface {
	CreateUser(data any, sessionContext mongo.SessionContext) (any, error)
	FindOneUserByEmail(email string, sessionContext mongo.SessionContext) (any, error)
}

type UserService struct {
	repository repo.Repository
}

func (us *UserService) CreateUser(data any, sessionContext mongo.SessionContext) (any, error) {
	return us.repository.UserRepository.Create(data, sessionContext)
}

func (us *UserService) FindOneUserByEmail(email string, sessionContext mongo.SessionContext) (any, error) {
	res, err := us.repository.UserRepository.FindOneByKey("email", email, sessionContext)
	if err != nil {
		return nil, &model.AppError{
			Source:     "UserService_FindUserByEmail",
			StatusCode: 500,
			Message:    "Internal server error",
			Err:        err,
		}
	}
	var user model.User
	if err := utils.MapToStruct(res.(map[string]any), &user); err != nil {
		return nil, &model.AppError{
			Source:     "UserService_FindUserByEmail",
			StatusCode: 500,
			Message:    "Internal server error",
			Err:        err,
		}
	}
	return user, nil
}

func GetUsereService(repository repo.Repository) UserServiceInterface {
	return &UserService{
		repository: repository,
	}
}
