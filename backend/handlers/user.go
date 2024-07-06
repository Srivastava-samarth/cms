package handlers

import (
	"cms/utils"
	"cms/database"
	"cms/models"
	"context"
	"time"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error{
	var user models.UserModel
	if err := c.BodyParser(user);err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"Cannot parse the json"})
	}
	password,_ := bcrypt.GenerateFromPassword([]byte(user.Password),14)
	user.Password = string(password)

	user.ID = primitive.NewObjectID()
	user.CreatedAt = time.Now().Unix()
	user.UpdatedAt = time.Now().Unix()

	collection := database.DB.Database("content-management-system").Collection("users")
	ctx,cancel := context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()

	_,err := collection.InsertOne(ctx,user);
	if err!=nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":"Error in creating the user"})
	}
	return c.JSON(user);
}

func Login(c *fiber.Ctx) error{
	var res models.Login
	if err := c.BodyParser(res); err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"Error in parsing the json"})
	}
	collection := database.DB.Database("conetnt-management-system").Collection("users")
	ctx,cancel := context.WithTimeout(context.Background(),5*time.Second);
	defer cancel();
	var user models.UserModel;
	err := collection.FindOne(ctx,bson.M{"email":res.Email}).Decode(&user)
	if err!=nil{
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error":"Invalid email or password"})
	} 
	err = bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(res.Password))
	if err!=nil{
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error":"Invalid email or password"})
	}
	token, err := utils.GenerateJWT(user.ID.Hex())
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Cannot generate token",
        })
	}
	return c.JSON(fiber.Map{"token":token,})
}