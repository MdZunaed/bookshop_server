package routes

import (
	"github.com/MdZunaed/bookshop/controller"
	"github.com/MdZunaed/bookshop/repo"
	"github.com/MdZunaed/bookshop/service"
	"github.com/MdZunaed/bookshop/utils"
	"github.com/gin-gonic/gin"
)

func RegisterBookRoutes(router *gin.RouterGroup) {
	repository := repo.GetRepository()
	bookService := service.GetBookService(*repository)
	responseService := utils.GetResponseService()
	bookController := controller.GetBookController(bookService, *responseService)

	router.GET(
		"/all-books",
		bookController.GetAllBooksBook,
	)

	router.GET(
		"/:bookId",
		bookController.GetBookById,
	)

	router.POST(
		"/create",
		bookController.CreateBook,
	)

	router.PUT(
		"/:bookId",
		bookController.UpdateBook,
	)

	router.DELETE(
		"/:bookId",
		bookController.DeleteBookById,
	)
}