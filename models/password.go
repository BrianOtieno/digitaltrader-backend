package models

import (
	"time"

	"digitaltrader/db"
)

type PasswordResetToken struct {
	Email     string    `json:"email" gorm:"primaryKey"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
}

// MigrateUser auto-migrates the User model, including new fields
func MigratePasswordResetToken() {
	db := db.InitDB()
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&PasswordResetToken{})
}

// CreatePasswordResetToken creates a new password reset token
func CreatePasswordResetToken(token *PasswordResetToken) error {
	db := db.InitDB()
	if err := db.Create(&token).Error; err != nil {
		return err
	}
	return nil
}

// GetPasswordResetToken retrieves a password reset token by email
func GetPasswordResetToken(token *PasswordResetToken, email string) error {
	db := db.InitDB()
	if err := db.Where("email = ?", email).First(&token).Error; err != nil {
		return err
	}
	return nil
}

// UpdatePasswordResetToken updates an existing password reset token
func UpdatePasswordResetToken(token *PasswordResetToken, email string) error {
	db := db.InitDB()
	if err := db.Where("email = ?", email).Updates(&token).Error; err != nil {
		return err
	}
	return nil
}

// DeletePasswordResetToken deletes a password reset token by email
func DeletePasswordResetToken(email string) error {
	db := db.InitDB()
	if err := db.Where("email = ?", email).Delete(&PasswordResetToken{}).Error; err != nil {
		return err
	}
	return nil
}

// GetAllPasswordResetTokens retrieves all password reset tokens
func GetAllPasswordResetTokens(tokens *[]PasswordResetToken) error {
	db := db.InitDB()
	if err := db.Find(&tokens).Error; err != nil {
		return err
	}
	return nil
}
