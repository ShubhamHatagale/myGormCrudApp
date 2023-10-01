// model/users.go

package model

import (
	"github.com/jinzhu/gorm"
)

type Users struct {
	gorm.Model
	Name     string  `json:"name"`
	Surname  string  `json:"surname"`
	Gender   string  `json:"Gender"`
	Company  string  `json:"company"`
	Password float64 `json:"password"`
}
