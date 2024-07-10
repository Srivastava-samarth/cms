package utils


import (
	"os"
	"time"
	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = os.Getenv("JWT_SECRET");

type Claims struct{
	UserID string `json:"user_id`
	jwt.StandardClaims
}

func GenerateJWT(userID string) (string, error) {
	expTime := time.Now().Add(24 * time .Hour)
	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	return token.SignedString(jwtSecret)
}

func ParseJWT(tokenString string) (*Claims,error){
	claims := &Claims{}

	token,err := jwt.ParseWithClaims(tokenString,claims,func(t *jwt.Token) (interface{}, error) {
		return jwtSecret,nil;
	})
	if err!=nil{
		return nil,err;
	}
	if !token.Valid{
		return nil,err;
	}
	return claims,nil;
}