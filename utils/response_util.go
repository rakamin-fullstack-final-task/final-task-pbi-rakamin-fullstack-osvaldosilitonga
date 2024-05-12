package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rakamin-fullstack-final-task/final-task-pbi-rakamin-fullstack-osvaldosilitonga/dto"
)

var (
	ApiOk = dto.WebResponse{
		Code:     http.StatusOK,
		Messages: "ok",
	}

	ApiDelete = dto.WebResponse{
		Code:     http.StatusOK,
		Messages: "deleted successfully",
	}

	ApiCreate = dto.WebResponse{
		Code:     http.StatusCreated,
		Messages: "created successfully",
	}

	ApiUpdate = dto.WebResponse{
		Code:     http.StatusOK,
		Messages: "updated successfully",
	}

	ApiBadRequest = dto.WebResponse{
		Code:     http.StatusBadRequest,
		Messages: "Bad Request",
	}

	ApiNotFound = dto.WebResponse{
		Code:     http.StatusNotFound,
		Messages: "not found",
	}

	ApiForbidden = dto.WebResponse{
		Code:     http.StatusForbidden,
		Messages: "forbidden",
	}

	ApiInternalServer = dto.WebResponse{
		Code:     http.StatusInternalServerError,
		Messages: "internal server error",
	}

	ApiUnauthorized = dto.WebResponse{
		Code:     http.StatusUnauthorized,
		Messages: "unauthorized",
	}
)

func Response(c *gin.Context, res *dto.WebResponse, data any, err string) {
	res.Data = data
	res.Errors = err

	c.JSON(res.Code, res)
}
