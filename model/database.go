package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func Database() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	if err = db.AutoMigrate(&Asset{}); err != nil {
		log.Println(err)
	}

	return db, err
}
