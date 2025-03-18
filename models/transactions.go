package models

import (
	"time"

	"digitaltrader/db"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// TransactionType represents the possible transaction types
type TransactionType string

const (
	TypeDeposit  TransactionType = "deposit"
	TypeWithdraw TransactionType = "withdraw"
)

type Transaction struct {
	ID          uint            `json:"id" gorm:"primaryKey;autoIncrement"`
	PayableID   uint            `json:"payable_id" gorm:"index:payable_type_payable_id_ind;index:payable_type_ind;index:payable_confirmed_ind;index:payable_type_confirmed_ind"`
	PayableType string          `json:"payable_type" gorm:"index:payable_type_payable_id_ind;index:payable_type_ind;index:payable_confirmed_ind;index:payable_type_confirmed_ind"`
	WalletID    uint            `json:"wallet_id"`
	Type        TransactionType `json:"type" gorm:"type:varchar(20);index;index:payable_type_ind;index:payable_type_confirmed_ind;check:type IN ('deposit','withdraw')"`
	Amount      float64         `json:"amount" gorm:"type:decimal(64,0)"`
	Confirmed   bool            `json:"confirmed" gorm:"index:payable_confirmed_ind;index:payable_type_confirmed_ind"`
	Meta        string          `json:"meta" gorm:"type:json;null"`
	UUID        string          `json:"uuid" gorm:"type:uuid;unique"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
}

// BeforeCreate hook to generate UUID
func (t *Transaction) BeforeCreate(tx *gorm.DB) (err error) {
	if t.UUID == "" {
		t.UUID = uuid.New().String()
	}
	return
}

// MigrateTransaction auto-migrates the Transaction model with indexes
func MigrateTransaction() {
	db := db.InitDB()
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Transaction{})

	// Ensure indexes are created (some databases might need this explicitly)
	db.Exec(`CREATE INDEX payable_type_payable_id_ind ON transactions (payable_type, payable_id)`)
	db.Exec(`CREATE INDEX payable_type_ind ON transactions (payable_type, payable_id, type)`)
	db.Exec(`CREATE INDEX payable_confirmed_ind ON transactions (payable_type, payable_id, confirmed)`)
	db.Exec(`CREATE INDEX payable_type_confirmed_ind ON transactions (payable_type, payable_id, type, confirmed)`)
}

// CreateTransaction creates a new transaction
func CreateTransaction(transaction *Transaction) error {
	db := db.InitDB()
	return db.Create(transaction).Error
}

// GetTransaction retrieves a transaction by ID
func GetTransaction(id uint) (*Transaction, error) {
	db := db.InitDB()
	var transaction Transaction
	if err := db.First(&transaction, id).Error; err != nil {
		return nil, err
	}
	return &transaction, nil
}

// UpdateTransaction updates an existing transaction
func UpdateTransaction(transaction *Transaction, id uint) error {
	db := db.InitDB()
	return db.Model(&Transaction{}).Where("id = ?", id).Updates(transaction).Error
}

// DeleteTransaction deletes a transaction by ID
func DeleteTransaction(id uint) error {
	db := db.InitDB()
	return db.Delete(&Transaction{}, id).Error
}

// GetAllTransactions retrieves all transactions
func GetAllTransactions() ([]Transaction, error) {
	db := db.InitDB()
	var transactions []Transaction
	if err := db.Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}
