package models

import (
	"time"

	"digitaltrader/db"
)

type Language struct {
	ID         uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name       string    `json:"name" gorm:"type:varchar(255)"`
	Cover      string    `json:"cover" gorm:"type:varchar(255)"`
	Content    string    `json:"content" gorm:"type:text"`
	IsDefault  int8      `json:"is_default" gorm:"type:tinyint"`
	ExtraField *string   `json:"extra_field" gorm:"type:text;null"`
	Status     int8      `json:"status" gorm:"type:tinyint"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// MigrateLanguage auto-migrates the Language model
func MigrateLanguage() {
	db := db.InitDB()
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Language{})
}

// CreateLanguage creates a new language
func CreateLanguage(language *Language) error {
	db := db.InitDB()
	return db.Create(language).Error
}

// GetLanguage retrieves a language by ID
func GetLanguage(id uint) (*Language, error) {
	db := db.InitDB()
	var language Language
	if err := db.First(&language, id).Error; err != nil {
		return nil, err
	}
	return &language, nil
}

// UpdateLanguage updates an existing language
func UpdateLanguage(language *Language, id uint) error {
	db := db.InitDB()
	return db.Model(&Language{}).Where("id = ?", id).Updates(language).Error
}

// DeleteLanguage deletes a language by ID
func DeleteLanguage(id uint) error {
	db := db.InitDB()
	return db.Delete(&Language{}, id).Error
}

// GetAllLanguages retrieves all languages
func GetAllLanguages() ([]Language, error) {
	db := db.InitDB()
	var languages []Language
	if err := db.Find(&languages).Error; err != nil {
		return nil, err
	}
	return languages, nil
}
