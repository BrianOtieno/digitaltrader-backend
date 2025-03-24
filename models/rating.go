package models

import (
	"time"

	"digitaltrader/db"
)

type Rating struct {
	ID         uint       `json:"id" gorm:"primaryKey;autoIncrement"`
	UID        *int       `json:"uid" gorm:"type:integer;null"`
	PID        *int       `json:"pid" gorm:"type:integer;null"`
	DID        *int       `json:"did" gorm:"type:integer;null"`
	SID        *int       `json:"sid" gorm:"type:integer;null"`
	Rate       *float64   `json:"rate" gorm:"type:decimal(10,2);null"`
	Msg        *string    `json:"msg" gorm:"type:text;null"`
	Way        *string    `json:"way" gorm:"type:varchar(255);null"`
	Timestamp  *time.Time `json:"timestamp" gorm:"type:date;null"`
	Status     int8       `json:"status" gorm:"type:tinyint;default:0"`
	ExtraField *string    `json:"extra_field" gorm:"type:text;null"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`

	// Relationships (assuming User, Product, Driver, and Store models exist)
	User    User    `gorm:"foreignKey:UID;references:ID;constraint:OnDelete:SET NULL"`
	Product Product `gorm:"foreignKey:PID;references:ID;constraint:OnDelete:SET NULL"`
	Driver  Driver  `gorm:"foreignKey:DID;references:ID;constraint:OnDelete:SET NULL"`
	Store   Store   `gorm:"foreignKey:SID;references:ID;constraint:OnDelete:SET NULL"`
}

// MigrateRating auto-migrates the Rating model
func MigrateRating() {
	db := db.InitDB()
	// Migrate dependent models first due to foreign key dependencies
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{}, &Product{}, &Driver{}, &Store{}, &Rating{})
}

// CreateRating creates a new rating
func CreateRating(rating *Rating) error {
	db := db.InitDB()
	return db.Create(rating).Error
}

// GetRating retrieves a rating by ID
func GetRating(id uint) (*Rating, error) {
	db := db.InitDB()
	var rating Rating
	if err := db.Preload("User").Preload("Product").Preload("Driver").Preload("Store").First(&rating, id).Error; err != nil {
		return nil, err
	}
	return &rating, nil
}

// UpdateRating updates an existing rating
func UpdateRating(rating *Rating, id uint) error {
	db := db.InitDB()
	return db.Model(&Rating{}).Where("id = ?", id).Updates(rating).Error
}

// DeleteRating deletes a rating by ID
func DeleteRating(id uint) error {
	db := db.InitDB()
	return db.Delete(&Rating{}, id).Error
}

// GetAllRatings retrieves all ratings
func GetAllRatings() ([]Rating, error) {
	db := db.InitDB()
	var ratings []Rating
	if err := db.Preload("User").Preload("Product").Preload("Driver").Preload("Store").Find(&ratings).Error; err != nil {
		return nil, err
	}
	return ratings, nil
}
