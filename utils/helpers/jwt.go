package helpers

import (
	"errors"
	"log"
	"os"
	"time"

	models "github.com/Besufikad17/minab_events/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

var claims jwt.MapClaims

func CreateToken(user models.User, rememberMe bool) (string, error) {
	var expiryDate int64

	if rememberMe {
		expiryDate = time.Now().Add(time.Hour * 24 * 7).Unix()
	} else {
		expiryDate = time.Now().Add(time.Hour * 24).Unix()
	}

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	jwtSecret := os.Getenv("JWT_SECRET")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":          user.ID,
			"firstName":   user.FirstName,
			"lastName":    user.LastName,
			"email":       user.Email,
			"phoneNumber": user.PhoneNumber,
			"exp":         expiryDate,
		})
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", nil
	}
	return tokenString, nil
}

func VerifyToken(tokenString string) error {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	jwtSecret := os.Getenv("JWT_SECRET")

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return errors.New("invalid token")
	}

	return nil
}
