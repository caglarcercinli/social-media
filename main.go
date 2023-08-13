package main

import (
	"example.com/social-media/controllers"
	"github.com/gin-gonic/gin"
)

func main(){
	router := gin.Default()
	router.GET("/users/:id", controllers.GetUserById)
	router.GET("/messages/:user_id", controllers.GetMessagesByUserId)
	router.Run("localhost:8080")
}
