package routes

import (
	"github.com/ArtiomStartev/Go-Cache-API/controller"
	"github.com/ArtiomStartev/Go-Cache-API/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/patrickmn/go-cache"
)

func Setup(app *fiber.App, cache *cache.Cache) {
	app.Get("/posts/:id", middleware.CacheMiddleware(cache), controller.GetPost)
}
