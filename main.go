// Define main as a module
package main

// Import needed libraries
import (
	"backend/my-little-kanvas/configs"
	"backend/my-little-kanvas/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// The start function of the module
func main() {
	// I guess this extracts the router object from gin.Default()
	router := gin.Default()

	// connect to mongodb cluster
	configs.ConnectDB()

	// setup routes
	router.Use(cors.Default())
	routes.TodoRoute(router)
	
	// Tell the server to run on port 8080
	router.Run("localhost:8080")
}
