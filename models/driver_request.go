package models

import (
	"time"

	"digitaltrader/db"
)

type DriverRequest struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	FirstName   string    `json:"first_name" gorm:"type:varchar(255)"`
	LastName    string    `json:"last_name" gorm:"type:varchar(255)"`
	Email       string    `json:"email" gorm:"type:varchar(255);unique"`
	Password    string    `json:"password" gorm:"type:varchar(255)"`
	CountryCode string    `json:"country_code" gorm:"type:varchar(255)"`
	Mobile      string    `json:"mobile" gorm:"type:varchar(255)"`
	Address     string    `json:"address" gorm:"type:text"`
	CityID      int       `json:"city" gorm:"type:integer"`
	Cover       *string   `json:"cover" gorm:"type:varchar(255);null"`
	Lat         *string   `json:"lat" gorm:"type:varchar(255);null"`
	Lng         *string   `json:"lng" gorm:"type:varchar(255);null"`
	Gender      *int8     `json:"gender" gorm:"type:tinyint;null"`
	ExtraField  *string   `json:"extra_field" gorm:"type:text;null"`
	Status      int8      `json:"status" gorm:"type:tinyint;default:1"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	// Relationship (assuming a City model exists)
	City City `gorm:"foreignKey:CityID;references:ID;constraint:OnDelete:RESTRICT"`
}

// MigrateDriverRequest auto-migrates the DriverRequest model
func MigrateDriverRequest() {
	db := db.InitDB()
	// Migrate City first due to foreign key dependency
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&City{}, &DriverRequest{})
}

// CreateDriverRequest creates a new driver request record
func CreateDriverRequest(driverRequest *DriverRequest) error {
	db := db.InitDB()
	return db.Create(driverRequest).Error
}

// GetDriverRequest retrieves a driver request record by ID
func GetDriverRequest(id uint) (*DriverRequest, error) {
	db := db.InitDB()
	var driverRequest DriverRequest
	if err := db.Preload("City").First(&driverRequest, id).Error; err != nil {
		return nil, err
	}
	return &driverRequest, nil
}

// UpdateDriverRequest updates an existing driver request record
func UpdateDriverRequest(driverRequest *DriverRequest, id uint) error {
	db := db.InitDB()
	return db.Model(&DriverRequest{}).Where("id = ?", id).Updates(driverRequest).Error
}

// DeleteDriverRequest deletes a driver request record by ID
func DeleteDriverRequest(id uint) error {
	db := db.InitDB()
	return db.Delete(&DriverRequest{}, id).Error
}

// GetAllDriverRequests retrieves all driver request records
func GetAllDriverRequests() ([]DriverRequest, error) {
	db := db.InitDB()
	var driverRequests []DriverRequest
	if err := db.Preload("City").Find(&driverRequests).Error; err != nil {
		return nil, err
	}
	return driverRequests, nil
}
