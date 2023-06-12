package models

import (
	"github.com/jinzhu/gorm"
)

/*Game Struct*/

type game struct {
	gorm.Model
	Nombre        string `gorm:"not null" json:"nombre"`
	Desarrollador string `gorm:"not null" json:"desarrollador"`
	Precio        int    `gorm:"not null" json:"precio"`
}
