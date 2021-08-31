package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ndaemy/gin-tutorial/routes"
)

func main() {
	router := gin.Default()
	routes.AddRoutes(router)

	router.Run(":5000")
}
