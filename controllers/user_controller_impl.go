package controllers

import (
	"context"
	"strconv"

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

func (uc *userControllerImpl) Update(c *gin.Context) {
	params := c.Param("id")
	id, err := strconv.Atoi(params)
	if err != nil {
		utils.Response(c, &utils.ApiBadRequest, nil, "invalid params id")
		return
	}

	var req dto.UserUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Response(c, &utils.ApiBadRequest, nil, err.Error())
		return
	}

	if req.Username != "" && len(req.Username) < 3 {
		utils.Response(c, &utils.ApiBadRequest, nil, "username is to short")
		return
	}
	if req.Password != "" && len(req.Password) < 6 {
		utils.Response(c, &utils.ApiBadRequest, nil, "password is to short")
		return
	}

	ctx := context.WithValue(c, "id", uint64(id))

	res, err := uc.userService.Update(ctx, &req)
	if err != nil {
		helpers.ErrorCheck(c, err)
		return
	}

	utils.Response(c, &utils.ApiOk, res, "")
}
