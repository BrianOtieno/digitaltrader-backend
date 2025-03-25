package models

import (
	"time"

	"digitaltrader/db"

	"gorm.io/gorm"
)

type Setting struct {
	ID                 uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	CurrencySymbol     string    `json:"currency_symbol" gorm:"type:varchar(255)"`
	CurrencySide       string    `json:"currency_side" gorm:"type:varchar(255)"`
	CurrencyCode       string    `json:"currency_code" gorm:"type:varchar(255)"`
	AppDirection       string    `json:"app_direction" gorm:"type:varchar(255)"`
	Logo               string    `json:"logo" gorm:"type:varchar(255)"`
	SMSName            string    `json:"sms_name" gorm:"type:varchar(255)"`
	SMSCreds           string    `json:"sms_creds" gorm:"type:text"`
	Delivery           int8      `json:"delivery" gorm:"type:tinyint"`
	FindType           int8      `json:"find_type" gorm:"type:tinyint;default:0"`
	MakeOrders         int8      `json:"make_orders" gorm:"type:tinyint;default:0"`
	ResetPwd           int8      `json:"reset_pwd" gorm:"type:tinyint;default:0"`
	UserLogin          int8      `json:"user_login" gorm:"type:tinyint;default:0"`
	StoreLogin         int8      `json:"store_login" gorm:"type:tinyint;default:0"`
	UserVerifyWith     int8      `json:"user_verify_with" gorm:"type:tinyint;default:0"`
	SearchRadius       float64   `json:"search_radius" gorm:"type:decimal(10,2);default:10"`
	DriverLogin        int8      `json:"driver_login" gorm:"type:tinyint;default:0"`
	WebLogin           int8      `json:"web_login" gorm:"type:tinyint;default:0"`
	LoginStyle         int8      `json:"login_style" gorm:"type:tinyint;default:1"`
	RegisterStyle      int8      `json:"register_style" gorm:"type:tinyint;default:1"`
	HomePageStyleApp   int8      `json:"home_page_style_app" gorm:"type:tinyint;default:1"`
	CountryModal       string    `json:"country_modal" gorm:"type:text"`
	WebCategory        string    `json:"web_category" gorm:"type:text"`
	DefaultCountryCode string    `json:"default_country_code" gorm:"type:varchar(255)"`
	DefaultCityID      *string   `json:"default_city_id" gorm:"type:varchar(255);null"`
	DefaultDeliveryZip *string   `json:"default_delivery_zip" gorm:"type:varchar(255);null"`
	Social             *string   `json:"social" gorm:"type:text;null"`
	AppColor           string    `json:"app_color" gorm:"type:text"`
	AppStatus          int8      `json:"app_status" gorm:"type:tinyint;default:1"`
	DriverAssign       int8      `json:"driver_assign" gorm:"type:tinyint;default:0"`
	FCMToken           *string   `json:"fcm_token" gorm:"type:text;null"`
	ExtraField         *string   `json:"extra_field" gorm:"type:text;null"`
	Status             int8      `json:"status" gorm:"type:tinyint"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

// MigrateSetting auto-migrates the Setting model
func MigrateSetting() {
	db := db.InitDB()
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Setting{})
}

// CreateSetting creates or updates the single settings record
func CreateSetting(setting *Setting) error {
	db := db.InitDB()
	var existing Setting
	if err := db.First(&existing).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return db.Create(setting).Error
		}
		return err
	}
	// If a record exists, update it instead
	return db.Model(&existing).Updates(setting).Error
}

// GetSetting retrieves the single settings record
func GetSetting() (*Setting, error) {
	db := db.InitDB()
	var setting Setting
	if err := db.First(&setting).Error; err != nil {
		return nil, err
	}
	return &setting, nil
}

// UpdateSetting updates the single settings record
func UpdateSetting(setting *Setting) error {
	db := db.InitDB()
	var existing Setting
	if err := db.First(&existing).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return db.Create(setting).Error
		}
		return err
	}
	return db.Model(&existing).Updates(setting).Error
}

// DeleteSetting deletes the single settings record (not typical for settings, but included for completeness)
func DeleteSetting() error {
	db := db.InitDB()
	return db.Delete(&Setting{}, 1).Error
}

// GetAllSettings retrieves all settings (though typically only one record exists)
func GetAllSettings() ([]Setting, error) {
	db := db.InitDB()
	var settings []Setting
	if err := db.Find(&settings).Error; err != nil {
		return nil, err
	}
	return settings, nil
}
