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

	category := Category{Name: "Banho"}
	db.Create(&category)

	db.Create(&Product{
		Name:       "D42",
		Price:      100.0,
		CategoryID: category.ID,
	})

	// products := []Product{
	// 	{Name: "test1", Price: 1.0},
	// 	{Name: "test2", Price: 2.0},
	// 	{Name: "test3", Price: 3.0},
	// }

	// db.Create(products)

	// var product Product
	// db.First(&product, 1)
	// db.First(&product, "name = ?", "test1")

	// db.Model(&product).Update("Price", 200)

	// db.Model(&product).Updates(Product{Price: 200, Name: "F42"})
	// db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Name": "F42"})

	// db.Delete(&product, 1)

	var products []Product
	// db.Limit(2).Offset(2).Find(&products)
	db.Preload("Category").Find(&products)
	for _, product := range products {
		fmt.Println(product.Category.Name)
	}
	fmt.Printf("TOTAL: %v\n", len(products))
}
