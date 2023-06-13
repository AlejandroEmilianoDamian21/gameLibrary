package storage

import (
	"github.com/AlejandroEmilianoDamian21/listGamesGO/models"

	"github.com/jinzhu/gorm"
)

/*JuegoDB struct*/
type JuegoDB struct {
	db *gorm.DB
}

/*NuevoJuegoDB Create a new storage user service*/
func NuevoJuegoDB() JuegoDB {
	nuevaDB := ConnectDB()
	nuevoServicio := JuegoDB{db: nuevaDB}

	return nuevoServicio
}

/*Obtener Juego del juego con el ID*/
func (j *JuegoDB) ObtenerJuego(id int) (*models.Juego, error) {
	var juego *models.Juego = new(models.Juego)
	// SELECT * FROM juegos WHERE id = 10;
	if err := j.db.First(&juego, id).Error; err != nil {
		return nil, err
	}
	return juego, nil
}

/*ObtenerJuegos  Obtener todos  los  juegos */
func (j *JuegoDB) ObtenerJuegos() ([]*models.Juego, error) {
	var juego []*models.Juego = []*models.Juego{new(models.Juego)}
	/*Select * from juegos*/

	if err := j.db.Find(&juego).Error; err != nil {
		return nil, err
	}
	return juego, nil
}
