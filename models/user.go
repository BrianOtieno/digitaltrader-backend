// models/user_model.go
package models

import (
	"digitaltrader/db"
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// LoginRequest is the data structure to receive login details
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type User struct {
	gorm.Model
	ID            uint   `json:"id" gorm:"primary_key:auto_increment"`
	FirstName     string `json:"firstname"`
	LastName      string `json:"lastname"`
	Username      string `json:"username" gorm:"unique"`
	Password      string `json:"password"` // Nullable for social logins
	ProfilePhoto  string `json:"profile_photo"`
	IsActive      bool   `json:"is_active" gorm:"default:true;not null"`
	IsDeleted     bool   `json:"is_deleted" gorm:"default:false;not null"`
	GoogleOAuthID string `json:"google_oauth_id" gorm:"unique;default:null"` // For Google login
	AppleOAuthID  string `json:"apple_oauth_id" gorm:"unique;default:null"`  // For Apple login
	Email         string `json:"email" gorm:"unique"`
}

// MigrateUser auto-migrates the User model, including new fields
func MigrateUser() {
	db := db.InitDB()
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{})
}

// HashPassword hashes the user's password before saving it
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CreateUser creates a new user in the database
func (user *User) CreateUser() error {
	// Initialize DB connection
	db := db.InitDB()

	// Hash the password before storing it
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	// Check if the user already exists based on username or email
	var existingUser User
	if err := db.Where("username = ?", user.Username).First(&existingUser).Error; err == nil {
		return fmt.Errorf("user already exists") // return custom error
	}

	// Create the user in the database
	if err := db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

// UpdateUser updates an existing user's information
func UpdateUser(user *User, id string) error {
	db := db.InitDB()
	if err := db.Where("id = ?", id).Updates(&user).Error; err != nil {
		return err
	}
	return nil
}

// GetAllUsers retrieves all users from the database
func GetAllUsers(users *[]User) error {
	db := db.InitDB()
	if err := db.Find(&users).Error; err != nil {
		return err
	}
	return nil
}

// GetUserByID retrieves a user by their ID
func GetUserByID(user *User, id string) error {
	db := db.InitDB()
	if err := db.Where("id = ?", id).Find(&user).Error; err != nil {
		return err
	}
	return nil
}

// // GetUserByUsername retrieves a user by their username
// func GetUserByUsername(user *User, username string) error {
// 	db := db.InitDB() // Initialize the database connection
// 	if err := db.Where("username = ?", username).Find(&user).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

// DeleteUser deletes a user from the database
func DeleteUser(user *User, id string) error {
	db := db.InitDB()
	if err := db.Where("id = ?", id).Delete(&user).Error; err != nil {
		return err
	}
	return nil
}

// FindFirstUser finds the first user by their ID
func FindFirstUser(user *User, id string) error {
	db := db.InitDB()
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		return err
	}
	return nil
}

// GetUserByUsername retrieves a user from the database by their username
func GetUserByUsername(user *User, username string) error {
	db := db.InitDB()
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return errors.New("user not found")
	}
	return nil
}

// GetUserByEmail retrieves a user by their email
func GetUserByEmail(user *User, email string) error {
	db := db.InitDB()
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return errors.New("user not found")
	}
	return nil
}

// CheckUserByUsername retrieves a user by username (email)
func CheckUserByUsername(username string) (*User, error) {
	db := db.InitDB()
	var user User
	err := db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// CreateNewUser creates a new user in the database
func CreateNewUser(firstName, lastName, username, profilePhoto string) (*User, error) {
	db := db.InitDB()
	newUser := User{
		FirstName:    firstName,
		LastName:     lastName,
		Username:     username,
		ProfilePhoto: profilePhoto,
	}
	if err := db.Create(&newUser).Error; err != nil {
		return nil, err
	}
	return &newUser, nil
}
