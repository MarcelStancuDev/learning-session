package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"chat-app/pkg/auth"
	"chat-app/pkg/models"
	"chat-app/pkg/websocket"
)

func main() {
	//Connect to database
	dsn := "host=localhost user=youruser password=yourpassword dbname=yourdb port=5432 sslmodel=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database: ", err)
	}

	//Auto migrate models
	db.AutoMigrate(&models.User{}, &models.Message{}, &models.ChatRoom{})

	//Initialize Gin router
	r := gin.Default()

	//User registration endpoint
	r.POST("/register", func(c *gin.Context) {
		auth.Register(c, db)
	})

	//User login endpoint
	r.POST("/login", func(c *gin.Context) {
		auth.Login(c, db)
	})

	//WebSocket endpoint for real-time messaging
	r.GET("/ws/:roomID", func(c *gin.Context) {
		websocket.HandleWebSocket(c, db)
	})

	//Start the server
	r.Run("")
}
