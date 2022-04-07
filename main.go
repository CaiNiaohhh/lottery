package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	. "lottry/utils"
	"time"
)

func main() {
	InitDB()
	if err := InitRedisClient(); err != nil {
		fmt.Println(fmt.Sprintf("conn redis fail, err is %v", err))
	}
	go func() {
		SyncRedisAndMysql(time.Second * 10)
	}()
	router := gin.Default()
	router.GET("/set", Set)
	router.GET("/get", Get)
	router.GET("/getLuckUserList", GetRedisData)
	router.Run()

}