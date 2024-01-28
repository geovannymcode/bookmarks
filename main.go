package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Geovanny0401/bookmarks/internal/config"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.GetConfig("config.json")
	if err != nil {
		log.Fatal(err)
	}
	r := gin.Default()
	r.GET("/hello", hello)
	log.Fatal(r.Run(fmt.Sprintf(":%d", cfg.ServerPort)))
}

func hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
	})
}
