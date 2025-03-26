package models

import (
	"time"

	"digitaltrader/db"

	"gorm.io/gorm"
)

type Page struct {
	ID         uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name       string    `json:"name" gorm:"type:varchar(255)"`
	Content    string    `json:"content" gorm:"type:text"`
	ExtraField *string   `json:"extra_field" gorm:"type:text;null"`
	Status     int8      `json:"status" gorm:"type:tinyint"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// MigratePage auto-migrates the Page model
func MigratePage() {
	db := db.InitDB()
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Page{})
}

// SeedPages seeds the pages table with initial data
func SeedPages() error {
	db := db.InitDB()

	// // Truncate the table (similar to Pages::truncate() in Laravel)
	// if err := db.Exec("TRUNCATE TABLE pages").Error; err != nil {
	// 	return err
	// }

	// List of pages to seed
	pages := []Page{
		{Name: "About us", Content: "About us", ExtraField: stringPtr("NA"), Status: 1},
		{Name: "Privacy", Content: "Privacy", ExtraField: stringPtr("NA"), Status: 1},
		{Name: "Terms & Conditions", Content: "Terms & Conditions", ExtraField: stringPtr("NA"), Status: 1},
		{Name: "Refund Policy", Content: "Refund Policy", ExtraField: stringPtr("NA"), Status: 1},
		{Name: "Frequently Asked Questions", Content: "Frequently Asked Questions", ExtraField: stringPtr("NA"), Status: 1},
		{Name: "Help", Content: "Help", ExtraField: stringPtr("NA"), Status: 1},
		{Name: "Legal Mentions", Content: "Legal Mentions", ExtraField: stringPtr("NA"), Status: 1},
		{Name: "Cookies", Content: "Cookies", ExtraField: stringPtr("NA"), Status: 1},
	}

	// Create each page record if it doesn't already exist
	for _, page := range pages {
		var existing Page
		if err := db.Where("name = ?", page.Name).First(&existing).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				// If not found, create the page
				if err := db.Create(&page).Error; err != nil {
					return err
				}
			} else {
				// Other errors
				return err
			}
		}
		// If found, skip duplicates (or update if desired)
	}

	return nil
}

// Helper function to create a string pointer
func stringPtr(s string) *string {
	return &s
}

// CreatePage creates a new page
func CreatePage(page *Page) error {
	db := db.InitDB()
	return db.Create(page).Error
}

// GetPage retrieves a page by ID
func GetPage(id uint) (*Page, error) {
	db := db.InitDB()
	var page Page
	if err := db.First(&page, id).Error; err != nil {
		return nil, err
	}
	return &page, nil
}

// UpdatePage updates an existing page
func UpdatePage(page *Page, id uint) error {
	db := db.InitDB()
	return db.Model(&Page{}).Where("id = ?", id).Updates(page).Error
}

// DeletePage deletes a page by ID
func DeletePage(id uint) error {
	db := db.InitDB()
	return db.Delete(&Page{}, id).Error
}

// GetAllPages retrieves all pages
func GetAllPages() ([]Page, error) {
	db := db.InitDB()
	var pages []Page
	if err := db.Find(&pages).Error; err != nil {
		return nil, err
	}
	return pages, nil
}
