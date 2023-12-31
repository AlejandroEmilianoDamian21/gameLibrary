package models

import (
	"time"

	"gorm.io/gorm"
)

type Genero struct {
	gorm.Model
	Nombre      string `gorm:"not null;unique" json:"nombre"`
	Descripcion string `gorm:"not null" json:"descripcion"`
	ImagenURL   string `gorm:"not null;default:'default.png'" json:"imagenUrl"`
}

func (u *Genero) AfterUpdate(tx *gorm.DB) (err error) {
	u.UpdatedAt = time.Now()
	return
}
