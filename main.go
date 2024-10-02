package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"yohanestatus/minecraft"
	"yohanestatus/ragnarok"
	"yohanestatus/database"
	"log"
)

func main() {
	app := fiber.New()

	// Menghubungkan ke database
	err := database.ConnectDB()
	if err != nil {
		log.Fatalf("Gagal terhubung ke database: %v", err)
	}
	defer database.CloseDB()

	// Menambahkan middleware logger
	app.Use(logger.New(logger.Config{
		Format: "${time} - ${ip} - ${url} - ${status} - ${method} - ${latency} - ${ua}\n",
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "Asia/Jakarta",
	}))

	app.Get("/status", handleStatus)

	// Menambahkan rute untuk redirect ke datenshi.pw
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("https://datenshi.pw", 301)
	})

	app.Listen(":3000")
}

func handleStatus(c *fiber.Ctx) error {
	games := c.Query("games")

	switch games {
	case "0":
		return minecraft.HandleMCStatus(c)
	case "1":
		return ragnarok.HandleRagnarokStatus(c)
	case "2":
		// Implementasi untuk game 2 bisa ditambahkan di sini
		return c.SendString("Implementasi untuk game 2 belum tersedia")
	default:
		return c.Status(400).SendString("Nilai games tidak valid")
	}
}
