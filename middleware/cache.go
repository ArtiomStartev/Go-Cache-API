package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/ArtiomStartev/Go-Cache-API/structs"
	"github.com/gofiber/fiber/v2"
	"github.com/patrickmn/go-cache"
	"net/http"
	"time"
)

func CacheMiddleware(cache *cache.Cache) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if c.Method() != http.MethodGet {
			// Only cache GET requests
			return c.Next()
		}

		// Generate a cache key based on the request path and query parameters
		cacheKey := c.Path() + "?" + c.Params("id")

		if data, found := cache.Get(cacheKey); found {
			c.Response().Header.Set("X-Cache-Status", "HIT")
			return c.Status(http.StatusOK).JSON(data)
		}

		c.Response().Header.Set("X-Cache-Status", "MISS")
		if err := c.Next(); err != nil {
			return err
		}

		var data structs.Post
		body := c.Response().Body()

		if err := json.Unmarshal(body, &data); err != nil {
			fmt.Println("Error unmarshalling response body: ", err)
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": "Internal Server Error",
			})
		}

		// Cache the response for 10 minutes
		cache.Set(cacheKey, data, time.Minute*10)

		return nil
	}
}
