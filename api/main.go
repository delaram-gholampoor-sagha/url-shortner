package api

import (
	"fmt"
	"log"
	"os"

	"github.com/Delaram-Gholampoor-Sagha/url-shortner/api/routes"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetUpRoutes(app *fiber.App) {
	app.Get("/:url", routes.ResolveURL)
	app.Post("/api/v1", routes.ShortenURL)
}

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err)
	}

	app := fiber.New()
	app.Use(logger.New())

	SetUpRoutes(app)

	log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}
