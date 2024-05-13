package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rakamin-fullstack-final-task/final-task-pbi-rakamin-fullstack-osvaldosilitonga/dto"
	"github.com/rakamin-fullstack-final-task/final-task-pbi-rakamin-fullstack-osvaldosilitonga/helpers"
	"github.com/rakamin-fullstack-final-task/final-task-pbi-rakamin-fullstack-osvaldosilitonga/service"
	"github.com/rakamin-fullstack-final-task/final-task-pbi-rakamin-fullstack-osvaldosilitonga/utils"
)

type photoControllerImpl struct {
	photoService service.PhotoService
}

func NewPhotoController(ps service.PhotoService) PhotoController {
	return &photoControllerImpl{
		photoService: ps,
	}
}

func (pc *photoControllerImpl) Upload(c *gin.Context) {
	// TODO: Apply middleware to set userID to gin.Context
	c.Set("id", "12") // Dummy data

	userId, exist := c.Get("id")
	if !exist {
		utils.Response(c, &utils.ApiBadRequest, nil, "user id not found")
		return
	}

	id, err := strconv.ParseUint(userId.(string), 10, 64)
	if err != nil {
		utils.Response(c, &utils.ApiBadRequest, nil, "invalid user id")
		return
	}

	var form dto.PhotoUploadFormRequest
	if err := c.ShouldBind(&form); err != nil {
		utils.Response(c, &utils.ApiBadRequest, nil, err.Error())
		return
	}

	uuid := uuid.New()

	form.Path = "./storage/uploads/" + uuid.String() + "-" + form.Image.Filename

	if err := c.SaveUploadedFile(form.Image, form.Path); err != nil {
		utils.Response(c, &utils.ApiInternalServer, nil, "something went wrong, fail when saving image")
		return
	}

	res, err := pc.photoService.Upload(c, id, &form)
	if err != nil {
		helpers.ErrorCheck(c, err)
		return
	}

	utils.Response(c, &utils.ApiOk, res, "")
}
