package models

import (
	"time"

	"digitaltrader/db"
)

type Driver struct {
	ID          uint       `json:"id" gorm:"primaryKey;autoIncrement"`
	FirstName   string     `json:"first_name" gorm:"type:varchar(255)"`
	LastName    string     `json:"last_name" gorm:"type:varchar(255)"`
	Email       string     `json:"email" gorm:"type:varchar(255);unique"`
	Password    string     `json:"password" gorm:"type:varchar(255)"`
	CountryCode string     `json:"country_code" gorm:"type:varchar(255)"`
	Mobile      string     `json:"mobile" gorm:"type:varchar(255)"`
	Address     string     `json:"address" gorm:"type:text"`
	Date        *time.Time `json:"date" gorm:"type:date;null"`
	CityID      int        `json:"city" gorm:"type:integer"` // Renamed to CityID for clarity
	Cover       *string    `json:"cover" gorm:"type:varchar(255);null"`
	Lat         *string    `json:"lat" gorm:"type:varchar(255);null"`
	Lng         *string    `json:"lng" gorm:"type:varchar(255);null"`
	Gender      *int8      `json:"gender" gorm:"type:tinyint;null"`
	Verified    *int8      `json:"verified" gorm:"type:tinyint;null"`
	FCMToken    *string    `json:"fcm_token" gorm:"type:text;null"`
	Current     *string    `json:"current" gorm:"type:varchar(255);null"`
	Others      *string    `json:"others" gorm:"type:text;null"`
	StripeKey   *string    `json:"stripe_key" gorm:"type:text;null"`
	ExtraField  *string    `json:"extra_field" gorm:"type:text;null"`
	Status      int8       `json:"status" gorm:"type:tinyint;default:1"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`

	// Relationship
	City City `gorm:"foreignKey:CityID;references:ID;constraint:OnDelete:RESTRICT"`
}

// MigrateDriver auto-migrates the Driver model
func MigrateDriver() {
	db := db.InitDB()
	// Migrate City first due to foreign key dependency
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&City{}, &Driver{})
}

// CreateDriver creates a new driver
func CreateDriver(driver *Driver) error {
	db := db.InitDB()
	return db.Create(driver).Error
}

// GetDriver retrieves a driver by ID
func GetDriver(id uint) (*Driver, error) {
	db := db.InitDB()
	var driver Driver
	if err := db.Preload("City").First(&driver, id).Error; err != nil {
		return nil, err
	}
	return &driver, nil
}

// UpdateDriver updates an existing driver
func UpdateDriver(driver *Driver, id uint) error {
	db := db.InitDB()
	return db.Model(&Driver{}).Where("id = ?", id).Updates(driver).Error
}

// DeleteDriver deletes a driver by ID
func DeleteDriver(id uint) error {
	db := db.InitDB()
	return db.Delete(&Driver{}, id).Error
}

// GetAllDrivers retrieves all drivers
func GetAllDrivers() ([]Driver, error) {
	db := db.InitDB()
	var drivers []Driver
	if err := db.Preload("City").Find(&drivers).Error; err != nil {
		return nil, err
	}
	return drivers, nil
}
