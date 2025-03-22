package models

import (
	"time"

	"digitaltrader/db"
)

type Payment struct {
	ID           uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name         string    `json:"name" gorm:"type:varchar(255)"`
	Env          int8      `json:"env" gorm:"type:tinyint"`
	Status       int8      `json:"status" gorm:"type:tinyint"`
	CurrencyCode string    `json:"currency_code" gorm:"type:varchar(255)"`
	Creds        *string   `json:"creds" gorm:"type:text;null"`
	ExtraField   *string   `json:"extra_field" gorm:"type:text;null"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// MigratePayment auto-migrates the Payment model
func MigratePayment() {
	db := db.InitDB()
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Payment{})
}

// CreatePayment creates a new payment
func CreatePayment(payment *Payment) error {
	db := db.InitDB()
	return db.Create(payment).Error
}

// GetPayment retrieves a payment by ID
func GetPayment(id uint) (*Payment, error) {
	db := db.InitDB()
	var payment Payment
	if err := db.First(&payment, id).Error; err != nil {
		return nil, err
	}
	return &payment, nil
}

// UpdatePayment updates an existing payment
func UpdatePayment(payment *Payment, id uint) error {
	db := db.InitDB()
	return db.Model(&Payment{}).Where("id = ?", id).Updates(payment).Error
}

// DeletePayment deletes a payment by ID
func DeletePayment(id uint) error {
	db := db.InitDB()
	return db.Delete(&Payment{}, id).Error
}

// GetAllPayments retrieves all payments
func GetAllPayments() ([]Payment, error) {
	db := db.InitDB()
	var payments []Payment
	if err := db.Find(&payments).Error; err != nil {
		return nil, err
	}
	return payments, nil
}
