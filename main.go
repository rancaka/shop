package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()

	f, err := os.Create("/var/log/shop/shop.gin.log")
	if err != nil {
		log.Fatalln(err)
	}
	gin.DefaultWriter = io.MultiWriter(f)

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		gin.Logger()
		c.JSON(http.StatusOK, gin.H{
			"App": "Shop",
		})
	})
	r.Run()
}
