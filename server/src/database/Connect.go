package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"h-vistars/server/src/models"
)

var DB *gorm.DB

func ConnectDatabase() {

	database, err := gorm.Open(sqlite.Open("h-vistars.db"), &gorm.Config{})

	if err != nil {
		panic("Ha ocurrido un error al conectarse a la base de datos")
	}

	err = database.AutoMigrate(&models.Beat{}, &models.User{}, &models.Role{}, &models.Country{}, &models.Genre{}, &models.License{}, &models.LicenseType{})
	if err != nil {
		return
	}

	DB = database
}
