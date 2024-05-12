package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rakamin-fullstack-final-task/final-task-pbi-rakamin-fullstack-osvaldosilitonga/dto"
	"github.com/rakamin-fullstack-final-task/final-task-pbi-rakamin-fullstack-osvaldosilitonga/helpers"
	"github.com/rakamin-fullstack-final-task/final-task-pbi-rakamin-fullstack-osvaldosilitonga/service"
	"github.com/rakamin-fullstack-final-task/final-task-pbi-rakamin-fullstack-osvaldosilitonga/utils"
)

type userControllerImpl struct {
	userService service.UserService
}

func NewUserController(us service.UserService) UserController {
	return &userControllerImpl{
		userService: us,
	}
}

func (uc *userControllerImpl) Register(c *gin.Context) {
	var req dto.UserRegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Response(c, &utils.ApiBadRequest, nil, err.Error())
		return
	}

	user, err := uc.userService.AddNewUser(c, &req)
	if err != nil {
		helpers.ErrorCheck(c, err)
		return
	}

	res := dto.UserRegisterResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
	}

	utils.Response(c, &utils.ApiCreate, res, "")
}

func (uc *userControllerImpl) Login(c *gin.Context) {
	var req dto.UserLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Response(c, &utils.ApiBadRequest, nil, err.Error())
		return
	}

	res, err := uc.userService.Login(c, &req)
	if err != nil {
		helpers.ErrorCheck(c, err)
		return
	}

	utils.Response(c, &utils.ApiOk, res, "")
}
