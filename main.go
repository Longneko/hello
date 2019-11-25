package main

import (
	"fmt"
	"sync"

	"github.com/Longneko/lamp/app/controller"
	"github.com/Longneko/lamp/app/lib/config"
	"github.com/Longneko/lamp/app/lib/database"
	"github.com/Longneko/lamp/app/models"
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

	// Init Router
	router := controller.NewDefaultRouter()

	wg := new(sync.WaitGroup)

	wg.Add(1)
	go func() {
		if err := router.Run(); err != nil {
			fmt.Printf("Error while runnig router: %s\n", err)
		}

		wg.Done()
	}()

	wg.Wait()
}
