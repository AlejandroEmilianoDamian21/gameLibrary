package models

import (
	"time"

	"gorm.io/gorm"
)

/*Game Struct*/

type Juego struct {
	gorm.Model
	/*// gorm.Model definition
		type Model struct {
	  	ID        uint           `gorm:"primaryKey"`
	  	CreatedAt time.Time
	  	UpdatedAt time.Time
	  	DeletedAt gorm.DeletedAt `gorm:"index"`
		}*/
	Nombre        string `gorm:"not null" json:"nombre"`
	Desarrollador string `gorm:"not null" json:"desarrollador"`
	Precio        int    `gorm:"not null" json:"precio"`
}

func (u *Juego) AfterUpdate(tx *gorm.DB) (err error) {
	u.UpdatedAt = time.Now()
	return
}
