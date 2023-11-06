package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

var db *gorm.DB

func main() {
	var err error
	db, err = gorm.Open(sqlite.Open("test.db"))
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Product{})

	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})
	r.GET("/create", create)
	r.GET("/first", first)
	r.GET("/last", last)
	r.GET("/update", update)
	r.GET("/del", del)

	r.Run(":8080")
}

func create(c *gin.Context) {
	db.Create(&Product{Code: "D42", Price: 100})
	var product Product
	db.Last(&product)
	c.IndentedJSON(http.StatusOK, &product)
}

func first(c *gin.Context) {
	var product Product
	db.First(&product)
	c.IndentedJSON(http.StatusOK, &product)
}

func last(c *gin.Context) {
	var product Product
	db.Last(&product)
	c.IndentedJSON(http.StatusOK, &product)
}

func update(c *gin.Context) {
	var product Product
	db.First(&product)
	db.Model(&product).Update("Price", 200)
	db.Model(&product).Updates(Product{Price: 999, Code: "F42"})
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
	c.IndentedJSON(http.StatusOK, &product)
}
func del(c *gin.Context) {
	var product Product
	db.First(&product)
	db.Delete(&product)
	c.IndentedJSON(http.StatusOK, &product)
}
