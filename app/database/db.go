package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {
	con, err := gorm.Open(mysql.Open(BuildDSN()), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
		return err
	}

	DB = con
	return nil
}
