package handlers

import (
	"cms/database"
	"cms/models"
	"context"
	"os"
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UploadMedia(c *fiber.Ctx) error{
	file,err := c.FormFile("file")
	if err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"Cannot get file"})
	}
	filename := primitive.NewObjectID().Hex() + filepath.Ext(file.Filename);
	filepath := filepath.Join("../media",filename);

	if err := c.SaveFile(file,filepath); err!=nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":"Cannot save file"})
	}

	media := &models.MediaModel{
		ID : primitive.NewObjectID().Hex(),
		Filename: filename,
		URL: "/media" + filename,
		CreatedAt: time.Now(),
	}

	collection := database.DB.Database("cms").Collection("media")
	ctx,cancel := context.WithTimeout(context.Background(),5*time.Second)
	defer cancel();

	_,err = collection.InsertOne(ctx,media);
	if err!=nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":"Cannot save the file to database"})
	}
	return c.JSON(media);
}

func GetMediaById(c *fiber.Ctx) error{
	ID := c.Params("id");
	collection := database.DB.Database("cms").Collection("media");
	ctx,cancel := context.WithTimeout(context.Background(),5*time.Second)
	defer cancel();

	var media models.MediaModel;
	err := collection.FindOne(ctx,bson.M{"_id":ID}).Decode(&media);
	if err!=nil{
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error":"Media not found"})
	}
	return c.JSON(media);
}

func DeleteMedia(c *fiber.Ctx) error{
	ID := c.Params("id");
	collection := database.DB.Database("cms").Collection("media");
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
	var media models.MediaModel
    err := collection.FindOne(ctx, bson.M{"_id": ID}).Decode(&media)
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Media not found"})
    }
	if err := os.Remove(filepath.Join("./media", media.Filename)); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot delete file",})
    }

    _, err = collection.DeleteOne(ctx, bson.M{"_id": ID})
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot delete media from database",})
    }

    return c.SendStatus(fiber.StatusNoContent)
}