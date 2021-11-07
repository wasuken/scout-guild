package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("./front/dist/*")
	r.Static("/assets", "./front/dist/")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	r.POST("/api/logs", func(c *gin.Context) {
	})
	r.Run()
}
