package models

type Product struct{
	Id int `gorm:"primaryKey" json:"id"`
	NameProduct string `gorm:"varchar(300)" json:"name_product"`
	Description string `gorm:"varchar(300)" json:"description"`
	Price int `gorm:"int(200)" json:"price"`
	CategoryId int `gorm:"foreignKey:CategoryId" json:"category_id"`
	Category Category 
}

type Category struct{
	Id int `gorm:"primaryKey" json:"id"`
	NameCategory string `gorm:"varchar(300)" json:"name_category"`
	Description string `gorm:"varchar(300)" json:"description"`
}