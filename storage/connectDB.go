package storage

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/AlejandroEmilianoDamian21/listGamesGO/models"
	"github.com/AlejandroEmilianoDamian21/listGamesGO/utils"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	//Autoload the env
	// _ "github.com/joho/godotenv/autoload"
)

var DB *gorm.DB

type Config struct {
	DBHost         string `mapstructure:"POSTGRES_HOST"`
	DBUserName     string `mapstructure:"POSTGRES_USER"`
	DBUserPassword string `mapstructure:"POSTGRES_PASSWORD"`
	DBName         string `mapstructure:"POSTGRES_DB"`
	DBPort         string `mapstructure:"POSTGRES_PORT"`

	JwtSecret    string        `mapstructure:"JWT_SECRET"`
	JwtExpiresIn time.Duration `mapstructure:"JWT_EXPIRED_IN"`
	JwtMaxAge    int           `mapstructure:"JWT_MAXAGE"`

	ClientOrigin string `mapstructure:"CLIENT_ORIGIN"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

func ConnectDB() *gorm.DB {
	var (
		host     = utils.AccessENV("DB_HOST")
		user     = utils.AccessENV("DB_USER")
		port     = utils.AccessENV("DB_PORT")
		password = utils.AccessENV("DB_PASSWORD")
		name     = utils.AccessENV("DB_NAME")
	)

	if host == "" {
		log.Fatalln("Error loading ENV")
		return nil
	}

	portInt, err := strconv.Atoi(port)

	if err != nil {
		log.Fatalln("Error en convertir el port :" + err.Error())
		return nil
	}

	//Connect to DB
	var DB *gorm.DB

	DB, err = gorm.Open("postgres", fmt.Sprintf("host=%s  port=%d user=%s password=%s dbname=%s sslmode=disable", host, portInt, user, password, name))

	//Check for errors in DB
	if err != nil {
		log.Fatalf("Error in connect the DB %v", err)
		return nil
	}
	if err := DB.DB().Ping(); err != nil {
		log.Fatalln("Error in make ping the DB" + err.Error())
		return nil
	}
	if DB.Error != nil {
		log.Fatalln("Any Error in connect the DB" + err.Error())
		return nil
	}
	log.Println("🚀 Connected Successfully to the Database")
	DB.AutoMigrate(&models.Juego{})
	DB.AutoMigrate(&models.User{})
	return DB
}
