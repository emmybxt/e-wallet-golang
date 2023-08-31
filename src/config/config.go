package config

import (
	"e-wallet/src/models"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var db *gorm.DB

func init() {
	err := loadEnv()

	if err != nil {
		log.Fatal(err)
	}
	postgresConnection := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s Timezone=Africa/Lagos",

		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSL_MODE"),
	)

	db, err := gorm.Open(postgres.Open(postgresConnection), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&models.User{}, &models.Wallet{})

	if err != nil {
		log.Fatal(err)
	}

}

func GetConnection() *gorm.DB {
	return db
}

func loadEnv() error {
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}
	return nil
}