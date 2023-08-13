package services

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func GetUserByIdFromDatabase(c *gin.Context)string {
    id := c.Param("id")
	conn, err := pgx.Connect(context.Background(), URL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	var username string

	err = conn.QueryRow(context.Background(), "select username from users where user_id=$1", id).Scan(&username)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
		os.Exit(1)
	}
	return username
}