package service

import (
	"github.com/MdZunaed/bookshop/repo"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookServiceInterface interface {
	CreateBook(data any, sessionContext mongo.SessionContext) (any, error)
	GetAllBooks(sessionContext mongo.SessionContext) (any, error)
	GetBookById(id string, sessionContext mongo.SessionContext) (any, error)
	UpdateBook(data any, sessionContext mongo.SessionContext) (any, error)
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
func (bs *BookService) UpdateBook(data any, sessionContext mongo.SessionContext) (any, error) {
	bookData := data.(map[string]any)
	id := bookData["_id"].(string)
	delete(bookData, "_id")
	return bs.repository.BookRepository.Update(id, bookData, sessionContext)
}
func (bs *BookService) DeleteBookById(id string, sessionContext mongo.SessionContext) (any, error) {
	return bs.repository.BookRepository.Delete(id, sessionContext)
}
