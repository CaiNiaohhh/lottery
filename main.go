package main

import (
	"github.com/gin-gonic/gin"
	. "utils/comm/utils"
)

func main() {
	InitDB()
	router := gin.Default()
	router.GET("/set", Set)
	router.GET("/get", Get)
	router.Run()

}