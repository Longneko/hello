package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"

	"golang.org/x/sync/errgroup"

	"github.com/Longneko/lamp/app/controller"
	"github.com/Longneko/lamp/app/lib/config"
	"github.com/Longneko/lamp/app/lib/database"
	"github.com/Longneko/lamp/app/models"
)

var (
	g errgroup.Group
)

func main() {
	// Init Config
	if err := config.InitConfig(); err != nil {
		panic(err)
	}

	// Init DB
	if err := database.InitDb(); err != nil {
		panic(err)
	}

	// Init DB Schema
	// TODO: move outside main.go and make configurable
	greetingsRepo, err := models.NewDefaultDbGreetingRepo()
	if err != nil {
		panic(err)
	}
	if err := greetingsRepo.CreateTable(); err != nil {
		panic(err)
	}

	// init server
	startServer(&g)
	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}


func startServer(g *errgroup.Group) {
	cfg := config.Get()

	gin.SetMode(cfg.Application.Mode)

	router := controller.NewDefaultRouter()
	server := &http.Server{
		// TODO: add config for server
		Addr:         cfg.Server.Address(),
		Handler:      router,
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
	}
	g.Go(func() error {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
		return err
	})
}
