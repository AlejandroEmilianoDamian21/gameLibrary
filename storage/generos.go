package storage

import (
	"log"

	"github.com/AlejandroEmilianoDamian21/listGamesGO/initializers"
	"github.com/AlejandroEmilianoDamian21/listGamesGO/models"
	"gorm.io/gorm"
)

type GeneroDB struct {
	db *gorm.DB
}

// NuevoGenero crea una nueva instancia del servicio GeneroDB.
// Retorna un objeto de tipo GeneroDB que se puede utilizar para interactuar con la base de datos de géneros.
func NuevoGeneroDB() GeneroDB {
	// Cargar la configuración del archivo .env
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("Failed to load config")
	}
	// Conectarse a la base de datos utilizando la configuración cargada
	nuevaBD := initializers.ConnectDB(&config)
	if nuevaBD == nil {
		log.Fatal("Failed to connect to the database")
	}

	// Crear una nueva instancia del servicio GeneroDB con la base de datos conectada
	nuevoServicio := GeneroDB{db: initializers.DB}

	// Retornar la instancia del servicio GeneroDB
	return nuevoServicio
}

/*ObtenerGeneros Obtener todos los generos*/
func (g *GeneroDB) ObtenerGeneros() ([]*models.Genero, error) {
	var genero []*models.Genero = []*models.Genero{new(models.Genero)}

	/*SELECT * FROM generos */
	if err := g.db.Find(&genero).Error; err != nil {
		return nil, err
	}
	return genero, nil
}
