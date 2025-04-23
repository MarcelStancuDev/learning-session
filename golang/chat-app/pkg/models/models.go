package models

import "gorm.io/gorm"

// User represents the user model
type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string
}

// Message represents the message model
type Message struct {
	gorm.Model
	Content string
	UserID  uint
	RoomID 	uint	
}

//ChatRoom represents the chat room model
type ChatRoom struct {
	gorm.Model
	Name string `gorm:"unique"`
}
