package models

import (
	"time"

	"digitaltrader/db"
)

type Order struct {
	ID             uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	UID            int       `json:"uid" gorm:"type:integer"`
	StoreID        string    `json:"store_id" gorm:"type:varchar(255)"`
	DateTime       time.Time `json:"date_time" gorm:"type:datetime"`
	PaidMethod     string    `json:"paid_method" gorm:"type:varchar(255)"`
	OrderTo        string    `json:"order_to" gorm:"type:varchar(255)"`
	Orders         string    `json:"orders" gorm:"type:text"`
	Notes          *string   `json:"notes" gorm:"type:text;null"`
	Address        *string   `json:"address" gorm:"type:text;null"`
	DriverID       *string   `json:"driver_id" gorm:"type:varchar(255);null"`
	Assignee       *string   `json:"assignee" gorm:"type:text;null"`
	Total          *float64  `json:"total" gorm:"type:decimal(10,2);null"`
	Tax            *float64  `json:"tax" gorm:"type:decimal(10,2);null"`
	GrandTotal     *float64  `json:"grand_total" gorm:"type:decimal(10,2);null"`
	Discount       *float64  `json:"discount" gorm:"type:decimal(10,2);null"`
	DeliveryCharge *float64  `json:"delivery_charge" gorm:"type:decimal(10,2);null"`
	WalletUsed     int8      `json:"wallet_used" gorm:"type:tinyint;default:0"`
	WalletPrice    *float64  `json:"wallet_price" gorm:"type:decimal(10,2);null"`
	CouponCode     *string   `json:"coupon_code" gorm:"type:text;null"`
	Extra          *string   `json:"extra" gorm:"type:text;null"`
	PayKey         *string   `json:"pay_key" gorm:"type:text;null"`
	Status         *string   `json:"status" gorm:"type:text;null"`
	PayStatus      int8      `json:"pay_status" gorm:"type:tinyint;default:0"`
	ExtraField     *string   `json:"extra_field" gorm:"type:text;null"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`

	// Relationships
	Store  Store  `gorm:"foreignKey:StoreID;references:ID;constraint:OnDelete:RESTRICT"`
	Driver Driver `gorm:"foreignKey:DriverID;references:ID;constraint:OnDelete:SET NULL"`
}

// MigrateOrder auto-migrates the Order model
func MigrateOrder() {
	db := db.InitDB()
	// Migrate Store and Driver first due to foreign key dependencies
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Store{}, &Driver{}, &Order{})
}

// CreateOrder creates a new order
func CreateOrder(order *Order) error {
	db := db.InitDB()
	return db.Create(order).Error
}

// GetOrder retrieves an order by ID
func GetOrder(id uint) (*Order, error) {
	db := db.InitDB()
	var order Order
	if err := db.Preload("Store").Preload("Driver").First(&order, id).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

// UpdateOrder updates an existing order
func UpdateOrder(order *Order, id uint) error {
	db := db.InitDB()
	return db.Model(&Order{}).Where("id = ?", id).Updates(order).Error
}

// DeleteOrder deletes an order by ID
func DeleteOrder(id uint) error {
	db := db.InitDB()
	return db.Delete(&Order{}, id).Error
}

// GetAllOrders retrieves all orders
func GetAllOrders() ([]Order, error) {
	db := db.InitDB()
	var orders []Order
	if err := db.Preload("Store").Preload("Driver").Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}
