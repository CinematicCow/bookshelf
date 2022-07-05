package db

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/Cinematiccow/bookshelf/models"
)

type database struct {
	host     string
	name     string
	user     string
	password string
	port     string
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: env file load error", err)
	}
}

var DB *gorm.DB

func ConnectDatabase() {
	database := database{
		host:     os.Getenv("POSTGRES_HOST"),
		name:     os.Getenv("POSTGRES_DB"),
		user:     os.Getenv("POSTGRES_USER"),
		password: os.Getenv("POSTGRES_PASSWORD"),
		port:     os.Getenv("POSTGRES_PORT"),
	}

	dsn := "host=" + database.host + " user=" + database.user + " password=" + database.password + " dbname=" + database.name + " port=" + database.port
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}
	log.Println("Connected to database")
	db.AutoMigrate(&models.Book{})

	DB = db

}
