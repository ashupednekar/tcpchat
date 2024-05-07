package chat

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDb() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Kolkata",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Could not connect to database: ", err)
	}
	return db
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&Users{}, &Groups{}, &Messages{})
}

type Users struct {
	ID     int    `gorm:"primaryKey"`
	Name   string `gorm:"uniqueIndex"`
	IP     string `gorm:"uniqueIndex"`
	Online bool
}

type Groups struct {
	ID    int    `gorm:"primaryKey"`
	Name  string `gorm:"uniqueIndex"`
	Users []int  `gorm:"type:json"`
}

type Messages struct {
	ID     int `gorm:"primaryKey"`
	FromID int
	From   Users `gorm:"constraint:OnDelete:CASCADE;"`
	ToID   int
	To     Groups `gorm:"constraint:OnDelete:CASCADE;"`
	Text   sql.NullString
	File   sql.NullByte
}
