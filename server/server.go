package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)




func main() {
	err : godotenv.Load(".env")
	if err != nil{
		log.Fatal(err)
	}

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Pro mike is a Go "})
	})
	app.Listen(":8000")
}