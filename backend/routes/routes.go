package routes

import (
	"cms/handlers"
	"cms/middleware"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App){
	api := app.Group("/api")

	//Auth routes
	api.Post("/register",handlers.Register);
	api.Post("login",handlers.Login);

	api.Use(middleware.Protected);

	//content routes
	api.Get("/contents",handlers.GetAllContent)
	api.Post("/content",handlers.CreateContent)
	api.Post("/content/:id",handlers.GetContentById)
	api.Put("/content/:id",handlers.UpdateContent)
	api.Delete("/content/:id",handlers.DeleteContent)

	//media Routes
	api.Post("/media/upload",handlers.UploadMedia)
	api.Get("/media/:id",handlers.GetMediaById)
	api.Delete("/media/:id",handlers.DeleteMedia)
}