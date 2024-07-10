package handlers


import (
	"cms/database"
	"cms/models"
	"context"
	"time"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	var contents []models.ContentModel;

	for response.Next(ctx){
		var content models.ContentModel
		response.Decode(&content)
		contents = append(contents,content)
	}
	return c.JSON(contents)
}

func CreateContent(c *fiber.Ctx) error{
	var content models.ContentModel;
	if err := c.BodyParser(content); err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"Cannot parse the content"})
	}
	content.ID = primitive.NewObjectID().Hex();
	content.AuthorID = c.Locals("userID").(string)
	content.CreatedAt = time.Now();
	content.UpdatedAt = time.Now();

	collection := database.DB.Database("cms").Collection("contents");
	ctx,cancel := context.WithTimeout(context.Background(),5*time.Second)
	defer cancel();

	_,err := collection.InsertOne(ctx,content);
	if err!=nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":"Cannot create content"})
	}
	return c.JSON(content);
}

func GetContentById(c *fiber.Ctx) error{
	ID := c.Params("id");
	collection := database.DB.Database("cms").Collection("contents")
	ctx,cancel := context.WithTimeout(context.Background(),5*time.Second)
	defer cancel();
	var content models.ContentModel;
	err := collection.FindOne(ctx,bson.M{"_id" : ID}).Decode(&content);
	if err!=nil{
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error":"Content not found"})
	}
	return c.JSON(content);
}

func UpdateContent(c *fiber.Ctx) error{
	ID := c.Params("id");
	userID := c.Locals("userID").(string);
	var content models.ContentModel;
	if err := c.BodyParser(content);err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"Cannot parse JSON"})
	}
	collection := database.DB.Database("cms").Collection("contents")
	ctx,cancel := context.WithTimeout(context.Background(),5*time.Second)
	defer cancel();

	var existingContent models.ContentModel;
	err := collection.FindOne(ctx,bson.M{"_id":ID}).Decode(&existingContent)
	if err!=nil{
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error":"Could not find the content"})
	}
	if existingContent.AuthorID != userID{
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "You are not authorized to update this content",})
	}
	content.UpdatedAt = time.Now();
	_,err = collection.UpdateOne(ctx,bson.M{"_id":ID},bson.M{"$set":ID})
	if err!=nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":"Cannot update content"})
	}
	return c.JSON(content);
}

func DeleteContent(c *fiber.Ctx) error{
	ID := c.Params("id");
	collection := database.DB.Database("cms").Collection("contents")
	ctx,cancel := context.WithTimeout(context.Background(),5*time.Second)
	defer cancel();

	var content models.ContentModel;
	err := collection.FindOne(ctx,bson.M{"_id":ID}).Decode(&content)
	if err!=nil{
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error":"Content not found"})
	}
	if content.AuthorID != c.Locals("userID").(string){
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "You are not authorized to delete this content",})
}

_,err = collection.DeleteOne(ctx,bson.M{"_id":ID})
if err!=nil{
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":"Cannot delet content"})
}
return c.JSON(fiber.Map{"message":"Deleted successfully"})
}

