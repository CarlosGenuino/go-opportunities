package router

import "github.com/gin-gonic/gin"

func Initialize() {

	//Initilizing gin Router
	router := gin.Default()

	//Initilizing the http routes
	initializeRoutes(router)

	//Run the Server
	router.Run(":8080")

}
