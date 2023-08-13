package services

import (
	"context"
	"fmt"
	"os"

	"example.com/social-media/models"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func GetUserByIdFromDatabase(c *gin.Context)(models.User, error) {
	var user models.User 
    id := c.Param("id")
	conn, err := pgx.Connect(context.Background(), URL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return user, err
	}
	defer conn.Close(context.Background())

	err = conn.QueryRow(context.Background(), "select user_id, username, password from users where user_id=$1", id).Scan(&user.User_id, &user.Username,&user.Password)
	fmt.Println(user)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return user, err
	}
	return user, err
}

func GetMessagesByUserIdFromDatabase(c *gin.Context)([]models.Message, error) {
	var messages []models.Message 
	user_id := c.Param("id")
	conn, err := pgx.Connect(context.Background(), URL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return messages, err
	}
	defer conn.Close(context.Background())
	//
	rows, err := conn.Query(context.Background(), "select message_id, sender_id, message_text from messages where sender_id=$1", user_id)
	defer rows.Close()
	for rows.Next() {
		var readMessage models.Message
		err = rows.Scan(&readMessage.Message_id,&readMessage.Sender_id,&readMessage.Message_text)
		if err != nil {
			fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
			return messages, err
		}
		messages = append(messages, readMessage )
	}
	return messages, err
}