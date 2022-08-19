package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"ppob-api.go/ppob-api/config"
	"ppob-api.go/ppob-api/controller"
	"ppob-api.go/ppob-api/repository"
	"ppob-api.go/ppob-api/service"
)

var (
	client         *mongo.Client             = config.ConnectDB()
	userRepository repository.UserRepository = repository.NewUserRepository(client)
	userService    service.UserService       = service.NewUserService(userRepository)
	userController controller.UserController = controller.NewUserController(userService)
)

func main() {
	r := gin.Default()

	router := r.Group("/api")
	{
		router.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})

		userRouter := router.Group("/users")
		{
			userRouter.HEAD("/:type/check/:phone", userController.CheckPhoneNumber)
			userRouter.POST("/register", userController.Register)
			userRouter.POST("/:phone/generate-otp", userController.GenerateOTP)
		}
	}

	r.Run()
}
