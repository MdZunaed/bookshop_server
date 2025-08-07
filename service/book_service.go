package service

import (
	"github.com/MdZunaed/bookshop/repo"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookServiceInterface interface {
	CreateBook(data any, sessionContext mongo.SessionContext) (any, error)
	GetAllBooks(sessionContext mongo.SessionContext) (any, error)
	GetBookById(id string, sessionContext mongo.SessionContext) (any, error)
	UpdateBook(id string, data map[string]any, sessionContext mongo.SessionContext) (any, error)
	DeleteBookById(id string, sessionContext mongo.SessionContext) (any, error)
}

type BookService struct {
	repository repo.Repository
}

func GetBookService(repository repo.Repository) BookServiceInterface {
	return &BookService{
		repository: repository,
	}
}

func (bs *BookService) CreateBook(data any, sessionContext mongo.SessionContext) (any, error) {
	return bs.repository.BookRepository.Create(data, sessionContext)
}
func (bs *BookService) GetAllBooks(sessionContext mongo.SessionContext) (any, error) {
	return bs.repository.BookRepository.FindAll(nil, sessionContext)
}
func (bs *BookService) GetBookById(id string, sessionContext mongo.SessionContext) (any, error) {
	return bs.repository.BookRepository.FindOne(id, sessionContext)
}
func (bs *BookService) UpdateBook(id string, data map[string]any, sessionContext mongo.SessionContext) (any, error) {
	return bs.repository.BookRepository.Update(id, data, sessionContext)
}
func (bs *BookService) DeleteBookById(id string, sessionContext mongo.SessionContext) (any, error) {
	return bs.repository.BookRepository.Delete(id, sessionContext)
}
