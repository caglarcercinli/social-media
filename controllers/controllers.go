package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserById(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "user1")
	return
}