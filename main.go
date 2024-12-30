package main

import (
	"fmt"
	"github.com/ArtiomStartev/Go-Cache-API/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/patrickmn/go-cache"
	"time"
)

func main() {
	app := fiber.New()
	cache := cache.New(time.Minute*10, time.Minute*20)

	routes.Setup(app, cache)

	if err := app.Listen(":8080"); err != nil {
		fmt.Println("Error starting server on port 8080: ", err)
		return
	}
}
