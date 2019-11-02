package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	templatesPath = "templates/"
)

func main() {
	router := initRouter()

	router.Run()
}

func initRouter() (router *gin.Engine) {
	router = gin.Default()
	router.LoadHTMLGlob(templatesPath + "/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(
			http.StatusOK,
			"index.tmpl",
			gin.H{
				"title": "Hello!",
			},
		)
	})

	return
}
