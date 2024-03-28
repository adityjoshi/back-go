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
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	database.InitDatabase()
	defer database.CloseDatabase()

	router := gin.Default()
	router.Use(cors.Default())
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("session", store))

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
