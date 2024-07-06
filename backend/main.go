package main

import (
	"cms/database"
	"cms/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New();
	app.Use(logger.New());
	database.Connect();
	app.Static("/media","./media");
	routes.Routes(app);
	app.Listen(":4000");
}