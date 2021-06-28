package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sogilis/poc-wiismile/controller"
	"github.com/sogilis/poc-wiismile/security"
	"log"
	"strconv"
)

const port = 8080

func main() {
	// Creates a router without any middleware by default
	r := gin.New()

	// CORS handling
	r.Use(security.CorsHandler())

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	// Route Handlers / Endpoints
	r.GET("/todoAPI/todos", controller.GetCatalogs)
	r.POST("/todoAPI/todo", controller.SaveCatalog)
	r.GET("/todoAPI/todo/:id", controller.GetCatalogById)
	r.DELETE("/todoAPI/todo/:id", controller.DeleteCatalog)

	log.Printf("todo-project running at 'http://localhost:%d'", port)
	err := r.Run(":" + strconv.Itoa(port))
	if err != nil {
		log.Fatal(err)
	}

}
