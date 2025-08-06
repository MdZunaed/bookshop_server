package controller

import (
	"fmt"

	"github.com/MdZunaed/bookshop/dto"
	"github.com/MdZunaed/bookshop/service"
	"github.com/MdZunaed/bookshop/utils"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService     service.AuthServiceInterface
	responseService utils.ResponseService
}

func (ac *AuthController) Login(ctx *gin.Context) {
	var loginDto dto.LoginDto
	if err := ctx.ShouldBindJSON(&loginDto); err != nil {
		ctx.Error(fmt.Errorf("400::%s::%s::%v", "Bad Request", err.Error(), err))
		return
	}
	data, err := ac.authService.Login(loginDto, nil)
	if err != nil {
		ctx.Error(err)
		return
	}
	ac.responseService.Success(ctx, 200, "Created successfully", data)
}

func GetAuthController(authService service.AuthServiceInterface, responseService utils.ResponseService) *AuthController {
	return &AuthController{
		authService:     authService,
		responseService: responseService,
	}
}
