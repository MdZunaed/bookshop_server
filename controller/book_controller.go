package controller

import (
	"github.com/MdZunaed/bookshop/dto"
	"github.com/MdZunaed/bookshop/model"
	"github.com/MdZunaed/bookshop/service"
	"github.com/MdZunaed/bookshop/utils"
	"github.com/gin-gonic/gin"
)

type BookController struct {
	bookService     service.BookServiceInterface
	responseService utils.ResponseService
}

func GetBookController(bookService service.BookServiceInterface, responseService utils.ResponseService) *BookController {
	return &BookController{
		bookService:     bookService,
		responseService: responseService,
	}
}

func (bc *BookController) CreateBook(ctx *gin.Context) {
	var bookDto model.Book
	if err := ctx.ShouldBindJSON(&bookDto); err != nil {
		ctx.Error(&model.AppError{
			Source:     "BookController_CreateBook",
			StatusCode: 400,
			Message:    "No body content found",
			Err:        err,
		})
		return
	}
	data, err := bc.bookService.CreateBook(bookDto, nil)
	if err != nil {
		ctx.Error(err)
		return
	}
	bc.responseService.Success(ctx, 200, "Created Book successfully", data)
}

func (bc *BookController) GetAllBooksBook(ctx *gin.Context) {
	data, err := bc.bookService.GetAllBooks(nil)
	if err != nil {
		ctx.Error(err)
		return
	}
	bc.responseService.Success(ctx, 200, "Success", data)
}

func (bc *BookController) GetBookById(ctx *gin.Context) {
	var bookId = ctx.Param("bookId")

	if bookId == "" {
		ctx.Error(&model.AppError{
			Source:     "BookController_GetBookById",
			StatusCode: 404,
			Message:    "Id is empty",
		})
		return
	}
	data, err := bc.bookService.GetBookById(bookId, nil)
	if err != nil {
		ctx.Error(err)
		return
	}
	bc.responseService.Success(ctx, 200, "Success", data)
}

func (bc *BookController) UpdateBook(ctx *gin.Context) {
	var bookDto dto.BookDto
	if err := ctx.ShouldBindJSON(&bookDto); err != nil {
		ctx.Error(&model.AppError{
			Source:     "BookController_UpdateBook",
			StatusCode: 400,
			Message:    "No body content found",
			Err:        err,
		})
		return
	}
	data, err := bc.bookService.UpdateBook(bookDto, nil)
	if err != nil {
		ctx.Error(err)
		return
	}
	bc.responseService.Success(ctx, 200, "Updated Book successfully", data)
}

func (bc *BookController) DeleteBookById(ctx *gin.Context) {
	var bookId = ctx.Param("bookId")

	if bookId == "" {
		ctx.Error(&model.AppError{
			Source:     "BookController_DeleteBookById",
			StatusCode: 404,
			Message:    "Id is empty",
		})
		return
	}
	data, err := bc.bookService.DeleteBookById(bookId, nil)
	if err != nil {
		ctx.Error(err)
		return
	}
	bc.responseService.Success(ctx, 200, "Deleted Book successfully", data)
}
