package models

import "gorm.io/gorm"

type Beat struct {
	gorm.Model
	Title             string `json:"title"`
	Author            string `json:"author" gorm:"foreignKey:Author"`
	Year              uint   `json:"year"`
	GenreID           uint   `json:"genreId" gorm:"foreignKey:GenreID"`
	Bpm               uint   `json:"bpm"`
	Key               string `json:"key"`
	DurationInSeconds uint   `json:"durationInSeconds"`
	Price             uint   `json:"price"`
	ThumbnailURL      string `json:"thumbnailUrl"`
}

type User struct {
	gorm.Model
	Username        string `json:"username"`
	AvatarURL       string `json:"avatarUrl"`
	Name            string `json:"name"`
	LastName        string `json:"lastName"`
	Email           string `json:"email" gorm:"unique"`
	Password        string `json:"password"`
	PhoneNumber     string `json:"phoneNumber" gorm:"unique"`
	Country         string `json:"country" gorm:"foreignKey:Country"`
	Beats           []Beat `json:"beats" gorm:"foreignKey:Author"`
	IsEmailVerified bool   `json:"isEmailVerified" gorm:"default:false"`
	Role            string `json:"role" gorm:"foreignKey:Role; default:standard"`
}

type Role struct {
	gorm.Model
	Name string `json:"name" gorm:"unique"`
}

type Country struct {
	gorm.Model
	Name string `json:"name" gorm:"unique"`
}

type Genre struct {
	gorm.Model
	Type string `json:"type" gorm:"unique"`
}

type License struct {
	gorm.Model
	BeatID uint   `json:"beatId" gorm:"foreignKey:BeatID"`
	UserID uint   `json:"userId" gorm:"foreignKey:UserID"`
	Type   string `json:"type" gorm:"foreignKey:Type"`
}

type LicenseType struct {
	gorm.Model
	Name string `json:"name"`
}
