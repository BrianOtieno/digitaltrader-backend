package models

import (
	"time"

	"digitaltrader/db"
)

type Store struct {
	ID              uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	UID             int       `json:"uid" gorm:"type:integer"`
	Name            string    `json:"name" gorm:"type:text"`
	Mobile          string    `json:"mobile" gorm:"type:varchar(255)"`
	Lat             string    `json:"lat" gorm:"type:varchar(255)"`
	Lng             string    `json:"lng" gorm:"type:varchar(255)"`
	Verified        *int8     `json:"verified" gorm:"type:tinyint;null"`
	Address         *string   `json:"address" gorm:"type:text;null"`
	Descriptions    *string   `json:"descriptions" gorm:"type:text;null"`
	Images          *string   `json:"images" gorm:"type:text;null"`
	Cover           *string   `json:"cover" gorm:"type:text;null"`
	Commission      *float64  `json:"commission" gorm:"type:decimal(10,2);null"`
	OpenTime        *string   `json:"open_time" gorm:"type:varchar(255);null"`
	CloseTime       *string   `json:"close_time" gorm:"type:varchar(255);null"`
	IsClosed        *int8     `json:"is_closed" gorm:"type:tinyint;null"`
	CertificateURL  *string   `json:"certificate_url" gorm:"type:varchar(255);null"`
	CertificateType *string   `json:"certificate_type" gorm:"type:varchar(255);null"`
	Rating          *float64  `json:"rating" gorm:"type:decimal(10,2);null"`
	TotalRating     *int      `json:"total_rating" gorm:"type:integer;null"`
	CID             *int      `json:"cid" gorm:"type:integer;null"`
	Zipcode         *string   `json:"zipcode" gorm:"type:text;null"`
	ExtraField      *string   `json:"extra_field" gorm:"type:text;null"`
	Status          int8      `json:"status" gorm:"type:tinyint"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// MigrateStore auto-migrates the Store model
func MigrateStore() {
	db := db.InitDB()
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Store{})
}

// CreateStore creates a new store
func CreateStore(store *Store) error {
	db := db.InitDB()
	return db.Create(store).Error
}

// GetStore retrieves a store by ID
func GetStore(id uint) (*Store, error) {
	db := db.InitDB()
	var store Store
	if err := db.First(&store, id).Error; err != nil {
		return nil, err
	}
	return &store, nil
}

// UpdateStore updates an existing store
func UpdateStore(store *Store, id uint) error {
	db := db.InitDB()
	return db.Model(&Store{}).Where("id = ?", id).Updates(store).Error
}

// DeleteStore deletes a store by ID
func DeleteStore(id uint) error {
	db := db.InitDB()
	return db.Delete(&Store{}, id).Error
}

// GetAllStores retrieves all stores
func GetAllStores() ([]Store, error) {
	db := db.InitDB()
	var stores []Store
	if err := db.Find(&stores).Error; err != nil {
		return nil, err
	}
	return stores, nil
}
