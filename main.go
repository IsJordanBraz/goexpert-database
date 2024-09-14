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
	Products []Product `gorm:"many2many:products_categories;"`
}

type Product struct {
	gorm.Model
	Name       string
	Price      float64
	Categories []Category `gorm:"many2many:products_categories;"`
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{}, &Category{})

	newCategory1 := Category{Name: "Banho"}
	db.Create(&newCategory1)

	newCategory2 := Category{Name: "Eletro"}
	db.Create(&newCategory2)

	newProduct := Product{
		Name:       "Shampo",
		Price:      100.0,
		Categories: []Category{newCategory1, newCategory2},
	}
	db.Create(&newProduct)

	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Find(&categories).Error
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		fmt.Println(category.Name)
		for _, product := range category.Products {
			fmt.Println("-", product.Name)
		}
	}
}
