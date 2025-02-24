package router

import "github.com/gin-gonic/gin"

func Init() {
	// initialize the gin router
	router := gin.Default()

	// initialize the routes
	initializeRoutes(router)

	// run the server
	router.Run(":8080") // listen and serve on
}
