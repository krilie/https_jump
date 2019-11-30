package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func main() {
	// 跳转的地址
	toUrl := os.Getenv("JUMP_TO")
	if toUrl == "" {
		toUrl = "https://localhost"
	}
	engine := gin.New()
	engine.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, toUrl)
	})
	err := engine.Run(":80")
	log.Println(err)
}
