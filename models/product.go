package models

import (
	"time"

	"digitaltrader/db"
)

type Product struct {
	ID            uint       `json:"id" gorm:"primaryKey;autoIncrement"`
	StoreID       int        `json:"store_id" gorm:"type:integer"`
	Cover         string     `json:"cover" gorm:"type:text"`
	Name          string     `json:"name" gorm:"type:text"`
	Images        string     `json:"images" gorm:"type:text"`
	OriginalPrice *float64   `json:"original_price" gorm:"type:decimal(10,2);null"`
	SellPrice     *float64   `json:"sell_price" gorm:"type:decimal(10,2);null"`
	Discount      *float64   `json:"discount" gorm:"type:decimal(10,2);null"`
	Kind          *int8      `json:"kind" gorm:"type:tinyint;null"`
	CateID        *int       `json:"cate_id" gorm:"type:integer;null"`
	SubCateID     *int       `json:"sub_cate_id" gorm:"type:integer;null"`
	InHome        *int8      `json:"in_home" gorm:"type:tinyint;null"`
	IsSingle      *int8      `json:"is_single" gorm:"type:tinyint;null"`
	HaveGram      *int8      `json:"have_gram" gorm:"type:tinyint;null"`
	Gram          *string    `json:"gram" gorm:"type:varchar(255);null"`
	HaveKg        *int8      `json:"have_kg" gorm:"type:tinyint;null"`
	Kg            *string    `json:"kg" gorm:"type:varchar(255);null"`
	HavePcs       *int8      `json:"have_pcs" gorm:"type:tinyint;null"`
	Pcs           *string    `json:"pcs" gorm:"type:varchar(255);null"`
	HaveLiter     *int8      `json:"have_liter" gorm:"type:tinyint;null"`
	Liter         *string    `json:"liter" gorm:"type:varchar(255);null"`
	HaveMl        *int8      `json:"have_ml" gorm:"type:tinyint;null"`
	Ml            *string    `json:"ml" gorm:"type:varchar(255);null"`
	Descriptions  *string    `json:"descriptions" gorm:"type:text;null"`
	KeyFeatures   *string    `json:"key_features" gorm:"type:text;null"`
	Disclaimer    *string    `json:"disclaimer" gorm:"type:text;null"`
	ExpDate       *time.Time `json:"exp_date" gorm:"type:date;null"`
	TypeOf        int8       `json:"type_of" gorm:"type:tinyint;default:2"` // 1 = veg, 0 = non-veg
	InOffer       int8       `json:"in_offer" gorm:"type:tinyint;default:2"`
	InStock       int8       `json:"in_stoke" gorm:"type:tinyint;default:0"`
	Rating        *float64   `json:"rating" gorm:"type:decimal(10,2);null"`
	TotalRating   *int       `json:"total_rating" gorm:"type:integer;null"`
	Variations    *string    `json:"variations" gorm:"type:text;null"`
	Size          *int8      `json:"size" gorm:"type:tinyint;null"`
	Status        int8       `json:"status" gorm:"type:tinyint;default:0"`
	ExtraField    *string    `json:"extra_field" gorm:"type:text;null"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`

	// Relationships
	Store       Store       `gorm:"foreignKey:StoreID;references:ID;constraint:OnDelete:RESTRICT"`
	Category    Category    `gorm:"foreignKey:CateID;references:ID;constraint:OnDelete:SET NULL"`
	SubCategory SubCategory `gorm:"foreignKey:SubCateID;references:ID;constraint:OnDelete:SET NULL"`
}

// MigrateProduct auto-migrates the Product model
func MigrateProduct() {
	db := db.InitDB()
	// Migrate Store, Category, and SubCategory first due to foreign key dependencies
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Store{}, &Category{}, &SubCategory{}, &Product{})
}

// CreateProduct creates a new product
func CreateProduct(product *Product) error {
	db := db.InitDB()
	return db.Create(product).Error
}

// GetProduct retrieves a product by ID
func GetProduct(id uint) (*Product, error) {
	db := db.InitDB()
	var product Product
	if err := db.Preload("Store").Preload("Category").Preload("SubCategory").First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

// UpdateProduct updates an existing product
func UpdateProduct(product *Product, id uint) error {
	db := db.InitDB()
	return db.Model(&Product{}).Where("id = ?", id).Updates(product).Error
}

// DeleteProduct deletes a product by ID
func DeleteProduct(id uint) error {
	db := db.InitDB()
	return db.Delete(&Product{}, id).Error
}

// GetAllProducts retrieves all products
func GetAllProducts() ([]Product, error) {
	db := db.InitDB()
	var products []Product
	if err := db.Preload("Store").Preload("Category").Preload("SubCategory").Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}
