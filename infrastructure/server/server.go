package server

import "github.com/gin-gonic/gin"

func Serve() {
	router := gin.Default()
	Route(router)
	router.Run(":8080")
}
