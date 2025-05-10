package models

import (
	"gorm.io/gorm"
)

// type Cart struct {
// 	gorm.Model
// 	Items []struct {
// 		Product  Product `json:"product"`
// 		Quantity int     `json:"quantity"`
// 	} `json:"items"`
// }

type CartItem struct {
	gorm.Model
	CartID    uint
	ProductID uint
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Quantity  int     `json:"quantity"`
}

type Cart struct {
	gorm.Model
	Items []CartItem `json:"items" gorm:"foreignKey:CartID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
