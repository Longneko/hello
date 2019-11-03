package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/Longneko/lamp/models"
)

const (
	templatesPath = "templates/"
)

func NewDefaultRouter() (router *gin.Engine) {
	router = gin.Default()

	router.LoadHTMLGlob(templatesPath + "/*")

	router.GET("/", redirectToHello)
	router.GET("/index", redirectToHello)

	router.GET("hello", hello)
	router.POST("hello", acceptHello)

	return
}

func redirectToHello(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/hello")
}

func hello(c *gin.Context) {
	repo, err := models.GetDefaultGreetingRepo()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.HTML(
		http.StatusOK,
		"index.tmpl",
		gin.H{
			"title":     "Hello!",
			"greetings": repo.GetSorted(true),
		},
	)
}

func acceptHello(c *gin.Context) {
	var greeting models.Greeting

	c.Bind(&greeting)
	greeting.Time = time.Now().UTC()

	if err := storeGreeting(greeting); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	redirectToHello(c)
}

func storeGreeting(g models.Greeting) error {
	repo, err := models.GetDefaultGreetingRepo()
	if err != nil {
		return err
	}
	repo.Store(g)
	return nil
}
