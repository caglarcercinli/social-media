package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main(){
	router := gin.Default()
	router.GET("/users/:id", getUserById)
	router.Run("localhost:8080")
}

func getUserById(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "user1")
	return
}