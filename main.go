package main

import (
	"log"
	"net/http"
	"time"

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
	server := &http.Server{
		// TODO: add config for server
		Addr:         ":8080",
		Handler:      controller.NewDefaultRouter(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	g.Go(func() error {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
		return err
	})
	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
