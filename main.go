package main

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{})
	db.Create(&Product{
		Name:  "D42",
		Price: 100.0,
	})

	products := []Product{
		{Name: "test1", Price: 1.0},
		{Name: "test2", Price: 2.0},
		{Name: "test3", Price: 3.0},
	}

	db.Create(products)

	// Read
	var product Product
	db.First(&product, 1)
	db.First(&product, "name = ?", "test1")

	// Update - update product's price to 200
	db.Model(&product).Update("Price", 200)

	// Update - update multiple fields
	db.Model(&product).Updates(Product{Price: 200, Name: "F42"}) // non-zero fields
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Name": "F42"})

	// Delete - delete product
	db.Delete(&product, 1)
}
