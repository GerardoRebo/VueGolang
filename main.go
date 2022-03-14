package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})
	r.GET("/no", func(c *gin.Context) {
		c.String(http.StatusOK, "VergasMames")
	})
	r.Run()
}
