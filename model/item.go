// model/item.go

package model

import (
	"github.com/jinzhu/gorm"
)

type Item struct {
	gorm.Model
	Name string `json:"name"`
	// Surname string  `json:"surname"`
	Price float64 `json:"price"`
}
