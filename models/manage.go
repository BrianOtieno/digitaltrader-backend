package models

import (
	"time"

	"digitaltrader/db"
)

type Manage struct {
	ID         uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	AppClose   int8      `json:"app_close" gorm:"type:tinyint;default:1"` // 1 = open, 0 = closed
	Message    *string   `json:"message" gorm:"type:text;null"`
	DateTime   *string   `json:"date_time" gorm:"type:varchar(255);null"`
	Status     int8      `json:"status" gorm:"type:tinyint;default:0"`
	ExtraField *string   `json:"extra_field" gorm:"type:text;null"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// MigrateManage auto-migrates the Manage model
func MigrateManage() {
	db := db.InitDB()
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Manage{})
}

// CreateManage creates a new manage record
func CreateManage(manage *Manage) error {
	db := db.InitDB()
	return db.Create(manage).Error
}

// GetManage retrieves a manage record by ID
func GetManage(id uint) (*Manage, error) {
	db := db.InitDB()
	var manage Manage
	if err := db.First(&manage, id).Error; err != nil {
		return nil, err
	}
	return &manage, nil
}

// UpdateManage updates an existing manage record
func UpdateManage(manage *Manage, id uint) error {
	db := db.InitDB()
	return db.Model(&Manage{}).Where("id = ?", id).Updates(manage).Error
}

// DeleteManage deletes a manage record by ID
func DeleteManage(id uint) error {
	db := db.InitDB()
	return db.Delete(&Manage{}, id).Error
}

// GetAllManages retrieves all manage records
func GetAllManages() ([]Manage, error) {
	db := db.InitDB()
	var manages []Manage
	if err := db.Find(&manages).Error; err != nil {
		return nil, err
	}
	return manages, nil
}
