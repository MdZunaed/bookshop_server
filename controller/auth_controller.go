package controller

import (
	"github.com/MdZunaed/bookshop/dto"
	"github.com/MdZunaed/bookshop/model"
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
		ctx.Error(&model.AppError{
			Source:     "AuthController_Login",
			StatusCode: 400,
			Message:    "Bad Request",
			Err:        err,
		})
		//ctx.Error(fmt.Errorf("%s::400::%s::%v","AuthController_Login", "Bad Request",  err))
		return
	}
	data, err := ac.authService.Login(loginDto, nil)
	if err != nil {
		if appErr, ok := err.(model.AppError); ok {
			ctx.Error(appErr)
			return
		}
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
