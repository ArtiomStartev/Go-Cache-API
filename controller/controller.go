package controller

import (
	"encoding/json"
	"fmt"
	"github.com/ArtiomStartev/Go-Cache-API/structs"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func GetPost(c *fiber.Ctx) error {
	var post structs.Post

	id := c.Params("id")
	if id == "" {
		fmt.Println("Error: Missing ID parameter")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Missing ID parameter",
		})
	}

	res, err := http.Get("https://jsonplaceholder.typicode.com/posts/" + id)
	if err != nil || res.StatusCode != http.StatusOK {
		fmt.Println("Error fetching post: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}
	defer res.Body.Close()

	if err = json.NewDecoder(res.Body).Decode(&post); err != nil {
		fmt.Println("Error decoding response body: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(post)
}
