package database

import (
	"fmt"
	models "kotaku/Models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost)/kotaku"), &gorm.Config{})

	if err != nil {
		panic("This DB Not Connected")
	} else {
		fmt.Println("Connected")
	}

	DB = db

	db.AutoMigrate(models.User{})
	db.AutoMigrate(models.InfoKota{})
	db.AutoMigrate(models.Voucher{})

}
