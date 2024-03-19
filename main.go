// package main

// import (
// 	"net/http"

// 	"BACKEND-GO/routes"

// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	router := gin.New()
// 	router.Use(gin.Logger())

// 	server := &http.Server{
// 		Addr:    ":2426",
// 		Handler: router,
// 	}
// 	routes.UserRoute(router)

// 	server.ListenAndServe()
// }

package main

import (
	"log"
	"net/http"

	"BACKEND-GO/routes" // Update with correct path

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Access environment variables
	//	jwtSecret := os.Getenv("JWTSECRET")

	router := gin.Default()
	routes.UserRoute(router)

	server := &http.Server{
		Addr:    ":2426",
		Handler: router,
	}

	log.Println("Server is running at :2426...")
	server.ListenAndServe()
}
