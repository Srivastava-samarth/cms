package handlers

import (
	"cms/database"
	"cms/models"
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllContent(c *fiber.Ctx) error{
	collection := database.DB.Database("content-management-system").Collection("contents")
	ctx,cancel := context.WithTimeout(context.Background(),10*time.Second)
	defer cancel();

	response,err := collection.Find(ctx,bson.M{})
	if err!=nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":"Cannot List Contents"})
	}
	defer response.Close(ctx);

	contents := make([]models.ContentModel, 0)

	for response.Next(ctx){
		var content models.ContentModel
		response.Decode(&content)
		contents = append(contents,content)
	}
	return c.JSON(contents)
}

func CreateContent(c *fiber.Ctx) error{
	
}

func GetContentById(c *fiber.Ctx) error{

}

func UpdateContent(c *fiber.Ctx) error{

}

func DeleteContent(c *fiber.Ctx) error{

}

