package config

import (
	"fmt"
	"go-api-jwt/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// configuration for connecting to database

func ConnectDatabase() *gorm.DB {
	username := "fauzil"
	password := "password"
	database := "database_movie"
	host := "tcp(127.0.0.1:3306)" // tcp(localhost:port)

	dsn := fmt.Sprintf("%v:%v@%v/%v?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error)
	}

	db.AutoMigrate(&models.Movie{}, &models.AgeRatingCategory{}, &models.User{})

	return db
}
