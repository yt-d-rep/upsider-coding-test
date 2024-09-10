package server

import "github.com/gin-gonic/gin"

func Serve() {
	router := gin.Default()
	route(router)
	router.Run(":8080")
}
