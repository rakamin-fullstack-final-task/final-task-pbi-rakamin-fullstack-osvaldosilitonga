package controllers

import "github.com/gin-gonic/gin"

type PhotoController interface {
	Upload(c *gin.Context)
}
