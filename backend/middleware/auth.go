package middleware


import (
	"cms/utils"

	"github.com/gofiber/fiber/v2"
)

func Protected() fiber.Handler{
	return func(c *fiber.Ctx) error {
		token := c.Get("Authorization")
		if token == ""{
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error":"Unauthorized"})
		}
		claims,err := utils.ParseJWT(token);
		if err!=nil{
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error":"Unauthorized"})
		}
		c.Locals("userID",claims.UserID)	
		return c.Next();
	}
}