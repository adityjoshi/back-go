package main

import (
	"log"
	"net/http"

	"BACKEND-GO/database"
	"BACKEND-GO/initiliazers"
	"BACKEND-GO/routes" // Update with correct path

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	initiliazers.LoadEnvVariable()
}
func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	database.InitDatabase()
	defer database.CloseDatabase()
	// Access environment variables
	//	jwtSecret := os.Getenv("JWTSECRET")

	router := gin.Default()
	router.Use(cors.Default())
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("session", store))

	//router.Static("/", "BACKEND-GO/Backend/frontend/dist")
	routes.UserRoute(router)
	routes.ComplaintRoutes(router)
	routes.StudentRoutes(router)

	server := &http.Server{
		Addr:    ":2426",
		Handler: router,
	}

	log.Println("Server is running at :2426...")
	server.ListenAndServe()
}

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
