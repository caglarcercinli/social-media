package controllers

import (
	"net/http"

	"example.com/social-media/services"
	"github.com/gin-gonic/gin"
)

func GetUserById(c *gin.Context) {
	user, err := services.GetUserByIdFromDatabase(c)
	if err!=nil{
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, user)
	return
}

func GetMessagesByUserId(c *gin.Context) {
	messages,err := services.GetMessagesByUserIdFromDatabase(c)
	if err!=nil || len(messages) == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "messages not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, messages)
	return
}