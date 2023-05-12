package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"h-vistars/server/src/models"
)

var DB *gorm.DB

func Connect() {
	dataSourceName := "root:2ZII@92p2vw%^KOMVN6t@tcp(localhost:3306)/hVistars?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dataSourceName))

	if err != nil {
		panic("Ha ocurrido un error al conectarse a la base de datos")
	}

	err = db.AutoMigrate(&models.Role{}, &models.Country{}, &models.Genre{}, &models.LicenseType{}, &models.User{}, &models.Beat{}, &models.License{})
	if err != nil {
		return
	}

	DB = db
}
