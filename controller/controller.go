package controller

import (
	"net/http"

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
	repo, err := models.NewDefaultDbGreetingRepo()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	greetings, err := repo.GetAll()
	if err != nil {
		if err != models.ErrRecordNotFound {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		err = nil // simply no one has said 'hello' yet T_T
	}

	c.HTML(
		http.StatusOK,
		"index.tmpl",
		gin.H{
			"title":     "Hello!",
			"greetings": greetings,
		},
	)
}

func acceptHello(c *gin.Context) {
	var greeting models.Greeting

	c.Bind(&greeting)

	if err := storeGreeting(greeting); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	redirectToHello(c)
}

func storeGreeting(g models.Greeting) error {
	repo, err := models.NewDefaultDbGreetingRepo()
	if err != nil {
		return err
	}

	return repo.Store(g)
}
