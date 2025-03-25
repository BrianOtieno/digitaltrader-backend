package models

import (
	"time"

	"digitaltrader/db"
)

type Complaint struct {
	ID           uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	UID          int       `json:"uid" gorm:"type:integer"`
	OrderID      int       `json:"order_id" gorm:"type:integer"`
	IssueWith    int8      `json:"issue_with" gorm:"type:tinyint"`
	DriverID     *int      `json:"driver_id" gorm:"type:integer;null"`
	StoreID      *int      `json:"store_id" gorm:"type:integer;null"`
	ProductID    *int      `json:"product_id" gorm:"type:integer;null"`
	ReasonID     *int      `json:"reason_id" gorm:"type:integer;null"`
	Title        *string   `json:"title" gorm:"type:text;null"`
	ShortMessage *string   `json:"short_message" gorm:"type:text;null"`
	Images       *string   `json:"images" gorm:"type:text;null"`
	ExtraField   *string   `json:"extra_field" gorm:"type:text;null"`
	Status       int8      `json:"status" gorm:"type:tinyint"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`

	// Relationships
	User    User    `gorm:"foreignKey:UID;references:ID;constraint:OnDelete:RESTRICT"`
	Order   Order   `gorm:"foreignKey:OrderID;references:ID;constraint:OnDelete:RESTRICT"`
	Driver  Driver  `gorm:"foreignKey:DriverID;references:ID;constraint:OnDelete:SET NULL"`
	Store   Store   `gorm:"foreignKey:StoreID;references:ID;constraint:OnDelete:SET NULL"`
	Product Product `gorm:"foreignKey:ProductID;references:ID;constraint:OnDelete:SET NULL"`
	Reason  Reason  `gorm:"foreignKey:ReasonID;references:ID;constraint:OnDelete:SET NULL"`
}

// Assuming a Reason model for reason_id (define if not already present)
type Reason struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string    `json:"name" gorm:"type:varchar(255)"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// MigrateComplaint auto-migrates the Complaint model
func MigrateComplaint() {
	db := db.InitDB()
	// Migrate dependent models first due to foreign key dependencies
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{}, &Order{}, &Driver{}, &Store{}, &Product{}, &Reason{}, &Complaint{})
}

// CreateComplaint creates a new complaint record
func CreateComplaint(complaint *Complaint) error {
	db := db.InitDB()
	return db.Create(complaint).Error
}

// GetComplaint retrieves a complaint record by ID
func GetComplaint(id uint) (*Complaint, error) {
	db := db.InitDB()
	var complaint Complaint
	if err := db.Preload("User").Preload("Order").Preload("Driver").Preload("Store").Preload("Product").Preload("Reason").First(&complaint, id).Error; err != nil {
		return nil, err
	}
	return &complaint, nil
}

// UpdateComplaint updates an existing complaint record
func UpdateComplaint(complaint *Complaint, id uint) error {
	db := db.InitDB()
	return db.Model(&Complaint{}).Where("id = ?", id).Updates(complaint).Error
}

// DeleteComplaint deletes a complaint record by ID
func DeleteComplaint(id uint) error {
	db := db.InitDB()
	return db.Delete(&Complaint{}, id).Error
}

// GetAllComplaints retrieves all complaint records
func GetAllComplaints() ([]Complaint, error) {
	db := db.InitDB()
	var complaints []Complaint
	if err := db.Preload("User").Preload("Order").Preload("Driver").Preload("Store").Preload("Product").Preload("Reason").Find(&complaints).Error; err != nil {
		return nil, err
	}
	return complaints, nil
}
