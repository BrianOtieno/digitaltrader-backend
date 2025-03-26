package models

import (
	"time"

	"digitaltrader/db"

	"gorm.io/gorm"
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

// SeedPayments seeds the payments table with initial data
func SeedPayments() error {
	db := db.InitDB()

	// // Truncate the table (similar to Payments::truncate() in Laravel)
	// if err := db.Exec("TRUNCATE TABLE payments").Error; err != nil {
	// 	return err
	// }

	// List of payment methods to seed
	payments := []Payment{
		{Name: "MPESA", Env: 1, Status: 1, CurrencyCode: "KSH"},
		{Name: "COD", Env: 1, Status: 1, CurrencyCode: "USD"},
		{Name: "Stripe", Env: 1, Status: 1, CurrencyCode: "USD"},
		{Name: "PayPal", Env: 1, Status: 1, CurrencyCode: "USD"},
		{Name: "PayTM", Env: 1, Status: 1, CurrencyCode: "INR"},
		{Name: "RazorPay", Env: 1, Status: 1, CurrencyCode: "INR"},
		{Name: "InstaMOJO", Env: 1, Status: 1, CurrencyCode: "INR"},
		{Name: "PayStack", Env: 1, Status: 1, CurrencyCode: "NGN"},
		{Name: "Flutterwave", Env: 1, Status: 1, CurrencyCode: "NGN"},
	}

	// Create each payment record if it doesn't already exist
	for _, payment := range payments {
		var existing Payment
		if err := db.Where("name = ?", payment.Name).First(&existing).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				// If not found, create the payment
				if err := db.Create(&payment).Error; err != nil {
					return err
				}
			} else {
				// Other errors
				return err
			}
		}
		// If found, you could optionally update it here if needed
		// For now, we skip duplicates
	}

	return nil
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
