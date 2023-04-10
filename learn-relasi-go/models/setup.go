package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DB() (*gorm.DB, error){
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/one-to-many"))
	if err != nil{
		return nil, err
	}

	db.AutoMigrate(&Product{}, &Category{})
	return db, nil
}