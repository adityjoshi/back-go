package main

import (
	"net/http"

	"BACKEND-GO/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.Use(gin.Logger())

	server := &http.Server{
		Addr:    ":2426",
		Handler: router,
	}
	routes.UserRoute(router)

	server.ListenAndServe()
}
