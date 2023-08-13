package controllers

import (
	"net/http"

	"example.com/social-media/services"
	"github.com/gin-gonic/gin"
)

func GetUserById(c *gin.Context) {
	name := services.GetUserByIdFromDatabase(c)
	c.IndentedJSON(http.StatusOK, name)
	return
}