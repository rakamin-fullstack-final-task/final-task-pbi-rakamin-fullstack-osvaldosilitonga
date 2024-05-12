package helpers

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rakamin-fullstack-final-task/final-task-pbi-rakamin-fullstack-osvaldosilitonga/utils"
)

func ErrorCheck(c *gin.Context, err error) {
	var errSplit []string = strings.Split(err.Error(), ",")

	switch errSplit[0] {
	case "400":
		utils.Response(c, &utils.ApiBadRequest, nil, errSplit[1])
	default:
		utils.Response(c, &utils.ApiInternalServer, nil, errSplit[1])
	}
}
