package routes

import (
	"github.com/MdZunaed/bookshop/controller"
	"github.com/MdZunaed/bookshop/repository"
	"github.com/MdZunaed/bookshop/service"
	"github.com/MdZunaed/bookshop/utils"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(router *gin.RouterGroup) {
	repository := repository.GetRepository()
	userService := service.GetUsereService(*repository)
	authService := service.GetAuthService(*repository, userService)
	responseService := utils.GetResponseService()
	authController := controller.GetAuthController(authService, *responseService)

	router.POST(
		"/create",
		authController.Login,
	)
}
