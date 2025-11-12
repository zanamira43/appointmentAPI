package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/zanamira43/appointment-api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Unable to load .env file %v", err)
	}
}

func Connect() error {

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbname)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	fmt.Println("database Successfully Connected")

	DB = db

	db.AutoMigrate(
		&models.User{},
		&models.Patient{},
		&models.TimeTable{},
		&models.Problem{},
		&models.Session{},
		&models.PaymentType{},
		&models.Payment{},
		&models.Settings{},
		&models.NoteBook{},
	)

	return nil
}
