package posts

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func list(c *gin.Context) {
	c.JSON(http.StatusOK, "list")
}

func write(c *gin.Context) {
	c.JSON(http.StatusOK, "write")
}

func read(c *gin.Context) {
	c.JSON(http.StatusOK, "read")
}

func update(c *gin.Context) {
	c.JSON(http.StatusOK, "update")
}

func remove(c *gin.Context) {
	c.JSON(http.StatusOK, "remove")
}
