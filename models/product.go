package models

import "gorm.io/gorm"

type ProductPayload struct {
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Price       float64 `json:"price" validate:"required,gt=0"`
	Link        string  `json:"link" validate:"required"`
	ImageData   string  `json:"image_data" validate:"required"` // Base64 string dengan prefix data URI
}

type Product struct {
	gorm.Model
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Link        string  `json:"link"`
	ImageURL    string  `json:"imageUrl"` // Hanya menyimpan URL/path gambar
}
