package controller

import (

	"github.com/MdZunaed/bookshop/model"
	"github.com/MdZunaed/bookshop/service"
	"github.com/MdZunaed/bookshop/utils"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService     service.UserServiceInterface
	responseService utils.ResponseService
}

func (uc *UserController) CreateUser(ctx *gin.Context) {
	var user model.NewUser
	if err := ctx.ShouldBindBodyWithJSON(&user); err != nil {
		ctx.Error(&model.AppError{
			Source: "UserController_CreateUser",
			StatusCode: 400,
			Message: "Bad Request",
			Err: err,
		})

		// ctx.Error(fmt.Errorf("%s::400::%s::%v","UserController_CreateUser", "Bad Request",  err))
		return
	}
	data, err := uc.userService.CreateUser(user, nil)
	if err != nil {
		ctx.Error(err)
		return
	}
	uc.responseService.Success(ctx, 200, "Created successfully", data)
}

func GetUserController(userService service.UserServiceInterface, responseService utils.ResponseService) *UserController {
	return &UserController{
		userService:     userService,
		responseService: responseService,
	}
}
