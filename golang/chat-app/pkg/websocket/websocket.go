package websocket

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"chat-app/pkg/models"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleWebSocket(c *gin.Context, db *gorm.DB) {
	roomID := c.Param("roomID")

	conn, err := upgrader.Upgrade(c.Writer, c.Writer, nil)
	if err != nil {
		log.Println("Failed to upgrade connection:", err)
		return
	}

	defer conn.Close()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}

		broadcastMessage(roomID, message)
	}
}

func broadcastMessage(roomID string, message []byte, db *gorm.DB) {
	log.Printf("Broadcasting message to room %s: %s", roomID, message)
}