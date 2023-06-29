package models

import (
	"time"

	"gorm.io/gorm"
)

type Genero struct {
	gorm.Model
	Name        string `gorm:"type:varchar(100); unique"`
	Description string
	ImageURL    string
}

func (u *Genero) AfterUpdate(tx *gorm.DB) (err error) {
	u.UpdatedAt = time.Now()
	return
}
