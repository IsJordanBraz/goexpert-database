package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       uint `gorm:"primarykey"`
	Name     string
	Products []Product
}

type Product struct {
	gorm.Model
	Name       string
	Price      float64
	CategoryID uint
	Category   Category
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{}, &Category{})

	newCategory := Category{Name: "Banho E Tosa"}
	db.Create(&newCategory)

	newProduct := Product{
		Name:       "Shampo",
		Price:      100.0,
		CategoryID: newCategory.ID,
	}
	db.Create(&newProduct)

	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Find(&categories).Error
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		for _, product := range category.Products {
			fmt.Println(product.Name, category.Name)
		}
	}
}
