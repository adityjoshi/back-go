// //

// package helpers

// import (
// 	"fmt"
// 	"os"
// 	"time"

// 	"github.com/golang-jwt/jwt/v5"
// )

// func jwtGenerator(userID int, userType string) (string, error) {
// 	// Retrieve the JWT secret dynamically
// 	jwtSecret := []byte(os.Getenv("JWTSECRET"))

// 	// Create a new JWT token with user payload
// 	token := jwt.New(jwt.SigningMethodHS256)
// 	claims := token.Claims.(jwt.MapClaims)
// 	claims["user"] = map[string]interface{}{
// 		"user_id": userID,
// 		"type":    userType,
// 	}
// 	claims["exp"] = time.Now().Add(time.Hour).Unix()

// 	// Sign the token with the secret key and return
// 	tokenString, err := token.SignedString(jwtSecret)
// 	if err != nil {
// 		return "", err
// 	}
// 	return tokenString, nil
// }

// func jwtDecoder(tokenString string) (jwt.MapClaims, error) {
// 	// Retrieve the JWT secret dynamically
// 	jwtSecret := []byte(os.Getenv("JWTSECRET"))

// 	// Parse and validate the JWT token
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		return jwtSecret, nil
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Extract and return the claims if the token is valid
// 	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
// 		return claims, nil
// 	} else {
// 		return nil, fmt.Errorf("invalid token")
// 	}
// }

package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtSecret = []byte(os.Getenv("JWTSECRET"))

func GenerateJWT(userID int, userType string, blockID uint) (string, error) {
	// Create a new JWT token with user payload
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user"] = map[string]interface{}{
		"user_id":  userID,
		"type":     userType,
		"block_id": blockID,
	}
	claims["exp"] = time.Now().Add(time.Hour).Unix()

	// Sign the token with the secret key and return
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func DecodeJWT(tokenString string) (jwt.MapClaims, error) {
	// Parse and validate the JWT token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}

	// Extract and return the claims if the token is valid
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}
