package models

import (
	"time"

	"digitaltrader/db"
)

type ChatRoom struct {
	ID           uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	UID          int       `json:"uid" gorm:"type:integer"`
	Participants string    `json:"participants" gorm:"type:varchar(255)"`
	ExtraField   *string   `json:"extra_field" gorm:"type:text;null"`
	Status       int8      `json:"status" gorm:"type:tinyint;default:1"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// MigrateChatRoom auto-migrates the ChatRoom model
func MigrateChatRoom() {
	db := db.InitDB()
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&ChatRoom{})
}

// CreateChatRoom creates a new chat room
func CreateChatRoom(chatRoom *ChatRoom) error {
	db := db.InitDB()
	return db.Create(chatRoom).Error
}

// GetChatRoom retrieves a chat room by ID
func GetChatRoom(id uint) (*ChatRoom, error) {
	db := db.InitDB()
	var chatRoom ChatRoom
	if err := db.First(&chatRoom, id).Error; err != nil {
		return nil, err
	}
	return &chatRoom, nil
}

// UpdateChatRoom updates an existing chat room
func UpdateChatRoom(chatRoom *ChatRoom, id uint) error {
	db := db.InitDB()
	return db.Model(&ChatRoom{}).Where("id = ?", id).Updates(chatRoom).Error
}

// DeleteChatRoom deletes a chat room by ID
func DeleteChatRoom(id uint) error {
	db := db.InitDB()
	return db.Delete(&ChatRoom{}, id).Error
}

// GetAllChatRooms retrieves all chat rooms
func GetAllChatRooms() ([]ChatRoom, error) {
	db := db.InitDB()
	var chatRooms []ChatRoom
	if err := db.Find(&chatRooms).Error; err != nil {
		return nil, err
	}
	return chatRooms, nil
}
