package models

import (
	"time"

	"digitaltrader/db"
)

type ChatMessage struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	RoomID      int       `json:"room_id" gorm:"type:integer"`
	UID         string    `json:"uid" gorm:"type:varchar(255)"`
	FromID      int       `json:"from_id" gorm:"type:integer"`
	Message     string    `json:"message" gorm:"type:text"`
	MessageType string    `json:"message_type" gorm:"type:varchar(255)"`
	Timestamp   *string   `json:"timestamp" gorm:"type:varchar(255);null"`
	ExtraField  *string   `json:"extra_field" gorm:"type:text;null"`
	Status      int8      `json:"status" gorm:"type:tinyint;default:1"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	// Relationship
	ChatRoom ChatRoom `gorm:"foreignKey:RoomID;references:ID;constraint:OnDelete:RESTRICT"`
}

// MigrateChatMessage auto-migrates the ChatMessage model
func MigrateChatMessage() {
	db := db.InitDB()
	// Migrate ChatRoom first due to foreign key dependency
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&ChatRoom{}, &ChatMessage{})
}

// CreateChatMessage creates a new chat message
func CreateChatMessage(chatMessage *ChatMessage) error {
	db := db.InitDB()
	return db.Create(chatMessage).Error
}

// GetChatMessage retrieves a chat message by ID
func GetChatMessage(id uint) (*ChatMessage, error) {
	db := db.InitDB()
	var chatMessage ChatMessage
	if err := db.Preload("ChatRoom").First(&chatMessage, id).Error; err != nil {
		return nil, err
	}
	return &chatMessage, nil
}

// UpdateChatMessage updates an existing chat message
func UpdateChatMessage(chatMessage *ChatMessage, id uint) error {
	db := db.InitDB()
	return db.Model(&ChatMessage{}).Where("id = ?", id).Updates(chatMessage).Error
}

// DeleteChatMessage deletes a chat message by ID
func DeleteChatMessage(id uint) error {
	db := db.InitDB()
	return db.Delete(&ChatMessage{}, id).Error
}

// GetAllChatMessages retrieves all chat messages
func GetAllChatMessages() ([]ChatMessage, error) {
	db := db.InitDB()
	var chatMessages []ChatMessage
	if err := db.Preload("ChatRoom").Find(&chatMessages).Error; err != nil {
		return nil, err
	}
	return chatMessages, nil
}
