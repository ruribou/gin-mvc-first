package main

import (
	"tutorial/controller"
)

func main() {
	router := controller.GetRouter()
	router.Run(":8080")
}