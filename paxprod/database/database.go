package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/nkolentcev/paxcontrol/paxprod/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBInstance struct {
	DB *gorm.DB
}

var DB DBInstance

func ConnectDB() {
	dsn := fmt.Sprintf("host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Europe/Moscow",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("failed connect postgres. \n", err)
		os.Exit(2)
	}

	log.Printf("db connected", time.Now().Format("2006-01-02 15:04:05"))
	db.Logger = logger.Default.LogMode(logger.Info)
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.BoardingPass{})

	DB = DBInstance{
		DB: db,
	}

}
