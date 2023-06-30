package storage

// import "fmt"

// func hi() {
// 	fmt.Println("Hola")
// }

import (
	"log"

	"github.com/AlejandroEmilianoDamian21/listGamesGO/initializers"
	"github.com/AlejandroEmilianoDamian21/listGamesGO/models"
	"gorm.io/gorm"
)

/*JuegoDB struct*/
type JuegoDB struct {
	db *gorm.DB
}

/*NuevoJuegoDB Create a new storage user service*/
func NuevoJuegoDB() JuegoDB {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("Failed to load config")
		// Manejar el error en caso de que ocurra al cargar la configuración
	}
	nuevaDB := initializers.ConnectDB(&config)
	if nuevaDB == nil {
		log.Fatal("Failed to connect to the database")
	}

	nuevoServicio := JuegoDB{db: initializers.DB}

	return nuevoServicio
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

/*Obtener Juego del juego con el ID*/
func (j *JuegoDB) ObtenerJuego(id int) (*models.Juego, error) {
	var juego *models.Juego = new(models.Juego)
	// SELECT * FROM juegos WHERE id = 10;
	if err := j.db.First(&juego, id).Error; err != nil {
		return nil, err
	}
	return juego, nil
}

/*Crear un juego*/
func (j *JuegoDB) CrearJuego(nuevoJuego *models.Juego) (*models.Juego, bool, error) {
	if err := j.db.Create(&nuevoJuego).Error; err != nil {
		return nil, false, err
	}
	return nuevoJuego, true, nil
}

/*ModificarJuego Modifica el juego*/
func (j *JuegoDB) ModificarJuego(nuevoJuego *models.Juego) (bool, error) {
	if err := j.db.Model(&models.Juego{}).Updates(&nuevoJuego).Error; err != nil {
		return false, err
	}
	return true, nil
}

/*EliminarJuego Elimina un Juego*/
func (j *JuegoDB) EliminarJuego(id int) (bool, error) {
	juego := &models.Juego{}

	if err := j.db.Delete(juego, id).Error; err != nil {
		return false, err
	}
	return true, nil
}
