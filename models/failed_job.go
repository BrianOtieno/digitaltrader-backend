package models

import (
	"time"

	"digitaltrader/db"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FailedJob struct {
	ID         uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	UUID       string    `json:"uuid" gorm:"type:uuid;unique"`
	Connection string    `json:"connection" gorm:"type:text"`
	Queue      string    `json:"queue" gorm:"type:text"`
	Payload    string    `json:"payload" gorm:"type:longtext"`
	Exception  string    `json:"exception" gorm:"type:longtext"`
	FailedAt   time.Time `json:"failed_at" gorm:"default:CURRENT_TIMESTAMP"`
}

// BeforeCreate hook to generate UUID
func (fj *FailedJob) BeforeCreate(tx *gorm.DB) (err error) {
	if fj.UUID == "" {
		fj.UUID = uuid.New().String()
	}
	return
}

// MigrateFailedJob auto-migrates the FailedJob model
func MigrateFailedJob() {
	db := db.InitDB()
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&FailedJob{})
}

// CreateFailedJob creates a new failed job record
func CreateFailedJob(failedJob *FailedJob) error {
	db := db.InitDB()
	return db.Create(failedJob).Error
}

// GetFailedJob retrieves a failed job by ID
func GetFailedJob(id uint) (*FailedJob, error) {
	db := db.InitDB()
	var failedJob FailedJob
	if err := db.First(&failedJob, id).Error; err != nil {
		return nil, err
	}
	return &failedJob, nil
}

// UpdateFailedJob updates an existing failed job
func UpdateFailedJob(failedJob *FailedJob, id uint) error {
	db := db.InitDB()
	return db.Model(&FailedJob{}).Where("id = ?", id).Updates(failedJob).Error
}

// DeleteFailedJob deletes a failed job by ID
func DeleteFailedJob(id uint) error {
	db := db.InitDB()
	return db.Delete(&FailedJob{}, id).Error
}

// GetAllFailedJobs retrieves all failed jobs
func GetAllFailedJobs() ([]FailedJob, error) {
	db := db.InitDB()
	var failedJobs []FailedJob
	if err := db.Find(&failedJobs).Error; err != nil {
		return nil, err
	}
	return failedJobs, nil
}
