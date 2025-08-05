package service

import (
	repo "github.com/MdZunaed/bookshop/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserServiceInterface interface {
	CreateUser(data any, sessionContext mongo.SessionContext) (any, error)
}

type UserService struct {
	repository repo.MongoRepoInterface
}

func (us *UserService) CreateUser(data any, sessionContext mongo.SessionContext) (any, error) {
	return us.repository.Create(data, sessionContext)
}

func GetUsereService(repository repo.MongoRepoInterface) UserServiceInterface {
	return &UserService{
		repository: repository,
	}
}
