package models

import (
	"time"

	"digitaltrader/db"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Wallet struct {
	ID            uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	HolderType    string    `json:"holder_type" gorm:"index:idx_holder_type_holder_id_slug,unique"`
	HolderID      uint      `json:"holder_id" gorm:"index:idx_holder_type_holder_id_slug,unique"`
	Name          string    `json:"name"`
	Slug          string    `json:"slug" gorm:"index;index:idx_holder_type_holder_id_slug,unique"`
	UUID          string    `json:"uuid" gorm:"type:uuid;unique"`
	Description   *string   `json:"description" gorm:"type:varchar(255);null"`
	Meta          *string   `json:"meta" gorm:"type:json;null"`
	Balance       float64   `json:"balance" gorm:"type:decimal(64,0);default:0"`
	DecimalPlaces uint16    `json:"decimal_places" gorm:"type:smallint unsigned;default:2"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`

	// Relationships
	Transactions []Transaction `gorm:"foreignKey:WalletID;constraint:OnDelete:RESTRICT"`
}

func (w *Wallet) BeforeCreate(tx *gorm.DB) (err error) {
	if w.UUID == "" {
		w.UUID = uuid.New().String()
	}
	return
}

// MigrateWallet auto-migrates the Wallet model
func MigrateWallet() {
	db := db.InitDB()
	// Migrate both tables, ensuring Wallet is created before Transaction due to FK
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Wallet{}, &Transaction{})
}

// CreateWallet creates a new wallet
func CreateWallet(wallet *Wallet) error {
	db := db.InitDB()
	return db.Create(wallet).Error
}

// GetWallet retrieves a wallet by ID
func GetWallet(id uint) (*Wallet, error) {
	db := db.InitDB()
	var wallet Wallet
	if err := db.Preload("Transactions").First(&wallet, id).Error; err != nil {
		return nil, err
	}
	return &wallet, nil
}

// UpdateWallet updates an existing wallet
func UpdateWallet(wallet *Wallet, id uint) error {
	db := db.InitDB()
	return db.Model(&Wallet{}).Where("id = ?", id).Updates(wallet).Error
}

// DeleteWallet deletes a wallet by ID
func DeleteWallet(id uint) error {
	db := db.InitDB()
	return db.Delete(&Wallet{}, id).Error
}

// GetAllWallets retrieves all wallets
func GetAllWallets() ([]Wallet, error) {
	db := db.InitDB()
	var wallets []Wallet
	if err := db.Preload("Transactions").Find(&wallets).Error; err != nil {
		return nil, err
	}
	return wallets, nil
}
