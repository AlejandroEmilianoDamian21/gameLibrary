package initializers

import (
	"fmt"
	"log"
	"os"

	"github.com/AlejandroEmilianoDamian21/listGamesGO/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	//Autoload the env
	// _ "github.com/joho/godotenv/autoload"
)

var DB *gorm.DB

func ConnectDB(config *Config) {
	var err error

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", config.DBHost, config.DBUserName, config.DBUserPassword, config.DBName, config.DBPort)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the Database! \n", err.Error())
		os.Exit(1)
	}

	DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	DB.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Running Migrations")
	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Migration Failed:  \n", err.Error())
		os.Exit(1)
	}

	log.Println("🚀 Connected Successfully to the Database")

}

// func ConnectDB() *gorm.DB {
// 	var (
// 		host     = utils.AccessENV("DB_HOST")
// 		user     = utils.AccessENV("DB_USER")
// 		port     = utils.AccessENV("DB_PORT")
// 		password = utils.AccessENV("DB_PASSWORD")
// 		name     = utils.AccessENV("DB_NAME")
// 	)

// 	if host == "" {
// 		log.Fatalln("Error loading ENV")
// 		return nil
// 	}

// 	portInt, err := strconv.Atoi(port)

// 	if err != nil {
// 		log.Fatalln("Error en convertir el port :" + err.Error())
// 		return nil
// 	}

// 	//Connect to DB
// 	var DB *gorm.DB

// 	DB, err = gorm.Open("postgres", fmt.Sprintf("host=%s  port=%d user=%s password=%s dbname=%s sslmode=disable", host, portInt, user, password, name))

// 	//Check for errors in DB
// 	if err != nil {
// 		log.Fatalf("Error in connect the DB %v", err)
// 		return nil
// 	}
// 	if err := DB.DB().Ping(); err != nil {
// 		log.Fatalln("Error in make ping the DB" + err.Error())
// 		return nil
// 	}
// 	if DB.Error != nil {
// 		log.Fatalln("Any Error in connect the DB" + err.Error())
// 		return nil
// 	}
// 	log.Println("🚀 Connected Successfully to the Database")
// 	DB.AutoMigrate(&models.Juego{})
// 	DB.AutoMigrate(&models.User{})
// 	return DB
// }
