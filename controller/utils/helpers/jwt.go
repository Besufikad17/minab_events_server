package helpers

import (
	"errors"
	"os"
	"strconv"
	"time"

	models "github.com/Besufikad17/minab_events/models"
	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(user models.User, rememberMe bool) (string, error) {
	var expiryDate int64

	if rememberMe {
		expiryDate = time.Now().Add(time.Hour * 24 * 7).Unix()
	} else {
		expiryDate = time.Now().Add(time.Hour * 24).Unix()
	}

	jwtSecret := os.Getenv("ACCESS_TOKEN_SECRET")
	if jwtSecret == "" {
		return "", errors.New("ACCESS_TOKEN_SECRET is not set")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":          user.ID,
			"firstName":   user.FirstName,
			"lastName":    user.LastName,
			"email":       user.Email,
			"phoneNumber": user.PhoneNumber,
			"exp":         expiryDate,
			"https://hasura.io/jwt/claims": map[string]interface{}{
				"x-hasura-default-role":  "user",
				"x-hasura-allowed-roles": [2]string{"user", "admin"},
				"x-hasura-user-id":       strconv.Itoa(user.ID),
			},
		})
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyToken(tokenString string) error {
	jwtSecret := os.Getenv("ACCESS_TOKEN_SECRET")
	if jwtSecret == "" {
		return errors.New("ACCESS_TOKEN_SECRET is not set")
	}

	token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
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
