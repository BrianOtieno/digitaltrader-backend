package models

import (
	"time"

	"digitaltrader/db"
)

type StoreRequest struct {
	ID           uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	FirstName    string    `json:"first_name" gorm:"type:varchar(255)"`
	LastName     string    `json:"last_name" gorm:"type:varchar(255)"`
	Email        string    `json:"email" gorm:"type:varchar(255);unique"`
	Password     string    `json:"password" gorm:"type:varchar(255)"`
	CountryCode  string    `json:"country_code" gorm:"type:varchar(255)"`
	Mobile       string    `json:"mobile" gorm:"type:varchar(255)"`
	Name         string    `json:"name" gorm:"type:text"`
	Lat          string    `json:"lat" gorm:"type:varchar(255)"`
	Lng          string    `json:"lng" gorm:"type:varchar(255)"`
	Address      *string   `json:"address" gorm:"type:text;null"`
	Descriptions *string   `json:"descriptions" gorm:"type:text;null"`
	Cover        *string   `json:"cover" gorm:"type:text;null"`
	OpenTime     *string   `json:"open_time" gorm:"type:varchar(255);null"`
	CloseTime    *string   `json:"close_time" gorm:"type:varchar(255);null"`
	CID          *int      `json:"cid" gorm:"type:integer;null"`
	Zipcode      *string   `json:"zipcode" gorm:"type:text;null"`
	ExtraField   *string   `json:"extra_field" gorm:"type:text;null"`
	Status       int8      `json:"status" gorm:"type:tinyint"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`

	// Relationship (assuming a City model exists)
	City City `gorm:"foreignKey:CID;references:ID;constraint:OnDelete:SET NULL"`
}

// MigrateStoreRequest auto-migrates the StoreRequest model
func MigrateStoreRequest() {
	db := db.InitDB()
	// Migrate City first due to foreign key dependency
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&City{}, &StoreRequest{})
}

// CreateStoreRequest creates a new store request record
func CreateStoreRequest(storeRequest *StoreRequest) error {
	db := db.InitDB()
	return db.Create(storeRequest).Error
}

// GetStoreRequest retrieves a store request record by ID
func GetStoreRequest(id uint) (*StoreRequest, error) {
	db := db.InitDB()
	var storeRequest StoreRequest
	if err := db.Preload("City").First(&storeRequest, id).Error; err != nil {
		return nil, err
	}
	return &storeRequest, nil
}

// UpdateStoreRequest updates an existing store request record
func UpdateStoreRequest(storeRequest *StoreRequest, id uint) error {
	db := db.InitDB()
	return db.Model(&StoreRequest{}).Where("id = ?", id).Updates(storeRequest).Error
}

// DeleteStoreRequest deletes a store request record by ID
func DeleteStoreRequest(id uint) error {
	db := db.InitDB()
	return db.Delete(&StoreRequest{}, id).Error
}

// GetAllStoreRequests retrieves all store request records
func GetAllStoreRequests() ([]StoreRequest, error) {
	db := db.InitDB()
	var storeRequests []StoreRequest
	if err := db.Preload("City").Find(&storeRequests).Error; err != nil {
		return nil, err
	}
	return storeRequests, nil
}
