package routes

import (
	"github.com/MdZunaed/bookshop/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoute(r *gin.Engine){
	router:= r.Group("/")

	// Error handler Middleware
	router.Use(middlewares.ErrorHandler())

	userRoutes:= router.Group("/users")
	{
		RegisterUserRoutes(userRoutes)
	}

	authRoutes:= router.Group("/auth")
	{
		RegisterAuthRoutes(authRoutes)
	}
}