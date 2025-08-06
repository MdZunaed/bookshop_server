package controller

import (
	"fmt"

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
		// ctx.JSON(400, gin.H{
		// 	"message": "Bad Request",
		// 	"error":   err.Error(),
		// })
		ctx.Error(fmt.Errorf("400::%s::%s::%v", "Bad Request", err.Error(), err))
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
