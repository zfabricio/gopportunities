package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Initialize Router
	r := gin.Default() // := --> Declarar uma variável
	// Initialize Routes
	initializeRoutes(r)
	// Run the server
	r.Run(":3000")
}
