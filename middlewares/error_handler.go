package middlewares

import (
	"errors"
	"fmt"

	"github.com/MdZunaed/bookshop/config"
	"github.com/MdZunaed/bookshop/model"
	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		if len(ctx.Errors) > 0 {
			rawErr := ctx.Errors[0].Err

			var appErr = new(model.AppError)
			if errors.As(rawErr, &appErr) {
				ctx.JSON(appErr.StatusCode, gin.H{
					"message": appErr.Message,
				})

				if config.GetEnvProperty("debug") != "" {
					fmt.Printf("\nSource: %s,\nStatus code: %d,\nMessage: %s,\nError: %v\n\n",
						appErr.Source, appErr.StatusCode, appErr.Message, appErr.Err)
				}
			} else {
				ctx.JSON(500, gin.H{
					"message": "Internal Server Error",
				})
			}
		}
	}
}

// func ErrorHander() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		ctx.Next()
// 		if len(ctx.Errors) > 0 {
// 			err := ctx.Errors[0]
// 			errMap := make(map[string]any, 0)
// 			if err != nil {
// 				splittedErr := strings.Split(err.Error(), "::")

// 				isDebugEnabled := config.GetEnvProperty("debug")
// 				if isDebugEnabled != "" && len(splittedErr) > 0 {
// 					fmt.Printf("Source: %s,\nStatus code: %s,\nMessage: %s,\nError: %v", splittedErr[0], splittedErr[1], splittedErr[2], splittedErr[3])
// 				}
// 				statusCode, err := strconv.Atoi(splittedErr[1])
// 				if err != nil {
// 					statusCode = 500
// 				}
// 				errMap["Status Code"] = splittedErr[1]
// 				errMap["message"] = splittedErr[2]
// 				ctx.JSON(statusCode, errMap)
// 			}
// 		}
// 	}
// }
