package model

type Product struct {
	Base
	Name           string           `json:"name" gorm:"type:varchar(100);not null"`
	CategoryID     uint             `json:"category_id" gorm:"not null"`
	Category       Category         `json:"category" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Price          float32          `json:"price" gorm:"not null"`
	Photo          string           `json:"photo" gorm:"type:varchar(255)"`
	Description    string           `json:"description" gorm:"type:text"`
	Discount       float32          `json:"discount" gorm:"default:0"`
	ProductVariant []ProductVariant `json:"product_variants" gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type ProductVariant struct {
	Base
	ProductID uint    `json:"product_id" gorm:"not null;index"`
	VariantID uint    `json:"variant_id" gorm:"not null;index"`
	Product   Product `json:"product" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Variant   Variant `json:"variant" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
