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

	userService := service.NewUserService(userRepository)

	userController := controllers.NewUserController(userService)

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
	}
}
