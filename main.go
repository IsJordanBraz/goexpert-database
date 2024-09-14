package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID   uint `gorm:"primarykey"`
	Name string
}

type SerialNumber struct {
	ID        uint `gorm:"primarykey"`
	Number    string
	ProductID uint
}

type Product struct {
	gorm.Model
	Name         string
	Price        float64
	CategoryID   uint
	Category     Category
	SerialNumber SerialNumber
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	category := Category{Name: "Banho E Tosa"}
	db.Create(&category)

	product := Product{
		Name:       "Shampo",
		Price:      100.0,
		CategoryID: category.ID,
	}
	db.Create(&product)

	db.Create(&SerialNumber{
		Number:    "12345",
		ProductID: product.ID,
	})

	var products []Product
	db.Preload("Category").Preload("SerialNumber").Find(&products)
	for _, product := range products {
		fmt.Println(product.SerialNumber.Number)
	}
	fmt.Printf("TOTAL: %v\n", len(products))
}
