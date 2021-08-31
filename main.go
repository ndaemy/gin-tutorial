package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ndaemy/gin-tutorial/models"
	"github.com/ndaemy/gin-tutorial/routes"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	app := gin.Default()

	// set db connection
	db, err := gorm.Open(sqlite.Open("db.sqlite"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.Post{})
	app.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	// add routers
	routes.AddRoutes(app)

	app.Run(":5000")
}
