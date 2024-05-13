package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rakamin-fullstack-final-task/final-task-pbi-rakamin-fullstack-osvaldosilitonga/controllers"
	"github.com/rakamin-fullstack-final-task/final-task-pbi-rakamin-fullstack-osvaldosilitonga/database"
	"github.com/rakamin-fullstack-final-task/final-task-pbi-rakamin-fullstack-osvaldosilitonga/repository"
	"github.com/rakamin-fullstack-final-task/final-task-pbi-rakamin-fullstack-osvaldosilitonga/service"
)

func Route(r *gin.Engine) {
	db := database.InitDB()

	userRepository := repository.NewUserRepository(db)
	photoRepository := repository.NewPhotoRepository(db)

	userService := service.NewUserService(userRepository)
	photoService := service.NewPhotoService(photoRepository)

	userController := controllers.NewUserController(userService)
	photoController := controllers.NewPhotoController(photoService)

	v1 := r.Group("/v1")
	{
		v1.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"code":    200,
				"message": "ok",
				"errors":  nil,
				"data":    nil,
			})
		})

		users := v1.Group("/users")
		{
			users.POST("/register", userController.Register)
			users.POST("/login", userController.Login)
			users.PATCH("/:id", userController.Update)
			users.DELETE("/:id", userController.Delete)
		}

		// Set a lower memory limit for multipart forms (default is 32 MiB)
		r.MaxMultipartMemory = 2 << 20 // 2 MiB

		photos := v1.Group("/photos")
		{
			photos.POST("/", photoController.Upload)
		}
	}
}
