package models

import (
	"digitaltrader/db"
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type ActivityLog struct {
	gorm.Model
	ID          uint      `json:"id" gorm:"primary_key:auto_increment"`
	EntityType  string    `json:"entity_type" binding:"required"` // E.g., "News", "Author"
	EntityID    uint      `json:"entity_id" binding:"required"`   // ID of the affected entity
	Action      string    `json:"action" binding:"required"`      // E.g., "CREATE", "UPDATE", "DELETE"
	PerformedBy uint      `json:"performed_by"`                   // User ID of the performer
	User        *User     `gorm:"foreignKey:PerformedBy;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	Details     string    `json:"details" gorm:"type:json"`        // JSON data about changes
	Timestamp   time.Time `json:"timestamp" gorm:"autoCreateTime"` // Time of the action
}

func MigrateActivityLog() {
	db := db.InitDB()
	db.AutoMigrate(&ActivityLog{})
}

func LogActivity(entityType string, entityID uint, action string, performedBy uint, details interface{}) error {
	db := db.InitDB()

	// Convert details to JSON
	detailsJSON, err := json.Marshal(details)
	if err != nil {
		return err
	}

	activityLog := ActivityLog{
		EntityType:  entityType,
		EntityID:    entityID,
		Action:      action,
		PerformedBy: performedBy,
		Details:     string(detailsJSON),
		Timestamp:   time.Now(),
	}

	return db.Create(&activityLog).Error
}

func GetActivityLogs(entityType string, entityID uint, logs *[]ActivityLog) error {
	db := db.InitDB()
	return db.Where("entity_type = ? AND entity_id = ?", entityType, entityID).Order("timestamp desc").Find(logs).Error
}
