package models

import (
	"time"

	"digitaltrader/db"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// TransferStatus represents the possible status values
type TransferStatus string

const (
	StatusExchange TransferStatus = "exchange"
	StatusTransfer TransferStatus = "transfer"
	StatusPaid     TransferStatus = "paid"
	StatusRefund   TransferStatus = "refund"
	StatusGift     TransferStatus = "gift"
)

type Transfer struct {
	ID         uint           `json:"id" gorm:"primaryKey"`
	FromType   string         `json:"from_type" gorm:"index:idx_from,unique"`
	FromID     uint           `json:"from_id" gorm:"index:idx_from,unique"`
	ToType     string         `json:"to_type" gorm:"index:idx_to,unique"`
	ToID       uint           `json:"to_id" gorm:"index:idx_to,unique"`
	Status     TransferStatus `json:"status" gorm:"type:varchar(20);default:'transfer';check:status IN ('exchange','transfer','paid','refund','gift')"`
	StatusLast TransferStatus `json:"status_last" gorm:"type:varchar(20);check:status_last IN ('exchange','transfer','paid','refund','gift')"`
	DepositID  uint           `json:"deposit_id" gorm:"not null"`
	WithdrawID uint           `json:"withdraw_id" gorm:"not null"`
	Discount   float64        `json:"discount" gorm:"type:decimal(64,0);default:0"`
	Fee        float64        `json:"fee" gorm:"type:decimal(64,0);default:0"`
	UUID       string         `json:"uuid" gorm:"type:uuid;unique"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`

	// Foreign key relationships
	Deposit  Transaction `gorm:"foreignKey:DepositID;references:ID;constraint:OnDelete:CASCADE"`
	Withdraw Transaction `gorm:"foreignKey:WithdrawID;references:ID;constraint:OnDelete:CASCADE"`
}

// BeforeCreate hook to generate UUID if not set
func (t *Transfer) BeforeCreate(tx *gorm.DB) (err error) {
	if t.UUID == "" {
		t.UUID = uuid.New().String()
	}
	return
}

// MigrateTransfer auto-migrates the Transfer model
func MigrateTransfer() {
	db := db.InitDB()
	db.AutoMigrate(&Transaction{}) // Ensure Transaction table exists first
	db.AutoMigrate(&Transfer{})
}

// CreateTransfer creates a new transfer record
func CreateTransfer(transfer *Transfer) error {
	db := db.InitDB()
	return db.Create(transfer).Error
}

// GetTransferByID retrieves a transfer record by ID
func GetTransferByID(id uint) (*Transfer, error) {
	db := db.InitDB()
	var transfer Transfer
	if err := db.Preload("Deposit").Preload("Withdraw").First(&transfer, id).Error; err != nil {
		return nil, err
	}
	return &transfer, nil
}

// GetAllTransfers retrieves all transfer records
func GetAllTransfers() ([]Transfer, error) {
	db := db.InitDB()
	var transfers []Transfer
	if err := db.Preload("Deposit").Preload("Withdraw").Find(&transfers).Error; err != nil {
		return nil, err
	}
	return transfers, nil
}

// UpdateTransfer updates an existing transfer record
func UpdateTransfer(transfer *Transfer, id uint) error {
	db := db.InitDB()
	return db.Model(&Transfer{}).Where("id = ?", id).Updates(transfer).Error
}

// DeleteTransfer deletes a transfer record by ID
func DeleteTransfer(id uint) error {
	db := db.InitDB()
	return db.Delete(&Transfer{}, id).Error
}
