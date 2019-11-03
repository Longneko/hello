package main

import (
	"github.com/Longneko/lamp/controller"
)

func main() {
	router := controller.NewDefaultRouter()

	router.Run()
}
