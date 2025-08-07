package routes

import (
	"github.com/MdZunaed/bookshop/controller"
	"github.com/MdZunaed/bookshop/repo"
	"github.com/MdZunaed/bookshop/service"
	"github.com/MdZunaed/bookshop/utils"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.RouterGroup) {
	repository := repo.GetRepository()
	userService := service.GetUsereService(*repository)
	responseService := utils.GetResponseService()
	userController := controller.GetUserController(userService, *responseService)

	router.POST(
		"/create",
		userController.CreateUser,
	)
}
