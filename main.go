package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	. "lottry/utils"
	"time"
)

func main() {
	router := gin.Default()
	router.Use(Recover_)
	InitDB()
	if err := InitRedisClient(); err != nil {
		fmt.Println(fmt.Sprintf("conn redis fail, err is %v", err))
	}
	go func() {
		SyncRedisAndMysql(time.Second * 1)
	}()
	router.GET("/set", Set)
	router.GET("/get", Get)
	router.GET("/getLuckUserList", GetRedisData)
	router.Run()

}

func Recover_(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			//打印错误堆栈信息
			log.Printf("panic: %v\n", r)
			// 打印报错栈
			//debug.PrintStack()
			//封装通用json返回
			c.JSON(200,gin.H{
				"code":"4444",
				"msg":"服务器内部错误",
			})
		}
	}()
	//加载完 defer recover，继续后续接口调用
	// 主要是因为捕获异常要放在最后面使用的，所以加上next阻塞，先执行其他
	c.Next()
}