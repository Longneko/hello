package main

import (
	"fmt"
	"sync"

	"github.com/Longneko/lamp/controller"
	"github.com/Longneko/lamp/models"
)

func main() {
	// init repo
	models.DefaultGreetingRepo = models.NewGreetingRepository()

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
