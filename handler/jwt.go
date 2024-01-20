package handler

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secretKey []byte

func generateRandomKey() (string, error) {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(key), nil
}

func InitJWTKey() {
	randomKey, err := generateRandomKey()
	if err != nil {
		fmt.Println("Error generating random key:", err)
		return
	}
	secretKey = []byte(randomKey)
	fmt.Println("Generated Random Key:", randomKey)

	// Generating a JWT token
	tokenString, err := generateJWTToken(1) // Assuming user ID is 1
	if err != nil {
		fmt.Println("Error generating JWT token:", err)
		return
	}
	fmt.Println("Generated JWT Token:", tokenString)
}

func generateJWTToken(userID uint) (string, error) {
	// Create a new token with a signing method
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims (payload) for the token
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = userID                               // Subject (user ID)
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() // Expiration time (1 hour from now)

	// Sign the token with the secret key
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func requiresJWTAuth(method, path string) bool {
	return (method == http.MethodPut || method == http.MethodDelete) && strings.HasPrefix(path, "/secure")
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the request requires JWT authentication based on method and path
		requiresAuth := requiresJWTAuth(r.Method, r.URL.Path)

		// If the request requires authentication, proceed with JWT validation
		if requiresAuth {
			// Get the token from the Authorization header
			tokenString := r.Header.Get("Authorization")
			if tokenString == "" {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			// Extract the token from the "Bearer " prefix
			tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

			// Validate the token
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				return secretKey, nil
			})
			if err != nil || !token.Valid {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
		}

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}
