package database

import (
	"fmt"
	"log"
	"strconv"

	"github.com/phawitpp/my-shortlink-api/config"
	"github.com/phawitpp/my-shortlink-api/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	p := config.LoadEnv("DB_PORT")
	port, err := strconv.ParseInt(p, 10, 32)

	if err != nil {
		log.Println("Error parsing port")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Bangkok", config.LoadEnv("DB_HOST"), config.LoadEnv("DB_USER"), config.LoadEnv("DB_PASSWORD"), config.LoadEnv("DB_NAME"), port)
	DB, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Database connected")

	DB.AutoMigrate(&model.LinkModel{})
	fmt.Println("Database migrated")
}
